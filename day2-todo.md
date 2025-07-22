# 📋 Day 2 Todo List - GoPay Clone Implementation

## 🔐 JWT Authentication Routes to Implement

### Missing Authentication Routes

- [x] `POST /public/users` - User registration with auto-wallet creation (in user service)
- [x] `POST /public/users/login` - User login with JWT token response
- [x] `GET /users/:id` - Get current user profile (JWT protected) - **serves as auth/profile**
- [ ] `POST /auth/refresh` - Refresh JWT token
- [ ] `POST /auth/logout` - Logout user (blacklist token)

### JWT Middleware Implementation

- [x] Add JWT token generation utility in `utils/jwt.go`
- [x] Add JWT token validation and parsing via echo-jwt
- [x] Apply JWT middleware to protected routes
- [x] JWT ownership checks in handlers (via `utils.CLaimJwt()`)

### Routes That Need JWT Protection

- [x] `PUT /users/:id` - Update User (only own profile)
- [x] `GET /users/:id` - Get User Profile (only own profile) - **serves as auth/profile**
- [x] `GET /users/:user_id/accounts` - Get User Accounts (only own accounts)
- [x] All `/accounts/*` routes - Account operations
- [x] All `/transactions/*` routes - Transaction operations
- [x] All `/qr/*` routes - QR operations
- [ ] All future `/orders/*` and `/rides/*` routes

---

## 🍕 New Model Routes to Implement

### Merchant Profile Routes (`/merchants`)

- [x] `POST /merchants/profile` - Create merchant profile
- [x] `PUT /merchants/profile` - Update merchant profile
- [x] `GET /merchants` - List all merchants/restaurants
- [x] `GET /merchants/:id` - Get specific merchant profile details

### Menu Item Routes (Under Merchants) (`/merchants/:merchant_id/menuitems`)

- [x] `GET /merchants/:merchant_id/menuitems` - Get all menu items for merchant
- [x] `GET /merchants/:merchant_id/menuitems/:id` - Get specific menu item
- [x] `POST /merchants/:merchant_id/menuitems` - Create menu item (merchant owner only)
- [x] `PUT /merchants/:merchant_id/menuitems/:id` - Update menu item (merchant owner only)
- [x] `DELETE /merchants/:merchant_id/menuitems/:id` - Delete menu item (merchant owner only)

### Browse Menu Routes (`/menuitems`)

- [ ] `GET /menuitems` - Browse all menu items (with merchant filter, category filter) ( search by menu item, but returns the restaurant that has that food)

### Order Routes (`/orders`)

- [ ] `POST /orders` - Create new order (contains multiple menu items from ONE merchant)
- [ ] `GET /orders` - Get user's orders
- [ ] `GET /orders/:id` - Get specific order details
- [ ] `PUT /orders/:id/status` - Update order status (merchant/driver only)
- [ ] `PUT /orders/:id/assign-driver` - Assign driver to order (system only)

### Ride Service Routes (`/ride`)

- [ ] `POST /ride/request` - Request a ride
- [ ] `GET /ride/requests` - Get user's ride requests
- [ ] `GET /ride/requests/:id` - Get specific ride details
- [ ] `PUT /ride/requests/:id/accept` - Accept ride (driver only)
- [ ] `PUT /ride/requests/:id/status` - Update ride status
- [ ] `GET /ride/drivers/nearby` - Find nearby available drivers

### Driver Profile Routes (`/drivers`)

- [ ] `POST /drivers/profile` - Create driver profile
- [ ] `GET /drivers/profile` - Get current driver profile
- [ ] `PUT /drivers/profile` - Update driver profile
- [ ] `PUT /drivers/status` - Update driver status (online/offline)
- [ ] `PUT /drivers/location` - Update driver current location

### Contact Routes (`/contacts`)

- [ ] `GET /contacts` - Get user's contacts
- [ ] `POST /contacts` - Add new contact
- [ ] `DELETE /contacts/:id` - Remove contact

---

## 🏗️ Backend Services & Handlers to Create

### New Service Files Needed

- [x] `services/merchant_service.go` - Merchant profile management
- [x] `services/menu_service.go` - Menu item management (under merchants)
- [ ] `services/order_service.go` - Food order management
- [ ] `services/ride_service.go` - Ride booking management
- [ ] `services/driver_service.go` - Driver profile management
- [ ] `services/contact_service.go` - Contact management
- [ ] `services/auth_service.go` - Authentication logic

### New Handler Files Needed

- [x] `handlers/merchant_handler.go` - Merchant profile endpoints
- [x] `handlers/menu_handler.go` - Menu item endpoints (under merchants)
- [ ] `handlers/order_handler.go` - Food order endpoints
- [ ] `handlers/auth_handler.go` - Authentication endpoints
- [ ] `handlers/ride_handler.go` - Ride booking endpoints
- [ ] `handlers/driver_handler.go` - Driver profile endpoints
- [ ] `handlers/contact_handler.go` - Contact endpoints

### New Route Files Needed

- [x] `routes/merchant_routes.go` - Merchant profile routes + menu routes
- [ ] `routes/order_routes.go` - Food order routes
- [ ] `routes/auth_routes.go` - Authentication routes
- [ ] `routes/ride_routes.go` - Ride booking routes
- [ ] `routes/driver_routes.go` - Driver profile routes
- [ ] `routes/contact_routes.go` - Contact routes

---

## 🗃️ Database & Migration Updates

### Missing Model Migrations

- [x] Add `models.DriverProfile` to AutoMigrate in main.go
- [x] Add `models.MerchantProfile` to AutoMigrate in main.go
- [x] Add `models.MenuItem` to AutoMigrate in main.go
- [x] Add `models.Order` to AutoMigrate in main.go
- [x] Add `models.OrderItem` to AutoMigrate in main.go
- [x] Add `models.Ride` to AutoMigrate in main.go

### Account System Enhancement

- [x] Create default accounts on user registration (main_balance + points) - **in user service**
- [x] Update account service to handle wallet types properly (AccountType enum)
- [x] Add wallet type validation in account operations

---

## 🔧 Refactoring & Improvements

### Code Structure Improvements

- [ ] Add proper error handling constants in `utils/errors.go`
- [ ] Create response DTOs for complex objects (hide sensitive fields)
- [ ] Add request/response logging middleware
- [ ] Add rate limiting middleware for API endpoints

### Validation Enhancements

- [ ] Create validators for new models in `validator/` directory:
  - [x] `merchant_validator.go`
  - [x] `menu_validator.go`
  - [ ] `order_validator.go`
  - [ ] `ride_validator.go`
  - [ ] `driver_validator.go`
  - [ ] `auth_validator.go`

### Transaction System Enhancement

- [ ] Update transaction service to handle service-specific transactions
- [ ] Add merchant info to transaction model/response
- [ ] Implement proper balance checking before transactions
- [ ] Add transaction rollback mechanism for failed operations

---

## 🧪 Testing & Documentation

### API Testing

- [ ] Create Postman collection for all endpoints
- [ ] Test authentication flow end-to-end
- [ ] Test order creation and payment flow
- [ ] Test ride booking and payment flow

### Documentation

- [ ] Update API documentation with new endpoints
- [ ] Create database schema diagram
- [ ] Document authentication flow
- [ ] Create deployment guide

---

## 🎯 Priority Order for Implementation

### Phase 1 (Essential for Demo)

1. ✅ JWT Authentication system
2. ✅ Contact management
3. ✅ Food ordering system (Order, MenuItem, Merchant)
4. ✅ Ride booking system (Ride, Driver)

### Phase 2 (Enhanced Features)

1. ✅ Multi-wallet account enhancements
2. ✅ Transaction analytics and history
3. ✅ Driver/Merchant profile management
4. ✅ Advanced error handling and logging

### Phase 3 (Polish)

1. ✅ Rate limiting and security
2. ✅ API documentation
3. ✅ Testing suite
4. ✅ Deployment preparation

---

**Target**: Complete Phase 1 by end of Day 2 to have a working demo for the internship application deadline (Day 7).
