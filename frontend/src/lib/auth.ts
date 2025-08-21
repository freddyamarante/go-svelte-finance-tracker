import { writable } from 'svelte/store';

export interface User {
	id: string;
	email: string;
	created_at: string;
	updated_at: string;
}

export interface AuthState {
	user: User | null;
	token: string | null;
	isAuthenticated: boolean;
	loading: boolean;
}

// Create the auth store
const createAuthStore = () => {
	const { subscribe, set, update } = writable<AuthState>({
		user: null,
		token: null,
		isAuthenticated: false,
		loading: false
	});

	// Get API URL from environment
	const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:6060';

	// Initialize auth state from localStorage
	const initAuth = () => {
		const token = localStorage.getItem('auth_token');
		const userStr = localStorage.getItem('auth_user');
		
		if (token && userStr) {
			try {
				const user = JSON.parse(userStr);
				set({
					user,
					token,
					isAuthenticated: true,
					loading: false
				});
			} catch (error) {
				console.error('Error parsing stored user data:', error);
				logout();
			}
		} else {
			// No stored auth data, set loading to false
			set({
				user: null,
				token: null,
				isAuthenticated: false,
				loading: false
			});
		}
	};

	// Login function
	const login = async (email: string, password: string) => {
		update(state => ({ ...state, loading: true }));

		try {
			const response = await fetch(`${API_BASE_URL}/auth/login`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			if (!response.ok) {
				const errorData = await response.json();
				throw new Error(errorData.error || 'Login failed');
			}

			const data = await response.json();
			
			// Store in localStorage
			localStorage.setItem('auth_token', data.token);
			localStorage.setItem('auth_user', JSON.stringify(data.user));

			// Update store
			set({
				user: data.user,
				token: data.token,
				isAuthenticated: true,
				loading: false
			});

			return { success: true };
		} catch (error) {
			update(state => ({ ...state, loading: false }));
			return { 
				success: false, 
				error: error instanceof Error ? error.message : 'Login failed' 
			};
		}
	};

	// Logout function
	const logout = () => {
		localStorage.removeItem('auth_token');
		localStorage.removeItem('auth_user');
		set({
			user: null,
			token: null,
			isAuthenticated: false,
			loading: false
		});
	};

	// Get auth headers for API calls
	const getAuthHeaders = () => {
		const token = localStorage.getItem('auth_token');
		return token ? { 'Authorization': `Bearer ${token}` } : {};
	};

	// Make authenticated API call
	const authenticatedFetch = async (url: string, options: RequestInit = {}) => {
		const token = localStorage.getItem('auth_token');
		
		if (!token) {
			throw new Error('No authentication token');
		}

		const response = await fetch(url, {
			...options,
			headers: {
				...options.headers,
				'Authorization': `Bearer ${token}`
			}
		});

		// If token is invalid, logout
		if (response.status === 401) {
			logout();
			throw new Error('Authentication expired');
		}

		return response;
	};

	return {
		subscribe,
		login,
		logout,
		initAuth,
		getAuthHeaders,
		authenticatedFetch
	};
};

export const auth = createAuthStore();

// Initialize auth on app start
if (typeof window !== 'undefined') {
	auth.initAuth();
} 