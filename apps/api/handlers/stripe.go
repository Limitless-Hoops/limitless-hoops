package handlers

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
	"github.com/stripe/stripe-go/v76/webhook"
)

func HandleStripeWebhook(c *fiber.Ctx) error {
	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	signature := c.Get("Stripe-Signature")

	payload := c.Body()

	// Verify signature
	event, err := webhook.ConstructEvent(payload, signature, webhookSecret)
	if err != nil {
		log.Printf("Webhook signature verification failed: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid signature")
	}

	log.Printf("Received Stripe Event: %s", event.Type)

	// Persist the event exactly as the client's spreadsheet tracks it
	stripeEvent := models.StripeEvent{
		StripeEventID:    event.ID,
		EventType:        string(event.Type),
		CreatedAtStripe:  time.Unix(event.Created, 0).Format(time.RFC3339),
		ReceivedAt:       time.Now(),
		ProcessingStatus: "pending",
		PayloadHash:      string(payload), // Storing payload briefly for debugging or parsing later
	}

	// Basic object ID extraction based on standard event types
	var obj map[string]interface{}
	if err := json.Unmarshal(event.Data.Raw, &obj); err == nil {
		if id, ok := obj["id"].(string); ok {
			stripeEvent.ObjectID = id
		}
	}

	// Save to database
	if err := database.DB.Create(&stripeEvent).Error; err != nil {
		log.Printf("Failed to save StripeEvent to DB: %v", err)
		// We still return 200 OK to Stripe so it doesn't endlessly retry if it's a DB issue on our end we can't fix now
	}

	// Task 6: Future scaffold for core events
	switch event.Type {
	case "checkout.session.completed":
		log.Println("Handling checkout.session.completed")
		var session struct {
			Metadata map[string]string `json:"metadata"`
			Customer string            `json:"customer"`
			Mode     string            `json:"mode"`
		}
		if err := json.Unmarshal(event.Data.Raw, &session); err == nil {
			playerID := session.Metadata["player_id"]
			planType := session.Metadata["plan_type"]
			
			if playerID != "" {
				membership := models.Membership{
					MembershipID:            "mem_" + time.Now().Format("20060102150405"),
					PersonID:                playerID,
					PlanType:                planType,
					StripeCustomerID:        session.Customer,
					StripeCheckoutSessionID: stripeEvent.ObjectID,
					PaymentStatus:           "paid",
					MembershipStatus:        "active",
					MembershipStart:         time.Now().Format(time.RFC3339),
				}
				database.DB.Create(&membership)
				log.Printf("Created membership for player %s", playerID)
			}
		}
	case "invoice.paid":
		log.Println("Handling invoice.paid")
		// TODO: Log to PaymentLedger, update Membership balance
	case "invoice.payment_failed":
		log.Println("Handling invoice.payment_failed")
		// TODO: Log to PaymentLedger, update Membership status to past_due
	case "customer.subscription.deleted":
		log.Println("Handling customer.subscription.deleted")
		// TODO: Cancel membership, free SeatInventory
	}

	// Mark processed in our DB
	database.DB.Model(&models.StripeEvent{}).
		Where("stripe_event_id = ?", event.ID).
		Updates(map[string]interface{}{
			"processing_status": "processed",
			"processed_at":      time.Now().Format(time.RFC3339),
		})

	return c.SendStatus(fiber.StatusOK)
}
