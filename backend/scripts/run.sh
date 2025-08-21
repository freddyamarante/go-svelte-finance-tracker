#!/bin/bash

# Default values
LOG_REQUESTS=${LOG_REQUESTS:-true}
LOG_SQL=${LOG_SQL:-false}
LOG_CORS=${LOG_CORS:-false}
ENV=${ENV:-development}
BACKEND_PORT=${BACKEND_PORT:-6060}
BACKEND_HOST=${BACKEND_HOST:-localhost}

echo "Starting backend with configuration:"
echo "  Environment: $ENV"
echo "  Port: $BACKEND_PORT"
echo "  Host: $BACKEND_HOST"
echo "  Log Requests: $LOG_REQUESTS"
echo "  Log SQL: $LOG_SQL"
echo "  Log CORS: $LOG_CORS"
echo ""

export LOG_REQUESTS
export LOG_SQL
export LOG_CORS
export ENV
export BACKEND_PORT
export BACKEND_HOST

go run . 