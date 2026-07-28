# Matrix QR Decomposition Platform

A multi-service backend application built with Go and TypeScript that performs QR matrix decomposition, computes matrix statistics, secures endpoints with JWT authentication, and runs through Docker Compose.

---

## Overview

This project demonstrates a service-oriented architecture where:

- A **Go API** performs matrix validation and QR decomposition using Gonum.
- A **Node.js API** computes statistical information from the decomposition results.
- Services communicate through HTTP.
- JWT authentication protects sensitive endpoints.
- Docker Compose orchestrates the entire system.

---

## Architecture

```text
Client
  │
  ▼
Go API (Fiber)
  │
  ├── JWT Authentication
  ├── Matrix Validation
  ├── QR Decomposition (Gonum)
  │
  ▼
Node API (Express + TypeScript)
  │
  ├── Statistics Calculation
  ├── Matrix Analysis
  │
  ▼
Combined Response
```

---

## Tech Stack

### Go Service

- Go
- Fiber
- Gonum
- JWT

### Node Service

- Node.js
- TypeScript
- Express

### Infrastructure

- Docker
- Docker Compose

---

## Features

### Authentication

- JWT token generation
- Protected QR decomposition endpoint
- Middleware-based authorization

### Matrix Validation

The API validates:

- Empty matrices
- Empty rows
- Non-rectangular matrices

### QR Decomposition

Uses Gonum to calculate:

- Q matrix
- R matrix

### Statistics

The Node service calculates:

- Maximum value
- Minimum value
- Sum
- Average
- Diagonal matrix detection

### Dockerized Deployment

Both services run inside Docker containers and communicate through Docker networking.

---

## Project Structure

```text
matrix-app/
│
├── go-api/
│   ├── cmd/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── services/
│   ├── Dockerfile
│   └── go.mod
│
├── node-api/
│   ├── src/
│   │   ├── controllers/
│   │   ├── services/
│   │   └── types/
│   ├── Dockerfile
│   └── package.json
│
├── docker-compose.yml
│
└── README.md
```

---

## Running with Docker

### Prerequisites

Install:

- Docker Desktop

### Build and Start

From the project root:

```bash
docker compose up --build
```

### Services

| Service  | URL                   |
| -------- | --------------------- |
| Go API   | http://localhost:8080 |
| Node API | http://localhost:3000 |

---

## Authentication

### Login

Endpoint:

```http
POST /login
```

Request:

```json
{
  "username": "admin",
  "password": "admin123"
}
```

Response:

```json
{
  "token": "your-jwt-token"
}
```

Save the token and use it in the Authorization header:

```text
Authorization: Bearer <token>
```

---

## QR Decomposition Endpoint

### Request

```http
POST /api/qr
```

Headers:

```text
Authorization: Bearer <token>
Content-Type: application/json
```

Body:

```json
{
  "matrix": [
    [1, 2],
    [3, 4],
    [5, 6]
  ]
}
```

### Example Response

```json
{
  "q": [
    [-0.169, 0.8971, 0.4082],
    [-0.5071, 0.276, -0.8165],
    [-0.8452, -0.345, 0.4082]
  ],
  "r": [
    [-5.9161, -7.4374],
    [0, 0.8281],
    [0, 0]
  ],
  "statistics": {
    "max": 0.8971,
    "min": -7.4374,
    "average": -0.8812,
    "sum": -13.2186,
    "hasDiagonalMatrix": false
  }
}
```

---

## Local Development

### Start the Go API

```bash
cd go-api
go run ./cmd
```

The Go API will run on:

```text
http://localhost:8080
```

### Start the Node API

```bash
cd node-api
npm install
npm run dev
```

The Node API will run on:

```text
http://localhost:3000
```

---

## Running Tests

### Execute Tests

```bash
cd go-api
go test ./...
```

### Coverage Report

```bash
go test ./... -cover
```

---

## API Flow

```text
POST /login
        │
        ▼
   JWT Token
        │
        ▼
POST /api/qr
        │
        ▼
JWT Middleware
        │
        ▼
Matrix Validation
        │
        ▼
QR Decomposition
        │
        ▼
Node Statistics API
        │
        ▼
Combined Response
```

---

## Future Improvements

Potential enhancements:

- PostgreSQL integration
- User management
- Calculation history
- Refresh tokens
- Swagger/OpenAPI documentation
- GitHub Actions CI/CD
- Frontend dashboard
- Role-based authorization

---
