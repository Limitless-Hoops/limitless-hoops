# Limitless Hoops - Cost & Architecture Analysis (100 Member Baseline)

When transitioning to a custom platform, there are unavoidable recurring infrastructure and processing costs. This document outlines the **high-end estimated costs** assuming a baseline of 100 active members paying an average of $200/month (totaling $20,000 processed monthly).

## 1. Payment Processing (Stripe vs PayPal)
Payment gateways take a cut of every transaction. If the business absorbs these fees instead of passing them to the customer, this will be the largest monthly expense.

### Option A: Credit Cards (High End Cost)
*   **Stripe Fee:** 2.9% + $0.30 per transaction.
*   **Cost for 100 Members:** $580 in percentages + $30 in fixed fees = **$610 / month**.
*   **Verdict:** This is the most expensive route, but offers the least friction for the customer.

### Option B: ACH / Bank Transfers (Recommended for Savings)
*   **Stripe Fee:** 0.8% (Capped at $5.00 maximum).
*   **Cost for 100 Members:** $200 * 0.8% = $1.60 per member = **$160 / month**.
*   **Verdict:** Switching customers to ACH saves the business **$450 every month** ($5,400/year). 

*Note: The business can also opt for a "Convenience Fee" model, where the 3% CC fee is added to the customer's cart, making payment processing effectively $0 for the business.*

---

## 2. Authentication (Identity Management)
We are adopting a **100% Password-less Strategy**. Users will log in via Google, Facebook, LinkedIn, or an Email One-Time-Passcode (OTP). 

### Custom OAuth & OTP (Go Backend)
*   **Cost:** **$0 / month**. 
*   **Pros:** Highly secure (no breached passwords), less friction for users, no vendor lock-in. Email OTPs cost fractions of a penny via AWS.
*   **Cons:** *Apple Sign-In is excluded for now as it requires a $99/yr Apple Developer fee.*

---

## 3. Hosting & Infrastructure (AWS)
Running a custom web app, database, and background workers requires server hosting. We will assume the high-end managed route for maximum reliability.

### AWS Managed Services (ECS + RDS)
*   **Estimated Cost:** **~$60 / month**.
*   **Breakdown:** 
    *   Managed PostgreSQL Database (RDS): ~$18/mo
    *   App Hosting (ECS / Fargate): ~$15/mo
    *   Load Balancer & Networking: ~$25/mo
*   **Verdict:** "Set it and forget it". Automated backups, highly secure, and scales automatically. 

---

## 4. Communications (Email & SMS)
Based on 100 members attending multiple events per week, we estimate heavy communication volume.

### SMS Text Reminders (High-End Volume)
*   **Assumption:** 4 event reminders per week * 4 weeks * 100 members = 1,600 texts. Let's budget for 3,000 texts/month to cover marketing and cancellations.
*   **Provider:** Twilio or AWS SNS (~$0.0079 per text).
*   **Estimated Cost:** **$24 / month**.

### Email Notifications (Receipts, Announcements)
*   **Assumption:** 10,000 emails/month.
*   **Provider:** AWS SES ($0.10 per 1,000 emails).
*   **Estimated Cost:** **$1 / month**.

---

## Summary of High-End Estimated Costs (100 Members)

Assuming the business absorbs ALL fees (no convenience fees passed to customers) and operates at the high-end of infrastructure and communication usage:

### Scenario 1: Everyone pays with Credit Card
*   **Payment Processing:** $610.00
*   **AWS Hosting:** $60.00
*   **SMS & Email:** $25.00
*   **Domain Name:** $1.00
*   **TOTAL ESTIMATED MONTHLY COST:** **$696.00 / month**

### Scenario 2: Everyone pays with ACH (Bank Transfer)
*   **Payment Processing:** $160.00
*   **AWS Hosting:** $60.00
*   **SMS & Email:** $25.00
*   **Domain Name:** $1.00
*   **TOTAL ESTIMATED MONTHLY COST:** **$246.00 / month**

*Conclusion: The platform itself costs roughly $86/mo to run on AWS at scale. The primary variable is how the business chooses to process and absorb the $20,000/mo in transaction volume.*
