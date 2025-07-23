# 🍕 GoFood Order Flow Simulation

## 📋 Current State Analysis

### ✅ What We Have:
- **User System**: Registration, login, JWT auth
- **Merchant System**: Restaurant profiles, menu items
- **Order System**: Basic order creation (without payment)
- **Driver System**: Complete CRUD for drivers
- **Account System**: Multi-wallet (main_balance, gopay_points, savings)
- **Transaction System**: Basic transaction recording

### ❌ What's Missing:
- **Payment Integration**: Order → Wallet deduction
- **Driver Assignment**: Auto-assign available drivers
- **Transaction Recording**: Proper transaction history for all parties
- **Order Status Flow**: Complete status management
- **Error Handling**: Payment failures, insufficient balance
- **Concurrency**: Handle multiple orders simultaneously

---

## 🔄 Complete Order Flow Design

### **Phase 1: Order Creation & Payment (User Side)**

```
USER PLACES ORDER
    ↓
1. Validate Order Request
    - Check merchant exists and is open
    - Check menu items availability
    - Calculate total (items + delivery fee)
    ↓
2. Check User Balance
    - Get user's main_balance
    - Verify sufficient funds (total_amount)
    - Handle insufficient balance error
    ↓
3. Reserve/Hold Payment (Optional - Advanced)
    - Create pending transaction
    - Hold amount in user's account
    - Prevent double spending
    ↓
4. Create Order Record
    - Save order with status: "pending"
    - Save order items
    - Generate order ID
    ↓
5. Process Payment
    - Deduct from user's main_balance
    - Create transaction record (type: "payment", category: "food")
    - Update order status: "confirmed"
    ↓
6. Find Available Driver
    - Query available drivers by vehicle type
    - Select closest driver (for now: random selection)
    - Assign driver to order
    ↓
7. Notify All Parties
    - User: Order confirmed, driver assigned
    - Merchant: New order received
    - Driver: New delivery assigned
    ↓
ORDER CREATED SUCCESSFULLY
```

### **Phase 2: Order Processing (Merchant Side)**

```
MERCHANT RECEIVES ORDER
    ↓
1. Order Status: "confirmed" → "cooking"
    - Merchant accepts and starts cooking
    - Update order status
    - Notify user and driver
    ↓
2. Order Status: "cooking" → "ready_for_pickup"
    - Food is ready
    - Update order status
    - Notify driver to pickup
```

### **Phase 3: Delivery Process (Driver Side)**

```
DRIVER RECEIVES ASSIGNMENT
    ↓
1. Order Status: "ready_for_pickup" → "delivery"
    - Driver picks up food from merchant
    - Update order status
    - Start delivery tracking
    ↓
2. Order Status: "delivery" → "completed"
    - Driver delivers to customer
    - Update order status
    - Process final transactions
```

### **Phase 4: Transaction Settlement (System Side)**

```
ORDER COMPLETED
    ↓
1. Calculate Earnings Distribution
    - Merchant earnings: order_total - delivery_fee - platform_fee
    - Driver earnings: delivery_fee - platform_fee
    - Platform earnings: platform_fee
    ↓
2. Create Settlement Transactions
    - Credit merchant account (if they have one)
    - Credit driver account (if they have one)
    - Record platform earnings
    ↓
3. Update All Balances
    - Merchant balance += earnings
    - Driver balance += delivery_fee
    - Platform balance += fees
    ↓
TRANSACTION COMPLETE
```

---

## 🏗️ Implementation Strategy

### **Option A: Sequential Processing (Recommended for MVP)**

**Pros:**
- Simple to implement and debug
- Easy error handling and rollback
- Clear transaction boundaries
- No race conditions

**Cons:**
- Slower processing
- Blocks on each step
- Less scalable

**Implementation:**
```go
func (s *OrderService) ProcessCompleteOrder(order *CreateOrderRequest) error {
    // 1. Validate order
    // 2. Check balance
    // 3. Create order
    // 4. Process payment
    // 5. Assign driver
    // 6. Return success
    
    // All in one database transaction for consistency
}
```

### **Option B: Asynchronous with Background Jobs (Advanced)**

**Pros:**
- Faster response to user
- Better scalability
- Can handle high volume
- Resilient to failures

**Cons:**
- Complex error handling
- Requires job queue system
- Harder to debug
- Eventual consistency

**Implementation:**
```go
func (s *OrderService) CreateOrder(order *CreateOrderRequest) error {
    // 1. Validate and create order quickly
    // 2. Queue background job for payment processing
    // 3. Queue background job for driver assignment
    // 4. Return order created (pending payment)
}
```

### **🎯 Recommended Approach: Sequential for Now**

Start with sequential processing to get the complete flow working, then optimize with async later.

---

## 📝 Detailed Implementation Tasks

### **1. Enhance Order Creation Handler**

```go
// Current: Basic order creation
// Need: Complete order processing with payment

func (h *OrderHandler) CreateOrder(c echo.Context) error {
    // 1. Validate request
    // 2. Check user balance
    // 3. Find available driver
    // 4. Process payment (deduct from wallet)
    // 5. Create order with assigned driver
    // 6. Create transaction records
    // 7. Return complete order details
}
```

### **2. Payment Integration Service**

```go
// New service: OrderPaymentService
type OrderPaymentService struct {
    accountService *AccountService
    transactionService *TransactionService
    orderService *OrderService
}

func (s *OrderPaymentService) ProcessOrderPayment(userID uint, orderTotal float64, orderID uint) error {
    // 1. Check balance
    // 2. Create transaction
    // 3. Update account balance
    // 4. Handle errors and rollback
}
```

### **3. Driver Assignment Service**

```go
// Enhance DriverService with assignment logic
func (s *DriverService) AssignDriverToOrder(orderID uint, vehicleType VehicleType) (*DriverProfile, error) {
    // 1. Find available drivers
    // 2. Select best driver (closest, highest rating, etc.)
    // 3. Assign driver to order
    // 4. Update driver status
}
```

### **4. Transaction Recording Enhancement**

```go
// Enhanced transaction creation for orders
func (s *TransactionService) CreateOrderTransaction(
    userID uint, 
    merchantID uint, 
    driverID uint, 
    orderID uint, 
    amount float64,
) error {
    // Create transactions for:
    // - User payment (debit)
    // - Merchant earning (credit) 
    // - Driver earning (credit)
    // - Platform fee (credit)
}
```

### **5. Error Handling & Rollback**

```go
// Handle various failure scenarios:
// - Insufficient balance
// - No available drivers
// - Payment processing failure
// - Order creation failure

func (s *OrderService) RollbackOrder(orderID uint) error {
    // 1. Reverse payment
    // 2. Release driver
    // 3. Delete order
    // 4. Restore balances
}
```

### **6. Order Status Management**

```go
// Complete status flow with validation
func (s *OrderService) UpdateOrderStatus(orderID uint, newStatus OrderStatus, userID uint) error {
    // 1. Validate status transition
    // 2. Check user authorization
    // 3. Update status
    // 4. Trigger side effects (payments, notifications)
}
```

---

## 🔄 Database Transaction Strategy

### **Single Transaction Approach (Recommended)**

```go
func (s *OrderService) ProcessCompleteOrder(req *CreateOrderRequest) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        // 1. Create order
        // 2. Process payment
        // 3. Assign driver
        // 4. Create transactions
        // All or nothing - rollback on any failure
    })
}
```

### **Multi-Transaction with Compensating Actions**

```go
func (s *OrderService) ProcessCompleteOrder(req *CreateOrderRequest) error {
    // 1. Create order (transaction 1)
    // 2. Process payment (transaction 2)
    // 3. Assign driver (transaction 3)
    // If any fails, run compensating actions to undo previous steps
}
```

---

## 🧪 Testing Strategy

### **Test Scenarios to Cover:**

1. **Happy Path**: Order → Payment → Driver Assignment → Completion
2. **Insufficient Balance**: Order fails at payment step
3. **No Available Drivers**: Order fails at assignment step
4. **Merchant Closed**: Order fails at validation step
5. **Menu Item Unavailable**: Order fails at validation step
6. **Payment Processing Error**: Order fails with rollback
7. **Concurrent Orders**: Multiple users ordering simultaneously
8. **Driver Busy**: Driver becomes unavailable during assignment

### **Test Data Needed:**

```go
// Test users with different balance levels
testUser1 := User{Balance: 100000} // Sufficient
testUser2 := User{Balance: 1000}   // Insufficient

// Test merchants with different statuses
testMerchant1 := Merchant{IsOpen: true}
testMerchant2 := Merchant{IsOpen: false}

// Test drivers with different availability
testDriver1 := Driver{Status: "online", VehicleType: "motorcycle"}
testDriver2 := Driver{Status: "offline", VehicleType: "car"}
```

---

## 🚀 Implementation Priority

### **Phase 1: Core Flow (High Priority)**
1. ✅ Payment integration with wallet deduction
2. ✅ Driver assignment logic
3. ✅ Transaction recording for all parties
4. ✅ Error handling and rollback

### **Phase 2: Enhanced Features (Medium Priority)**
1. ⚡ Real-time status updates
2. 📱 Notification system
3. 🗺️ Driver location tracking
4. 💰 Dynamic pricing

### **Phase 3: Production Ready (Low Priority)**
1. 🔄 Asynchronous processing
2. 📊 Analytics and reporting
3. 🛡️ Advanced fraud detection
4. ⚖️ Load balancing

---

## 📋 Next Steps

1. **Start with Sequential Implementation**
2. **Focus on Payment Integration First**
3. **Add Driver Assignment Logic**
4. **Implement Comprehensive Error Handling**
5. **Test Each Component Thoroughly**
6. **Optimize for Concurrency Later**

**Target**: Complete order flow working end-to-end with payment and driver assignment!
