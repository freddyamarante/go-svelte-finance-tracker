#!/bin/bash

# Setup environment file from example
if [ ! -f .env ]; then
    echo "Creating .env file from env.example..."
    cp env.example .env
    echo "✅ .env file created successfully!"
    echo "📝 Edit .env file to customize your configuration"
else
    echo "⚠️  .env file already exists"
    echo "📝 Edit .env file to customize your configuration"
fi

echo ""
echo "Available environment variables:"
echo "  BACKEND_PORT     - Backend server port (default: 8080)"
echo "  FRONTEND_PORT    - Frontend server port (default: 5173)"
echo "  LOG_REQUESTS     - Enable HTTP request logging (default: true)"
echo "  LOG_SQL          - Enable SQL query logging (default: false)"
echo "  LOG_CORS         - Enable CORS configuration logging (default: false)"
echo "  VITE_API_URL     - Backend API URL for frontend (default: http://localhost:8080)"
echo "" 