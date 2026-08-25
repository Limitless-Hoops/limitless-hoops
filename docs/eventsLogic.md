# Limitless Hoops - Events & Scheduling Logic

This document defines the types of events hosted by Limitless Hoops, their scheduling mechanics, and the minimum membership tier required to access them. 

*(For a full breakdown of what each tier includes and its pricing, see `membershipTiers.md`)*

## Event Categories & Definitions

### 1. Workouts, Practices, & Open Play (High Frequency)
These form the core weekly schedule and are accessible to all active members.
*   **Games:** Local league games or scrimmages.
    *   **Minimum Tier:** Recreation
*   **Skill Workouts (M/W/F):** General skill development drills.
    *   **Minimum Tier:** Recreation
*   **Team Practices (Tue/Thu):** Team strategy and organized practice.
    *   **Minimum Tier:** Recreation
*   **Shooting Workouts (Sun):** Specialized shooting drills.
    *   **Minimum Tier:** Recreation
*   **Open Gym (Sat):** Unstructured playtime.
    *   **Minimum Tier:** Recreation

### 2. Specialized Training (Clinics & Workshops)
These are targeted, advanced training sessions requiring a slightly higher commitment.
*   **Skill Clinics:** Specialized deep-dive clinics.
    *   **Minimum Tier:** Seasonal
*   **PnR (Pick and Roll) Workshops:** Tactical workshops focusing on the PnR.
    *   **Minimum Tier:** Seasonal

### 3. Camps & Assessments (Premium Events)
These are premium, high-value events restricted to full-year or volunteer members.
*   **Combines:** Ability and agility assessments.
    *   **Minimum Tier:** Limitless
*   **Winter Camp:** Multi-day winter seasonal intensive camp.
    *   **Minimum Tier:** Limitless
*   **Summer Camp:** Multi-day summer seasonal intensive camp.
    *   **Minimum Tier:** Limitless

---

## Technical Implications for Event Engine
*   **Tier Gating (Hard Requirement):** When an event is created in the admin panel, it must be assigned a `minimum_tier_required`. The backend must check the user's current active subscription tier against this value before allowing a booking. 
*   **UI Discoverability:** Events that a user does not have access to (e.g., a Recreation user viewing a Combine) should still be visible on the public and member calendars, but the booking button must be replaced with an "Upgrade Required" CTA to incentivize tier upgrades.
