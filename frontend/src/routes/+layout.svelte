<script lang="ts">
	import '../app.css';
	import { auth } from '$lib/auth';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	
	let { children } = $props();
	
	// Use derived to subscribe to auth store
	let authState = $derived($auth);

	// Handle authentication state changes
	$effect(() => {
		if (authState && !authState.loading && !authState.isAuthenticated && $page.url.pathname !== '/login') {
			goto('/login');
		}
	});

	function handleLogout() {
		auth.logout();
		goto('/login');
	}
</script>

{#if !authState || authState.loading}
	<div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
		<div class="text-center">
			<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto"></div>
			<p class="mt-4 text-gray-600">Loading...</p>
		</div>
	</div>
{:else if !authState.isAuthenticated && $page.url.pathname !== '/login'}
	<div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
		<div class="text-center">
			<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto"></div>
			<p class="mt-4 text-gray-600">Redirecting to login...</p>
		</div>
	</div>
{:else}
	{#if authState.isAuthenticated && $page.url.pathname !== '/login'}
		<!-- Navigation header for authenticated users -->
		<header class="bg-white shadow-sm border-b border-gray-200">
			<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div class="flex justify-between items-center h-16">
					<div class="flex items-center">
						<h1 class="text-xl font-semibold text-gray-900">Finance Tracker</h1>
					</div>
					<div class="flex items-center space-x-4">
						{#if authState.user}
							<span class="text-sm text-gray-700">Welcome, {authState.user.email}</span>
						{/if}
						<button
							on:click={handleLogout}
							class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors duration-200"
						>
							<svg class="h-4 w-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
							</svg>
							Logout
						</button>
					</div>
				</div>
			</div>
		</header>
	{/if}

	<main class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
		{@render children()}
	</main>
{/if}
