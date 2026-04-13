# User API

A RESTful API built with Go for user management. Features CRUD operations with a repository pattern architecture and PostgreSQL database.

## Tech Stack

- Go 1.24
- Chi (router)
- PostgreSQL 16
- pgx (database driver)
- Tern (migrations)

## Project Structure

```
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── response.go
│   └── user/
│       ├── dto.go
│       ├── handler.go
│       ├── model.go
│       └── repository.go
├── migrations/
│   ├── 001_create_users_table.sql
│   ├── tern.conf
│   └── tern.conf.example
├── .env.example
├── compose.yaml
└── go.mod
```

## Setup

1. Clone the repository

2. Copy environment files:
```bash
cp .env.example .env
cp migrations/tern.conf.example migrations/tern.conf
```

3. Configure your `.env` and `migrations/tern.conf` with your database credentials

4. Start PostgreSQL:
```bash
docker compose up -d
```

5. Run migrations:
```bash
cd migrations
tern migrate
```

6. Run the API:
```bash
go run ./cmd/api/
```

Server starts at `http://localhost:8080`

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | List all users |
| GET | `/users/{id}` | Get user by ID |
| POST | `/users` | Create new user |
| PUT | `/users/{id}` | Update user |
| DELETE | `/users/{id}` | Delete user |

## Usage Examples

**Create user:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@email.com"}'
```

**List users:**
```bash
curl http://localhost:8080/users
```

**Get user:**
```bash
curl http://localhost:8080/users/{id}
```

**Update user:**
```bash
curl -X PUT http://localhost:8080/users/{id} \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john.doe@email.com"}'
```

**Delete user:**
```bash
curl -X DELETE http://localhost:8080/users/{id}
```