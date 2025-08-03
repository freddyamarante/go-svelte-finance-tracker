# 🔐 Automated Authentication Setup

This project now includes automated authentication with a sample user that's created automatically when you start the application.

## 🚀 Quick Start

### Option 1: Use the setup script (Recommended)
```bash
./scripts/setup-auth.sh
```

This script will:
- Create the data directory for persistent storage
- Start the Docker containers
- Test the sample user login
- Show you the credentials

### Option 2: Manual setup
```bash
# Create data directory
mkdir -p data

# Start the application
docker-compose -f docker-compose.dev.yml up --build
```

## 📋 Sample User Credentials

The following user is automatically created when the database is first initialized:

- **Email**: `demo@example.com`
- **Password**: `demo123`

## 🧪 Testing in Postman

### 1. Login with sample user
```
POST http://localhost:8080/auth/login
Content-Type: application/json

{
  "email": "demo@example.com",
  "password": "demo123"
}
```

### 2. Use the JWT token for protected endpoints
```
GET http://localhost:8080/transactions
Authorization: Bearer <your-jwt-token>
```

## 🔧 How it works

1. **Persistent Database**: The database is stored in `./data/finance.db` and persists between container restarts
2. **Automatic Sample Data**: When the database is first created, a sample user and transactions are automatically added
3. **JWT Authentication**: The backend uses JWT tokens for stateless authentication
4. **Protected Routes**: All transaction endpoints require authentication

## 📁 File Structure

```
├── data/                    # Persistent database storage
│   └── finance.db          # SQLite database file
├── scripts/
│   └── setup-auth.sh       # Automated setup script
├── docker-compose.dev.yml  # Updated with persistent database
└── .gitignore              # Updated to exclude data directory
```

## 🔄 Database Persistence

- **Development**: Database file is stored in `./data/finance.db`
- **Production**: Uses Docker named volume `finance_db`
- **Sample Data**: Automatically created on first run
- **User Data**: Persists between container restarts

## 🛠️ Troubleshooting

### Sample user not working?
1. Check if the database exists: `ls -la data/`
2. View backend logs: `docker-compose -f docker-compose.dev.yml logs backend`
3. Restart containers: `docker-compose -f docker-compose.dev.yml restart`

### Database reset
To start fresh with a new database:
```bash
rm -rf data/
docker-compose -f docker-compose.dev.yml down
docker-compose -f docker-compose.dev.yml up --build
```

## 🔒 Security Notes

- The sample user is for development/testing only
- In production, users should register their own accounts
- JWT tokens expire after 24 hours
- Passwords are hashed using bcrypt
- All sensitive endpoints require authentication

## 📚 API Endpoints

### Public Endpoints
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `POST /auth/logout` - Logout user

### Protected Endpoints (require JWT token)
- `GET /auth/profile` - Get user profile
- `PUT /auth/profile` - Update user profile
- `PUT /auth/change-password` - Change password
- `GET /transactions` - Get user transactions
- `POST /transactions` - Create new transaction 