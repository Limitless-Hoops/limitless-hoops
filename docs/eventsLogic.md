# Limitless Hoops - Events & Scheduling Logic

This document outlines the types of events hosted by Limitless Hoops and the business logic governing who can attend them, how much they cost, and how they map to membership tiers.

## Event Categories

Based on the current schedule, events fall into several major categories:

### 1. Workouts & Open Play (High Frequency)
*   **Shooting Workouts**
*   **Skill Workouts**
*   **Open Gyms**
*   **Morning Boot Camps** (Seasonal)

### 2. Team & Elite Events
*   **Team Practices** *(Restricted: Elite Team Only)*
*   **Team Strategy Sessions**

### 3. Specialized Training (Clinics & Workshops)
*   **Skill Clinics**
*   **Pick n Roll Workshops**

### 4. Competitions & Assessments
*   **Local League Games**
*   **Travel Tournaments** *(Requires small fee)*
*   **Ability & Agility Combines** *(Requires small fee)*
*   **3-PT Tourney / Free Throw Rally / All-Star Game** (Quarterly special events)

### 5. Camps (Seasonal)
*   **Summer Camps**
*   **Winter Camps**

---

## Business Rules for Booking & Access

### 1. General Access Rule
By default, **every member can attend everything**, regardless of their membership tier (Recreation, Seasonal, Limitless). 

### 2. Elite Team Restriction
The only current hard restriction is **Team Practices**. These are strictly for players who have made the Elite Travel Team (who also pay the $500/year Elite fee).

### 3. Attendance Fees (Skin in the Game)
While most workouts and open gyms are fully covered by the monthly/quarterly/annual membership tuition, the client requires **small fees for Tournaments and Combines**. 
*   *Purpose:* To ensure accountability and that players actually show up when they book a spot for high-stakes or limited-capacity events.

### 4. Future Tiered Capabilities (The "Wishlist" Feature)
The client specifically noted: *"If we had tiered capabilities, I'd love to make it more events available for those who pay the whole year."*

**System Design Implications for this Feature:**
Because we are building this custom, we can absolutely support this. The event creation schema needs to support:
*   **Tier Gating:** Admins can set minimum required tiers (e.g., "Limitless Tier Only" for certain premium clinics).
*   **Tiered Pricing:** An event can have dynamic pricing based on the user's tier. For example, a Combine might be:
    *   Limitless Members: $0 (Free)
    *   Seasonal Members: $10
    *   Recreation Members: $25
    *   Non-Members: $50
*   **Early Access:** Allow higher tiers to book spots 48 hours before lower tiers (valuable for camps and clinics with strict player maximums).
