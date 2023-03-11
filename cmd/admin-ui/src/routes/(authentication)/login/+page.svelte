<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button, Input, Label, Spinner } from 'flowbite-svelte';

	let loggingIn: boolean = false;

	const toggleLogin = () => {
		loggingIn = true;
		setTimeout(() => {
			loggingIn = false;
			goto('/');
		}, 3000);
	};
</script>

<div class="flex flex-row items-center justify-center min-h-screen bg-slate-50 w-full">
	<div class="bg-white shadow-md rounded-md px-8 py-10 max-w-lg w-full">
		<h1 class="text-3xl font-bold text-zinc-800 text-center mb-6">Login</h1>
		<div class="space-y-4 mb-8">
			<div>
				<Label class="mb-2">Your Email:</Label>
				<Input placeholder="name@example.com" />
			</div>
			<div>
				<Label class="mb-2">Your Password:</Label>
				<Input placeholder="••••••••" type="password" />
			</div>
		</div>
		<Button on:click={toggleLogin} bind:disabled={loggingIn} class="disabled:cursor-auto w-full">
			{#if loggingIn}
				<Spinner class="mr-3" size="4" />
			{/if}
			Login
		</Button>
	</div>
</div>
