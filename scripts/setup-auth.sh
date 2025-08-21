#!/bin/bash

echo "🔐 Setting up automated authentication for Finance Tracker"
echo "========================================================"

# Create data directory if it doesn't exist
if [ ! -d "data" ]; then
    echo "📁 Creating data directory for persistent database..."
    mkdir -p data
fi

echo "✅ Data directory ready"

echo ""
echo "🚀 Starting the application with automated sample user..."
echo ""
echo "📋 Sample user credentials will be automatically created:"
echo "   Email: demo@example.com"
echo "   Password: demo123"
echo ""
echo "🔄 Restarting Docker containers..."

# Stop existing containers
docker-compose -f docker-compose.dev.yml down

# Start containers with fresh build
docker-compose -f docker-compose.dev.yml up --build -d

echo ""
echo "⏳ Waiting for backend to start..."
sleep 5

echo ""
echo "🧪 Testing sample user login..."

# Test the sample user login
response=$(curl -s -X POST http://localhost:6060/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"demo@example.com","password":"demo123"}')

if echo "$response" | grep -q "token"; then
    echo "✅ Sample user login successful!"
    echo "🎉 Authentication is now automated!"
    echo ""
    echo "📝 You can now use these credentials in Postman:"
    echo "   POST http://localhost:6060/auth/login"
    echo "   Body: {\"email\":\"demo@example.com\",\"password\":\"demo123\"}"
else
    echo "❌ Sample user login failed. Please check the logs:"
    echo "   docker-compose -f docker-compose.dev.yml logs backend"
fi

echo ""
echo "🌐 Backend is running at: http://localhost:6060"
echo "🎨 Frontend is running at: http://localhost:5173" 