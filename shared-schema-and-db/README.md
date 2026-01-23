# Multi-Tenant Notes API - Shared Schema Approach

A Go-based multi-tenant notes/announcements API implementing the **Shared Schema** multi-tenancy pattern using PostgreSQL, Fiber v3, and GORM.


## 🎯 Overview

This project demonstrates a multi-tenant application where multiple organizations (tenants) share the same database and schema. Each tenant's data is isolated using a `tenant_id` discriminator column, ensuring data security and proper isolation.

### Shared Schema Pattern

This implementation uses the **Shared Schema** approach where:

- **Single Database**: All tenants share one database
- **Single Schema**: All tables are shared across tenants
- **Tenant Discriminator**: Each row has a `tenant_id` column to identify ownership
- **Row-Level Isolation**: Queries automatically filter by `tenant_id`

**Pros:**
- Cost-effective (single database instance)
- Easy maintenance and updates
- Efficient resource utilization

**Cons:**
- Requires careful query filtering
- All tenants affected by database issues
- Scaling limitations for very large datasets

## Features

- **Multi-tenant architecture** with shared schema
- **JWT-based authentication** with tenant awareness
- **Announcement/Notes management** per tenant
- **Automatic tenant isolation** via middleware
- **RESTful API** with Fiber v3
- **Database migrations** with auto-seeding
- **Hexagonal architecture** (Ports & Adapters)


### Installation

```bash
git clone <repository-url>
```


3. Set up PostgreSQL database:
```bash
docker compose up -d
```


4. Run the application:
```bash
cd shared-schema-and-db
go run cmd/main.go
```

The server will start on `http://localhost:3000`


## 📡 API Endpoints

### Health Check

```http
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

### Authentication

#### Login

```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "sinan.engin@beyaz-futbol.com",
  "password": "password123",
  "tenant": "beyaz-futbol"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
}
```

### Announcements (Protected Routes)

All announcement endpoints require JWT authentication via `Authorization: Bearer <token>` header.

#### Get All Announcements

```http
GET /api/notes
Authorization: Bearer <token>
```

**Response:**
```json
{
  "announcements": [
    {
      "id": 1,
      "title": "Yeni Sezon Başlıyor!",
      "content": "Sevgili futbolseverler...",
      "created_at": "2024-01-20T10:30:00Z",
      "user_id": 1,
      "tenant_id": 1
    }
  ]
}
```

#### Create Announcement

```http
POST /api/notes
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Yeni Duyuru",
  "content": "Duyuru içeriği..."
}
```

**Response:**
```json
{
  "message": "announcement created successfully",
  "announcement": {
    "id": 10,
    "title": "Yeni Duyuru",
    "content": "Duyuru içeriği...",
    "created_at": "2024-01-23T12:00:00Z",
    "user_id": 1,
    "tenant_id": 1
  }
}
```

#### Delete Announcement

```http
DELETE /api/notes/:id
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "announcement deleted successfully"
}
```

## 🗄 Database Schema

### Tenants Table

```sql
CREATE TABLE tenants (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL UNIQUE
);
```

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id) ON DELETE CASCADE
);
```

### Announcements Table

```sql
CREATE TABLE announcements (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id) ON DELETE CASCADE
);
```


### 1. Beyaz Futbol 🏆
A Turkish football discussion show.

**Users:**
- `sinan.engin@beyaz-futbol.com` - Sinan Engin
- `ahmet.cakar@beyaz-futbol.com` - Ahmet Çakar
- `ertem.sener@beyaz-futbol.com` - Ertem Şener

**Password:** `password123`

### 2. Hell Kitchen 👨‍🍳
A cooking competition show.

**Users:**
- `gordon.ramsay@hell-kitchen.com` - Gordon Ramsay
- `mehmet.yalcinkaya@hell-kitchen.com` - Mehmet Yalçınkaya
- `sofia.fehn@hell-kitchen.com` - Sofia Fehn

**Password:** `password123`

### 3. Mentalist 🔍
A crime investigation series.

**Users:**
- `patrick.jane@mentalist.com` - Patrick Jane
- `kimball.cho@mentalist.com` - Kimball Cho
- `teresa.lisbon@mentalist.com` - Teresa Lisbon

**Password:** `password123`
