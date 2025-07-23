# 🔄 Order Status Update Authorization Rules

## 📋 Current Implementation Status

✅ **Basic Status Update**: `PUT /orders/:id/status` endpoint is implemented but **without authorization checks**
❌ **Authorization Logic**: Not implemented yet - any authenticated user can update any order status

---

## 🎯 Required Authorization Rules

### 🔐 Who Can Update What Status

#### **👤 User (Customer)**

- **Can Update**:
  - `pending` → `cancelled` (only if order hasn't been confirmed yet)
- **Cannot Update**: Any other status transitions
- **Authorization Check**: `order.UserID == loggedInUserId`
- **Business Rule**: Users can only cancel their own pending orders

#### **🏪 Merchant (Restaurant Owner)**

- **Can Update**:
  - `pending` → `confirmed` (accept the order)
  - `confirmed` → `cooking` (start preparing food)
  - `cooking` → `ready_for_pickup` (food is ready)
- **Cannot Update**: `delivery`, `completed`, `cancelled`
- **Authorization Check**: `merchant.UserId == loggedInUserId` (via order.MerchantID)
- **Business Rule**: Merchants control the food preparation pipeline

#### **🚗 Driver**

- **Can Update**:
  - `ready_for_pickup` → `delivery` (picked up the order)
  - `delivery` → `completed` (delivered to customer)
- **Cannot Update**: `pending`, `confirmed`, `cooking`, `cancelled`
- **Authorization Check**: `order.DriverID == loggedInUserId`
- **Business Rule**: Drivers control the delivery pipeline

---

## 🚫 Status Transition Rules (No Reverse Updates)

### **Valid Forward Transitions Only**

```
pending → confirmed → cooking → ready_for_pickup → delivery → completed
    ↓
cancelled (only from pending)
```

### **Invalid Transitions** (Should be blocked)

- Any backward transitions (e.g., `cooking` → `confirmed`)
- Skipping steps (e.g., `pending` → `cooking`)
- Invalid role transitions (e.g., driver updating to `cooking`)

---

## 🛠️ Implementation Plan

### 1. **Create Authorization Helper Function**

```go
func (h *OrderHandler) validateStatusUpdate(order *models.Order, userID uint, newStatus models.OrderStatus) error {
    // Check if user has permission to update this order
    // Check if status transition is valid
    // Return error if not authorized or invalid transition
}
```

### 2. **Update Order Handler**

- Add authorization check before calling service
- Fetch order first to determine current status and ownership
- Validate both permission and transition logic

### 3. **Service Layer Updates**

- Add status transition validation in `OrderService.UpdateOrderStatus`
- Prevent invalid status flows at database level

### 4. **Route Protection**

- Ensure proper JWT middleware is applied
- Consider role-based middleware for different user types

---

## 📝 Current Order Status Constants

```go
const (
    OrderPending   OrderStatus = "pending"
    OrderConfirmed OrderStatus = "confirmed"
    OrderCooking   OrderStatus = "cooking"
    OrderDelivery  OrderStatus = "delivery"
    OrderCompleted OrderStatus = "completed"
    OrderCancelled OrderStatus = "cancelled"
)
```

**Note**: May need to add `ready_for_pickup` status between `cooking` and `delivery` for better workflow.

---

## 🎯 Priority Level

**Priority**: Medium-High
**Complexity**: Medium (requires understanding relationships between User, Order, Merchant, Driver)
**Dependencies**:

- Driver system implementation
- Role-based authentication
- Proper error handling for unauthorized access

---

## 🔍 Testing Scenarios

### Test Cases to Implement:

1. **User tries to cancel confirmed order** → Should fail
2. **Merchant tries to mark order as delivered** → Should fail
3. **Driver tries to update non-assigned order** → Should fail
4. **Valid status progression by correct roles** → Should succeed
5. **Invalid backward status transitions** → Should fail

---

**Next Steps**: Implement when driver system is complete and role-based authentication is needed for production use.
