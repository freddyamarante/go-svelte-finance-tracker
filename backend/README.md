# Finance Tracker Backend

A Go backend for the Finance Tracker application built with Gin, GORM, and SQLite.

## Project Structure

```
backend/
├── internal/                 # Private application code
│   ├── config/              # Configuration management
│   │   └── config.go        # Environment variables and app config
│   ├── database/            # Database operations
│   │   └── database.go      # Database connection and initialization
│   ├── handlers/            # HTTP request handlers
│   │   ├── health.go        # Health check endpoint
│   │   └── transactions.go  # Transaction CRUD operations
│   ├── models/              # Data models
│   │   └── transaction.go   # Transaction model definition
│   └── server/              # HTTP server setup
│       └── server.go        # Server configuration and routing
├── scripts/                 # Utility scripts
│   └── run.sh              # Development run script
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── Dockerfile              # Production Docker image
├── Dockerfile.dev          # Development Docker image
├── .air.toml              # Hot reload configuration
├── .dockerignore          # Docker build exclusions
└── README.md              # This file
```

## Architecture

This backend follows clean architecture principles with clear separation of concerns:

### Layers

1. **Models** (`internal/models/`)
   - Data structures and business entities
   - GORM model definitions

2. **Database** (`internal/database/`)
   - Database connection management
   - Migration and seeding logic
   - Database operations abstraction

3. **Handlers** (`internal/handlers/`)
   - HTTP request/response handling
   - Input validation
   - Business logic coordination

4. **Server** (`internal/server/`)
   - HTTP server setup
   - Middleware configuration
   - Route definitions

5. **Config** (`internal/config/`)
   - Environment variable management
   - Application configuration
   - Type-safe configuration structs

### Key Features

- **Clean Architecture**: Separation of concerns with clear boundaries
- **Dependency Injection**: Loose coupling between components
- **Error Handling**: Comprehensive error handling and logging
- **Configuration Management**: Environment-based configuration
- **Database Abstraction**: GORM for database operations
- **CORS Support**: Configurable CORS policies
- **Logging**: Structured logging with configurable levels

## Development

### Prerequisites

- Go 1.24+
- SQLite (for local development)

### Local Development

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Run with hot reload:**
   ```bash
   ./scripts/run.sh
   ```

3. **Run directly:**
   ```bash
   go run main.go
   ```

### Environment Variables

See `env.example` for available configuration options:

- `BACKEND_PORT`: Server port (default: 6060)
- `BACKEND_HOST`: Server host (default: localhost)
- `DB_PATH`: Database file path (default: finance.db)
- `LOG_REQUESTS`: Enable HTTP request logging (default: true)
- `LOG_SQL`: Enable SQL query logging (default: false)
- `LOG_CORS`: Enable CORS configuration logging (default: false)

### API Endpoints

- `GET /health` - Health check
- `GET /transactions` - List all transactions
- `POST /transactions` - Create a new transaction

## Docker

### Development
```bash
docker-compose -f ../docker-compose.dev.yml up --build
```

### Production
```bash
docker-compose up --build
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/handlers
```

## Database

The application uses SQLite for simplicity. The database file is created automatically on first run.

### Schema

```sql
CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    description TEXT NOT NULL,
    amount REAL NOT NULL,
    type TEXT NOT NULL
);
```

### Sample Data

The application automatically seeds sample data on first run:
- Initial Savings: $1000 (income)
- Groceries: $50 (expense)
- Salary: $2500 (income)
- Gas: $45 (expense)
- Freelance Work: $300 (income)

## Contributing

1. Follow Go best practices and conventions
2. Add tests for new functionality
3. Update documentation for API changes
4. Use meaningful commit messages

## License

This project is part of the Finance Tracker application. 