# Limitless Hoops Tiers & Billing Logic

This document serves as the source of truth for Membership Tiers, Pricing, and the core Business/Billing Logic for the application.

## Core Billing Paradigm: "Club Dues" (Strict Contracts)

The billing architecture is built on a **Financed Contract / Club Dues** model, rather than a standard month-to-month SaaS model.
* Users are purchasing a fixed-term contract (2 months, 4 months, or 1 year).
* Installment options exist solely as a convenience to finance the total cost of the tier. They do *not* represent month-to-month subscriptions.
* Because payments are front-loaded (e.g., paying for 12 months in 4 months), the application strictly enforces the rules below to prevent users from gaming the system and to protect the club's upfront overhead costs.

### General Billing Rules
1. **Grace Periods & Repossession:** If an installment payment fails, Stripe will automatically retry the card for **7 days** (Dunning). If the payment still fails after the 7-day grace period, the subscription is canceled by Stripe. The application will listen for this cancellation webhook and **immediately revoke gym access and forfeit the user's roster spot.** No further balance is owed or collected after repossession (Clean Break).
2. **Cancellations (Clean Break):** If a user voluntarily cancels their installment plan in the app, their access is revoked immediately and **all future automated charges are canceled**. They owe nothing further for the remaining contract balance. The app will warn them: *"Canceling your payment plan will forfeit your roster spot and gym access immediately. Past payments are non-refundable per club policy, but you will not be billed for any remaining installments."*
3. **No Pro-Rated Refunds:** Payments already made are non-refundable. They cover the club's upfront overhead. Users do not get "fractional time" if they stop paying early.
4. **Upgrades:** Users can upgrade mid-contract. The system will calculate the remaining balance: `(Cost of New Tier) - (Amount Already Paid)`. The user is then billed for this difference.
5. **Waitlist Priority:** If not on a recurring charge, a waitlist member may immediately take the user's roster spot when the term ends.

---

## Membership Tiers (Pricing & Amenities)

### 1. Recreation
* **Term:** 2 Months
* **Option 1 (Pay in Full):** $500
* **Option 2 (Installments):** $100 x 5 Weekly Installments
* **Amenities Included:**
  * Games
  * M/W/F Skill Workouts
  * Tue/Thu Practices
  * Sunday Shooting
  * Saturday Open Gym

### 2. Seasonal
* **Term:** 4 Months
* **Option 1 (Pay in Full):** $1,000
* **Option 2 (Installments):** $500 x 2 Monthly Installments
* **Amenities Included:**
  * *All Recreation amenities plus:*
  * Skill Clinics
  * PnR Workshops

### 3. Limitless
* **Term:** 1 Year
* **Option 1 (Pay in Full):** $3,000
* **Option 2 (Installments):** $750 x 4 Monthly Installments
* **Amenities Included:**
  * *All Seasonal amenities plus:*
  * Winter Camp
  * Summer Camp
  * Combines

### 4. Volunteer
* **Term:** 2 Months (Resets every 2 months)
* **Price:** $200 – $500 (Exact amount varies)
* **Billing Type:** One-time refundable deposit
* **Rules & Logistics:**
  * Through 12th grade.
  * 4 spots per grade.
  * First come, first served.
  * Requires 90 minutes of volunteering per week.
  * Deposit is refunded at the end of the 2 months, or rolled over if they decide to continue volunteering for another term.
* **Amenities Included:**
  * *Full Limitless Access* (Access to all events, clinics, and camps during the 2-month term).

---

## Technical Implementation Details
* **Source of Truth:** Stripe handles all billing schedules and recurring logic. 
* **Automated Revocation:** The backend relies entirely on Stripe Webhooks (`customer.subscription.deleted` or `invoice.payment_failed` after the dunning period) to revoke access. There is no custom "proration" or "fractional time" math handled by the backend.
* **Upgrades:** Upgrades require generating a new Stripe invoice/checkout for the calculated difference in tier costs.
