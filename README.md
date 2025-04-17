# 🐳 Go URL Shortener with PostgreSQL (Dockerized)

This is a backend API for a URL shortener service built in Go, using PostgreSQL as the database and JWT for authentication. It supports automatic DB migrations and is fully dockerized.

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

### 2. Run with Docker

```bash
docker-compose up --build
```

---

## 📡 API Endpoints

---

### 🔐 Auth Routes

#### 📌 POST `/auth/register`
**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Success Response:**
```json
{
  "id": 1,
  "email": "user@example.com"
}
```

#### 📌 POST `/auth/login`
**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Success Response:**
```json
{
  "token": "your_jwt_token"
}
```

---

### 🔗 URL Routes (Protected with JWT)

All requests require a valid `Authorization: Bearer <token>` header.

#### 📌 POST `/urls/shorten`
**Request Body:**
```json
{
  "original_url": "https://example.com",
  "custom_key": "custom123" // optional
}
```

**Response:**
```json
{
  "short_url": "/custom123"
}
```

#### 📌 GET `/urls/`
**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "original_url": "https://example.com",
    "short_key": "custom123",
    "visits": 10,
    "created_at": "2025-04-17T10:00:00Z"
  },
  ...
]
```

#### 📌 GET `/urls/{id}`
**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "original_url": "https://example.com",
  "short_key": "custom123",
  "visits": 10,
  "created_at": "2025-04-17T10:00:00Z"
}
```

#### 📌 PUT `/urls/{id}`
**Request Body:**
```json
{
  "original_url": "https://updated-example.com",
  "custom_key": "" // optional
}
```

**Response:**
```json
{
  "data": "OK"
}
```

#### 📌 DELETE `/urls/{id}`
**Response:** `200 OK`

#### 📌 GET `/urls/{id}/stats`
**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "original_url": "https://example.com",
  "short_key": "custom123",
  "visits": 10,
  "created_at": "2025-04-17T10:00:00Z"
}
```

---

### 🚀 Public Route

#### 📌 GET `/url/{shortKey}`
**Redirects to the original URL**  
**Example:** `/url/custom123` → redirects to `https://example.com`

---

## 🔒 Authentication

- Use the `/auth/login` route to receive a JWT token.
- Send the token in the `Authorization` header for protected routes:
```
Authorization: Bearer your_token_here
```


---

## 📜 License

MIT