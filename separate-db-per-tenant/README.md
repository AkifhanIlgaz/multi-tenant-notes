# Multi-Tenant Notes API — Separate Schema Approach

A Go-based multi-tenant notes/announcements API implementing the Separate Schema multi-tenancy pattern using PostgreSQL, Fiber v3, and GORM.

## 🎯 Overview

This project demonstrates a multi-tenant application where multiple organizations (tenants) share the same database but have isolated schemas. Each tenant gets its own PostgreSQL schema (e.g., `beyaz-futbol`), and application code sets the schema per request to ensure strict data isolation.

### Separate Schema Pattern

This implementation uses the Separate Schema approach where:

- Single Database: All tenants share one database instance
- Separate Schemas: Each tenant has its own schema with the same set of tables
- Schema Routing: The application sets `search_path` dynamically per request
- Strong Isolation: Data is isolated at the schema level

Pros:
- Stronger isolation than shared tables
- Easier to enforce tenant-specific constraints
- Safer migrations (can target per tenant)

Cons:
- More complex routing and connection management
- Schema proliferation (many schemas to maintain)
- Migrations must be schema-aware

## ✨ Features

- Separate schema per tenant (e.g., `<slug>`)
- JWT-based authentication with tenant-awareness
- Announcement/Notes management per tenant
- Per-request schema selection via transaction wrapper
- RESTful API with Fiber v3
- Migrations and seed data per tenant
- Hexagonal architecture (Ports & Adapters)

## 📦 Project Structure

Key directories and responsibilities:

- `cmd/`: Application entrypoint
- `internal/core/`: Domain layer
  - `models/`: Entities (Tenant, User, Announcement)
  - `ports/`: Repository interfaces
  - `service/`: Business logic
- `internal/adapters/`: Adapters layer
  - `api/`: HTTP handlers, middleware, routes
  - `db/`: Database connector, config, repositories, migrations, seed
    - `migrations/`: Schema creation and auto-migrate logic
    - `repositories/`: GORM repositories using per-request schema routing

## 🚀 Getting Started

1. Clone the repository:
   - `git clone <repository-url>`

2. Start PostgreSQL:
   - `docker compose up -d`

3. Run the application:
   - `cd shared-db-separate-schema`
   - `go run cmd/main.go`

Server starts at `http://localhost:3000`.

## 🗄 Database and Schemas



Tables inside each tenant schema:
- `users`
- `announcements`

Migrations logic:
- Create schema if not exists: `CREATE SCHEMA IF NOT EXISTS tenant_<slug>`
- Set schema for operations: `SET search_path TO tenant_<slug>`
- Run `AutoMigrate` for `users` and `announcements` inside that schema

Seeding logic:
- Before inserting tenant-specific data (users, announcements), set `search_path` to the tenant schema
- Use distinct demo users per tenant


## 📡 API Endpoints

Health Check:
- `GET /health`
- Response: `{"status":"ok"}`

Authentication:
- `POST /api/auth/login`
- Body:
  ```
  {
    "email": "sinan.engin@beyaz-futbol.com",
    "password": "password123",
    "tenant": "beyaz-futbol"
  }
  ```
- Response includes JWT with tenant awareness; middleware places `tenant_slug` in request context

Announcements (Protected):
- `GET /api/announcements`
- `POST /api/announcements`
- `DELETE /api/announcements/:id`
- All require `Authorization: Bearer <token>` and route queries to the tenant's schema

## 🔧 Configuration

Database:
- Configure connection settings under `internal/adapters/db/config`
- Connection is shared; schema is set per request via transaction wrapper

Migrations:
- Under `internal/adapters/db/migrations`
- Ensure migrations accept `tenantSlug` and set schema before `AutoMigrate`

Repositories:
- Under `internal/adapters/db/repositories`
- Wrap each repository method in the transaction wrapper to set `search_path` using the context’s `tenant_slug`

## 🧪 Demo Tenants and Users

Tenants (example slugs):
1. Beyaz Futbol — slug: `beyaz_futbol`
   - Users: `sinan.engin@beyaz_futbol.com`, `ahmet.cakar@beyaz_futbol.com`, `ertem.sener@beyaz_futbol.com`
   - Password: `password123`

2. Hell Kitchen — slug: `hell_kitchen`
   - Users: `gordon.ramsay@hell_kitchen.com`, `mehmet.yalcinkaya@hell_kitchen.com`, `sofia.fehn@hell_kitchen.com`
   - Password: `password123`

3. Mentalist — slug: `mentalist`
   - Users: `patrick.jane@mentalist.com`, `kimball.cho@mentalist.com`, `teresa.lisbon@mentalist.com`
   - Password: `password123`



## 🛠 Tech Stack

- Language: Go
- Web Framework: Fiber v3
- ORM: GORM
- Database: PostgreSQL
- Architecture: Hexagonal (Ports & Adapters)
- Auth: JWT
