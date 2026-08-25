# Limitless Hoops Booking System 2.0 - Business Logic & Frontend Views

## 1. User Accounts & Roles
Accounts are absolutely necessary to solve the current pain points. We need distinct experiences for different types of users:

*   **Customers (Parents/Players):** 
    *   Need multiple login options (Email, Google, Apple, Facebook).
    *   **The "Household" Concept:** One account needs to manage multiple players (kids), and multiple adults (split households) need to be able to manage the same player.
    *   Must be able to manage their own memberships (pause, resume, cancel) without going to a third-party site.
*   **Coaches:** 
    *   Need to view class rosters easily from their phones.
    *   Need to receive specific alerts for only the classes they are assigned to.
*   **Admins:** 
    *   Need full control over mass event creation, venue management (uncapped, >20), global reporting, and refunds.

## 2. Memberships, Pricing & Payments
The payment system needs to be brought in-house to reduce friction.

*   **Tiers & Strict Contracts:** Support for 2-month, 4-month, and 1-year commitments (Volunteer, Recreation, Seasonal, Limitless). See `membershipTiers.md` for the strict "Club Dues" billing logic, grace periods, and cancellation rules.
*   **Dynamic Pricing:** The system must recognize the user's tier and apply the correct price (sometimes free) to an event.
*   **Integrated Checkout:** Payments happen directly on the site via Stripe.
*   **Cart System (High Priority):** Customers must be able to select multiple dates/events, add them to a cart, and check out all at once.
*   **Billing Transparency:** Automated heads-up emails before recurring charges happen. Centralized view for admins to see all recurring payments.
*   **Upgrades & Credits:** Ability to calculate mid-contract upgrades (New Tier - Amount Paid = Balance Owed) and apply promo codes.

## 3. The Booking & Event Engine
Events are complex and need flexible rules.

*   **Event Types:** Recurring events, one-time events, and "TBD" events (where time/location are genuinely to be determined, not faked).
*   **Visibility & Caps:** Control when an event becomes visible, when booking closes, and the maximum capacity (triggering a waitlist).
*   **Extras:** Ability to upsell add-ons (carpools, hotels, gear) at checkout for specific events.
*   **Flexibility:** Admins must be able to edit a schedule/series even if a single event in that series is already booked. Customers must be able to cancel a single instance of a recurring booking without canceling the whole series.

---

## Proposed Frontend Views (Initial Draft)

To build this frontend-first, here are the primary screens/views we should define:

### Customer Views
1.  **Public Event Discovery:** Filterable calendar/list view of upcoming events (Main vs. Special events). Embeddable on the marketing site.
2.  **Event Detail / Booking:** Shows event info, pricing (based on user state), and limited public roster (who's going).
3.  **Shopping Cart & Checkout:** The new multi-date checkout experience with integrated payment, promo codes, and up-sells (extras).
4.  **Customer Dashboard:**
    *   **Household Hub:** Manage linked adults and children.
    *   **Schedule:** Upcoming bookings with iCal sync and options to cancel specific dates.
    *   **Membership & Billing:** View tier, pause/cancel membership, view past invoices, and manage payment methods.

### Coach Views
1.  **Coach Dashboard (Mobile-First):** Shows their assigned schedule.
2.  **Mobile Roster:** View attendees, emergency contacts, and waitlist for a specific class.

### Admin Views
1.  **Global Dashboard:** High-level metrics, upcoming payments, and quick actions.
2.  **Mass Event Builder:** The interface for generating months of events at once, managing blackouts/exceptions, and assigning coaches.
3.  **Venue Management:** Add and manage unlimited locations.
4.  **Financial Hub:** View all recurring payments, issue refunds directly, and generate monthly printable reports.
5.  **Communications Center:** Send targeted SMS/Email to specific event rosters.