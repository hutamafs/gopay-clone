# 📋 Day 2 Todo List - GoPay Clone Implementation

## 🔐 JWT Authentication Routes to Implement

### Missing Authentication Routes

- [x] `POST /auth/register` - User registration with auto-wallet creation
- [x] `POST /auth/login` - User login with JWT token response
- [x] `GET /auth/profile` - Get current user profile (JWT protected)
- [ ] `POST /auth/refresh` - Refresh JWT token
- [ ] `POST /auth/logout` - Logout user (blacklist token)

### JWT Middleware Implementation

- [ ] Create JWT middleware function in `middleware/auth.go`
- [ ] Add JWT token generation utility in `utils/jwt.go`
- [ ] Add JWT token validation and parsing
- [ ] Apply JWT middleware to protected routes

### Routes That Need JWT Protection

- [ ] `PUT /users/:id` - Update User (only own profile)
- [ ] `DELETE /users/:id` - Delete User (only own profile)
- [ ] `GET /users/:user_id/accounts` - Get User Accounts (only own accounts)
- [ ] All `/accounts/*` routes - Account operations
- [ ] All `/transactions/*` routes - Transaction operations
- [ ] All `/qr/*` routes - QR operations
- [ ] All future `/orders/*` and `/rides/*` routes

---

## 🍕 New Model Routes to Implement

### Food Service Routes (`/food`)

- [ ] `GET /food/merchants` - List all restaurants/merchants
- [ ] `GET /food/merchants/:id` - Get merchant details with menu
- [ ] `GET /food/merchants/:id/menu` - Get merchant menu items
- [ ] `POST /food/orders` - Create food order
- [ ] `GET /food/orders` - Get user's food orders
- [ ] `GET /food/orders/:id` - Get specific order details
- [ ] `PUT /food/orders/:id/status` - Update order status (merchant/driver only)

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

### Merchant Profile Routes (`/merchants`)

- [ ] `POST /merchants/profile` - Create merchant profile
- [ ] `GET /merchants/profile` - Get current merchant profile
- [ ] `PUT /merchants/profile` - Update merchant profile
- [ ] `POST /merchants/menu` - Add menu item
- [ ] `PUT /merchants/menu/:id` - Update menu item
- [ ] `DELETE /merchants/menu/:id` - Delete menu item

### Contact Routes (`/contacts`)

- [ ] `GET /contacts` - Get user's contacts
- [ ] `POST /contacts` - Add new contact
- [ ] `DELETE /contacts/:id` - Remove contact

---

## 🏗️ Backend Services & Handlers to Create

### New Service Files Needed

- [ ] `services/order_service.go` - Food order management
- [ ] `services/ride_service.go` - Ride booking management
- [ ] `services/driver_service.go` - Driver profile management
- [ ] `services/merchant_service.go` - Merchant profile management
- [ ] `services/menu_service.go` - Menu item management
- [ ] `services/contact_service.go` - Contact management
- [ ] `services/auth_service.go` - Authentication logic

### New Handler Files Needed

- [ ] `handlers/auth_handler.go` - Authentication endpoints
- [ ] `handlers/order_handler.go` - Food order endpoints
- [ ] `handlers/ride_handler.go` - Ride booking endpoints
- [ ] `handlers/driver_handler.go` - Driver profile endpoints
- [ ] `handlers/merchant_handler.go` - Merchant profile endpoints
- [ ] `handlers/contact_handler.go` - Contact endpoints

### New Route Files Needed

- [ ] `routes/auth_routes.go` - Authentication routes
- [ ] `routes/order_routes.go` - Food order routes
- [ ] `routes/ride_routes.go` - Ride booking routes
- [ ] `routes/driver_routes.go` - Driver profile routes
- [ ] `routes/merchant_routes.go` - Merchant profile routes
- [ ] `routes/contact_routes.go` - Contact routes

---

## 🗃️ Database & Migration Updates

### Missing Model Migrations

- [ ] Add `models.DriverProfile` to AutoMigrate in main.go
- [ ] Add `models.MerchantProfile` to AutoMigrate in main.go
- [ ] Add `models.MenuItem` to AutoMigrate in main.go
- [ ] Add `models.Order` to AutoMigrate in main.go
- [ ] Add `models.OrderItem` to AutoMigrate in main.go
- [ ] Add `models.Ride` to AutoMigrate in main.go

### Account System Enhancement

- [ ] Create default accounts on user registration (main_balance + points)
- [ ] Update account service to handle wallet types properly
- [ ] Add wallet type validation in account operations

---

## 🔧 Refactoring & Improvements

### Code Structure Improvements

- [ ] Add proper error handling constants in `utils/errors.go`
- [ ] Create response DTOs for complex objects (hide sensitive fields)
- [ ] Add request/response logging middleware
- [ ] Add rate limiting middleware for API endpoints

### Validation Enhancements

- [ ] Create validators for new models in `validator/` directory:
  - [ ] `order_validator.go`
  - [ ] `ride_validator.go`
  - [ ] `driver_validator.go`
  - [ ] `merchant_validator.go`
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
