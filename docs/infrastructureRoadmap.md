# Infrastructure Setup & Onboarding Roadmap

This document serves as a step-by-step checklist for setting up the third-party infrastructure for Limitless Hoops. It is broken down by phase so you only create accounts and incur costs when absolutely necessary.

---

## Phase 1: Foundation (Do This Now)
*You are currently doing this. This sets up the base cloud environment.*

- [ ] **Create GCP Account & Project**
  - Create a new project in Google Cloud Console named "Limitless Hoops".
  - Attach a billing account.
- [ ] **Enable Core GCP APIs**
  - Enable **Cloud Run API** (for hosting the app).
  - Enable **Cloud SQL Admin API** (for the Postgres database).
- [ ] **Setup Authentication**
  - Enable **Identity Platform** (or Firebase Auth).
  - Configure Google as a Sign-In provider.
  - Enable Email/Passwordless or Phone OTP providers.

---

## Phase 2: Company Inboxes (Do This When You Need an @limitlesshoops.com Email)
*This replaces AWS WorkMail. Do this when you are ready to stop using a personal Gmail account for business.*

- [ ] **Create Google Workspace Account**
  - Sign up for Google Workspace "Business Starter" ($6/mo).
  - Create your primary human inbox (e.g., `premo@limitlesshoops.com` or `admin@limitlesshoops.com`).
- [ ] **Configure Free Aliases (Crucial Cost-Saver)**
  - Go to the Google Workspace Admin Console.
  - Add **Aliases** to your primary inbox for:
    - `billing@limitlesshoops.com`
    - `support@limitlesshoops.com`
    - `info@limitlesshoops.com`
  - *Note: You do not pay extra for these. Any emails sent to these addresses will automatically route into your primary inbox.*

---

## Phase 3: Automated App Emails (Do This When Coding the Backend)
*This handles all transactional emails like receipts, event updates, and Email OTPs.*

- [ ] **Create Resend Account**
  - Sign up for the free tier (3,000 emails/mo).
- [ ] **Verify Domain**
  - Add your domain (`limitlesshoops.com`).
  - Resend will give you a set of DNS records (TXT/DKIM). Add these to your Domain Registrar or GCP Cloud DNS. This proves you own the domain and keeps emails out of spam folders.
- [ ] **Configure the Go Backend**
  - Generate a Resend API Key.
  - Set the "From" address in your Go code to `noreply@limitlesshoops.com` or `billing@limitlesshoops.com`.
  - *Note: Because you set up the alias in Phase 2, if a user clicks "Reply" to a receipt sent by Resend, it will safely land in your Google Workspace inbox.*

---

## Phase 4: Automated SMS (Do This 3-4 Weeks Before Launch)
*This handles SMS OTPs and text message alerts to parents/coaches. Start this early due to US telecom compliance laws.*

- [ ] **Create Twilio Account**
  - Sign up and purchase a local US Phone Number.
- [ ] **A2P 10DLC Compliance Registration (CRITICAL)**
  - Navigate to the **Twilio Trust Hub**.
  - Register your business as a **Brand** (requires EIN/Tax ID).
  - Register a **Campaign** for your Brand. 
  - Select use cases like "Account Notifications" and "2FA".
  - *Warning: Telecom carriers manually review these. It can take 2-4 weeks for approval. SMS delivery to US numbers will fail or be heavily filtered if this is not complete.*
- [ ] **Configure the Go Backend**
  - Generate Twilio Account SID and Auth Token.
  - Integrate into the backend notification service.
