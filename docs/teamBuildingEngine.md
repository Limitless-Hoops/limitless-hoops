# Team Building & Referral Engine (2027 Plan)

To support the 2027 Team-Building plan, the system requires a robust Teams and Affiliate/Referral architecture. This moves the app from a simple "Booking Calendar" to a full "Franchise Management" tool.

*For the overall org hierarchy that governs teams (city-scoped vs. national), see `architecture.md` → The Data Hierarchy.*

---

## 1. Team Structure

### City Teams (Organization-Scoped)
Each city (Organization) has its own set of teams, divided by grade. Teams have `organization_id` set to their city's org.
*   **Regions (Denver launch):** Aurora, Broomfield, Centennial, Denver
*   **Grade Divisions (Per Region):** 1st–4th, 5th–6th, 7th–8th, 9th–12th
*   **Capacity:** 12 Active Players + 1 Backup per team.
*   **Volunteer constraint:** Maximum 4 volunteers per team.
*   **Player assignment is unique within a city:** A 5th grader in Denver can only be on the Denver 5th/6th grade team. There is no competition between coaches for the same player.

### National / Company-Wide Teams (Future)
For company-wide all-star teams, travel squads, or showcase rosters, a Team record can have `organization_id = null`. These teams draw players from multiple cities and are managed by a `super_admin`. See `architecture.md`.

---

## 2. Team Leaders & Coaches (Staff)
*   **Leader Assignment:** A user (Adult) can be assigned as the `Leader` of a specific Team via the `UserRole` table (`role: coach`, `organization_id: [city]`). See `accountSchema.md`.
*   **Lifecycle:**
    *   *Phase 1 (Jan – Aug):* Recruiting / Community building (Managing WhatsApp, tracking leads).
    *   *Phase 2 (Aug onward):* Option to convert to "Paid Coach" at $40/hour.
*   **Coach Responsibilities:**
    *   Coaches serve as the figurehead for their respective WhatsApp group.
    *   Administrative tasks (recruiting, emails, follow-ups, roster management) are handled by the central system/admin.
    *   Coaches play on Sunday nights to build relationships with prospects.
    *   M/W/F Skill Workouts are rotated between 4 coaches throughout the month.

---

## 3. Referral & Payout Engine (The Affiliate System)
The system must track exactly *who* drove a player to officially enroll to calculate payouts.
*Note: Referral rewards are ONLY paid out for Limitless tier enrollments.*

### Payout Rules:
*   **Coaches:** Receive **$200** for each Limitless-tier player who joins their team.
*   **Cross-Grade / Cross-Location / Cross-City Referrals:** If a coach or family refers someone outside their own team (different grade, different region, or even a different city), the referring person receives **$100**.

### Technical Implementation:
*   Every User account has a unique `Referral_Code` field (see `accountSchema.md`).
*   During checkout, the system logs the `Referred_By_User_ID` from the referral code entered.
*   Referral payouts are only triggered on **Limitless tier** enrollments.
*   The referral system is scoped per-organization — a coach's $200 payout applies within their city. Cross-city referrals earn the flat $100 regardless of org.
*   An admin dashboard (`/admin/referrals`) tallies these payouts so the business can issue checks/credits to coaches and families at the appropriate time.
