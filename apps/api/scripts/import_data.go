package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/limitless-hoops/api/database"
	"github.com/limitless-hoops/api/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func parseTime(val string) time.Time {
	if val == "" {
		return time.Time{}
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return time.Now() // fallback
		}
	}
	return t
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	f, err := excelize.OpenFile("../../docs/LH Recruiting Machine.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Helper to get rows as array of maps based on headers
	getMapRows := func(sheetName string) ([]map[string]string, error) {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, err
		}
		if len(rows) < 2 {
			return nil, nil // empty or just headers
		}
		headers := rows[0]
		var result []map[string]string
		for _, row := range rows[1:] {
			rowMap := make(map[string]string)
			for i, val := range row {
				if i < len(headers) {
					rowMap[headers[i]] = val
				}
			}
			result = append(result, rowMap)
		}
		return result, nil
	}

	// 1. Teams
	log.Println("Importing Teams...")
	teamsData, _ := getMapRows("Teams")
	for _, row := range teamsData {
		team := models.Team{
			TeamCode:    row["TeamCode"],
			Area:        row["Area"],
			TeamName:    row["TeamName"],
			GradeBand:   row["GradeBand"],
			StartGrade:  row["StartGrade"],
			EndGrade:    row["EndGrade"],
			Slug:        row["Slug"],
			PageURL:     row["PageURL"],
			WhatsAppURL: row["WhatsAppURL"],
			CoachID:     row["CoachID"],
			Status:      row["Status"],
		}
		database.DB.Save(&team)
	}

	// 2. People -> Household + User + Player
	log.Println("Importing People...")
	peopleData, _ := getMapRows("People")
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, row := range peopleData {
			personID := row["PersonID"] // THIS MUST BE THE PLAYER ID
			if personID == "" {
				continue
			}

			// Create Household
			householdID := uuid.New().String()
			household := models.Household{ID: householdID}
			if err := tx.Save(&household).Error; err != nil {
				return err
			}

			// Create User (Parent)
			userID := uuid.New().String()
			user := models.User{
				ID:        userID,
				FirstName: row["ParentName"], // Will have to split name later if needed, leaving as First for now
			}
			if err := tx.Save(&user).Error; err != nil {
				return err
			}

			// Seed UserEmail for the parent (primary, verified)
			if emailVal := row["Email"]; emailVal != "" {
				userEmail := models.UserEmail{
					ID:             uuid.New().String(),
					UserID:         userID,
					Email:          emailVal,
					IsPrimary:      true,
					IsVerified:     true,
					SourceProvider: "manual",
					CreatedAt:      time.Now(),
				}
				if err := tx.Save(&userEmail).Error; err != nil {
					return err
				}
				for _, cat := range []string{"billing", "events", "announcements", "security"} {
					pref := models.NotificationPreference{
						ID:         uuid.New().String(),
						UserID:     userID,
						TargetType: "email",
						TargetID:   userEmail.ID,
						Category:   cat,
						Enabled:    true,
					}
					if err := tx.Save(&pref).Error; err != nil {
						return err
					}
				}
			}

			// Seed UserPhone for the parent (primary, verified)
			if phoneVal := row["Phone"]; phoneVal != "" {
				userPhone := models.UserPhone{
					ID:         uuid.New().String(),
					UserID:     userID,
					Phone:      phoneVal,
					IsPrimary:  true,
					IsVerified: true,
					CreatedAt:  time.Now(),
				}
				if err := tx.Save(&userPhone).Error; err != nil {
					return err
				}
				for _, cat := range []string{"billing", "events", "announcements", "security"} {
					pref := models.NotificationPreference{
						ID:         uuid.New().String(),
						UserID:     userID,
						TargetType: "phone",
						TargetID:   userPhone.ID,
						Category:   cat,
						Enabled:    true,
					}
					if err := tx.Save(&pref).Error; err != nil {
						return err
					}
				}
			}

			// Link User to Household
			hu := models.HouseholdUser{
				HouseholdID: householdID,
				UserID:      userID,
				Role:        "Primary",
			}
			if err := tx.Save(&hu).Error; err != nil {
				return err
			}

			// Create Player (assigning personID so all legacy references match)
			player := models.Player{
				ID:                  personID,
				HouseholdID:         householdID,
				FirstName:           row["PlayerName"],
				Grade2027_28:        row["Grade2027_28"],
				ReferenceSchoolYear: row["ReferenceSchoolYear"],
				GraduationYear:      row["GraduationYear"],
				PublicDisplayName:   row["PublicDisplayName"],
				ReferralCode:        row["ReferralCode"],
				Venmo:               row["Venmo"],
				Status:              row["Status"],
				Notes:               row["Notes"],
			}
			if err := tx.Save(&player).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal("Failed to import People:", err)
	}

	// 3. Applications
	log.Println("Importing Applications...")
	appData, _ := getMapRows("Applications")
	for _, row := range appData {
		app := models.Application{
			ApplicationID:            row["ApplicationID"],
			PersonID:                 row["PersonID"],
			TeamCode:                 row["TeamCode"],
			Grade2027_28:             row["Grade2027_28"],
			Path:                     row["Path"],
			VerificationStatus:       row["VerificationStatus"],
			VerificationToken:        row["VerificationToken"],
			VerificationExpiresAt:    row["VerificationExpiresAt"],
			VerificationReminderSent: row["VerificationReminderSent"],
			RosterStatus:             row["RosterStatus"],
			SeatID:                   row["SeatID"],
			PriorityTimestamp:        row["PriorityTimestamp"],
			PaymentStatus:            row["PaymentStatus"],
			PaymentPlan:              row["PaymentPlan"],
			PaymentConfirmedAt:       row["PaymentConfirmedAt"],
			SelectedReferrerCode:     row["SelectedReferrerCode"],
			NextBestAction:           row["NextBestAction"],
			SequenceIndex:            row["SequenceIndex"],
			NextSequenceAt:           row["NextSequenceAt"],
			MonthlyConfirmStatus:     row["MonthlyConfirmStatus"],
			LastConfirmedAt:          row["LastConfirmedAt"],
			RemovedAt:                row["RemovedAt"],
			Notes:                    row["Notes"],
			MembershipID:             row["MembershipID"],
			PaymentContext:           row["PaymentContext"],
			PaymentAccessTokenHash:   row["PaymentAccessTokenHash"],
			PaymentLinkStatus:        row["PaymentLinkStatus"],
			PaymentLinkSentAt:        row["PaymentLinkSentAt"],
			PromoCode:                row["PromoCode"],
			PromoType:                row["PromoType"],
		}
		database.DB.Save(&app)
	}

	// 4. Memberships
	log.Println("Importing Memberships...")
	memData, _ := getMapRows("Memberships")
	for _, row := range memData {
		mem := models.Membership{
			MembershipID:            row["MembershipID"],
			ApplicationID:           row["ApplicationID"],
			PersonID:                row["PersonID"],
			TeamCode:                row["TeamCode"],
			PlayerName:              row["PlayerName"],
			ParentName:              row["ParentName"],
			ParentEmail:             row["ParentEmail"],
			PaymentContext:          row["PaymentContext"],
			PlanType:                row["PlanType"],
			PaymentSource:           row["PaymentSource"],
			AnnualPrice:             row["AnnualPrice"],
			InstallmentAmount:       row["InstallmentAmount"],
			InstallmentsPerYear:     row["InstallmentsPerYear"],
			MembershipStart:         row["MembershipStart"],
			MembershipEnd:           row["MembershipEnd"],
			RenewalDate:             row["RenewalDate"],
			CurrentMembershipYear:   row["CurrentMembershipYear"],
			StripeCustomerID:        row["StripeCustomerID"],
			StripeSubscriptionID:    row["StripeSubscriptionID"],
			StripeCheckoutSessionID: row["StripeCheckoutSessionID"],
			DefaultPaymentMethodID:  row["DefaultPaymentMethodID"],
			NextInstallmentNumber:   row["NextInstallmentNumber"],
			NextPaymentDate:         row["NextPaymentDate"],
			AmountPaidCurrentYear:   row["AmountPaidCurrentYear"],
			BalanceCurrentYear:      row["BalanceCurrentYear"],
			PaymentStatus:           row["PaymentStatus"],
			MembershipStatus:        row["MembershipStatus"],
			LastSuccessfulPaymentAt: row["LastSuccessfulPaymentAt"],
			LastFailedPaymentAt:     row["LastFailedPaymentAt"],
			GraceExpiresAt:          row["GraceExpiresAt"],
			CancelNowAt:             row["CancelNowAt"],
			CancelAtRenewal:         row["CancelAtRenewal"],
			CancelledAt:             row["CancelledAt"],
			Notes:                   row["Notes"],
			BillingMode:             row["BillingMode"],
			ExternalBillingProvider: row["ExternalBillingProvider"],
			PromoCode:               row["PromoCode"],
			ExternalBillingStatus:   row["ExternalBillingStatus"],
			BillingAccessTokenHash:  row["BillingAccessTokenHash"],
			BillingLinkExpiresAt:    row["BillingLinkExpiresAt"],
			BillingLinkStatus:       row["BillingLinkStatus"],
		}
		database.DB.Save(&mem)
	}
	
	// 5. SeatInventory
	log.Println("Importing SeatInventory...")
	seatData, _ := getMapRows("SeatInventory")
	for _, row := range seatData {
		seat := models.SeatInventory{
			SeatID:        row["SeatID"],
			TeamCode:      row["TeamCode"],
			Grade2027_28:  row["Grade2027_28"],
			SeatType:      row["SeatType"],
			SlotNo:        row["SlotNo"],
			SeatStatus:    row["SeatStatus"],
			ApplicationID: row["ApplicationID"],
			PersonID:      row["PersonID"],
			HoldExpiresAt: row["HoldExpiresAt"],
			Notes:         row["Notes"],
		}
		database.DB.Save(&seat)
	}

	log.Println("✅ Data migration complete!")
}
