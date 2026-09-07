# Limitless Hoops - Events & Scheduling Logic

This document defines the types of events hosted by Limitless Hoops, their scheduling mechanics, and the minimum membership tier required to access them for free. Events without a free tier can be purchased A La Carte by anyone.

*(For a full breakdown of what each tier includes and its pricing, see `membershipTiers.md`)*

## Event Categories & Definitions

### 1. Core Schedule (High Frequency)
These form the core weekly schedule. Accessible to all active members as part of their tier.
*   **Games:** Local league games or scrimmages.
    *   **Minimum Tier (Free):** Recreation
*   **Skill Workouts (M/W/F):** General skill development drills.
    *   **Minimum Tier (Free):** Recreation
*   **Team Practices (Tue/Thu):** Team strategy and organized practice.
    *   **Minimum Tier (Free):** Recreation
*   **Shooting Workouts (Sun):** Specialized shooting drills.
    *   **Minimum Tier (Free):** Recreation
*   **Team Strategy (Mon):** Tactical strategy session.
    *   **Minimum Tier (Free):** Recreation. **A La Carte:** $5

### 2. Specialized Training (Clinics & Workshops)
Included in Seasonal and above. Purchasable A La Carte for Recreation members or the public.
*   **Skill Clinics (Thu):** Specialized deep-dive clinics — Rebounding, Offensive Skill, Shooting, Inside Scoring, Defensive Skill, Ball Handling.
    *   **Minimum Tier (Free):** Seasonal. **A La Carte:** $20
*   **PnR (Pick-and-Roll) Workshops (Tue):** Tactical workshops focusing on the PnR.
    *   **Minimum Tier (Free):** Seasonal. **A La Carte:** $10

### 3. Premium Events (Limitless-Only or A La Carte)
Included free in Limitless. Purchasable A La Carte for lower tiers or the public.
*   **Shot Guru (Fri):** Specialized Friday shooting instruction.
    *   **Minimum Tier (Free):** Limitless. **A La Carte:** $25
*   **Combines (Wed):** Ability and agility assessments.
    *   **Minimum Tier (Free):** Limitless. **A La Carte:** $15
*   **Winter Camp:** Multi-day winter seasonal intensive camp.
    *   **Minimum Tier (Free):** Limitless. **A La Carte:** TBD
*   **Summer Camp:** Multi-day summer seasonal intensive camp.
    *   **Minimum Tier (Free):** Limitless. **A La Carte:** TBD

### 4. Private Sessions (Never included in any tier)
*   **Private Sessions (Sat/Sun):** 1-on-1 training. Not included in any membership tier.
    *   **A La Carte:** $50

---

## Waitlist & Roster Priority Rule
*(Source: updated flyer + client confirmed leapfrog mechanic. Full strategic context in `membershipTiers.md` → Enrollment Timeline)*

When a roster spot opens up, or when a grade's capacity is approaching full, priority for that spot is determined as follows:

| Priority | Who | Condition |
|---|---|---|
| 1st | Paying members on recurring installments | Already committed, cannot be displaced |
| 2nd | Paying members (Pay in Full) | Spot secured at time of payment |
| 3rd | Unpaid prospects (WhatsApp list) | First-come, first-served — but can be leapfrogged |

*   **The Leapfrog Rule:** An unpaid prospect's place in the queue is NOT permanent. A new family who commits to a paid annual plan will jump ahead of all unpaid prospects, regardless of when the prospect joined the list.
*   Once a family makes any payment, their spot is secured and cannot be leapfrogged by anyone.
*   **Technical implication:** The waitlist must sort by `payment_status` first (paid before unpaid), then by `is_recurring` (recurring before one-time), then by `join_date` as the final tiebreaker.

---

## Technical Implications for Event Engine
*   **Tier Gating & A La Carte (Hard Requirement):** When an event is created in the admin panel, it must be assigned a `minimum_tier_required` and an `a_la_carte_price`. See `eventSchema.md` for the data model.
*   **Dynamic Checkout:** The backend checks the user's active subscription tier. If they meet `minimum_tier_required`, the event is $0. Otherwise, they are charged `a_la_carte_price`.
*   **UI Discoverability:** All events must be visible to the public. Events a user doesn't have free access to should display the A La Carte price with an "Upgrade for Free Access" CTA alongside the paid booking button.
*   **Duplicate Prevention:** The database enforces a unique constraint on `Event_Instance_ID` + `Player_ID`. See `eventSchema.md`.
