# 🎟️ Indico Backend - Voucher Management System

REST API Backend for voucher management system built with Go, Gin, GORM, and PostgreSQL. This project is part of **Take Home Test Indico By Telkomsel**.

---

## 📋 Project Description

Backend application for managing vouchers with complete CRUD features, JWT authentication, CSV upload/export, pagination, search, and sorting. All delete operations use **soft delete** to maintain data integrity.

---

## 🚀 Technologies Used

- **Go 1.24+** - Programming Language
- **Gin** - Web Framework
- **GORM** - ORM for Database
- **PostgreSQL** - Database
- **JWT (golang-jwt/jwt/v5)** - Authentication & Authorization
- **bcrypt (golang.org/x/crypto)** - Password Hashing
- **go-playground/validator/v10** - Request Validation
- **godotenv** - Environment Variable Management
- **Air** - Live Reload for Development

---

## ✨ Main Features

### 1. 🔐 Authentication API (Dummy)

- **POST** `/login` - Login with dummy credentials (accepts any username/password)
- JWT Token with **5 minutes** expiration
- Middleware for token validation in request headers

### 2. 📝 Voucher CRUD API

- **GET** `/vouchers` - List vouchers with pagination, search, sorting
- **GET** `/vouchers/get-by-id/:id` - Get voucher by ID
- **POST** `/vouchers` - Create new voucher
- **PUT** `/vouchers/:id` - Update voucher (partial update)
- **DELETE** `/vouchers/:id` - Soft delete voucher

### 3. 📊 Advanced Features

- **Pagination** - Support page & page_size
- **Search** - Search by code, name, description
- **Sorting** - Sort by id, code, name, discount, created_at (asc/desc)
- **Filter** - Filter by is_active status

### 4. 📁 CSV Operations

- **POST** `/vouchers/upload-csv` - Bulk upload vouchers from CSV
- **GET** `/vouchers/export` - Export all vouchers to CSV

### 5. 🕒 Readable Time Format

- All timestamps automatically formatted to Indonesian language
- Format: "Tuesday, December 24, 2025"
- Auto convert to WIB timezone (UTC+7)

---

## 🛠️ Setup & Installation

### Prerequisites

- Go 1.24 atau lebih tinggi
- PostgreSQL 12+
- Git

### 1. Clone Repository

```bash
git clone https://github.com/rifqi142/indico-be.git
cd indico-be
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Database

```bash
# Login ke PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE indico_db;

# Exit
\q
```

### 4. Configuration

```bash
# Copy .env.example to .env
cp .env.example .env

# Edit .env according to your configuration
```

**File `.env`:**

```env
# Application
APP_NAME=indico-be
APP_ENV=development
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=indico_db
DB_SSL_MODE=disable

# JWT
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRATION=5m

# Server
SERVER_READ_TIMEOUT=10s
SERVER_WRITE_TIMEOUT=10s
```

### 5. Run Application

**Option 1: Using Air (Hot Reload - Recommended)**

```bash
# Install Air
go install github.com/air-verse/air@latest

# Run with hot reload
air
```

**Option 2: Standard Go Run**

```bash
go run cmd/server/main.go
```

**Option 3: Build & Run**

```bash
# Build
go build -o bin/server cmd/server/main.go

# Run
./bin/server
```

Server will run at: `http://localhost:8080`

---

## 📚 API Documentation

### Base URL

```
http://localhost:8080
```

### Health Check

```bash
GET /health
```

### 1. Authentication

#### Login (Dummy)

```bash
POST /login
Content-Type: application/json

{
  "username": "admin",
  "password": "password"
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "type": "Bearer"
  }
}
```

**Note:** Login accepts any username/password combination and will generate JWT token that expires in **5 minutes**.

---

### 2. Vouchers (Protected - Requires JWT Token)

**Headers for all voucher endpoints:**

```
Authorization: Bearer YOUR_JWT_TOKEN
```

#### Get All Vouchers (with Pagination, Search, Sorting)

**Basic Request:**

```bash
GET /vouchers
```

**With Query Parameters:**

```bash
GET /vouchers?page=1&page_size=10&search=WELCOME&sort_by=discount&sort_order=desc
```

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `page` | integer | No | Page number (default: 1) |
| `page_size` | integer | No | Items per page (default: 10, max: 100) |
| `search` | string | No | Search by code, name, or description |
| `sort_by` | string | No | Sort field: id, code, name, discount, created_at |
| `sort_order` | string | No | Sort order: asc, desc (default: asc) |
| `is_active` | boolean | No | Filter by active status |

**Response:**

```json
{
  "status": "success",
  "message": "Vouchers retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "code": "WELCOME2025",
        "name": "Welcome Bonus 2025",
        "description": "Special discount for new customers",
        "discount": 25.0,
        "max_usage": 100,
        "used_count": 0,
        "valid_from": "Wednesday, January 1, 2025",
        "valid_until": "Wednesday, December 31, 2025",
        "is_active": true,
        "created_at": "Tuesday, December 24, 2025",
        "updated_at": "Tuesday, December 24, 2025"
      }
    ],
    "pagination": {
      "current_page": 1,
      "page_size": 10,
      "total_pages": 2,
      "total_items": 15
    }
  }
}
```

#### Get Voucher by ID

```bash
GET /vouchers/get-by-id/1
```

#### Create Voucher

```bash
POST /vouchers
Content-Type: application/json

{
  "code": "NEWYEAR2025",
  "name": "New Year Special",
  "description": "Happy New Year discount",
  "discount": 30.0,
  "max_usage": 100,
  "valid_from": "2025-01-01T00:00:00Z",
  "valid_until": "2025-01-31T23:59:59Z",
  "is_active": true
}
```

**Field Validation:**

- `code`: required, min=3, max=50, unique
- `name`: required, min=3, max=255
- `description`: optional
- `discount`: required, 0-100
- `max_usage`: required, min=1
- `valid_from`: required, ISO 8601 format
- `valid_until`: required, must be after valid_from
- `is_active`: optional (default: true)

#### Update Voucher

```bash
PUT /vouchers/1
Content-Type: application/json

{
  "name": "Updated Name",
  "discount": 35.0,
  "is_active": false
}
```

**Note:** All fields are optional. Only send fields you want to update.

#### Delete Voucher (Soft Delete)

```bash
DELETE /vouchers/1
```

---

### 3. CSV Operations

#### Upload CSV

```bash
POST /vouchers/upload-csv
Content-Type: multipart/form-data

file: sample_vouchers.csv
```

**CSV Format:**

```csv
code,name,description,discount,max_usage,valid_from,valid_until,is_active
TESTCSV01,Test Voucher,Description,10.00,50,2025-01-01,2025-12-31,true
```

**Response:**

```json
{
  "status": "success",
  "message": "CSV uploaded successfully",
  "data": {
    "success_count": 10,
    "failed_count": 2,
    "errors": ["Row 3: duplicate key value", "Row 5: invalid date format"]
  }
}
```

#### Export CSV

```bash
GET /vouchers/export
```

**Response:** File download `vouchers_export_YYYYMMDD_HHMMSS.csv`

---

## 📦 Database Schema

### Vouchers Table

| Column      | Type          | Constraints      | Description                 |
| ----------- | ------------- | ---------------- | --------------------------- |
| id          | SERIAL        | PRIMARY KEY      | Auto-increment ID           |
| code        | VARCHAR(50)   | UNIQUE, NOT NULL | Voucher code                |
| name        | VARCHAR(255)  | NOT NULL         | Voucher name                |
| description | TEXT          | -                | Voucher description         |
| discount    | DECIMAL(10,2) | NOT NULL         | Discount percentage (0-100) |
| max_usage   | INTEGER       | NOT NULL         | Maximum usage count         |
| used_count  | INTEGER       | DEFAULT 0        | Current usage count         |
| valid_from  | TIMESTAMP     | NOT NULL         | Start validity date         |
| valid_until | TIMESTAMP     | NOT NULL         | End validity date           |
| is_active   | BOOLEAN       | DEFAULT TRUE     | Active status               |
| created_at  | TIMESTAMP     | DEFAULT NOW()    | Creation timestamp          |
| updated_at  | TIMESTAMP     | DEFAULT NOW()    | Last update timestamp       |
| deleted_at  | TIMESTAMP     | NULL             | Soft delete timestamp       |

**Indexes:**

- `idx_vouchers_code` on `code`
- `idx_vouchers_is_active` on `is_active`
- `idx_vouchers_deleted_at` on `deleted_at`
- `idx_vouchers_valid_from` on `valid_from`
- `idx_vouchers_valid_until` on `valid_until`

---

## 📄 License

This project is created for **Take Home Test Indico By Telkomsel** purposes.

---

## 👨‍💻 Author

**Muhammad Rifqi Setiawan**

- GitHub: [@rifqi142](https://github.com/rifqi142)
- LinkedIn: [Muhammad Rifqi Setiawan](https://www.linkedin.com/in/muhrifqis/)
