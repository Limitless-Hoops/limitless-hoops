# Limitless Hoops Booking System 2.0

This repository contains the next-generation booking and management system for Limitless Hoops. The architecture is designed to support a robust web-first experience, paving the way for a dedicated native mobile app in the future.

## Finalized Tech Stack

### Frontend (Web Apps & Landing Page)
*   **Framework:** React
*   **Language:** TypeScript
*   **Styling:** TailwindCSS
*   **Data Fetching/State:** TanStack Query
*   **Future Mobile App:** Flutter & Dart (to be built against the stabilized Go API)

### Backend (API & Business Logic)
*   **Language:** Go
*   **Framework:** Fiber
*   **Database ORM:** GORM
*   **Future Caching:** Redis (deferred until scaling demands it)

### DevOps & Infrastructure
*   **Containerization:** Docker
*   **CI/CD:** GitHub Actions
*   **Infrastructure as Code:** Terraform
*   **Cloud Provider:** AWS

### Third-Party Integrations
*   **Payments:** Stripe (prioritizing ACH / Bank Transfers for recurring billing to reduce fees, with CC available)
*   **AI Enhancements:** TBD (Web helpers / support automation)

## Project Strategy
1.  **Web-First:** Build the responsive React web app to immediately solve business bottlenecks (Shopping Cart checkout, household management, admin dashboard).
2.  **API Stabilization:** Use the React app to battle-test the Go API.
3.  **Mobile Expansion:** Once the web app is live and the API is stable, build the native mobile app in Flutter.
