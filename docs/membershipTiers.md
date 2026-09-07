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
  * Team Strategy

### 2. Seasonal
* **Term:** 4 Months
* **Option 1 (Pay in Full):** $1,000
* **Option 2 (Installments):** $500 x 2 Monthly Installments
* **Amenities Included:**
  * *All Recreation amenities plus:*
  * Skill Clinics
  * PnR Workshops

### 3. Limitless
* **Term:** Rolling 365 days from the date of signup. This is NOT a calendar-year or school-year structure. A member who signs up in February is covered through the following February. There is no fixed start or end season — the gym never stops, and neither does the membership.
* **Early Bird (Before July):** $1,500/year (or 4 monthly installments of $375). This rate is locked in for life — same billing behavior as the standard rate, just at 50% off permanently.
* **Standard (After July):** Gradually increases to $3,000/year (or 4 monthly installments of $750).
* **Referral Rewards:** Limitless is the *only* tier that pays referral rewards.
* **Amenities Included:**
  * *All Seasonal amenities plus:*
  * Winter Camp
  * Summer Camp
  * Combines
  * Shot Guru
  * All Extras (except Private Sessions) at no extra cost.

#### Limitless Installment Billing Cycle (Annual)
For members on the installment plan, the billing cycle works as follows across a full year:

| Phase | Duration | What Happens |
|---|---|---|
| **Billing Phase** | Months 1–4 | 4 monthly payments charged. Access is active. |
| **Access Phase** | Months 5–12 | No charges. Access continues (already paid for). |
| **Renewal** | Month 12 Anniversary | Member is prompted to renew. They choose Pay in Full or 4 installments again for the next 365-day window. |

*Note: The "Access Phase" is not free — the member already paid for it in the first 4 months. It is simply the non-billing period of the annual contract.*

### 4. Volunteer
* **Term:** 1 Year (Requires 2-year commitment for refund)
* **Price:** $200 – $500 (Exact amount varies)
* **Billing Type:** One-time refundable deposit
* **Rules & Logistics:**
  * Through 12th grade.
  * 4 spots per grade.
  * First come, first served.
  * Requires 90 minutes of volunteering per week.
  * Volunteers must stay at least 2 years to receive their deposit back.
* **Amenities Included:**
  * *Full Limitless Access* (Access to all events, clinics, and camps).

---

## Enrollment Timeline (2027)

### Why the Timeline Is Structured This Way
The goal is 192 players committed by August so September/October functions as a real, full league. Families will not naturally commit to September basketball in January — they need a reason to act early. The Early Bird pricing ($1,500) is that reason. It also positions the annual plan at a psychologically palatable price point (~$375/month), comparable to what other programs charge for a single month, making year-round membership viable even for families who participate in multiple sports.

The progressive opening of shorter plans is intentional and must be enforced by the app. If all three plans were available from January, the vast majority of families would choose the shortest commitment possible, eliminating the early cash flow and league-building momentum the business needs. **The app must not allow a user to purchase a Recreation or Seasonal tier before its designated opening date.**

### Timeline
*   **Before July:** Only the Limitless tier is available (Early Bird rate). Joining the WhatsApp prospect list is free and secures a place in line, but does NOT guarantee a roster spot.
*   **July/August:** Limitless pricing gradually increases toward the full $3,000 annual rate.
*   **August:** 4-Month (Seasonal) option opens at full pricing.
*   **Late August:** 2-Month (Recreation) option opens at full pricing, timed so that 2 months of access carries families through the September/October league centerpiece.

### WhatsApp Prospect List & Leapfrog Rule
The prospect waitlist operates on a modified priority system during the recruiting phase:
*   Joining the WhatsApp group is free and places a family in line.
*   A prospect's position in line is **NOT permanent**. If a grade's roster is filling up and someone behind them in the queue commits to an annual plan, that paying member leapfrogs unpaid prospects.
*   Once a family pays for any tier, their roster spot is secured and cannot be leapfrogged.
*   **Technical implication:** The waitlist must track both `join_date` and `payment_status` for each prospect. Sorting must place paying/recurring members before unpaid prospects, then use `join_date` as a tiebreaker within each group.

---

## Technical Implementation Details
* **Source of Truth:** Stripe handles all billing schedules and recurring logic.
* **Rolling Start Date:** Each member's term starts on their individual signup date, not a fixed calendar date. The backend must store the member's `term_start_date` and `term_end_date` (start + 365 days) and use these to manage access, not a shared season calendar.
* **Installment Billing Schedule:** The 4-month billing phase followed by an 8-month non-billing access period cannot be handled by a standard Stripe monthly subscription. This requires **Stripe Subscription Schedules** — a schedule is created with 4 billing phases, then a pause, with renewal queued at the 12-month anniversary.
* **Automated Revocation:** The backend relies entirely on Stripe Webhooks (`customer.subscription.deleted` or `invoice.payment_failed` after the dunning period) to revoke access. There is no custom "proration" or "fractional time" math handled by the backend.
* **Upgrades:** Upgrades require generating a new Stripe invoice/checkout for the calculated difference in tier costs.
