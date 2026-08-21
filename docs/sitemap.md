# Limitless Hoops Sitemap

This document outlines the routing architecture for the Limitless Hoops application ecosystem.

## 1. Booking App (`booking-react`)
This is the core application for users to manage their bookings, events, and checkout.

### Public Routes (No login required)
*   `/login` - **Sign In** (Redirects here if unauthenticated)
*   `/register` - **Sign Up**
*   `/events` - **Events Calendar** (Public view of available events)
*   `/events/:eventId` - **Event Details** (Dynamic route for specific event information)

### Protected Routes (Requires User login)
*   `/dashboard` - **User Dashboard** (Overview of upcoming bookings and recent activity)
*   `/account` - **Account Management** (Settings, password change, payment methods)
*   `/cart` - **Cart** (Review items before purchase)
*   `/checkout` - **Checkout** (Stripe payment integration)
*   `/checkout/success` - **Order Confirmation** (Redirect after successful payment)
*   `/checkout/cancel` - **Payment Failed/Canceled** (Redirect if user cancels Stripe payment)

### Admin / Coach Routes (Requires elevated permissions)
*   `/admin/dashboard` - **Admin Dashboard** (Overview of all system activity)
*   `/admin/events` - **Event Management** (Create, edit, delete events)
*   `/coach/dashboard` - **Coach Dashboard** (Overview of coach's specific classes/events)
*   `/coach/roster` - **Event Roster** (View attendees for a specific event)

### Global Routes
*   `*` - **404 Not Found** (Catch-all for invalid URLs)

---

## 2. Marketing Site / Landing Page (Future App)
This will be a completely separate application serving as the public face of the business.

### Public Routes
*   `/` - **Home Page** (Marketing, value proposition, calls to action)
*   `/about` - **About Us** 
*   `/contact` - **Contact / Support**
*   `/terms` - **Terms of Service**
*   `/privacy` - **Privacy Policy**

*Note: The marketing site will link directly to the Booking App (e.g., via "Book Now" buttons pointing to `app.limitlesshoops.com/login` or `app.limitlesshoops.com/events`).*
