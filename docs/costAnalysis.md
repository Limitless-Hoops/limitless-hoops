# Limitless Hoops - Cost & Architecture Analysis

When transitioning to a custom platform, there are unavoidable recurring infrastructure and processing costs. This document outlines the **estimated costs**, emphasizing that operational costs scale linearly with the number of active members and transaction volumes. 

## 1. Stripe Payment Processing
Payment gateways take a cut of every transaction. If the business absorbs these fees instead of passing them to the customer, this will be the largest variable expense.

### Option A: Credit Cards (High End Cost)
*   **Stripe Fee:** 2.9% + $0.30 per transaction.
*   **Example (Per $1,000 processed):** $29.00 in percentage fees + $0.30 per charge.
*   **Verdict:** This is the most expensive route, but offers the least friction for the customer.

### Option B: ACH / Bank Transfers (Recommended for Savings)
*   **Stripe Fee:** 0.8% (Capped at $5.00 maximum per transaction).
*   **Example (Per $1,000 processed):** $8.00 in percentage fees (assuming a single $1k charge).
*   **Verdict:** Switching customers to ACH yields significant savings. For every $1,000 processed via ACH instead of CC, the business saves roughly $21.

*Note: The business can also opt for a "Convenience Fee" model, where the ~3% CC fee is added to the customer's cart, making payment processing effectively $0 for the business.*

---

## 2. Authentication (Identity Management)
We are adopting a **100% Password-less Strategy**. Users will log in via Google, Facebook, LinkedIn, or an Email One-Time-Passcode (OTP). 

### Google Identity / Custom OTP
*   **Cost:** **$0 / month**. (Firebase Auth / Identity Platform has a massive free tier).
*   **Pros:** Highly secure (no breached passwords), less friction for users, built-in Google SSO support.
*   **Cons:** *Apple Sign-In is excluded for now as it requires a $99/yr Apple Developer fee.*

---

## 3. Hosting & Infrastructure (GCP)
Running a custom web app, database, and background workers requires server hosting. We are using a serverless approach with Google Cloud Platform (GCP) to minimize costs while allowing infinite scale.

### GCP Serverless (Cloud Run + Cloud SQL)
*   **Estimated Base Cost:** **~$15 - $20 / month**.
*   **Breakdown:** 
    *   Managed PostgreSQL Database (Cloud SQL Micro): ~$15/mo
    *   App Hosting (Cloud Run): $0 to $5/mo (Scales to zero, massive free tier)
    *   Load Balancer & Networking: $0/mo (Included securely with Cloud Run)
*   **Verdict:** Highly cost-effective for early-stage products. Cloud Run scales down to exactly $0 when there is no traffic, avoiding the "always-on" tax of traditional cloud architecture.

---

## 4. Communications (Email & SMS)
Communication costs scale directly with the size of the roster and frequency of events.

### SMS Text Reminders
*   **Provider:** Twilio (~$0.0079 per text).
*   **Estimated Cost:** ~$7.90 per 1,000 text messages sent.

### Email Notifications (Receipts, OTPs, Announcements)
*   **Provider:** Resend.
*   **Estimated Cost:** **$0/mo** (First 3,000 emails per month are free). Once passing 3,000/mo, the cost is a flat $20/mo for up to 50,000 emails.

---

## Summary of Costs

The platform itself has a very low base operating cost (roughly $25-$30/mo to run on GCP at base scale, including Twilio SMS and domains). 

The primary variable expense is entirely dependent on how much revenue is processed through the platform and the payment methods chosen:
*   **Credit Card Heavy:** Approximately 3% of gross revenue goes to Stripe.
*   **ACH Heavy:** Approximately 0.8% of gross revenue goes to Stripe.
*   **Convenience Fee Model:** Customers absorb the Stripe fees, netting the business 100% of the listed price.
