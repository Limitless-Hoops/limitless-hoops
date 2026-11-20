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
*   **First_Name:** String (Collected during onboarding)
*   **Last_Name:** String (Collected during onboarding)
*   **DOB:** String (Collected during onboarding — used for age verification on legal waivers)
*   **Stripe_Customer_ID:** (To manage their saved payment methods)
*   **Referral_Code:** (Unique string — used for the affiliate/referral tracking system)
*   **Created_At / Updated_At:** Timestamps

> **Note:** `Email` and `Phone` are **NOT** single fields on the User record. They are stored in the `user_emails` and `user_phones` tables below to support multi-provider identity and multiple contacts per user.

## 3. User Emails
Stores one or more email addresses per user. Supports multi-provider OAuth where different social accounts may bring different emails.
*   **ID:** (UUID)
*   **User_ID:** Links to the User.
*   **Email:** (Unique across the entire table)
*   **Is_Primary:** Boolean — exactly one `true` per user. This email is the identity anchor and cannot be deleted. It does **not** control which email receives notifications (see `notification_preferences`).
*   **Is_Verified:** Boolean — must be `true` before the email can be used for Email OTP logins.
*   **Source_Provider:** (String: `'manual'`, `'google'`, `'facebook'`, `'twitter'`) — Non-manual emails are read-only in the UI and can only be removed by disconnecting the linked social provider.
*   **Label:** (Optional user-defined label, e.g. "Work", "Personal")
*   **Created_At:** Timestamp

## 4. User Phones
Stores one or more phone numbers per user.
*   **ID:** (UUID)
*   **User_ID:** Links to the User.
*   **Phone:** (Unique across the entire table)
*   **Is_Primary:** Boolean — exactly one `true` per user. Cannot be deleted.
*   **Is_Verified:** Boolean — must be `true` before the phone can be used for SMS OTP logins.
*   **Label:** (Optional user-defined label, e.g. "Cell", "Home")
*   **Created_At:** Timestamp

## 5. Notification Preferences
Controls which emails and phones receive which categories of notifications. Decoupled from the "primary" concept — any contact can receive any notification type independently.
*   **ID:** (UUID)
*   **User_ID:** Links to the User.
*   **Target_Type:** (Enum: `'email'` or `'phone'`)
*   **Target_ID:** FK → `user_emails.id` or `user_phones.id`
*   **Category:** (Enum: `'billing'`, `'events'`, `'announcements'`, `'security'`)
*   **Enabled:** Boolean (Default `true` when a contact is first added)

**Business Rules:**
- The `security` category cannot be disabled on the primary contact. This prevents account lockout.
- Seeded automatically for all 4 categories when a new email or phone is added.

## 6. UserRole (Role Assignment)
Roles are NOT a single field on the User record. They are stored in a join table so a user can hold different roles in different organizations (e.g., a Coach in Denver, a Customer in Chicago).
*   **User_ID:** Links to the User.
*   **Role:** (Enum: `super_admin`, `city_admin`, `coach`, `customer`)
*   **Organization_ID:** (Nullable) — Links to the Organization this role applies to. `null` is only valid for `super_admin` and means company-wide access.

*See `architecture.md` → Auth & Role Hierarchy for the full breakdown of what each role can do.*

## 7. Player (The Athlete)
This is the kid actually attending the workouts. They do not have login credentials.
*   **ID:** (UUID)
*   **Organization_ID:** Links to the Organization (city) this player belongs to.
*   **First_Name:** String
*   **Last_Name:** String
*   **Date_Of_Birth:** Date (Crucial for age-restricted events)
*   **Current_Grade:** Integer (Crucial for grade-based volunteer spots and age restrictions)
*   **Medical_Notes:** Text (Allergies, etc.)
*   **Prospect_Status:** (Enum: None, Waitlisted, Intent_Confirmed) — Tracks players in the pipeline who have not yet paid.

## 8. Household Management (The Bridge)
To support split-households (e.g., divorced parents who share custody), we use a join table that connects Users to Players. This allows Mom and Dad to both see "Johnny's" schedule from their separate accounts.
*   **User_ID:** Links to the Adult
*   **Player_ID:** Links to the Kid
*   **Relationship:** (Enum: Primary Guardian, Secondary Guardian, Emergency Contact)
*   **Can_Manage_Billing:** Boolean (Can this specific adult change the kid's membership tier?)

## 9. Subscriptions (Billing State)
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
    *   *(Note: "Paused" is NOT a valid status. These are strict contracts and cannot be paused.)*
*   **Auto_Renew:** Boolean (Default: `true`.)
*   **Referred_By_User_ID:** (Nullable) The User whose Referral_Code was entered at checkout.
*   **Current_Period_End:** Timestamp
*   **Stripe_Subscription_ID:** Links directly to Stripe's billing engine.
