# Environment Configuration

This project now uses a centralized environment configuration system where all environment variables are managed at the root level and shared between backend and frontend services.

## Quick Setup

1. **Create your environment file:**
   ```bash
   ./scripts/setup-env.sh
   ```

2. **Edit the configuration:**
   ```bash
   nano .env
   ```

3. **Start the application:**
   ```bash
   # Development
   docker-compose -f docker-compose.dev.yml up
   
   # Production
   docker-compose up
   ```

## Environment Variables Structure

### Application Configuration
| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `ENV` | string | `development` | Environment (development/production) |
| `APP_NAME` | string | `finance-tracker` | Application name |

### Backend Configuration
| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `BACKEND_PORT` | string | `8080` | Backend server port |
| `BACKEND_HOST` | string | `localhost` | Backend server host |
| `DB_PATH` | string | `finance.db` | SQLite database file path |
| `LOG_REQUESTS` | boolean | `true` | Enable HTTP request logging |
| `LOG_SQL` | boolean | `false` | Enable SQL query logging |
| `LOG_CORS` | boolean | `false` | Enable CORS configuration logging |

### Frontend Configuration
| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `VITE_API_URL` | string | `http://localhost:8080` | Backend API URL |
| `VITE_APP_NAME` | string | `Finance Tracker` | Application name for frontend |
| `FRONTEND_PORT` | string | `5173` | Frontend development port |
| `FRONTEND_HOST` | string | `localhost` | Frontend development host |

### CORS Configuration
| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `CORS_ALLOW_ORIGINS` | string | `*` | Comma-separated allowed origins |
| `CORS_ALLOW_METHODS` | string | `GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS` | Comma-separated allowed methods |
| `CORS_ALLOW_HEADERS` | string | `Origin,Content-Length,Content-Type,Accept,Authorization,X-Requested-With` | Comma-separated allowed headers |
| `CORS_MAX_AGE` | integer | `43200` | CORS preflight cache time (seconds) |

### Docker Configuration
| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `DOCKER_NETWORK` | string | `finance-network` | Docker network name |
| `BACKEND_CONTAINER_NAME` | string | `finance-backend` | Backend container name |
| `FRONTEND_CONTAINER_NAME` | string | `finance-frontend` | Frontend container name |

## Configuration Examples

### Development Environment
```bash
# .env
ENV=development
BACKEND_PORT=8080
FRONTEND_PORT=5173
LOG_REQUESTS=true
LOG_SQL=false
LOG_CORS=false
VITE_API_URL=http://localhost:8080
```

### Production Environment
```bash
# .env
ENV=production
BACKEND_PORT=8080
FRONTEND_PORT=3000
LOG_REQUESTS=false
LOG_SQL=false
LOG_CORS=false
VITE_API_URL=https://api.yourdomain.com
```

### Debug Environment
```bash
# .env
ENV=development
BACKEND_PORT=8080
FRONTEND_PORT=5173
LOG_REQUESTS=true
LOG_SQL=true
LOG_CORS=true
VITE_API_URL=http://localhost:8080
```

## Usage Patterns

### Local Development
```bash
# 1. Setup environment
./scripts/setup-env.sh

# 2. Start backend only
cd backend && ./scripts/run.sh

# 3. Start frontend only
cd frontend && bun run dev

# 4. Or start both with Docker
docker-compose -f docker-compose.dev.yml up
```

### Docker Development
```bash
# Start with custom environment
LOG_REQUESTS=false docker-compose -f docker-compose.dev.yml up

# Start with different ports
BACKEND_PORT=9000 FRONTEND_PORT=3001 docker-compose -f docker-compose.dev.yml up
```

### Production Deployment
```bash
# Build and run production
docker-compose up --build

# Run in background
docker-compose up -d

# Check logs
docker-compose logs -f
```

## Environment File Management

### Creating Environment Files
```bash
# Create from example
cp env.example .env

# Create production environment
cp env.example .env.production

# Create staging environment
cp env.example .env.staging
```

### Using Different Environment Files
```bash
# Use specific environment file
docker-compose --env-file .env.production up

# Override specific variables
BACKEND_PORT=9000 docker-compose up
```

## Security Considerations

### Sensitive Data
- Never commit `.env` files to version control
- Use `.env.example` for documentation
- Store secrets in secure environment managers

### Production Security
```bash
# Production .env example
ENV=production
LOG_REQUESTS=false
LOG_SQL=false
LOG_CORS=false
CORS_ALLOW_ORIGINS=https://yourdomain.com
VITE_API_URL=https://api.yourdomain.com
```

## Troubleshooting

### Common Issues

1. **Environment variables not loading:**
   ```bash
   # Check if .env file exists
   ls -la .env
   
   # Verify file permissions
   chmod 600 .env
   ```

2. **Docker not picking up changes:**
   ```bash
   # Rebuild containers
   docker-compose down
   docker-compose up --build
   ```

3. **Port conflicts:**
   ```bash
   # Check what's using the port
   lsof -i :8080
   
   # Change port in .env
   BACKEND_PORT=9000
   ```

### Validation
```bash
# Validate environment configuration
./scripts/validate-env.sh
```

## Best Practices

1. **Use descriptive variable names** (e.g., `BACKEND_PORT` instead of `PORT`)
2. **Group related variables** with clear sections
3. **Provide sensible defaults** for all variables
4. **Document all variables** in README files
5. **Use environment-specific files** for different deployments
6. **Never commit sensitive data** to version control

## Migration from Old Configuration

If you were using the old backend-only environment configuration:

1. **Backup your old configuration:**
   ```bash
   cp backend/.env backend/.env.backup
   ```

2. **Create new root-level configuration:**
   ```bash
   ./scripts/setup-env.sh
   ```

3. **Migrate your settings:**
   - `PORT` → `BACKEND_PORT`
   - `HOST` → `BACKEND_HOST`
   - Add frontend configuration as needed

4. **Test the new configuration:**
   ```bash
   docker-compose -f docker-compose.dev.yml up
   ``` 