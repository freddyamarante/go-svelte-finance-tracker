# Docker Setup for Finance Tracker

This project has been dockerized for easy deployment and development. The setup includes both production and development configurations.

## Quick Start

### Production Build
```bash
# Build and run the entire application
docker-compose up --build

# Access the application
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
```

### Development Build
```bash
# Run with hot reloading for development
docker-compose -f docker-compose.dev.yml up --build

# Access the application
# Frontend: http://localhost:5173
# Backend API: http://localhost:8080
```

## Architecture

The application consists of two main services:

### Backend (Go + Gin)
- **Port**: 8080
- **Database**: SQLite (file-based)
- **Framework**: Gin with GORM
- **Features**: REST API for transactions

### Frontend (Svelte + Vite)
- **Port**: 3000 (production) / 5173 (development)
- **Framework**: SvelteKit
- **Styling**: Tailwind CSS
- **Build Tool**: Vite

## Docker Files Structure

```
├── backend/
│   ├── Dockerfile          # Production backend image
│   ├── Dockerfile.dev      # Development backend with hot reload
│   ├── .air.toml          # Hot reload configuration
│   └── .dockerignore      # Backend build exclusions
├── frontend/
│   ├── Dockerfile          # Production frontend image
│   ├── Dockerfile.dev      # Development frontend with hot reload
│   ├── nginx.conf         # Nginx configuration for production
│   └── .dockerignore      # Frontend build exclusions
├── docker-compose.yml      # Production orchestration
├── docker-compose.dev.yml  # Development orchestration
└── DOCKER_README.md       # This file
```

## Development Workflow

### 1. Start Development Environment
```bash
docker-compose -f docker-compose.dev.yml up --build
```

### 2. Make Changes
- Backend changes will automatically reload thanks to Air
- Frontend changes will automatically reload thanks to Vite
- No need to restart containers for code changes

### 3. View Logs
```bash
# All services
docker-compose -f docker-compose.dev.yml logs -f

# Specific service
docker-compose -f docker-compose.dev.yml logs -f backend
docker-compose -f docker-compose.dev.yml logs -f frontend
```

## Production Deployment

### 1. Build Production Images
```bash
docker-compose up --build
```

### 2. Run in Background
```bash
docker-compose up -d
```

### 3. Stop Services
```bash
docker-compose down
```

## Environment Variables

### Frontend
- `VITE_API_URL`: Backend API URL (default: http://localhost:8080)

### Backend
- `GIN_MODE`: Gin mode (debug/release)

## Database

The application uses SQLite for simplicity. The database file (`finance.db`) is created automatically and persists in the backend container.

For production, consider:
- Using a proper database like PostgreSQL
- Setting up database migrations
- Implementing backup strategies

## Troubleshooting

### Common Issues

1. **Port Already in Use**
   ```bash
   # Check what's using the port
   lsof -i :8080
   lsof -i :3000
   
   # Stop conflicting services
   docker-compose down
   ```

2. **Build Failures**
   ```bash
   # Clean up and rebuild
   docker-compose down
   docker system prune -f
   docker-compose up --build
   ```

3. **Permission Issues**
   ```bash
   # Fix file permissions
   sudo chown -R $USER:$USER .
   ```

### Logs and Debugging
```bash
# View real-time logs
docker-compose logs -f

# Access container shell
docker-compose exec backend sh
docker-compose exec frontend sh

# Check container status
docker-compose ps
```

## Performance Optimization

### Production Optimizations
- Multi-stage builds for smaller images
- Nginx for serving static files
- Proper caching headers
- Security headers

### Development Optimizations
- Volume mounts for fast file access
- Hot reloading for both frontend and backend
- Shared networks for service communication

## Security Considerations

- Frontend runs behind Nginx with security headers
- Backend CORS is configured for development
- Database file permissions are managed
- No sensitive data in images

## Next Steps

For production deployment, consider:
1. Setting up a reverse proxy (Nginx/Traefik)
2. Implementing SSL/TLS certificates
3. Setting up monitoring and logging
4. Database migration strategy
5. CI/CD pipeline integration 