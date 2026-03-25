# User API

A RESTful API built with Go for user management. Features CRUD operations with a repository pattern architecture, designed for easy database integration.

## Tech Stack

- Go 1.24
- Chi (router)
- UUID (google/uuid)

## Project Structure

```
├── internal/
│   ├── api/
│   │   └── response.go
│   └── user/
│       ├── dto.go
│       ├── handler.go
│       ├── model.go
│       └── repository.go
└── m.go
```

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | List all users |
| GET | `/users/{id}` | Get user by ID |
| POST | `/users` | Create new user |
| PUT | `/users/{id}` | Update user |
| DELETE | `/users/{id}` | Delete user |

## Running

```bash
go run main.go
```

Server starts at `http://localhost:8080`

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

## Roadmap

- [ ] Database integration (replace in-memory map)