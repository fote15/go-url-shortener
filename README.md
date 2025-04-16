
# 🐳 Go App with Postgres (Dockerized)

This project is a Go backend API with:

- PostgreSQL as a database
- Automatic DB migrations on startup
- Clean multi-stage Docker build
- `.env` support for config

---

## 📦 Requirements

- Docker
- Docker Compose

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

---

## 📡 API Endpoints

### 🔐 Auth Routes

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| POST   | `/auth/register` | Register a new user     |
| POST   | `/auth/login`    | Login and get JWT token |

### 🔗 URL Routes (JWT Protected)

| Method | Endpoint              | Description                          |
|--------|-----------------------|--------------------------------------|
| POST   | `/urls/shorten`       | Create a new shortened URL           |
| GET    | `/urls/`              | List user's shortened URLs           |
| GET    | `/urls/{id}`          | Get a specific URL (by ID)           |
| PUT    | `/urls/{id}`          | Edit an existing shortened URL       |
| DELETE | `/urls/{id}`          | Delete a shortened URL               |
| GET    | `/urls/{id}/stats`    | Get visit statistics for a short URL |

### 🚀 Public Routes

| Method | Endpoint             | Description                         |
|--------|----------------------|-------------------------------------|
| GET    | `/url/{shortKey}`    | Redirect to the original long URL   |
