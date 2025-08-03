# Finance Tracker

A finance tracking application built with Svelte frontend and Go backend.

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Git

### Setup & Run

1. **Clone and setup environment:**
   ```bash
   git clone <repository-url>
   cd go-svelte-finance-tracker
   ./scripts/setup-env.sh
   ```

2. **Start the application:**
   ```bash
   # Development
   docker-compose -f docker-compose.dev.yml up
   
   # Production
   docker-compose up
   ```

3. **Access the application:**
   - Frontend: http://localhost:5173 (dev) or http://localhost:3000 (prod)
   - Backend API: http://localhost:8080

## Development

### Local Development (without Docker)
```bash
# Backend
cd backend && ./scripts/run.sh

# Frontend (in another terminal)
cd frontend && bun run dev
```

### Environment Configuration
- Edit `.env` file to customize ports, logging, and API URLs
- See `ENVIRONMENT_README.md` for detailed configuration options

### Available Commands
```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up

# Start production environment
docker-compose up

# Stop all services
docker-compose down

# View logs
docker-compose logs -f

# Rebuild containers
docker-compose up --build
```

## Project Structure
```
├── backend/          # Go API server
├── frontend/         # Svelte application
├── scripts/          # Setup and utility scripts
├── docker-compose.yml        # Production Docker setup
├── docker-compose.dev.yml    # Development Docker setup
└── .env              # Environment configuration
```

For detailed environment configuration, see `ENVIRONMENT_README.md`.