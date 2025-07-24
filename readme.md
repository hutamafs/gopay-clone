# 🚀 GoPay Super App Clone

A comprehensive multi-service super app backend built with Go, featuring wallet management, food delivery, and driver services. This project demonstrates enterprise-level API architecture with real-world business logic.

## 🌟 **Features**

### 💰 **Multi-Wallet System**

- Multiple account types per user (main_balance, points, savings)
- Real-time balance management
- Secure transaction processing
- Complete audit trails

### 🍕 **GoFood Service (Food Delivery)**

- Restaurant/merchant management
- Menu items with categories and availability
- Complete order lifecycle management
- Real-time status tracking
- Driver assignment and management
- Role-based order status updates

### 🚗 **Driver Management**

- Complete CRUD operations
- Vehicle type management (motorcycle, car)
- Real-time availability tracking
- Automatic order assignment

### 🔐 **Security & Authentication**

- JWT-based authentication
- Role-based access control (Customer, Merchant, Driver)
- Input validation and sanitization
- Secure password handling

### 💳 **Transaction System**

- Atomic money transfers
- Service-specific transaction tracking (food, ride)
- Complete financial audit trails
- Transaction status management

## 🏗️ **Architecture**

### **Project Structure**

```
gopay-clone/
├── config/          # Database configuration
├── handlers/        # HTTP request handlers
├── migrations/      # Database migrations
├── models/         # Data models and structs
├── routes/         # API route definitions
├── services/       # Business logic layer
├── utils/          # Helper functions and utilities
├── validator/      # Input validation logic
└── main.go         # Application entry point
```

### **Tech Stack**

- **Backend**: Go 1.19+
- **Framework**: Echo v4
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT
- **Validation**: Custom validators
- **Deployment**: Render/Railway/Heroku ready

## 🚀 **Quick Start**

### **Prerequisites**

```bash
- Go 1.19 or higher
- PostgreSQL 12+
- Git
```

### **Installation**

```bash
# Clone repository
git clone https://github.com/yourusername/gopay-clone.git
cd gopay-clone

# Install dependencies
go mod tidy

# Set up environment variables
cp .env.example .env
# Edit .env with your database credentials and JWT secret

# Run the application
go run main.go
```

### **Environment Variables**

```env
DATABASE_URL=postgresql://username:password@localhost:5432/gopay_db
JWT_SECRET=your-super-secure-jwt-secret-at-least-32-characters
PORT=8080
APP_ENV=development
```

### **Database Setup**

```bash
# Create PostgreSQL database
createdb gopay_db

# Migrations run automatically on startup
go run main.go
```

## 📡 **API Documentation**

### **Base URL**

```
Local: http://localhost:8080/api/v1
Production: https://your-app.render.com/api/v1
```

### **Authentication**

All protected endpoints require JWT token in header:

```
Authorization: Bearer <your-jwt-token>
```

### **API Endpoints**

#### **🔐 Authentication**

```http
POST /api/v1/public/register     # Register new user
POST /api/v1/public/login        # User login
```

#### **👤 User Management**

```http
GET    /api/v1/users/profile     # Get user profile
PUT    /api/v1/users/profile     # Update user profile
GET    /api/v1/users/:id/orders  # Get user's orders
```

#### **🏪 Merchant Management**

```http
POST   /api/v1/public/merchants                         # Register merchant
GET    /api/v1/merchants                                # List all merchants
GET    /api/v1/merchants/:id                            # Get merchant details
PUT    /api/v1/merchants/:id                            # Update merchant profile
GET    /api/v1/merchants/:merchant_id/menu-item         # Get merchant's menu
POST   /api/v1/merchants/:merchant_id/menu-item         # Add menu item
PUT    /api/v1/merchants/:merchant_id/menu-item/:id     # Update menu item
DELETE /api/v1/merchants/:merchant_id/menu-items/:id    # Delete menu item
GET    /api/v1/menus/menu-items                         # Get all menu item
```

#### **💰 Account & Wallet**

```http
POST   /api/v1/accounts                                      # Create new account
GET    /api/v1/:user_id/accounts                             # Get user accounts
GET    /api/v1/accounts/:account_id/balance                  # Get account balance
GET    /api/v1/accounts/:account_id/detail                   # Get account detail
GET    /api/v1/accounts/:account_id/transactions             # Get account transaction history
PUT    /api/v1/accounts/:account_id                          # Update account detail
```

#### **💳 Transactions**

```http
GET    /api/v1/transactions               # Get user transactions
POST   /api/v1/transactions               # Create transaction
GET    /api/v1/transactions/:id           # Get transaction details
```

#### **📦 Orders (GoFood)**

```http
POST   /api/v1/public/orders              # Create new order
GET    /api/v1/orders/:id                 # Get order details
PUT    /api/v1/orders/:id/status          # Update order status
```

#### **🚗 Driver Management**

```http
GET    /api/v1/public/drivers             # List all drivers
POST   /api/v1/public/drivers             # Register driver
GET    /api/v1/drivers/available          # Get available drivers
GET    /api/v1/drivers/:id                # Get driver details
PUT    /api/v1/drivers/profile            # Update driver profile
PUT    /api/v1/drivers/status             # Update driver status
PUT    /api/v1/drivers/location           # Update driver location
DELETE /api/v1/drivers/profile            # Delete driver profile
```

## 🔄 **Business Flows**

### **Order Flow**

1. **Customer places order** → Validates menu items & calculates total
2. **Balance check** → Ensures sufficient wallet balance
3. **Driver assignment** → Finds and assigns available driver
4. **Payment processing** → Deducts from customer, credits merchant
5. **Order creation** → Creates order with all items and relationships
6. **Status tracking** → Real-time updates throughout delivery

### **Status Flow**

```
pending → confirmed → cooking → ready → delivery → completed
    ↓
cancelled (customer can cancel if pending)
```

### **Role-Based Status Updates**

- **Customer**: Can only cancel pending orders
- **Merchant**: pending → confirmed → cooking → ready
- **Driver**: ready → delivery → completed

## 🧪 **Testing**

### **Sample API Calls**

#### **Register User**

```bash
curl -X POST http://localhost:8080/api/v1/public/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "phone": "1234567890"
  }'
```

#### **Create Order**

```bash
curl -X POST http://localhost:8080/api/v1/public/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "merchant_id": 1,
    "delivery_address": "123 Main St, City",
    "order_items": [
      {
        "menu_item_id": 1,
        "quantity": 2,
        "notes": "Extra spicy"
      }
    ]
  }'
```

#### **Update Order Status**

```bash
curl -X PUT http://localhost:8080/api/v1/orders/1/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <merchant-jwt-token>" \
  -d '{
    "status": "confirmed"
  }'
```

## 🚀 **Deployment**

### **Render (Recommended)**

1. Connect GitHub repository to Render
2. Create Web Service with:
   - **Build Command**: `go build -o main .`
   - **Start Command**: `./main`
3. Add PostgreSQL database
4. Set environment variables

### **Docker**

```dockerfile
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

## 🗃️ **Database Schema**

### **Key Models**

- **User**: Authentication and profile data
- **Account**: Multi-wallet system (main_balance, points, etc.)
- **Merchant**: Restaurant/store information
- **MenuItem**: Menu items with pricing and availability
- **Order**: Order details with items and status
- **Transaction**: Financial records with audit trails
- **DriverProfile**: Driver information and vehicle details

### **Relationships**

- User → Multiple Accounts (1:n)
- User → Multiple Orders (1:n)
- Merchant → Multiple MenuItems (1:n)
- Order → Multiple OrderItems (1:n)
- Transaction → Sender/Receiver Accounts (n:1)

## 🔒 **Security Features**

- **JWT Authentication** with secure token handling
- **Password Hashing** using bcrypt
- **Input Validation** on all endpoints
- **SQL Injection Prevention** through GORM
- **Role-Based Access Control** for different user types
- **Transaction Atomicity** to prevent financial inconsistencies

## 🎯 **Key Achievements**

✅ **Complete Food Delivery System** - End-to-end order management
✅ **Real Money Transfers** - Actual wallet balance changes
✅ **Multi-Role Authorization** - Customer/Merchant/Driver workflows
✅ **Production-Ready Code** - Error handling, validation, security
✅ **Scalable Architecture** - Clean separation of concerns
✅ **Database Transactions** - ACID compliance for financial operations

## 🚀 **Future Enhancements**

- [ ] QR Payment System
- [ ] Real-time notifications
- [ ] GoRide service (ride-hailing)
- [ ] Admin dashboard
- [ ] WebSocket integration for live updates
- [ ] Advanced driver selection algorithms
- [ ] Mobile app integration

## 🤝 **Contributing**

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 **Acknowledgments**

- Echo framework for excellent HTTP routing
- GORM for powerful ORM capabilities
- PostgreSQL for robust data storage
- JWT for secure authentication

---

**Built with ❤️ and Go** 🚀

_This project demonstrates backend development with real-world business logic, making it perfect for portfolio showcase and production deployment._
