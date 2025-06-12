---
# 🏥 Golang Clinic Management System with JWT Authentication

A **role-based web backend** built using **Golang**, **Gin**, **PostgreSQL**, and **JWT Auth**, allowing **Receptionists** and **Doctors** to securely manage patient records via RESTful APIs. ✨

---

## 📌 Key Features

| Feature | Description |
|--------|-------------|
| 🔐 JWT Authentication | Secure login for receptionists & doctors with role-based access |
| 👩‍⚕️ Role-Based Access | Receptionist: Full CRUD, Doctor: Read + Update only |
| 📋 RESTful APIs | Built using Gin framework and tested on Postman |
| 🗂 PostgreSQL Integration | Efficient DB operations using `sqlx` |
| 🛡 Middleware Protection | All patient routes are guarded by JWT role checks |
| 🚀 Ready for Deployment | Modular structure, scalable for production |



## 📬 API Documentation (Postman)

🧪 View full working API collection with test cases here:

👉 **[View Postman Docs](https://documenter.getpostman.com/view/31280959/2sB2x6kC9r)**

---

## 👨‍💻 Tech Stack

| Tech          | Usage                |
|---------------|----------------------|
| Go (Golang)   | Core backend logic   |
| Gin           | HTTP routing & server|
| PostgreSQL    | Database layer       |
| SQLX          | Query handling       |
| JWT           | Token-based auth     |
| Postman       | API Testing          |

---

## 📂 Project Structure

```

golang-clinic-app/
├── config/             # DB config
├── controllers/        # Auth & patient logic
├── middleware/         # JWT auth middleware
├── models/             # Struct definitions
├── routes/             # Route registration
├── cmd/
│   └── main.go         # App entry point
└── go.mod

````

---

## 🚦 API Endpoints Summary

### 👤 Auth

| Method | Endpoint     | Description              |
|--------|--------------|--------------------------|
| POST   | /signup      | Register new user        |
| POST   | /login       | Login & get JWT token    |

### 🏥 Patients (Secured via JWT)

| Role        | Method | Endpoint          | Action            |
|-------------|--------|-------------------|-------------------|
| Receptionist| POST   | /api/patients     | Create Patient    |
| Both        | GET    | /api/patients     | View All Patients |
| Both        | PUT    | /api/patients/:id | Update Patient    |
| Receptionist| DELETE | /api/patients/:id | Delete Patient    |

---

## 🧪 Testing with Postman

- ✅ Add **Authorization: Bearer `<token>`** in headers
- ✅ Set **Content-Type: application/json**
- 🔐 Only authorized roles can access their endpoints
- ❌ Invalid or unauthorized requests return 403/401

---

## 🧠 Learning & Highlights

- Gained deep hands-on with **JWT-based auth system**
- Learned to implement **role-based middleware**
- Created clean Golang REST APIs using **Gin**
- Applied **PostgreSQL best practices** with SQLX
- Built a **fully testable backend**—ideal for portfolios & interviews

---

## 🚀 Getting Started

```bash
# 1. Clone the repo
git clone https://github.com/gouravpandey009/clinic-app-api.git
cd clinic-app

# 2. Setup PostgreSQL & .env
# Create database & user, then connect in config

# 3. Run the server
go run cmd/main.go
````

---

Great! Your documentation is already **outstanding**—clean, professional, and interview-ready. To make it **even more impressive** for recruiters or tech reviewers, you can now **add a new section on Render Deployment** to highlight real-world hosting + DevOps exposure.

---

## ☁️ Live Deployment (Render)

This project is fully deployed on **Render**, demonstrating production readiness.

| 🔗 **Live API Base URL** | `https://<your-app-name>.onrender.com` |
| ------------------------ | -------------------------------------- |

> 📌 Replace `<your-app-name>` with your actual Render subdomain URL.

---

### ⚙️ Deployment Steps (Render)

1. **Push code to GitHub**
2. **Create a new Web Service** on [Render](https://render.com)
3. **Set Environment Variables**

   * `DB_URL=postgresql://...&sslmode=require` *(use your PostgreSQL Render URL)*
4. **Build Command**

   ```bash
   go build -o app ./cmd
   ```
5. **Start Command**

   ```bash
   ./app
   ```
6. **Deploy** and access your REST API on the public URL provided by Render

---

## 🙋‍♂️ Author

👨‍💻 **Gourav Pandey**
🔗 [LinkedIn](https://www.linkedin.com/in/gouravpandey09/) 

---
