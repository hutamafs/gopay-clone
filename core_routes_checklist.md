# ✅ Core Routes To-Do Checklist (gopay Clone)

## 1. 🛡️ Authentication

- [ ] `POST /auth/register` - Register a new user
- [ ] `POST /auth/login` - Authenticate and return JWT
- [ ] `GET /auth/profile` - Get current logged-in user's profile

---

## 2. 💰 Account Management

- [x] `POST /accounts` - Create a new account
- [x] `GET /users/:id/accounts` - List all accounts for a user
- [x] `GET /accounts/:account_id/balance` - Get account balance
- [x] `PUT /accounts/:account_id` - Update account (name, balance)

---

## 3. 🔁 Transfers & Transactions

- [x] `POST /transactions` - Transfer money between accounts
- [x] `GET /accounts/:id/transactions` - View account transaction history
- [x] `GET /transactions/:id` - Get single transaction detail

---

## 4. 📲 QR Code Payment

- [ ] `POST /qr/generate` - Generate a QR code to receive payment
- [ ] `POST /qr/scan` - Scan QR code and process payment

---

## 5. 🧑‍🤝‍🧑 Contacts

- [ ] `GET /contacts` - Get list of saved contacts
- [ ] `POST /contacts` - Add a new contact

---

## 6. 📊 Analytics (Optional/Bonus)

- [ ] `GET /analytics/spending` - Daily/monthly spending breakdown
- [ ] `GET /analytics/category` - Grouped by category

---

## 7. 🧪 Utility & Dev

- [ ] Add request/response validation
- [ ] Add JWT middleware to protected routes
- [ ] Connect handler ↔ service ↔ repository

---

_This checklist tracks only core backend functionality. Add-on features like notifications, real-time WebSockets, admin dashboards can follow later._
