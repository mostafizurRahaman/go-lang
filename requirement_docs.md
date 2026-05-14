# Complete System Flow

# System Platforms

| User Type | Platform      |
| --------- | ------------- |
| Customer  | Mobile App    |
| Employee  | Mobile App    |
| Admin     | Web Dashboard |

---

# 1. Customer App Flow

# Step 1: Authentication

Customer:

- Registers/Login
- Completes Profile
- Saves Pickup Addresses

---

# Step 2: Customer Dashboard

Customer sees:

- Live Scrap Prices
- Sell Scrap Metal
- Sell Scrap Car
- Order History
- Notifications

---

# 2. Scrap Metal Selling Flow

## Step 1: Create Scrap Request

Customer:

- Selects Scrap Material Types
- Enters Estimated Weight
- Uploads Photos
- Adds Notes

---

## Step 2: Logistics Selection

Customer chooses:

### Option A: Drop-off

- Customer will deliver materials manually

### Option B: Pickup

- Customer selects saved address
- Pickup fee may apply

---

## Step 3: Submit Request

Customer submits RFQ.

### System Actions

- Create Request ID
- Notify Admin Dashboard
- Save uploaded files

**Status:** `Pending Review`

---

# 3. Scrap Car Selling Flow

## Step 1: VIN Entry

Customer enters VIN number.

### System Actions

VIN API fetches:

- Make
- Model
- Year
- Engine Details

---

## Step 2: Vehicle Checklist

Customer fills:

- Battery Available
- Wheels Available
- Catalytic Converter Available
- Missing Parts
- Damage Notes

Customer uploads:

- Vehicle Photos

---

## Step 3: Pickup/Drop-off Selection

Customer chooses:

- Drop-off OR
- Pickup

If pickup:

- Select address

---

## Step 4: Submit Request

### System Actions

- Save VIN Data
- Create RFQ (Request for Quote)
- Notify Admin

**Status:** `Pending Review`

---

# 4. Admin Dashboard Flow

# Step 1: New Request Notification

Admin dashboard receives:

- New Quote Request Alert

Admin opens request.

---

# Step 2: Review Request

Admin sees:

- Customer Details
- VIN Information
- Material Details
- Uploaded Images
- Pickup/Drop-off Selection

---

# Step 3: Pricing & Calculator

Admin enters:

- Scrap value
- Vehicle part values
- Weight calculations
- Pickup fee adjustments

### System Calculates

```text id="rk4s9q"
Final Offer =
Scrap Value
- Pickup Fee
= Net Offer
```

---

# Step 4: Send Offer

Admin clicks:

- Send Offer

### System Actions

- Notify Customer App
- Save Offer History

**Status:** `Offer Sent`

---

# 5. Customer Offer Flow

# Step 1: Customer Receives Offer

Customer sees:

- Scrap Value
- Pickup Fee
- Final Offer Amount

---

# Step 2: Customer Decision

## Option A: Accept Offer

### System Actions

- Create Order
- Generate Order ID

**Status:** `Accepted`

---

## Option B: Reject Offer

Order closes.

**Status:** `Rejected`

---

## Option C: Negotiate

Customer sends counter request.

**Status:** `Negotiation`

Admin reviews again.

---

# 6. Admin Assignment Flow

# Step 1: Assign Employee

Admin:

- Selects Employee/Driver
- Sets Pickup Schedule

### System Actions

- Notify Employee App
- Notify Customer App

**Status:** `Assigned`

---

# 7. Employee App Flow

# Step 1: Employee Dashboard

Employee sees:

- Assigned Orders
- Customer Information
- Pickup Address
- Pickup Schedule

---

# Step 2: Employee Contacts Customer

Employee can:

- Chat with Customer

Example:

> “I will arrive around 3 PM.”

---

# 8. Physical Inspection Flow (Outside System)

Inspection happens manually outside the system.

Employee physically checks:

- Scrap condition
- Vehicle condition
- Missing parts
- Material quantity

---

# 9. Offer Adjustment Flow (Optional)

## If Everything Matches

Employee proceeds with pickup.

---

## If Condition Is Different

Employee informs Admin manually.

### Admin Actions

- Recalculate Offer
- Update Offer in Dashboard

### System Actions

- Send Updated Offer to Customer

**Status:** `Offer Updated`

---

# 10. Customer Revised Offer Flow

Customer receives:

- Updated Offer

Customer can:

- Accept OR
- Reject

If accepted:

- Pickup continues

If rejected:

- Order cancelled

---

# 11. Pickup Completion Flow

# Step 1: Employee Collects Scrap

Employee:

- Loads materials/vehicle
- Confirms pickup

---

# Step 2: Cash Payment

Employee gives customer cash manually.

Employee updates:

- Payment Collected

---

# Step 3: Employee Completes Order

Employee clicks:

- Mark as Completed

### System Actions

- Save payment logs
- Notify Admin
- Notify Customer

**Status:** `Completed`

---

# 12. Customer Finalization Flow

Customer receives:

- Pickup Completed Notification
- Payment Confirmation
- Receipt/Transaction Record

---

# 13. Admin Monitoring Dashboard

Admin dashboard includes:

## Order Management

- Pending Requests
- Active Orders
- Completed Orders
- Cancelled Orders

---

## Employee Management

- Assign Employees
- Track Task Status
- Employee Availability

---

## Offer Management

- Create Offers
- Update Offers
- Negotiations

---

## Reports & Analytics

- Revenue Reports
- Scrap Volume Reports
- Employee Performance
- Completed Pickups
- Offer Conversion Rates

---

# 14. Customer App Features

## Main Features

- Authentication
- Scrap Selling
- VIN Decoder
- Offer Viewing
- Order Tracking
- Notifications
- Chat with Employee
- Order History

---

# 15. Employee App Features

## Main Features

- Login
- Assigned Orders
- Customer Contact
- Pickup Status Update
- Cash Collection Confirmation
- Chat System

---

# 16. Admin Dashboard Features

## Main Features

- Quote Management
- Offer Calculator
- Employee Assignment
- Customer Management
- Reports & Analytics
- Notifications
- Order Monitoring

---

# 17. Complete Status Lifecycle

```text id="i7ktm4"
Pending Review
       ↓
Offer Sent
       ↓
Accepted
       ↓
Assigned
       ↓
Offer Updated (Optional)
       ↓
Pickup Completed
       ↓
Completed
```

---

# 18. Core Integrations

## Required APIs

- VIN Decoder API
- Push Notification Service
- Email Service

---

# 19. Simplified System Architecture

## Customer App

- React Native / Flutter

## Employee App

- React Native / Flutter

## Admin Dashboard

- React / Next.js Web Dashboard

## Backend

- Node.js API Server

## Database

- MongoDB

## Storage

- AWS S3 / Cloud Storage 
- Deployment
