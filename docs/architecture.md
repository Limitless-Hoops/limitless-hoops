# Limitless Hoops - System Architecture

This document is the source of truth for the overall system design, subdomain map, folder structure, multi-tenancy model, and auth hierarchy. All other docs should be read in the context of this architecture.

---

## Core Design Principles

1. **Build for the city. Design for the nation.** Every data model decision must support multi-city expansion from day one, even if Denver is the only active city for now.
2. **Subdomains by tenant, not by feature.** Separate subdomains exist for separate *audiences* (cities, admins, public) — not for separating features like payments vs. bookings.
3. **Auth follows the org.** A user's role is always scoped to an organization. The same person can be a Coach in Denver and a Customer in Chicago.
4. **Nullable org = company-wide.** Any piece of data with `organization_id = null` belongs to the company as a whole (national events, national teams, national leagues).
5. **The monolith problem is solved by architecture, not subdomains.** The Go backend is structured as independent service packages. Each can be extracted into a true microservice when scale demands it — without changing the API surface.

---

## The Data Hierarchy

```
Company (Limitless Hoops)
    ├── [company-wide] Teams, Leagues, Tournaments   (organization_id = null)
    └── Organization / City
            ├── Teams
            │     └── Players
            ├── Events / Schedules
            ├── Coaches
            └── Subscriptions / Waitlists
```

*   **Company:** The single top-level entity. There is one Company record. Exists to scope national/cross-city data.
*   **Organization:** A city market (Denver, Chicago, Atlanta, etc.). All city-level data is scoped to an Organization via `organization_id`.
*   **National Scope:** Events, teams, and leagues with `organization_id = null` belong to the Company and are visible to all relevant cities.

---

## Subdomain Map

| Subdomain | Audience | Purpose | Phase |
|---|---|---|---|
| `limitlesshoops.com` | Public | Marketing site, landing pages | Exists now (static HTML) |
| `[city].limitlesshoops.com` | Families, coaches | City-specific app — accounts, payments, schedule, bookings, teams, recruiting, curriculum | Phase 1+ |
| `admin.limitlesshoops.com` | Super admins, city admins | Manage all orgs, create national events, manage grandfathered pricing, financial reports | Phase 1+ |
| `play.limitlesshoops.com` | Public | National tournament brackets, cross-city standings, public results — great marketing surface | Phase 3+ |
| `shop.limitlesshoops.com` | Public | E-commerce / merch — separate product, no membership required, may use third-party platform | Phase 3+ |

### Current Active Subdomain (Phase 1)
`denver.limitlesshoops.com` — Denver is the first and only Organization. The `[city]` pattern is in place from day one so adding Chicago later requires zero code changes: create an Organization record, spin up the subdomain, done.

---

## Folder Structure

```
limitless-hoops/
├── apps/
│   ├── app/                    → [city].limitlesshoops.com
│   │   └── modules/            → Feature modules (independently developed)
│   │       ├── auth/
│   │       ├── payments/
│   │       ├── bookings/       → Phase 2
│   │       ├── teams/          → Phase 2
│   │       ├── leagues/        → Phase 3
│   │       ├── recruiting/     → Phase 3
│   │       └── curriculum/     → Phase 3
│   ├── admin/                  → admin.limitlesshoops.com
│   ├── play/                   → play.limitlesshoops.com (Phase 3+)
│   ├── shop/                   → shop.limitlesshoops.com (Phase 3+)
│   └── api/                    → Go backend (serves all subdomains)
│       └── services/
│           ├── orgs/           → Organization & Company management
│           ├── auth/
│           ├── payments/
│           ├── bookings/       → Phase 2
│           ├── teams/          → Phase 2
│           ├── leagues/        → Phase 3
│           ├── recruiting/     → Phase 3
│           └── curriculum/     → Phase 3
└── packages/
    ├── ui/                     → Shared component library (used by app + admin)
    ├── utils/                  → Shared utility functions
    └── types/                  → Shared TypeScript types & interfaces
```

> **Note on current naming:** `apps/booking-react` should be renamed to `apps/app` and `apps/go-backend` should be renamed to `apps/api`. The folder structure above reflects the target state.

---

## Auth & Role Hierarchy

Roles are always scoped to an Organization. A user can hold different roles in different organizations.

| Role | Org Scope | Can Do |
|---|---|---|
| `super_admin` | All orgs (Company-level) | Everything — manage all cities, create national events, set any pricing |
| `city_admin` | One org | Manage that city's teams, coaches, events, pricing, members |
| `coach` | One team within one org | View their roster, manage their assigned team |
| `customer` | One or more orgs (as a member) | Book events, manage payments, view their household's schedule |

### How This Is Stored
The role is NOT a single field on the User record. It is a separate `UserRole` join table:

```
UserRole:
  User_ID
  Role         (Enum: super_admin, city_admin, coach, customer)
  Organization_ID  (Nullable — null means company-wide, only valid for super_admin)
```

This allows a parent to be a `customer` in Denver and also a `coach` in a different context, without any schema changes.

---

## Multi-Tenancy Implementation

### Subdomain Resolution (How the app knows which city it is)
When a request arrives at `denver.limitlesshoops.com`, the Go backend extracts the subdomain prefix (`denver`), looks up the matching Organization by `slug`, and attaches the `organization_id` to the request context. Every subsequent database query is automatically scoped to that org.

### National / Cross-City Queries
When a city app needs to display national events (e.g., a tournament their team is entered in), the query fetches:
```
Events WHERE organization_id = 'denver' OR organization_id IS NULL
        AND player is on roster
```
National content surfaces automatically within the city app. Families never need to navigate to a different subdomain to see their full picture.

### Adding a New City (Future)
1. Insert a new row into the `organizations` table (name, slug, city, state).
2. Create DNS entry: `chicago.limitlesshoops.com`
3. Assign a `city_admin` user for that org.
4. Done. No code changes required.

---

## Backend Service Boundaries

Each service in `apps/api/services/` is a self-contained Go package with its own:
- Handler (HTTP routes)
- Service (business logic)
- Repository (database queries)

They start as one deployable binary. When any single service needs independent scaling (e.g., the curriculum service under heavy video load), it can be extracted into its own binary and deployed separately without touching the others. The API surface (routes, request/response shapes) does not change.
