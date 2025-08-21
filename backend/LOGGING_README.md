# Logging Configuration

This backend now supports configurable logging through environment variables. This allows you to control what gets logged in different environments.

## Environment Variables

### Logging Controls

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `LOG_REQUESTS` | boolean | `true` | Enable/disable HTTP request logging |
| `LOG_SQL` | boolean | `false` | Enable/disable SQL query logging |
| `LOG_CORS` | boolean | `false` | Enable/disable CORS configuration logging |

### Server Configuration

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `PORT` | string | `6060` | Server port |
| `HOST` | string | `localhost` | Server host |
| `ENV` | string | `development` | Environment (development/production) |

### Database Configuration

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `DB_PATH` | string | `finance.db` | SQLite database file path |

### CORS Configuration

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `CORS_ALLOW_ORIGINS` | string | `*` | Comma-separated list of allowed origins |
| `CORS_ALLOW_METHODS` | string | `GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS` | Comma-separated list of allowed methods |
| `CORS_ALLOW_HEADERS` | string | `Origin,Content-Length,Content-Type,Accept,Authorization,X-Requested-With` | Comma-separated list of allowed headers |
| `CORS_MAX_AGE` | integer | `43200` | CORS preflight cache time in seconds |

## Usage Examples

### Minimal Logging (Production)
```bash
export LOG_REQUESTS=false
export LOG_SQL=false
export LOG_CORS=false
export ENV=production
go run .
```

### Verbose Logging (Debug)
```bash
export LOG_REQUESTS=true
export LOG_SQL=true
export LOG_CORS=true
export ENV=development
go run .
```

### Using the Run Script
```bash
# Default (development with request logging only)
./scripts/run.sh

# Production mode
ENV=production LOG_REQUESTS=false ./scripts/run.sh

# Debug mode with all logging
LOG_REQUESTS=true LOG_SQL=true LOG_CORS=true ./scripts/run.sh
```

### Docker Environment

#### Development
```bash
docker-compose -f docker-compose.dev.yml up
```
- `LOG_REQUESTS=true` - Shows HTTP requests
- `LOG_SQL=false` - No SQL logging
- `LOG_CORS=false` - No CORS logging

#### Production
```bash
docker-compose up
```
- `LOG_REQUESTS=false` - No HTTP request logging
- `LOG_SQL=false` - No SQL logging
- `LOG_CORS=false` - No CORS logging

## What Gets Logged

### When `LOG_REQUESTS=true`
- HTTP method and path
- Response status code
- Response time
- Request size
- User agent

### When `LOG_SQL=true`
- SQL queries executed by GORM
- Query parameters
- Query execution time

### When `LOG_CORS=true`
- CORS configuration at startup
- CORS policy details

## Performance Impact

- **Request logging**: Minimal impact, useful for monitoring
- **SQL logging**: Moderate impact, useful for debugging queries
- **CORS logging**: Minimal impact, only logs at startup

## Best Practices

### Development
```bash
LOG_REQUESTS=true LOG_SQL=false LOG_CORS=false
```
- Enable request logging to see API usage
- Disable SQL logging unless debugging database issues
- Disable CORS logging unless debugging CORS issues

### Production
```bash
LOG_REQUESTS=false LOG_SQL=false LOG_CORS=false
```
- Disable all logging for maximum performance
- Use external monitoring tools instead

### Debugging
```bash
LOG_REQUESTS=true LOG_SQL=true LOG_CORS=true
```
- Enable all logging when debugging issues
- Remember to disable in production

## Environment File

Copy `env.example` to `.env` and modify as needed:

```bash
cp env.example .env
```

Then edit `.env` with your preferred settings.

## Docker Override

You can override any environment variable when running Docker:

```bash
# Override specific variables
LOG_REQUESTS=false docker-compose up

# Use a different environment file
docker-compose --env-file .env.production up
``` 