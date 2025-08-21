# 🔐 Frontend Authentication

This document describes the frontend authentication implementation using SvelteKit and JWT tokens.

## 🏗️ Architecture

### **Authentication Flow**
1. **Login Page** (`/login`) - User enters credentials
2. **JWT Token Storage** - Token stored in localStorage
3. **Protected Routes** - All routes except `/login` require authentication
4. **Automatic Redirects** - Unauthenticated users redirected to login
5. **Token Validation** - API calls include JWT token in Authorization header

### **Key Components**

#### **Auth Store** (`src/lib/auth.ts`)
- Manages authentication state using Svelte stores
- Handles login/logout operations
- Provides authenticated API fetch function
- Persists authentication in localStorage

#### **Layout** (`src/routes/+layout.svelte`)
- Protects all routes except `/login`
- Shows loading states during authentication checks
- Provides navigation header with logout button
- Handles automatic redirects

#### **Login Page** (`src/routes/login/+page.svelte`)
- Modern, responsive login form
- Shows demo credentials for testing
- Handles login errors and loading states
- Redirects to dashboard on successful login

#### **Transactions Component** (`src/lib/Transactions.svelte`)
- Uses authenticated API calls
- Handles authentication errors
- Automatically redirects to login if token expires

## 🚀 Features

### **Security**
- JWT tokens with 24-hour expiration
- Automatic token validation on API calls
- Secure localStorage storage
- Automatic logout on token expiration

### **User Experience**
- Persistent login sessions
- Loading states and error handling
- Responsive design with Tailwind CSS
- Smooth transitions and animations

### **Developer Experience**
- TypeScript support
- Clean separation of concerns
- Reusable authentication utilities
- Easy to extend and modify

## 🧪 Testing

### **Demo Credentials**
- **Email**: `demo@example.com`
- **Password**: `demo123`

### **Testing Flow**
1. Navigate to `http://localhost:5173`
2. You'll be redirected to `/login`
3. Enter demo credentials
4. You'll be redirected to the dashboard
5. View your transactions
6. Use the logout button to sign out

## 🔧 Configuration

### **Environment Variables**
```env
VITE_API_URL=http://localhost:6060
```

### **API Endpoints Used**
- `POST /auth/login` - User authentication
- `GET /transactions` - Fetch user transactions (protected)

## 📁 File Structure

```
src/
├── lib/
│   ├── auth.ts              # Authentication store and utilities
│   └── Transactions.svelte  # Updated to use authenticated API
├── routes/
│   ├── +layout.svelte       # Main layout with auth protection
│   ├── +page.svelte         # Dashboard page
│   ├── +page.ts             # Page load function
│   └── login/
│       ├── +page.svelte     # Login form
│       └── +page.ts         # Login page load function
```

## 🎨 UI Components

### **Login Form**
- Clean, modern design
- Email and password fields
- Loading spinner during authentication
- Error message display
- Demo credentials display

### **Navigation Header**
- App title and branding
- User email display
- Logout button with icon
- Responsive design

### **Dashboard**
- Transaction summary cards
- Transaction list with hover effects
- Loading and empty states
- Error handling with retry buttons

## 🔄 State Management

### **Auth State**
```typescript
interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  loading: boolean;
}
```

### **Store Methods**
- `auth.login(email, password)` - Authenticate user
- `auth.logout()` - Sign out user
- `auth.authenticatedFetch(url, options)` - Make authenticated API calls
- `auth.getAuthHeaders()` - Get Authorization headers

## 🚨 Error Handling

### **Authentication Errors**
- Invalid credentials
- Network errors
- Token expiration
- Server errors

### **User Feedback**
- Loading spinners
- Error messages
- Success redirects
- Automatic retry options

## 🔮 Future Enhancements

### **Planned Features**
- User registration page
- Password reset functionality
- Remember me option
- Multi-factor authentication
- Session management
- Profile management

### **Security Improvements**
- Token refresh mechanism
- CSRF protection
- Rate limiting
- Secure cookie storage
- Audit logging 