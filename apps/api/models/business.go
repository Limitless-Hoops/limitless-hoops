package models

import (
	"time"
)

type Team struct {
	TeamCode    string `gorm:"primaryKey;type:varchar(255)" json:"team_code"`
	Area        string `gorm:"type:varchar(255)" json:"area"`
	TeamName    string `gorm:"type:varchar(255)" json:"team_name"`
	GradeBand   string `gorm:"type:varchar(255)" json:"grade_band"`
	StartGrade  string `gorm:"type:varchar(255)" json:"start_grade"`
	EndGrade    string `gorm:"type:varchar(255)" json:"end_grade"`
	Slug        string `gorm:"type:varchar(255)" json:"slug"`
	PageURL     string `gorm:"type:text" json:"page_url"`
	WhatsAppURL string `gorm:"type:text" json:"whatsapp_url"`
	CoachID     string `gorm:"type:varchar(255)" json:"coach_id"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
}

type Application struct {
	ApplicationID            string    `gorm:"primaryKey;type:varchar(255)" json:"application_id"`
	PersonID                 string    `gorm:"type:varchar(255);index" json:"person_id"` // Matches Player.ID
	TeamCode                 string    `gorm:"type:varchar(255);index" json:"team_code"`
	Grade2027_28             string    `gorm:"type:varchar(255)" json:"grade_2027_28"`
	Path                     string    `gorm:"type:varchar(255)" json:"path"`
	VerificationStatus       string    `gorm:"type:varchar(255)" json:"verification_status"`
	VerificationToken        string    `gorm:"type:varchar(255)" json:"verification_token"`
	VerificationExpiresAt    string    `gorm:"type:varchar(255)" json:"verification_expires_at"` // Keeping string if legacy stored it weirdly, but ideally time.Time
	VerificationReminderSent string    `gorm:"type:varchar(255)" json:"verification_reminder_sent"`
	RosterStatus             string    `gorm:"type:varchar(255)" json:"roster_status"`
	SeatID                   string    `gorm:"type:varchar(255)" json:"seat_id"`
	PriorityTimestamp        string    `gorm:"type:varchar(255)" json:"priority_timestamp"`
	PaymentStatus            string    `gorm:"type:varchar(255)" json:"payment_status"`
	PaymentPlan              string    `gorm:"type:varchar(255)" json:"payment_plan"`
	PaymentConfirmedAt       string    `gorm:"type:varchar(255)" json:"payment_confirmed_at"`
	SelectedReferrerCode     string    `gorm:"type:varchar(255)" json:"selected_referrer_code"`
	NextBestAction           string    `gorm:"type:varchar(255)" json:"next_best_action"`
	SequenceIndex            string    `gorm:"type:varchar(255)" json:"sequence_index"`
	NextSequenceAt           string    `gorm:"type:varchar(255)" json:"next_sequence_at"`
	MonthlyConfirmStatus     string    `gorm:"type:varchar(255)" json:"monthly_confirm_status"`
	LastConfirmedAt          string    `gorm:"type:varchar(255)" json:"last_confirmed_at"`
	RemovedAt                string    `gorm:"type:varchar(255)" json:"removed_at"`
	Notes                    string    `gorm:"type:text" json:"notes"`
	MembershipID             string    `gorm:"type:varchar(255)" json:"membership_id"`
	PaymentContext           string    `gorm:"type:varchar(255)" json:"payment_context"`
	PaymentAccessTokenHash   string    `gorm:"type:text" json:"payment_access_token_hash"`
	PaymentLinkStatus        string    `gorm:"type:varchar(255)" json:"payment_link_status"`
	PaymentLinkSentAt        string    `gorm:"type:varchar(255)" json:"payment_link_sent_at"`
	PromoCode                string    `gorm:"type:varchar(255)" json:"promo_code"`
	PromoType                string    `gorm:"type:varchar(255)" json:"promo_type"`
	CreatedAt                time.Time `json:"created_at"`
}

type Membership struct {
	MembershipID            string    `gorm:"primaryKey;type:varchar(255)" json:"membership_id"`
	ApplicationID           string    `gorm:"type:varchar(255);index" json:"application_id"`
	PersonID                string    `gorm:"type:varchar(255);index" json:"person_id"`
	TeamCode                string    `gorm:"type:varchar(255)" json:"team_code"`
	PlayerName              string    `gorm:"type:varchar(255)" json:"player_name"`
	ParentName              string    `gorm:"type:varchar(255)" json:"parent_name"`
	ParentEmail             string    `gorm:"type:varchar(255)" json:"parent_email"`
	PaymentContext          string    `gorm:"type:varchar(255)" json:"payment_context"`
	PlanType                string    `gorm:"type:varchar(255)" json:"plan_type"`
	PaymentSource           string    `gorm:"type:varchar(255)" json:"payment_source"`
	AnnualPrice             string    `gorm:"type:varchar(255)" json:"annual_price"`
	InstallmentAmount       string    `gorm:"type:varchar(255)" json:"installment_amount"`
	InstallmentsPerYear     string    `gorm:"type:varchar(255)" json:"installments_per_year"`
	MembershipStart         string    `gorm:"type:varchar(255)" json:"membership_start"`
	MembershipEnd           string    `gorm:"type:varchar(255)" json:"membership_end"`
	RenewalDate             string    `gorm:"type:varchar(255)" json:"renewal_date"`
	CurrentMembershipYear   string    `gorm:"type:varchar(255)" json:"current_membership_year"`
	StripeCustomerID        string    `gorm:"type:varchar(255)" json:"stripe_customer_id"`
	StripeSubscriptionID    string    `gorm:"type:varchar(255)" json:"stripe_subscription_id"`
	StripeCheckoutSessionID string    `gorm:"type:varchar(255)" json:"stripe_checkout_session_id"`
	DefaultPaymentMethodID  string    `gorm:"type:varchar(255)" json:"default_payment_method_id"`
	NextInstallmentNumber   string    `gorm:"type:varchar(255)" json:"next_installment_number"`
	NextPaymentDate         string    `gorm:"type:varchar(255)" json:"next_payment_date"`
	AmountPaidCurrentYear   string    `gorm:"type:varchar(255)" json:"amount_paid_current_year"`
	BalanceCurrentYear      string    `gorm:"type:varchar(255)" json:"balance_current_year"`
	PaymentStatus           string    `gorm:"type:varchar(255)" json:"payment_status"`
	MembershipStatus        string    `gorm:"type:varchar(255)" json:"membership_status"`
	LastSuccessfulPaymentAt string    `gorm:"type:varchar(255)" json:"last_successful_payment_at"`
	LastFailedPaymentAt     string    `gorm:"type:varchar(255)" json:"last_failed_payment_at"`
	GraceExpiresAt          string    `gorm:"type:varchar(255)" json:"grace_expires_at"`
	CancelNowAt             string    `gorm:"type:varchar(255)" json:"cancel_now_at"`
	CancelAtRenewal         string    `gorm:"type:varchar(255)" json:"cancel_at_renewal"`
	CancelledAt             string    `gorm:"type:varchar(255)" json:"cancelled_at"`
	Notes                   string    `gorm:"type:text" json:"notes"`
	BillingMode             string    `gorm:"type:varchar(255)" json:"billing_mode"`
	ExternalBillingProvider string    `gorm:"type:varchar(255)" json:"external_billing_provider"`
	PromoCode               string    `gorm:"type:varchar(255)" json:"promo_code"`
	ExternalBillingStatus   string    `gorm:"type:varchar(255)" json:"external_billing_status"`
	BillingAccessTokenHash  string    `gorm:"type:text" json:"billing_access_token_hash"`
	BillingLinkExpiresAt    string    `gorm:"type:varchar(255)" json:"billing_link_expires_at"`
	BillingLinkStatus       string    `gorm:"type:varchar(255)" json:"billing_link_status"`
}

type SeatInventory struct {
	SeatID        string `gorm:"primaryKey;type:varchar(255)" json:"seat_id"`
	TeamCode      string `gorm:"type:varchar(255);index" json:"team_code"`
	Grade2027_28  string `gorm:"type:varchar(255)" json:"grade_2027_28"`
	SeatType      string `gorm:"type:varchar(255)" json:"seat_type"`
	SlotNo        string `gorm:"type:varchar(255)" json:"slot_no"`
	SeatStatus    string `gorm:"type:varchar(255)" json:"seat_status"`
	ApplicationID string `gorm:"type:varchar(255)" json:"application_id"`
	PersonID      string `gorm:"type:varchar(255)" json:"person_id"`
	HoldExpiresAt string `gorm:"type:varchar(255)" json:"hold_expires_at"`
	Notes         string `gorm:"type:text" json:"notes"`
}

type PaymentLedger struct {
	LedgerID                string    `gorm:"primaryKey;type:varchar(255)" json:"ledger_id"`
	MembershipID            string    `gorm:"type:varchar(255);index" json:"membership_id"`
	ApplicationID           string    `gorm:"type:varchar(255)" json:"application_id"`
	PersonID                string    `gorm:"type:varchar(255)" json:"person_id"`
	TeamCode                string    `gorm:"type:varchar(255)" json:"team_code"`
	PaymentContext          string    `gorm:"type:varchar(255)" json:"payment_context"`
	PaymentSource           string    `gorm:"type:varchar(255)" json:"payment_source"`
	StripeEventID           string    `gorm:"type:varchar(255)" json:"stripe_event_id"`
	StripeCustomerID        string    `gorm:"type:varchar(255)" json:"stripe_customer_id"`
	StripeCheckoutSessionID string    `gorm:"type:varchar(255)" json:"stripe_checkout_session_id"`
	StripePaymentIntentID   string    `gorm:"type:varchar(255)" json:"stripe_payment_intent_id"`
	StripeInvoiceID         string    `gorm:"type:varchar(255)" json:"stripe_invoice_id"`
	StripeChargeID          string    `gorm:"type:varchar(255)" json:"stripe_charge_id"`
	StripeSubscriptionID    string    `gorm:"type:varchar(255)" json:"stripe_subscription_id"`
	MembershipYear          string    `gorm:"type:varchar(255)" json:"membership_year"`
	InstallmentNumber       string    `gorm:"type:varchar(255)" json:"installment_number"`
	Amount                  string    `gorm:"type:varchar(255)" json:"amount"`
	Currency                string    `gorm:"type:varchar(255)" json:"currency"`
	Status                  string    `gorm:"type:varchar(255)" json:"status"`
	PaymentDate             string    `gorm:"type:varchar(255)" json:"payment_date"`
	ManualReason            string    `gorm:"type:varchar(255)" json:"manual_reason"`
	EnteredBy               string    `gorm:"type:varchar(255)" json:"entered_by"`
	FailureCode             string    `gorm:"type:varchar(255)" json:"failure_code"`
	Notes                   string    `gorm:"type:text" json:"notes"`
	CreatedAt               time.Time `json:"created_at"`
}

type StripeEvent struct {
	StripeEventID    string    `gorm:"primaryKey;type:varchar(255)" json:"stripe_event_id"`
	EventType        string    `gorm:"type:varchar(255)" json:"event_type"`
	CreatedAtStripe  string    `gorm:"type:varchar(255)" json:"created_at_stripe"`
	ReceivedAt       time.Time `json:"received_at"`
	ProcessedAt      string    `gorm:"type:varchar(255)" json:"processed_at"`
	ProcessingStatus string    `gorm:"type:varchar(255)" json:"processing_status"`
	ObjectID         string    `gorm:"type:varchar(255)" json:"object_id"`
	ApplicationID    string    `gorm:"type:varchar(255)" json:"application_id"`
	MembershipID     string    `gorm:"type:varchar(255)" json:"membership_id"`
	AttemptCount     string    `gorm:"type:varchar(255)" json:"attempt_count"`
	LastError        string    `gorm:"type:text" json:"last_error"`
	PayloadHash      string    `gorm:"type:varchar(255)" json:"payload_hash"`
}
