# Limitless Hoops - MVP Scope & Delivery Plan

## Overview
The application will be delivered in two phases. Phase 1 (the "Payment App") must be ready by **January 2027** to capture early-bird Limitless signups as coaches begin recruiting. Phase 2 (the full "Booking App") follows after, adding the event calendar and all associated tooling.

Both parties (developer and client) confirmed this scope on September 6, 2026.

---

## Phase 1: Payment App (Target: January 2027)

The goal of Phase 1 is to get families onto a Stripe-managed recurring payment plan from day one of recruiting, eliminating the need for the client to build and then migrate a temporary payment system on his end.

### In Scope

*   **Account Creation:** Passwordless login (OAuth + Email OTP). Household setup (parents + players).
*   **Tier Selection & Checkout:** Users select the Limitless tier (the only tier available before July) and choose Pay in Full or the 4-installment plan. Stripe handles the checkout.
*   **Installment Billing Schedule:** Stripe Subscription Schedules handle the 4-month billing phase + 8-month access phase + annual renewal. No custom billing math on the backend.
*   **Auto-Renewal Logic:** See full mechanic below.
*   **Renewal Reminder Emails:** Automated email notifications at the defined windows before each renewal.
*   **Prospect Waitlist:** Families can join a waitlist. Waitlist position is visible to the family. Leapfrog rule is enforced (paying members jump unpaid prospects).
*   **Grandfathered Pricing:** Admin can assign any custom Stripe price to an existing member's subscription, preserving their historical rate regardless of what new members pay.
*   **Admin: Manual Tier Assignment:** Admin can manually set or override a member's active tier (for migrating existing ~50 customers who will self-register and then be upgraded by an admin).
*   **Card on File (Existing Members):** Existing members are sent a link to save their card via Stripe Setup Intent (no charge required). Admin then creates their grandfathered subscription manually.

### Out of Scope (Phase 2)
*   Event calendar and public schedule
*   Event booking / cart system
*   A La Carte purchasing
*   Coach dashboard and mobile roster
*   Mass event builder
*   Automated roster management
*   iCal sync

---

## Phase 2: Booking App (Target: TBD)

Phase 2 layers the full event and booking engine on top of the payment infrastructure built in Phase 1. All the business logic for events, tiers, access gating, A La Carte pricing, and coach tooling documented in `eventsLogic.md`, `eventSchema.md`, and `businessLogic.md` applies to this phase.

---

## Auto-Renewal Mechanic (Phase 1 Core Feature)

Recurring payments are **ON by default** for all installment and annual plan members. The following notification and opt-out windows govern what happens leading up to each renewal:

| Timing | Trigger | Action |
|---|---|---|
| **T-3 weeks** | Renewal is already turned OFF | Warn the family: *"Your spot will be opened to the waitlist in 1 week unless you turn renewal back on."* |
| **T-2 weeks** | All members with upcoming charges | Send renewal reminder: *"Your next charge is in 2 weeks. Here's how to opt out."* |
| **Opt-out confirmed** | Family turns off renewal at any point | Immediately notify the waitlist that a spot will open on [renewal date]. First waitlisted family to complete payment secures the spot. |
| **Renewal date** | Renewal ON | Stripe charges automatically. Access continues uninterrupted. |
| **Renewal date** | Renewal OFF | Stripe cancels. Access is revoked. Roster spot is released. |

### Simultaneous Waitlist Grabs
When a spot opens and multiple waitlist families attempt to claim it at the same time, the backend enforces a **database-level lock on the roster spot**. Only the first completed Stripe payment wins. All others are immediately notified that the spot was taken and they remain on the waitlist.

*The client acknowledged this edge case and confirmed it is a "good problem to eventually have" — meaning over-engineering this on day one is not required, but the lock mechanism must be in place.*

---

## Grandfathered Pricing Notes
*   Stripe natively supports per-customer custom pricing. A grandfathered member paying $1,200/year simply has a Stripe subscription created at $1,200 — completely independent of any public-facing price.
*   Existing members (~50 total) will self-register in the new app. An Admin then manually assigns their tier and creates their Stripe subscription at their historical rate.
*   No automated data import from the client's Google Sheets is planned for Phase 1. See `businessLogic.md` → Data Migration & Launch Strategy.
