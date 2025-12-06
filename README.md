# URL Shortener Backend

A production-ready REST API for URL shortening built with Go, Fiber, and PostgreSQL.

## 🔧 Tech Stack

| Component | Technology |
|-----------|-----------|
| Language | Go 1.23 |
| Framework | Fiber v2 |
| Database | PostgreSQL 16 |
| ORM | GORM |
| Container | Docker |
| Migration | GORM AutoMigrate |

## 📌 Features

### ✓ URL Shortening API
- Accepts long URLs via `POST /api/shorten`
- Validates URL format
- Generates unique 7-character short codes
- Persists data in PostgreSQL

### ✓ URL Redirection
- `GET /:code` endpoint
- Looks up original URL
- Redirects using HTTP 307
- Tracks visit count

### ✓ Reuse Existing URLs
- Returns the same short code if URL was already shortened

### ✓ Health Endpoint
- `GET /health` returns `{ status: "ok" }`

### ✓ Docker-ready
- Multi-stage Dockerfile
- Runs with PostgreSQL via docker-compose

## 📁 Project Structure

```
backend/
  cmd/server/main.go
  internal/
    config/          # Env loader
    db/              # PostgreSQL connection
    models/          # URL model
    repositories/    # DB operations
    services/        # Business logic
    handlers/        # Fiber controllers
    routes/          # Route registration
    utils/           # Short code generator + helpers
  Dockerfile
```

## ⚙️ Environment Variables

Create a `.env` file in the backend directory:

```env
APP_PORT=8080
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=url_user
POSTGRES_PASSWORD=url_pass
POSTGRES_DB=url_shortener
BASE_SHORT_URL=http://localhost:8080
```

## 🚀 Getting Started

### Running with Docker Compose
```bash
docker compose up --build
```

The backend will be available at `http://localhost:8080`

### Running Locally
```bash
cd backend
go run cmd/server/main.go
```

## 🧪 API Endpoints

### Shorten a URL
**Request:**
```bash
POST http://localhost:8080/api/shorten

{
  "url": "https://google.com"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "short_url": "http://localhost:8080/AbC123x",
    "short_code": "AbC123x",
    "original_url": "https://google.com"
  }
}
```

### Redirect to Original URL
**Request:**
```bash
GET http://localhost:8080/AbC123x
```

**Response:** Redirects to the original URL (HTTP 307)

### Health Check
**Request:**
```bash
GET http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok"
}
```

## 🔗 Service URLs

| Service | URL |
|---------|-----|
| Backend API | http://localhost:8080 |
| Backend Health | http://localhost:8080/health |
| PostgreSQL | localhost:5432 |

## 🎯 Future Improvements

- Custom short code support (vanity URLs)
- Click analytics dashboard
- URL expiration support
- Authentication for protected URLs
- QR code generation
