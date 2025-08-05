# Ticket Booking System 🎫

(Distributed) system for booking tickets 

Building this for fun and learning more about microservices

## Project Architecture

```
ticket-booking-system/
├── 📁 .github/
│   └── 📁 workflows/
│       └── 📄 ci.yml                    # GitHub Actions CI/CD pipeline
├── 📁 api-gateway/                      # API Gateway service
│   ├── 📄 Dockerfile
│   ├── 📄 go.mod
│   └── 📁 internal/
│       ├── 📁 config/
│       ├── 📁 handler/
│       ├── 📁 middleware/
│       └── 📁 service/
├── 📁 deployments/                      # Deployment configurations
│   ├── 📁 docker/
│   └── 📁 k8s/
├── 📁 proto/                           # Protocol Buffers definitions
│   ├── 📄 user.proto
│   ├── 📄 event.proto
│   └── 📄 booking.proto
├── 📁 services/                        # Microservices
│   ├── 📁 user-service/                # User management & authentication
│   │   ├── 📄 Dockerfile
│   │   ├── 📄 go.mod
│   │   ├── 📄 go.sum
│   │   ├── 📁 cmd/
│   │   │   └── 📄 main.go              # Application entry point
│   │   ├── 📁 docs/                    # API documentation
│   │   ├── 📁 internal/                # Private application code
│   │   │   ├── 📁 config/              # Configuration management
│   │   │   │   └── 📄 config.go
│   │   │   ├── 📁 database/            # Database migrations
│   │   │   │   └── 📁 migrations/
│   │   │   │       ├── 📄 000001_init_schema.up.sql
│   │   │   │       └── 📄 000001_init_schema.down.sql
│   │   │   ├── 📁 domain/              # Business logic & entities
│   │   │   │   └── 📄 user.go
│   │   │   ├── 📁 handler/             # HTTP handlers
│   │   │   │   ├── 📄 handler.go
│   │   │   │   ├── 📄 auth.go
│   │   │   │   ├── 📄 user.go
│   │   │   │   └── 📄 middleware.go
│   │   │   ├── 📁 repository/          # Data access layer
│   │   │   │   ├── 📄 repository.go    # Repository interfaces
│   │   │   │   ├── 📁 postgres/        # PostgreSQL implementation
│   │   │   │   │   ├── 📄 postgres.go
│   │   │   │   │   ├── 📄 user.go
│   │   │   │   │   └── 📄 auth.go
│   │   │   │   └── 📁 redis/           # Redis implementation
│   │   │   │       └── 📄 redis.go
│   │   │   └── 📁 service/             # Business logic services
│   │   │       ├── 📄 service.go       # Service interfaces
│   │   │       ├── 📄 user.go          # User service implementation
│   │   │       └── 📄 auth.go          # Auth service implementation
│   │   ├── 📁 pkg/                     # Public packages
│   │   └── 📁 test/                    # Integration tests
│   ├── 📁 event-service/               # Event management service
│   │   ├── 📄 Dockerfile
│   │   ├── 📄 go.mod
│   │   ├── 📁 cmd/
│   │   │   └── 📄 main.go
│   │   ├── 📁 internal/
│   │   │   ├── 📁 config/
│   │   │   ├── 📁 domain/
│   │   │   ├── 📁 handler/
│   │   │   ├── 📁 repository/
│   │   │   └── 📁 service/
│   │   └── 📁 test/
│   └── 📁 booking-service/             # Booking workflow service
│       ├── 📄 Dockerfile
│       ├── 📄 go.mod
│       ├── 📁 cmd/
│       │   └── 📄 main.go
│       ├── 📁 internal/
│       │   ├── 📁 config/
│       │   ├── 📁 domain/
│       │   ├── 📁 handler/
│       │   ├── 📁 repository/
│       │   └── 📁 service/
│       └── 📁 test/
├── 📄 .gitignore
├── 📄 docker-compose.yml               # Local development environment
├── 📄 go.mod                          # Root Go module
├── 📄 Makefile                        # Build & deployment commands
└── 📄 README.md                       # This file
```

## Tech Stack 

- **Backend**: Go 1.24.5
- **Databases**: PostgreSQL 15, Redis 7
- **Message Queue**: Apache Kafka
- **Containerization**: Docker & Docker Compose
- **API Gateway**: Custom implementation
- **Authentication**: JWT tokens
- **Testing**: Go testing framework
- **CI/CD**: GitHub Actions
- **Monitoring**: Health checks & logging

## Services Architecture

### 🏗️ Service Structure (Clean Architecture)

Each microservice follows **Clean Architecture** principles:

```
📁 internal/
├── 📁 config/          # Configuration & environment variables
├── 📁 domain/          # Business entities & logic (no dependencies)
├── 📁 repository/      # Data access layer (interfaces + implementations)
├── 📁 service/         # Business logic services
└── 📁 handler/         # HTTP handlers & routing
``` 
## ToDo 

- [x] Project layout & clean architecture
- [x] Docker Compose setup
- [x] Makefile with essential commands
- [x] GitHub Actions CI/CD pipeline
- [x] User service repository layer
- [ ] Complete user service implementation
- [ ] Implement authentication & JWT tokens
- [ ] Build event service
- [ ] Build booking service
- [ ] Implement API gateway
- [ ] Add comprehensive testing
- [ ] Add monitoring & observability
- [ ] Production deployment setup

## Done  

- [x] Project layout following clean architecture pattern
- [x] Makefile with database and migration commands
- [x] Database schema for user-service
- [x] Repository layer interfaces for user-service
- [x] Docker Compose with all infrastructure services
- [x] CI/CD pipeline with testing, building, and security scanning
- [x] Service separation and proper folder structure