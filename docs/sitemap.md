# Limitless Hoops Sitemap

This document outlines the routing architecture across all subdomains. For the reasoning behind this structure, see `architecture.md` → Subdomain Map.

---

## `limitlesshoops.com` — Marketing Site
Public-facing. No login required. Currently static HTML. Eventually a Next.js or similar static site.

*   `/` — Home / Landing Page
*   `/about` — About Us
*   `/contact` — Contact / Support
*   `/pricing` — Public pricing overview (links into the city app for actual checkout)
*   `/teams` — Overview of the team structure and locations
*   `/terms` — Terms of Service
*   `/privacy` — Privacy Policy

---

## `[city].limitlesshoops.com` — City App
The primary authenticated platform. The same codebase deployed per city. The subdomain prefix (`denver`, `chicago`, etc.) resolves to an `Organization` record in the database and scopes all data automatically.

**Current active instance:** `denver.limitlesshoops.com`

### Public Routes (No login required)
*   `/` — City home / welcome page
*   `/events` — Public event calendar (all events visible; booking requires login)
*   `/events/:eventId` — Event detail (info, A La Carte pricing, public roster count)
*   `/pricing` — City-specific tier pricing and amenities
*   `/teams` — City team roster and coach directory (public)
*   `/waitlist` — Join the prospect waitlist for a specific team/grade
*   `/login` — Sign In (passwordless)
*   `/register` — Sign Up

### Protected Routes (Requires login)
*   `/dashboard` — User dashboard (upcoming events, membership status, household summary)
*   `/household` — Manage linked children and co-guardians
*   `/account` — Account settings, referral code/link, payment methods
*   `/membership` — View active tier, installment schedule, auto-renew toggle, cancel/upgrade, past invoices
*   `/cart` — Review items before purchase *(Phase 2)*
*   `/checkout` — Stripe payment integration
*   `/checkout/success` — Order confirmation
*   `/checkout/cancel` — Payment failed or canceled

### Coach Routes (Requires `coach` role in this org)
*   `/coach` — Coach dashboard — assigned schedule (mobile-first)
*   `/coach/roster/:eventInstanceId` — Event roster: attendees, emergency contacts, waitlist

---

## `admin.limitlesshoops.com` — Admin Dashboard
Separate app. Requires elevated permissions. Manages all organizations from one place.

### Super Admin Routes (Requires `super_admin` role)
*   `/` — Global dashboard: metrics across all cities, upcoming charges, quick actions
*   `/orgs` — Manage all city organizations (create new city, activate/deactivate)
*   `/orgs/:orgId` — Drill into a specific city's data
*   `/events/national` — Create and manage company-wide events, tournaments, leagues
*   `/reports` — Cross-city financial reports and enrollment summaries

### City Admin Routes (Requires `city_admin` or `super_admin` role)
*   `/[orgId]/members` — View and manage all members, subscriptions, and prospect pipeline for a city
*   `/[orgId]/members/:userId` — Member detail: tier, payment history, household, manual tier override
*   `/[orgId]/events` — Event management: create, edit, delete events and series
*   `/[orgId]/events/builder` — Mass event builder (generate months of events at once)
*   `/[orgId]/coaches` — Manage coaches and team assignments
*   `/[orgId]/teams` — Team management (rosters, grade divisions)
*   `/[orgId]/referrals` — Referral and payout dashboard
*   `/[orgId]/venues` — Add and manage locations
*   `/[orgId]/communications` — Send targeted SMS/Email to specific rosters or tiers
*   `/[orgId]/finances` — Recurring payment overview, issue refunds, generate reports

---

## `play.limitlesshoops.com` — National Results (Phase 3+)
Public-facing. No login required. Displays cross-city tournament brackets, standings, and results. Great marketing surface for expansion into new cities.

*   `/` — National overview: active tournaments, featured teams
*   `/tournaments` — All national tournaments
*   `/tournaments/:tournamentId` — Bracket, schedule, and results for a specific tournament
*   `/leagues` — National league standings
*   `/teams` — National team directory (all-star, travel squads)

---

## `shop.limitlesshoops.com` — E-Commerce (Phase 3+)
Separate product. No Limitless membership required to purchase. May be implemented on a third-party platform (e.g., Shopify) rather than custom-built.

*   `/` — Store home
*   `/products` — All merchandise
*   `/products/:productId` — Product detail and purchase
