package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type CreateCheckoutRequest struct {
	PlayerID  string `json:"player_id"`
	PlanType  string `json:"plan_type"` // e.g., "player_early_upfront", "player_early_installments"
}

func CreateCheckoutSession(c *fiber.Ctx) error {
	var req CreateCheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 1. Validate Player exists
	var player models.Player
	if err := database.DB.First(&player, "id = ?", req.PlayerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}
	
	// Get household ID of the authenticated user
	sub, _ := c.Locals("user_id").(string)
	var household models.Household
	database.DB.First(&household, "owner_id = ?", sub)

	// 2. Configure Pricing Logic
	var priceData *stripe.CheckoutSessionLineItemPriceDataParams
	var mode string

	switch req.PlanType {
	case "player_early_upfront":
		mode = string(stripe.CheckoutSessionModePayment)
		priceData = &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String("usd"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name: stripe.String("Player Roster Spot (Early Pricing)"),
			},
			UnitAmount: stripe.Int64(150000), // $1500.00
		}
	case "player_early_installments":
		mode = string(stripe.CheckoutSessionModeSubscription)
		priceData = &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String("usd"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name: stripe.String("Player Roster Spot (15 Weekly Installments)"),
			},
			UnitAmount: stripe.Int64(10000), // $100.00
			Recurring: &stripe.CheckoutSessionLineItemPriceDataRecurringParams{
				Interval: stripe.String("week"),
			},
		}
	case "volunteer_200_upfront":
		mode = string(stripe.CheckoutSessionModePayment)
		priceData = &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String("usd"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name: stripe.String("Volunteer Spot Deposit"),
			},
			UnitAmount: stripe.Int64(20000), // $200.00
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid plan type"})
	}

	// 3. Create Stripe Checkout Session
	domain := "http://localhost:5173" // Frontend URL for dev
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode: stripe.String(mode),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: priceData,
				Quantity:  stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(domain + "/dashboard?success=true"),
		CancelURL:  stripe.String(domain + "/account?canceled=true"),
		Metadata: map[string]string{
			"player_id":    player.ID,
			"household_id": household.ID,
			"plan_type":    req.PlanType,
		},
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create checkout session"})
	}

	return c.JSON(fiber.Map{"url": s.URL})
}
