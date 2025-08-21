<script lang="ts">
	import { onMount } from 'svelte';
	import { auth } from '$lib/auth';
	import { goto } from '$app/navigation';

	interface Transaction {
		id: string;
		description: string;
		amount: number;
		type: 'income' | 'expense';
	}

	// Get API URL from environment or use default
	const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:6060';

	let transactions: Transaction[] = [];
	let loading = true;
	let error: string | null = null;

	async function fetchTransactions() {
		try {
			loading = true;
			error = null;
			
			const response = await auth.authenticatedFetch(`${API_BASE_URL}/transactions`);

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			transactions = await response.json();
		} catch (err) {
			if (err instanceof Error && err.message === 'Authentication expired') {
				// Redirect to login if authentication expired
				goto('/login');
				return;
			}
			error = err instanceof Error ? err.message : 'Failed to fetch transactions';
			console.error('Error fetching transactions:', err);
		} finally {
			loading = false;
		}
	}

	function formatCurrency(amount: number): string {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD'
		}).format(amount);
	}

	function getTotalIncome(): number {
		return transactions.filter((t) => t.type === 'income').reduce((sum, t) => sum + t.amount, 0);
	}

	function getTotalExpenses(): number {
		return transactions.filter((t) => t.type === 'expense').reduce((sum, t) => sum + t.amount, 0);
	}

	function getBalance(): number {
		return getTotalIncome() - getTotalExpenses();
	}

	onMount(() => {
		fetchTransactions();
	});
</script>

<div class="mx-auto max-w-5xl p-5 font-sans">
	<div class="mb-7 flex items-center justify-between">
		<h2 class="text-2xl font-bold text-gray-800">Personal Finance Tracker</h2>
		<button
			class="cursor-pointer rounded border-none bg-blue-500 px-4 py-2 font-medium text-white transition-colors duration-200 disabled:cursor-not-allowed disabled:bg-gray-400 hover:bg-blue-600"
			on:click={fetchTransactions}
			disabled={loading}
		>
			{loading ? 'Loading...' : 'Refresh'}
		</button>
	</div>

	{#if error}
		<div class="relative mb-4 rounded border border-red-300 bg-red-100 px-4 py-3 text-red-700">
			<p>Error: {error}</p>
			<button
				class="mt-2 cursor-pointer rounded border-none bg-red-500 px-3 py-1 text-white hover:bg-red-600 transition-colors duration-200"
				on:click={fetchTransactions}>Try Again</button
			>
		</div>
	{/if}

	{#if loading}
		<div class="p-10 text-center text-gray-600">
			<div class="inline-flex items-center">
				<div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600 mr-3"></div>
				<p>Loading transactions...</p>
			</div>
		</div>
	{:else if transactions.length === 0}
		<div class="p-10 text-center text-gray-600">
			<div class="max-w-md mx-auto">
				<svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
				</svg>
				<p class="text-lg font-medium mb-2">No transactions found</p>
				<p class="text-sm text-gray-500">Your transactions will appear here once you add some.</p>
			</div>
		</div>
	{:else}
		<!-- Summary Cards -->
		<div class="mb-7 grid grid-cols-1 gap-5 md:grid-cols-3">
			<div class="rounded-lg border border-gray-200 bg-white p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
				<h3 class="text-sm font-semibold uppercase text-gray-600 tracking-wide">Total Income</h3>
				<p class="text-2xl font-bold text-green-600">{formatCurrency(getTotalIncome())}</p>
			</div>
			<div class="rounded-lg border border-gray-200 bg-white p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
				<h3 class="text-sm font-semibold uppercase text-gray-600 tracking-wide">Total Expenses</h3>
				<p class="text-2xl font-bold text-red-600">{formatCurrency(getTotalExpenses())}</p>
			</div>
			<div class="rounded-lg border border-gray-200 bg-white p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
				<h3 class="text-sm font-semibold uppercase text-gray-600 tracking-wide">Balance</h3>
				<p class="text-2xl font-bold {getBalance() >= 0 ? 'text-green-600' : 'text-red-600'}">
					{formatCurrency(getBalance())}
				</p>
			</div>
		</div>

		<!-- Transactions List -->
		<div class="overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm">
			<h3 class="border-b border-gray-200 bg-gray-50 p-6 text-lg font-semibold text-gray-800">
				Recent Transactions
			</h3>
			<div class="overflow-y-auto">
				{#each transactions as transaction}
					<div class="flex items-center justify-between border-b border-gray-100 p-6 hover:bg-gray-50 transition-colors duration-150 last:border-b-0">
						<div class="flex-1">
							<h4 class="text-base font-semibold text-gray-800">{transaction.description}</h4>
							<p class="text-sm text-gray-500 mt-1">ID: {transaction.id}</p>
						</div>
						<div class="flex flex-col items-end gap-1">
							<span
								class="text-lg font-semibold {transaction.type === 'income'
									? 'text-green-600'
									: 'text-red-600'}"
							>
								{transaction.type === 'income' ? '+' : '-'}{formatCurrency(transaction.amount)}
							</span>
							<span class="text-sm text-gray-500 capitalize">{transaction.type}</span>
						</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>
