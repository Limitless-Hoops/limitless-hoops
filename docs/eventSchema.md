# Event & Scheduling Engine - Data Schema (Draft)

To build a scheduling engine, we need to separate the "Idea" of an event from the "Actual occurrences" on the calendar. This prevents massive data duplication and makes editing much easier.

## 1. Event Template (The "Idea")
This defines what the event is, but not *when* it happens.
*   **ID:** (UUID)
*   **Name:** e.g., "Elite Pick n Roll Workshop"
*   **Description:** "Mastering 12 efficient counters..."
*   **Category:** (Enum: Workout, Clinic, Camp, Combine, Game)
*   **Venue_ID:** Links to a locations table.
*   **Coach_ID:** Links to a staff/admin table.
*   **Default_Capacity:** e.g., 20 players.
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

## 3. Access & Pricing Rules
This determines who can see/book the event and what it costs. We tie this to the Template, but allow Instances to override.
*   **Event_ID:** Links to Template or Instance.
*   **Tier_ID:** Links to Membership Tiers (Recreation, Seasonal, Limitless, Non-Member).
*   **Price:** Decimal (e.g., $0.00, $15.00).
*   **Booking_Window_Start:** How many days before the event can this tier book? (Allows for early-access for premium tiers).

## 4. Bookings (The Roster)
This tracks who is actually going.
*   **ID:** (UUID)
*   **Event_Instance_ID:** The specific date.
*   **Player_ID:** The kid attending.
*   **Booked_By_Account_ID:** The parent who clicked "checkout".
*   **Status:** (Enum: Booked, Cancelled, Waitlisted, Attended, No-Show)
*   **Amount_Paid:** What did they actually pay at checkout?
*   **Transaction_ID:** Link to Stripe.
