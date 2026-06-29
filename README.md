<div align="center">

# 🅿️ SpotSync

### Smart Parking Zone Reservation System

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Echo](https://img.shields.io/badge/Echo-v5-4A90D9?style=for-the-badge&logo=go&logoColor=white)](https://echo.labstack.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Neon-336791?style=for-the-badge&logo=postgresql&logoColor=white)](https://neon.tech/)
[![JWT](https://img.shields.io/badge/JWT-Auth-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)](https://jwt.io/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)

**A production-ready RESTful API for managing parking zones and reservations with role-based access control, built with clean architecture principles.**

[🌐 Live API](#-live-url) · [📖 API Docs](#-api-endpoints) · [🚀 Quick Start](#-quick-start) · [🏗️ Architecture](#%EF%B8%8F-architecture)

---

</div>

## 🌐 Live URL

```
https://spot-sync-golang-be8141cedb62.herokuapp.com/api/v1
```

> Hit the root URL to verify the API is live:
> ```bash
> curl https://spotsync.onrender.com/api/v1
> # → { "message": "Welcome to SportSync" }
> ```

---

## ✨ Features

| Feature | Description |
|---|---|
| 🔐 **Authentication** | Secure JWT-based registration & login with bcrypt password hashing |
| 🛡️ **Role-Based Access** | `admin` and `driver` roles with middleware-enforced permissions |
| 🅿️ **Zone Management** | Full CRUD for parking zones (General, EV Charging, Covered) |
| 📋 **Reservations** | Create, view, and cancel parking reservations with capacity enforcement |
| 📊 **Availability Tracking** | Real-time available spot calculation per zone |
| ✅ **Input Validation** | Request validation using `go-playground/validator` |
| 🗃️ **Auto-Migration** | Automatic database schema management via GORM |
| 🩺 **Health Check** | Dedicated `/health` endpoint for monitoring |

---

## 🛠️ Tech Stack

```mermaid
graph LR
    A["🌐 Client"] -->|HTTP/JSON| B["🚀 Echo v5"]
    B -->|Middleware| C["🔐 JWT Auth"]
    B -->|Routes| D["📦 Handlers"]
    D -->|DTOs| E["⚙️ Services"]
    E -->|Models| F["🗄️ Repositories"]
    F -->|GORM| G["🐘 PostgreSQL"]

    style A fill:#1a1a2e,stroke:#e94560,color:#fff
    style B fill:#16213e,stroke:#0f3460,color:#fff
    style C fill:#533483,stroke:#0f3460,color:#fff
    style D fill:#0f3460,stroke:#e94560,color:#fff
    style E fill:#16213e,stroke:#533483,color:#fff
    style F fill:#1a1a2e,stroke:#0f3460,color:#fff
    style G fill:#336791,stroke:#fff,color:#fff
```

| Layer | Technology | Purpose |
|---|---|---|
| **Language** | Go 1.26 | Core runtime |
| **Framework** | Echo v5 | High-performance HTTP routing |
| **Database** | PostgreSQL (Neon) | Cloud-native serverless Postgres |
| **ORM** | GORM v1.31 | Object-relational mapping & migrations |
| **Auth** | JWT (golang-jwt/v5) | Stateless token authentication |
| **Hashing** | bcrypt (x/crypto) | Secure password hashing |
| **Validation** | go-playground/validator v10 | Struct-level request validation |
| **Config** | godotenv | Environment variable management |
| **Deployment** | Render | Cloud hosting with Procfile |

---

## 🏗️ Architecture

SpotSync follows **Clean Architecture** with clear separation of concerns across well-defined layers:

```mermaid
graph TB
    subgraph CLIENT["🌐 Client Layer"]
        REQ["HTTP Request"]
    end

    subgraph MIDDLEWARE["🛡️ Middleware Layer"]
        JWT["JWT Authentication"]
        ROLE["Role Authorization"]
        VAL["Request Validator"]
    end

    subgraph HANDLER["📦 Handler Layer"]
        AH["AuthHandler"]
        ZH["ZoneHandler"]
        RH["ReservationHandler"]
    end

    subgraph SERVICE["⚙️ Service Layer — Business Logic"]
        AS["AuthService"]
        ZS["ZoneService"]
        RS["ReservationService"]
    end

    subgraph REPOSITORY["🗄️ Repository Layer — Data Access"]
        UR["UserRepository"]
        ZR["ZoneRepository"]
        RR["ReservationRepository"]
    end

    subgraph DATABASE["🐘 Database"]
        PG["PostgreSQL"]
    end

    REQ --> JWT --> ROLE --> VAL
    VAL --> AH & ZH & RH
    AH --> AS
    ZH --> ZS
    RH --> RS
    AS --> UR
    ZS --> ZR
    RS --> RR & ZR
    UR & ZR & RR --> PG

    style CLIENT fill:#0d1117,stroke:#58a6ff,color:#fff
    style MIDDLEWARE fill:#161b22,stroke:#f78166,color:#fff
    style HANDLER fill:#161b22,stroke:#3fb950,color:#fff
    style SERVICE fill:#161b22,stroke:#d2a8ff,color:#fff
    style REPOSITORY fill:#161b22,stroke:#79c0ff,color:#fff
    style DATABASE fill:#336791,stroke:#fff,color:#fff
```

### 📁 Project Structure

```
SpotSync/
├── cmd/
│   └── main.go               # Application entry point
├── config/
│   ├── config.go              # Environment config loader
│   └── db.go                  # Database connection & migration
├── dto/
│   ├── auth_dto.go            # Auth request/response DTOs
│   ├── zone_dto.go            # Zone request/response DTOs
│   ├── reservation_dto.go     # Reservation request/response DTOs
│   └── response.go            # Generic API response wrapper
├── handler/
│   ├── auth_handler.go        # Auth HTTP handlers
│   ├── zone_handler.go        # Zone HTTP handlers
│   └── reservation_handler.go # Reservation HTTP handlers
├── middleware/
│   └── jwt_middleware.go      # JWT auth & role-based access
├── models/
│   ├── user.go                # User GORM model
│   ├── parking_zone.go        # ParkingZone GORM model
│   └── reservation.go         # Reservation GORM model
├── repository/
│   ├── user_repo.go           # User data access
│   ├── zone_repo.go           # Zone data access
│   ├── reservation_repo.go    # Reservation data access
│   └── errors.go              # Repository error definitions
├── routes/
│   └── routes.go              # Route registration
├── service/
│   ├── auth_service.go        # Auth business logic
│   ├── zone_service.go        # Zone business logic
│   └── reservation_service.go # Reservation business logic
├── utils/
│   ├── jwt.go                 # JWT token generation & validation
│   ├── password.go            # bcrypt hash utilities
│   └── errors.go              # Application-level errors
├── .env                       # Environment variables (not committed)
├── .gitignore
├── go.mod
├── go.sum
└── Procfile                   # Deployment config
```

---

## 🗄️ Database Schema

```mermaid
erDiagram
    USERS {
        uint id PK
        string name
        string email UK
        string password
        string role "driver | admin"
        timestamp created_at
        timestamp updated_at
    }

    PARKING_ZONES {
        uint id PK
        string name
        string type "general | ev_charging | covered"
        int total_capacity
        decimal price_per_hour
        timestamp created_at
        timestamp updated_at
    }

    RESERVATIONS {
        uint id PK
        uint user_id FK
        uint zone_id FK
        string license_plate
        string status "active | cancelled"
        timestamp created_at
        timestamp updated_at
    }

    USERS ||--o{ RESERVATIONS : "makes"
    PARKING_ZONES ||--o{ RESERVATIONS : "has"
```

---

## 📡 API Endpoints

Base URL: `/api/v1`

### 🔓 Public Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1` | Welcome message |
| `GET` | `/health` | Health check |

### 🔐 Authentication

| Method | Endpoint | Description | Body |
|---|---|---|---|
| `POST` | `/api/v1/auth/register` | Register a new user | `{ name, email, password, role? }` |
| `POST` | `/api/v1/auth/login` | Login & get JWT token | `{ email, password }` |

### 🅿️ Parking Zones

| Method | Endpoint | Auth | Role | Description |
|---|---|---|---|---|
| `GET` | `/api/v1/zones` | ❌ | — | List all parking zones |
| `GET` | `/api/v1/zones/:id` | ❌ | — | Get zone by ID |
| `POST` | `/api/v1/zones` | ✅ | `admin` | Create a new zone |
| `PUT` | `/api/v1/zones/:id` | ✅ | `admin` | Update a zone |
| `DELETE` | `/api/v1/zones/:id` | ✅ | `admin` | Delete a zone |

### 📋 Reservations

| Method | Endpoint | Auth | Role | Description |
|---|---|---|---|---|
| `POST` | `/api/v1/reservations` | ✅ | any | Create a reservation |
| `GET` | `/api/v1/reservations/my-reservations` | ✅ | any | Get my reservations |
| `DELETE` | `/api/v1/reservations/:id` | ✅ | any | Cancel a reservation |
| `GET` | `/api/v1/reservations` | ✅ | `admin` | Get all reservations |

### 📬 Request & Response Examples

<details>
<summary><strong>POST</strong> <code>/api/v1/auth/register</code></summary>

**Request:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123",
  "role": "driver"
}
```

**Response** `201 Created`:
```json
{
  "status": "success",
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "driver",
    "created_at": "2026-06-29T12:00:00Z",
    "updated_at": "2026-06-29T12:00:00Z"
  }
}
```
</details>

<details>
<summary><strong>POST</strong> <code>/api/v1/auth/login</code></summary>

**Request:**
```json
{
  "email": "john@example.com",
  "password": "secret123"
}
```

**Response** `200 OK`:
```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "driver"
    }
  }
}
```
</details>

<details>
<summary><strong>POST</strong> <code>/api/v1/zones</code> — Admin Only</summary>

**Headers:** `Authorization: Bearer <token>`

**Request:**
```json
{
  "name": "Zone A - Ground Floor",
  "type": "general",
  "total_capacity": 50,
  "price_per_hour": 5.00
}
```

**Response** `201 Created`:
```json
{
  "status": "success",
  "message": "Parking zone created successfully",
  "data": {
    "id": 1,
    "name": "Zone A - Ground Floor",
    "type": "general",
    "total_capacity": 50,
    "available_spots": 50,
    "price_per_hour": 5.00,
    "created_at": "2026-06-29T12:00:00Z"
  }
}
```
</details>

<details>
<summary><strong>POST</strong> <code>/api/v1/reservations</code></summary>

**Headers:** `Authorization: Bearer <token>`

**Request:**
```json
{
  "zone_id": 1,
  "license_plate": "ABC-1234"
}
```

**Response** `201 Created`:
```json
{
  "status": "success",
  "message": "Reservation confirmed successfully",
  "data": {
    "id": 1,
    "license_plate": "ABC-1234",
    "status": "active",
    "zone": {
      "id": 1,
      "name": "Zone A - Ground Floor",
      "type": "general"
    },
    "created_at": "2026-06-29T12:00:00Z"
  }
}
```
</details>

---

## 🔑 Authentication Flow

```mermaid
sequenceDiagram
    participant C as 🌐 Client
    participant A as 🚀 API Server
    participant DB as 🐘 PostgreSQL

    Note over C,DB: Registration Flow
    C->>A: POST /api/v1/auth/register
    A->>A: Validate & hash password (bcrypt)
    A->>DB: Insert user record
    DB-->>A: User created
    A-->>C: 201 — User data

    Note over C,DB: Login Flow
    C->>A: POST /api/v1/auth/login
    A->>DB: Find user by email
    DB-->>A: User record
    A->>A: Verify password & generate JWT
    A-->>C: 200 — JWT token + user data

    Note over C,DB: Protected Request
    C->>A: GET /api/v1/reservations/my-reservations
    Note right of C: Authorization: Bearer <token>
    A->>A: Validate JWT & extract claims
    A->>A: Check role permissions
    A->>DB: Query reservations
    DB-->>A: Reservation list
    A-->>C: 200 — Reservations data
```

---

## 🚀 Quick Start

### Prerequisites

- [Go 1.26+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/) (or a [Neon](https://neon.tech/) cloud instance)
- [Git](https://git-scm.com/)

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/SkillexSJ/SpotSync.git
cd SpotSync
```

### 2️⃣ Configure Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgresql://user:password@host:5432/dbname?sslmode=require
JWT_SECRET=your-super-secret-jwt-key
PORT=8080
```

| Variable | Required | Description | Default |
|---|---|---|---|
| `DATABASE_URL` | ✅ | PostgreSQL connection string | — |
| `JWT_SECRET` | ✅ | Secret key for signing JWT tokens | — |
| `PORT` | ❌ | Server port | `8080` |

### 3️⃣ Install Dependencies

```bash
go mod download
```

### 4️⃣ Run the Server

```bash
go run cmd/main.go
```

You should see:

```
✅ Connected to database successfully!
✅ Database migration completed!
✅ Database ping successful!
──────────────────────────────────────────
🚀 SpotSync server running on port 8080
──────────────────────────────────────────
```

### 5️⃣ Verify

```bash
curl http://localhost:8080/api/v1
# → { "message": "Welcome to SportSync" }

curl http://localhost:8080/health
# → { "status": "ok", "message": "SpotSync API is running" }
```

---

## 🧪 Test with cURL

```bash
# Register an admin
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Admin","email":"admin@test.com","password":"admin123","role":"admin"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"admin123"}'

# Create a zone (use token from login response)
curl -X POST http://localhost:8080/api/v1/zones \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <YOUR_TOKEN>" \
  -d '{"name":"Zone A","type":"general","total_capacity":50,"price_per_hour":5.00}'

# List all zones (public)
curl http://localhost:8080/api/v1/zones
```

---

## 🚢 Deployment

SpotSync is configured for **Render** deployment with the included `Procfile`:

```
web: cmd
```

Simply connect your GitHub repo to [Render](https://render.com/) and set the environment variables in the dashboard.

---

## 👨‍💻 Author

**SkillexSJ**

- GitHub: [@SkillexSJ](https://github.com/SkillexSJ)

---

<div align="center">

**Built with ❤️ in Go**

</div>
