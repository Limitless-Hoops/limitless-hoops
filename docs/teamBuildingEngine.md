# Team Building & Referral Engine (2027 Plan)

To support the 2027 Team-Building plan, the system requires a robust Teams and Affiliate/Referral architecture. This moves the app from a simple "Booking Calendar" to a full "Franchise Management" tool.

## 1. Regional Team Structure
The database needs a rigid hierarchy for teams to track the 16 target squads.
*   **Regions:** Aurora, Broomfield, Centennial, Denver
*   **Grade Divisions (Per Region):** 1st–4th, 5th–6th, 7th–8th, 9th–12th
*   **Capacity constraint:** 12 Active Players + 1 Backup per team.
*   **Volunteer constraint:** Maximum 4 volunteers per team.

## 2. Team Leaders & Coaches (Staff)
*   **Leader Assignment:** A user (Adult) can be assigned as the `Leader` of a specific Team.
*   **Lifecycle:** 
    *   *Phase 1 (Jan - Aug):* Recruiting / Community building (Managing WhatsApp, tracking leads).
    *   *Phase 2 (Aug onward):* Option to convert to "Paid Coach" at $40/hour.
*   **Coach Responsibilities (If converted):**
    *   Tue/Thu Team Practices
    *   1x M/W/F Skill Workout
    *   Weekend Games

## 3. Referral & Payout Engine (The Affiliate System)
The system must track exactly *who* drove a player to officially enroll to calculate payouts.

### Payout Rules:
*   **Base Payout per Regular Player:** $300 total pool.
*   **If Leader recruits to their OWN team:** Leader gets the full $300 (Technically $200 for the team spot + $100 referral).
*   **If Leader recruits to ANOTHER team:** Recruiting Leader gets $100. The receiving Team's Leader gets $200.
*   **If a Parent/Player refers someone:** The referring family gets $100. The receiving Team's Leader gets $200.
*   *Exclusion:* Volunteer players do NOT trigger the $200 Team Leader payout.

### Technical Implementation:
*   Every User account needs a unique `Referral_Code` or `Referral_Link`.
*   During checkout (especially during the July Pricing Window), the system logs the `Referred_By_User_ID`.
*   An admin dashboard is required to tally these payouts so the business can issue the checks/credits to the Leaders and Families in August once enrollments are locked.
