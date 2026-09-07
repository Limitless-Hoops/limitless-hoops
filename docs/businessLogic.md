# Limitless Hoops Booking System 2.0 - Business Logic & Frontend Views

*This document covers feature-level business rules. For the overall system architecture, subdomain map, multi-tenancy model, and folder structure, see `architecture.md`.*

## 1. User Accounts & Roles
Accounts are absolutely necessary to solve the current pain points. We need distinct experiences for different types of users:

*   **Customers (Parents/Players):** 
    *   Need multiple login options (Email, Google, Apple, Facebook).
    *   **The "Household" Concept:** One account needs to manage multiple players (kids), and multiple adults (split households) need to be able to manage the same player.
    *   Must be able to manage their own memberships (cancel, upgrade) without going to a third-party site. *(Note: Memberships cannot be paused. These are strict contracts. See `membershipTiers.md`)*
*   **Coaches:** 
    *   Need to view class rosters easily from their phones.
    *   Need to receive specific alerts for only the classes they are assigned to.
*   **Admins:** 
    *   Need full control over mass event creation, venue management (uncapped, >20), global reporting, and refunds.

## 2. Memberships, Pricing & Payments
The payment system needs to be brought in-house to reduce friction.

*   **Tiers & Strict Contracts:** Support for 2-month, 4-month, and 1-year commitments (Volunteer, Recreation, Seasonal, Limitless). See `membershipTiers.md` for the strict "Club Dues" billing logic, grace periods, and cancellation rules.
*   **Phased Delivery:** The app is being delivered in two phases. Phase 1 (Payment App, target January 2027) covers accounts, billing, and the waitlist. Phase 2 adds the full event/booking engine. See `mvpScope.md` for the full breakdown.
*   **Public Pricing:** All tier and event pricing must be visible to the public without requiring an account.
*   **Dynamic Pricing:** The system must recognize the user's tier and apply the correct price (sometimes free) to an event.
*   **Integrated Checkout:** Payments happen directly on the site via Stripe.
*   **Cart System (High Priority — Phase 2):** Customers must be able to select multiple dates/events, add them to a cart, and check out all at once.
*   **Billing Transparency:** Send an automated 1-week heads-up email before any recurring charges happen. Centralized view for admins to see all recurring payments.
*   **Auto-Renewal (ON by Default):** All recurring subscriptions default to auto-renewing. Families receive notification windows before each renewal and can opt out. See `mvpScope.md` for the full 3-week / 2-week notification window mechanic and waitlist notification rules.
*   **Grandfathered Pricing:** Admin can assign any custom Stripe price to an existing member, preserving their historical rate. Existing members save their card via Stripe Setup Intent (no charge). See `mvpScope.md`.
*   **Prospect Follow-ups:** Automated monthly check-ins with prospects (waitlist/unpaid) to confirm intent. Once payment clears, their roster spot is marked as secured.
*   **Upgrades & Credits:** Ability to calculate mid-contract upgrades (New Tier - Amount Paid = Balance Owed) and apply promo codes.

## 3. The Booking & Event Engine
Events are complex and need flexible rules.

*   **Public Schedule:** The full event schedule must be public-facing.
*   **Event Types:** Recurring events, one-time events, and "TBD" events (where time/location are genuinely to be determined, not faked).
*   **Visibility & Caps:** Control when an event becomes visible, when booking closes, and the maximum capacity.
*   **Waitlist Priority (Confirmed):** Recurring payment members retain priority for their roster spot. When a spot opens, recurring members on the waitlist are contacted before one-time/pay-in-full members, regardless of when they joined the waitlist. See `eventsLogic.md` for full details.
*   **Duplicate Prevention:** A single player (child) cannot be booked into the exact same event instance twice, even if the parent adds it to their cart multiple times.
*   **Automated Notifications:** Any time an admin edits a roster or edits an event's details (time, location), the system must auto-email everyone currently booked for that event.
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
    *   **Membership & Billing:** View tier, cancel/upgrade membership, view past invoices, and manage payment methods.

---

## 4. Data Migration & Launch Strategy
**Decision (confirmed):** There will be no automated import from the client's existing Google Sheets roster into the new app.

*   **Rationale:** The existing customer base is small (~50 customers). The operational overhead of building a one-time import tool is not justified at this stage.
*   **Process:** Existing customers will create their own accounts in the new app. Once they are in, an Admin will manually assign their correct tier via the Admin Dashboard based on what they have already paid.
*   **Revisit:** If the roster grows significantly before launch, this decision should be revisited.

### Coach Views
1.  **Coach Dashboard (Mobile-First):** Shows their assigned schedule.
2.  **Mobile Roster:** View attendees, emergency contacts, and waitlist for a specific class.

### Admin Views
1.  **Global Dashboard:** High-level metrics, upcoming payments, and quick actions.
2.  **Mass Event Builder:** The interface for generating months of events at once, managing blackouts/exceptions, and assigning coaches.
3.  **Venue Management:** Add and manage unlimited locations.
4.  **Financial Hub:** View all recurring payments, issue refunds directly, and generate monthly printable reports.
5.  **Communications Center:** Send targeted SMS/Email to specific event rosters.