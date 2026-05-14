---

### 1. User Flow (The Customer Journey)

1. **Dashboard/Landing:**
   - Views live rates (e.g., $0.12/lb Steel).
   - **Option A (Sell Metals):** Enters material types and weights.
   - **Option B (Sell Car):** Enters VIN. System auto-fills details (Make,
     Model, Year). Customer completes the "Condition Checklist"
     (Battery/Wheels/Converter status).
2. **Logistics Selection:**
   - Selects **Drop-off** (Calculates price based on weight).
   - Selects **Pickup** (Opens Address Book -> Selects saved address).
3. **Request Submission:** User submits RFQ. Status: `Pending Admin Review`.
4. **Offer Review:**
   - Receives Notification (Email/Push).
   - Views Breakdown (Scrap Value - Logistics Fee = Net Offer).
   - **Action:** Clicks "Accept" or "Negotiate/Reject."
5. **Execution:**
   - Status shifts to `Offer Accepted`.
   - Opens "Order Details" page: Sees "Chat with Driver" button and "Live Map"
     (when assigned).
6. **Finalization:** Receives payment confirmation and marks order as
   `Completed`.

---

### 2. Admin Flow (The Operational Hub)

1. **Notification Hub:** Receives "New Quote Request" alert.
2. **The Calculator (The Engine):**
   - Opens Quote Request.
   - **VIN Integration:** Sees vehicle specs + Customer's "Condition Checklist."
   - **Calculation:** Inputs values into the built-in Calculator (Logic mirrored
     from your Excel).
   - **Logistics Adjustment:** If "Pickup" is selected, the Admin toggles the
     "Pickup Fee" input. The system renders the `Final Offer`.
3. **Sending Offer:** Admin clicks "Send Offer."
4. **Assignment:** Once accepted, Admin assigns an Employee from the "Active
   Drivers" list.
5. **Monitoring:** Oversees the "Live Order Dashboard" where all active pickups
   are plotted on a map.

---

### 3. Employee Flow (The Logistics Side)

1. **Login:** Views assigned tasks for the day.
2. **Preparation:** Clicks "Start Trip." Status updates to `En Route` (Customer
   sees live location).
3. **Communication:** Accesses the "Chat" to contact the customer (e.g., "I'm 10
   minutes away, please ensure the gate is open").
4. **Site Inspection (Critical Step):**
   - Arrives at the location.
   - Inspects vehicle/metal. If the condition is worse than the customer
     described, the employee clicks "Update Inspection."
   - **System Logic:** If updated, the Admin receives an alert to "Re-calculate"
     the offer before the employee loads the goods.
5. **Completion:** Confirms pickup. Status moves to `Received`.

---

### 4. Logic & Flow Visualization

| Stage          | Trigger          | System Action                 | Status Update     |
| :------------- | :--------------- | :---------------------------- | :---------------- |
| **Submission** | Customer Submits | VIN API Decode + Logic Calc   | `Pending`         |
| **Review**     | Admin Inputs     | Apply Calculator + Pickup Fee | `Offer Sent`      |
| **Acceptance** | Customer Clicks  | Create Order ID               | `Accepted`        |
| **Logistics**  | Admin Assigns    | Notify Driver                 | `Driver Assigned` |
| **En Route**   | Driver Starts    | Enable GPS Tracking           | `En Route`        |
| **Inspection** | Driver Arrives   | Verify items against Quote    | `In Progress`     |
| **Closing**    | Driver Completes | Trigger Payment               | `Completed`       |

---

### Key Technical Considerations for Success

- **The "Calculator" Component:** Do not just write this as a simple form. Build
  it as a **reusable Component** in your code. This way, if you update the Excel
  sheet logic later, you only change the code in one place, and it updates for
  the Admin, the Quote preview, and the Final Invoice.
- **Address Book Management:** Ensure the address book allows for **Geocoding**
  (converting the address to Latitude/Longitude). This is required for the
  "Pickup Fee" calculation (e.g., you could automate the fee based on distance:
  `$2/mile`).
- **The Chat System:** Keep the chat **bound to the Order ID**. When an order is
  marked `Completed`, the chat should be archived, and the "Message" button
  should disappear so users can't spam employees after hours.

**Does this flow cover all the scenarios you imagined, or should we add a
specific "Payment" workflow (e.g., Bank Transfer vs. Cash on Delivery)?**
