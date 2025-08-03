<script lang="ts">
	import { onMount } from 'svelte';

	interface Transaction {
		id: string;
		description: string;
		amount: number;
		type: 'income' | 'expense';
	}

	// Get API URL from environment or use default
	const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

	let transactions: Transaction[] = [];
	let loading = true;
	let error: string | null = null;

	async function fetchTransactions() {
		try {
			loading = true;
			error = null;
			const response = await fetch(`${API_BASE_URL}/transactions`);

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			transactions = await response.json();
		} catch (err) {
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
			class="cursor-pointer rounded border-none bg-blue-500 px-4 py-2 font-medium text-white transition-colors duration-200 disabled:cursor-not-allowed disabled:bg-gray-400"
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
				class="mt-2 cursor-pointer rounded border-none bg-red-500 px-3 py-1 text-white"
				on:click={fetchTransactions}>Try Again</button
			>
		</div>
	{/if}

	{#if loading}
		<div class="p-10 text-center text-gray-600">
			<p>Loading transactions...</p>
		</div>
	{:else if transactions.length === 0}
		<div class="p-10 text-center text-gray-600">
			<p>No transactions found.</p>
		</div>
	{:else}
		<!-- Summary Cards -->
		<div class="mb-7 grid grid-cols-1 gap-5 md:grid-cols-3">
			<div class="rounded border border-gray-200 bg-white p-5 shadow">
				<h3 class="text-sm font-semibold uppercase text-gray-600">Total Income</h3>
				<p class="text-lg font-bold text-green-500">{formatCurrency(getTotalIncome())}</p>
			</div>
			<div class="rounded border border-gray-200 bg-white p-5 shadow">
				<h3 class="text-sm font-semibold uppercase text-gray-600">Total Expenses</h3>
				<p class="text-lg font-bold text-red-500">{formatCurrency(getTotalExpenses())}</p>
			</div>
			<div class="rounded border border-gray-200 bg-white p-5 shadow">
				<h3 class="text-sm font-semibold uppercase text-gray-600">Balance</h3>
				<p class="text-lg font-bold {getBalance() >= 0 ? 'text-green-500' : 'text-red-500'}">
					{formatCurrency(getBalance())}
				</p>
			</div>
		</div>

		<!-- Transactions List -->
		<div class="overflow-hidden rounded border border-gray-200 bg-white shadow">
			<h3 class="border-b border-gray-200 bg-gray-100 p-5 text-lg font-semibold text-gray-800">
				Recent Transactions
			</h3>
			<div class="overflow-y-auto">
				{#each transactions as transaction}
					<div class="flex items-center justify-between border-b border-gray-200 p-4">
						<div>
							<h4 class="text-base font-semibold text-gray-800">{transaction.description}</h4>
							<p class="text-sm text-gray-500">ID: {transaction.id}</p>
						</div>
						<div class="flex flex-col items-end gap-1">
							<span
								class="text-base font-semibold {transaction.type === 'income'
									? 'text-green-500'
									: 'text-red-500'}"
							>
								{transaction.type === 'income' ? '+' : '-'}{formatCurrency(transaction.amount)}
							</span>
							<span class="text-sm text-gray-500">{transaction.type}</span>
						</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>
