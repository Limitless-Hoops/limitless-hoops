# Event & Scheduling Engine - Data Schema (Draft)

To build a scheduling engine, we need to separate the "Idea" of an event from the "Actual occurrences" on the calendar. This prevents massive data duplication and makes editing much easier.

*For how events are scoped to a city vs. the whole company (national events), see `architecture.md` → The Data Hierarchy.*

---

## 1. Event Template (The "Idea")
This defines what the event is, but not *when* it happens.
*   **ID:** (UUID)
*   **Organization_ID:** (Nullable) — Links to the Organization (city) this event belongs to. `null` means this is a company-wide / national event visible to all cities. See `architecture.md`.
*   **Name:** e.g., "Elite Pick n Roll Workshop"
*   **Description:** "Mastering 12 efficient counters..."
*   **Category:** (Enum: Workout, Clinic, Camp, Combine, Game, Tournament, Private_Session)
*   **Venue_ID:** Links to a locations table.
*   **Coach_ID:** Links to a staff/admin table.
*   **Default_Capacity:** e.g., 20 players.
*   **Minimum_Tier_Required:** (Enum: None, Recreation, Seasonal, Limitless) — The lowest tier that gets free access to this event. "None" means it is purchasable by anyone including non-members.
*   **A_La_Carte_Price:** Decimal — The price charged to users who do NOT meet the `Minimum_Tier_Required`. Set to `$0.00` if the event is always free. (See `eventsLogic.md` for per-event pricing.)
*   **Is_Active:** Boolean (Can admins still use this template?)

## 2. Event Instance (The "Actual Date")
This represents a specific date and time on the calendar. It inherits rules from the Template but can be overridden.
*   **ID:** (UUID)
*   **Template_ID:** Links back to the Template.
*   **Start_Time:** (Timestamp with Timezone)
*   **End_Time:** (Timestamp with Timezone)
*   **Status:** (Enum: Scheduled, Cancelled, Completed)
*   **Override_Capacity:** (Optional) If this specific day needs a different cap.
*   **Override_Coach_ID:** (Optional) If there is a substitute coach for this day.
*   **Series_ID:** (Optional) If this is part of a recurring batch, they share an ID so we can say "Delete all future instances in this series".

## 3. Bookings (The Roster)
This tracks who is actually going.
*   **ID:** (UUID)
*   **Event_Instance_ID:** The specific date.
*   **Player_ID:** The kid attending.
*   **Booked_By_User_ID:** The parent who clicked "checkout".
*   **Status:** (Enum: Booked, Cancelled, Waitlisted, Attended, No-Show)
*   **Amount_Paid:** What did they actually pay at checkout?
*   **Transaction_ID:** Link to Stripe.

*(Note: The combination of `Event_Instance_ID` + `Player_ID` must be unique at the database level. This is the primary enforcement mechanism for the duplicate booking prevention rule. See `businessLogic.md`)*

---

## National Events (Cross-City)
When a city app queries events to display to a family, it fetches both city-scoped and national events in a single query:

```sql
SELECT * FROM event_templates
WHERE organization_id = '[current_city_org_id]'
   OR organization_id IS NULL
```

National events (tournaments, combines, all-star games) are created by a `super_admin` in `admin.limitlesshoops.com` and automatically surface in every city's schedule. No manual syncing required.
