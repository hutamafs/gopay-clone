# 🔐 JWT Implementation Status - GoPay Clone

## Current Route Analysis

### ✅ EXISTING ROUTES (Need JWT Protection)

#### User Routes (`/users`)

- [ ] `GET /users` - Get all users (**needs admin protection**)
- [ ] `GET /users/:id` - Get user by ID (**needs JWT + ownership check**)
- [ ] `PUT /users/:id` - Update user (**needs JWT + ownership check**)
- [ ] `DELETE /users/:id` - Delete user (**needs JWT + ownership check**)
- [ ] `GET /users/:user_id/accounts` - Get user accounts (**needs JWT + ownership check**)

#### Account Routes (`/accounts`)

- [ ] `POST /accounts` - Create account (**needs JWT**)
- [ ] `GET /accounts/:account_id/balance` - Get balance (**needs JWT + ownership check**)
- [ ] `PUT /accounts/:account_id` - Update account (**needs JWT + ownership check**)
- [ ] `GET /accounts/:account_id/detail` - Get account detail (**needs JWT + ownership check**)
- [ ] `GET /accounts/:account_id/transactions` - Get account transactions (**needs JWT + ownership check**)

#### Transaction Routes (`/transactions`)

- [ ] `POST /transactions` - Create transaction (**needs JWT**)
- [ ] `GET /transactions/:transaction_id` - Get transaction detail (**needs JWT + ownership check**)

#### QR Routes (`/qr`)

- [ ] `POST /qr` - Create QR code (**needs JWT**)
- [ ] `PUT /qr/:qr_id` - Scan QR code (**needs JWT**)

---

## ❌ MISSING ROUTES (Need to be Created)

### Authentication Routes (`/auth`) - **PRIORITY**

- [ ] `POST /auth/register` - Register new user + auto-create wallets
- [ ] `POST /auth/login` - Login user + return JWT token
- [ ] `GET /auth/profile` - Get current user profile (JWT protected)
- [ ] `POST /auth/refresh` - Refresh JWT token
- [ ] `POST /auth/logout` - Logout user

### Contact Routes (`/contacts`)

- [ ] `GET /contacts` - Get user's contacts (JWT protected)
- [ ] `POST /contacts` - Add new contact (JWT protected)
- [ ] `DELETE /contacts/:id` - Remove contact (JWT protected)

### Order/Food Routes (`/food` or `/orders`)

- [ ] `GET /food/merchants` - List restaurants (public)
- [ ] `GET /food/merchants/:id` - Get merchant details (public)
- [ ] `POST /food/orders` - Create food order (JWT protected)
- [ ] `GET /food/orders` - Get user's orders (JWT protected)
- [ ] `GET /food/orders/:id` - Get order details (JWT protected)
- [ ] `PUT /food/orders/:id/status` - Update order status (JWT protected + role check)

### Ride Routes (`/rides`)

- [ ] `POST /rides/request` - Request ride (JWT protected)
- [ ] `GET /rides` - Get user's rides (JWT protected)
- [ ] `GET /rides/:id` - Get ride details (JWT protected)
- [ ] `PUT /rides/:id/accept` - Accept ride (JWT protected + driver role)
- [ ] `PUT /rides/:id/status` - Update ride status (JWT protected)

### Driver Profile Routes (`/drivers`)

- [ ] `POST /drivers/profile` - Create driver profile (JWT protected)
- [ ] `GET /drivers/profile` - Get driver profile (JWT protected)
- [ ] `PUT /drivers/profile` - Update driver profile (JWT protected)
- [ ] `PUT /drivers/status` - Update status (JWT protected + driver role)

### Merchant Profile Routes (`/merchants`)

- [ ] `POST /merchants/profile` - Create merchant profile (JWT protected)
- [ ] `GET /merchants/profile` - Get merchant profile (JWT protected)
- [ ] `PUT /merchants/profile` - Update merchant profile (JWT protected)
- [ ] `POST /merchants/menu` - Add menu item (JWT protected + merchant role)

---

## 🛠️ Implementation Steps for JWT

### 1. Create JWT Utilities

```bash
# Files to create:
- utils/jwt.go           # JWT token generation/validation
- middleware/auth.go     # JWT middleware
```

### 2. Create Authentication System

```bash
# Files to create:
- handlers/auth_handler.go
- services/auth_service.go
- routes/auth_routes.go
- validator/auth_validator.go
```

### 3. Apply JWT Middleware to Existing Routes

```go
// Example: Update routes/user_routes.go
users.GET("/:id", middleware.JWTMiddleware(), userHandler.GetUserById)
users.PUT("/:id", middleware.JWTMiddleware(), userHandler.UpdateUser)
// etc...
```

### 4. Add User Context to Handlers

```go
// Update handlers to get current user from JWT context
func (h *UserHandler) UpdateUser(c echo.Context) error {
    currentUserID := c.Get("user_id").(uint)
    paramUserID := c.Param("id")

    // Check ownership
    if currentUserID != uint(paramUserID) {
        return utils.UnauthorizedResponse(c, "Cannot update other user's profile")
    }
    // ... rest of handler
}
```

---

## 🎯 Priority Implementation Order

### Phase 1: Essential Auth (Day 2 Morning)

1. ✅ JWT utilities and middleware
2. ✅ Auth routes (register, login, profile)
3. ✅ Protect existing user management routes

### Phase 2: Core Protection (Day 2 Afternoon)

1. ✅ Protect all account operations
2. ✅ Protect all transaction operations
3. ✅ Protect QR operations

### Phase 3: New Feature Routes (Day 2 Evening)

1. ✅ Contact management routes
2. ✅ Basic order routes preparation
3. ✅ Basic ride routes preparation

---

## 🔒 Security Considerations

### Ownership Checks Needed

- Users can only access their own accounts/transactions
- Users can only update their own profiles
- Drivers can only accept rides/update ride status
- Merchants can only manage their own menus/orders

### Role-Based Access

- Driver-only routes (accept rides, update driver status)
- Merchant-only routes (menu management, order status updates)
- Admin routes (if implementing admin features)

### Token Management

- JWT expiration handling
- Refresh token mechanism
- Logout token blacklisting (optional)

This gives you a clear roadmap for implementing JWT across your entire API!
