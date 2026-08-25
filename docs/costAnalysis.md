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

### Custom OAuth & OTP (Go Backend)
*   **Cost:** **$0 / month**. 
*   **Pros:** Highly secure (no breached passwords), less friction for users, no vendor lock-in. Email OTPs cost fractions of a penny via AWS.
*   **Cons:** *Apple Sign-In is excluded for now as it requires a $99/yr Apple Developer fee.*

---

## 3. Hosting & Infrastructure (AWS)
Running a custom web app, database, and background workers requires server hosting. We will assume the high-end managed route for maximum reliability. The base infrastructure has fixed minimum costs, but scales automatically.

### AWS Managed Services (ECS + RDS)
*   **Estimated Base Cost:** **~$60 / month**.
*   **Breakdown:** 
    *   Managed PostgreSQL Database (RDS): ~$18/mo
    *   App Hosting (ECS / Fargate): ~$15/mo
    *   Load Balancer & Networking: ~$25/mo
*   **Verdict:** "Set it and forget it". Automated backups, highly secure, and scales automatically. As traffic grows, AWS will dynamically provision more resources, leading to a linear but marginal increase in costs.

---

## 4. Communications (Email & SMS)
Communication costs scale directly with the size of the roster and frequency of events.

### SMS Text Reminders (High-End Volume)
*   **Provider:** Twilio or AWS SNS (~$0.0079 per text).
*   **Estimated Cost:** ~$7.90 per 1,000 text messages sent.

### Email Notifications (Receipts, Announcements)
*   **Provider:** AWS SES ($0.10 per 1,000 emails).
*   **Estimated Cost:** Highly affordable. 10,000 emails costs $1.00.

---

## Summary of Costs

The platform itself has a very low base operating cost (roughly $86/mo to run on AWS at base scale, including domain names and minor communications). 

The primary variable expense is entirely dependent on how much revenue is processed through the platform and the payment methods chosen:
*   **Credit Card Heavy:** Approximately 3% of gross revenue goes to Stripe.
*   **ACH Heavy:** Approximately 0.8% of gross revenue goes to Stripe.
*   **Convenience Fee Model:** Customers absorb the Stripe fees, netting the business 100% of the listed price.
