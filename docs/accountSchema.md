# Household & Account Engine - Data Schema (Draft)

Youth sports organizations require a more complex account structure than standard apps because the person logging in and paying (the Parent) is rarely the person actually attending the event (the Player). Furthermore, modern families require split-household support (e.g., Mom and Dad have separate logins but manage the same kid).

*For the overall multi-tenancy and org hierarchy that governs this schema, see `architecture.md`.*

---

## 0. Company
The single top-level entity. There is exactly one Company record. All organizations (cities) belong to it. Exists to scope national/cross-city data that is not tied to any one city.
*   **ID:** (UUID)
*   **Name:** e.g., "Limitless Hoops"

## 1. Organization (City / Market)
Represents a city or regional market. Every piece of city-specific data is scoped to an Organization via `organization_id`. See `architecture.md` for how subdomain resolution maps to this record.
*   **ID:** (UUID)
*   **Company_ID:** Links to the Company.
*   **Name:** e.g., "Denver"
*   **Slug:** e.g., `"denver"` — Used for subdomain routing (`denver.limitlesshoops.com`).
*   **City:** String
*   **State:** String
*   **Is_Active:** Boolean (Is this market live?)

## 2. User (The Adult / Guardian)
This is the person who creates an account, logs in, and pays.
*(Note: We are using a 100% Password-less strategy. Login and Registration use the same unified flow. Users authenticate via Google OAuth, Email OTP, or SMS OTP.)*
*   **ID:** (UUID)
*   **First_Name:** String (Collected during onboarding if not provided by Auth)
*   **Last_Name:** String (Collected during onboarding if not provided by Auth)
*   **Email:** (Unique, Nullable if they created account strictly via SMS — though email should be highly encouraged for billing receipts)
*   **OAuth_Provider:** (Enum: `Google`, `Email_OTP`, `SMS_OTP`)
*   **OAuth_ID:** (String — The unique ID from the provider, or the email/phone used for OTP)
*   **Phone_Number:** (Crucial for SMS alerts and SMS OTP logins. Collected during onboarding if using Google/Email auth)
*   **Stripe_Customer_ID:** (To manage their saved payment methods)
*   **Referral_Code:** (Unique string — used for the affiliate/referral tracking system. See `teamBuildingEngine.md`)

## 3. UserRole (Role Assignment)
Roles are NOT a single field on the User record. They are stored in a join table so a user can hold different roles in different organizations (e.g., a Coach in Denver, a Customer in Chicago).
*   **User_ID:** Links to the User.
*   **Role:** (Enum: `super_admin`, `city_admin`, `coach`, `customer`)
*   **Organization_ID:** (Nullable) — Links to the Organization this role applies to. `null` is only valid for `super_admin` and means company-wide access.

*See `architecture.md` → Auth & Role Hierarchy for the full breakdown of what each role can do.*

## 4. Player (The Athlete)
This is the kid actually attending the workouts. They do not have login credentials.
*   **ID:** (UUID)
*   **Organization_ID:** Links to the Organization (city) this player belongs to.
*   **First_Name:** String
*   **Last_Name:** String
*   **Date_Of_Birth:** Date (Crucial for age-restricted events)
*   **Current_Grade:** Integer (Crucial for grade-based volunteer spots and age restrictions)
*   **Medical_Notes:** Text (Allergies, etc.)
*   **Prospect_Status:** (Enum: None, Waitlisted, Intent_Confirmed) — Tracks players in the pipeline who have not yet paid. "Intent_Confirmed" means the player responded to a monthly check-in and still intends to participate. See `businessLogic.md` for automated check-in rules.

## 5. Household Management (The Bridge)
To support split-households (e.g., divorced parents who share custody), we use a join table that connects Users to Players. This allows Mom and Dad to both see "Johnny's" schedule from their separate accounts.
*   **User_ID:** Links to the Adult
*   **Player_ID:** Links to the Kid
*   **Relationship:** (Enum: Primary Guardian, Secondary Guardian, Emergency Contact)
*   **Can_Manage_Billing:** Boolean (Can this specific adult change the kid's membership tier?)

## 6. Subscriptions (Billing State)
Because one parent might pay for two kids, and each kid might be on a different tier, subscriptions are tied to the *Player*, but paid by the *User*.
*   **ID:** (UUID)
*   **Player_ID:** The kid receiving the service.
*   **Organization_ID:** The city this subscription belongs to.
*   **Paying_User_ID:** The adult footing the bill.
*   **Tier_ID:** The membership level (Recreation, Seasonal, Limitless, Volunteer)
*   **Status:** (Enum: Active, Past_Due, Cancelled)
    *   `Active`: Subscription is in good standing.
    *   `Past_Due`: A payment failed; within the 7-day Dunning grace period.
    *   `Cancelled`: Contract ended or was terminated. Access is revoked.
    *   *(Note: "Paused" is NOT a valid status. These are strict contracts and cannot be paused. See `membershipTiers.md`)*
*   **Auto_Renew:** Boolean (Default: `true`. Families can toggle this off via the app within the defined opt-out windows. See `mvpScope.md`.)
*   **Referred_By_User_ID:** (Nullable) The User whose Referral_Code was entered at checkout.
*   **Current_Period_End:** Timestamp (When the current term ends or the next charge fires)
*   **Stripe_Subscription_ID:** Links directly to Stripe's billing engine.
