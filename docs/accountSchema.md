# Household & Account Engine - Data Schema (Draft)

Youth sports organizations require a more complex account structure than standard apps because the person logging in and paying (the Parent) is rarely the person actually attending the event (the Player). Furthermore, modern families require split-household support (e.g., Mom and Dad have separate logins but manage the same kid).

## 1. User (The Adult / Guardian)
This is the person who creates an account, logs in, and pays. 
*(Note: We are using a 100% Password-less strategy for security and friction reduction. Users log in via OAuth or Email OTP).*
*   **ID:** (UUID)
*   **Email:** (Unique)
*   **OAuth_Provider:** (Enum: Google, Facebook, LinkedIn, Email_OTP)
*   **OAuth_ID:** (String - The unique ID from the provider)
*   **Phone_Number:** (Crucial for the SMS alerts)
*   **Stripe_Customer_ID:** (To manage their saved payment methods)
*   **Role:** (Enum: Customer, Coach, Admin)

## 2. Player (The Athlete)
This is the kid actually attending the workouts. They do not have login credentials.
*   **ID:** (UUID)
*   **First_Name:** String
*   **Last_Name:** String
*   **Date_Of_Birth:** Date (Crucial for age-restricted events)
*   **Current_Grade:** Integer (Crucial for the "Pre-3rd Grade 50% off" discount rule)
*   **Medical_Notes:** Text (Allergies, etc.)
*   **Active_Tier_ID:** Links to the Membership Tier they are currently on.

## 3. Household Management (The Bridge)
To support split-households (e.g., divorced parents who share custody), we use a "join table" that connects Users to Players. This allows Mom and Dad to both see "Johnny's" schedule from their separate accounts.
*   **User_ID:** Links to the Adult
*   **Player_ID:** Links to the Kid
*   **Relationship:** (Enum: Primary Guardian, Secondary Guardian, Emergency Contact)
*   **Can_Manage_Billing:** Boolean (Can this specific adult change the kid's membership tier?)

## 4. Subscriptions (Billing State)
Because one parent might pay for two kids, and each kid might be on a different tier, subscriptions are tied to the *Player*, but paid by the *User*.
*   **ID:** (UUID)
*   **Player_ID:** The kid receiving the service.
*   **Paying_User_ID:** The adult footing the bill.
*   **Tier_ID:** The membership level (Recreation, Seasonal, etc.)
*   **Status:** (Enum: Active, Paused, Past_Due, Cancelled)
*   **Current_Period_End:** Timestamp (When the next charge happens)
*   **Stripe_Subscription_ID:** Links directly to Stripe's billing engine.
