# Multi-Tenant Notes API — Separate Schema Approach


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
