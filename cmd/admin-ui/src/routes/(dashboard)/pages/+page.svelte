<script>
	import { Button, Dropzone, Progressbar } from 'flowbite-svelte';

	let progressInt = 0;
	$: progress = progressInt.toString();

	const onUpload = () => {
		setTimeout(() => {
			progressInt = 25;
			setTimeout(() => {
				progressInt = 75;
				setTimeout(() => {
					progressInt = 100;
				}, 2000);
			}, 1000);
		}, 500);
	};
</script>

<h1 class="text-2xl mb-2">Upload new site</h1>

<p class="mb-6">This will override the current site</p>

<div class="max-w-lg">
	<Dropzone id="dropzone">
		<svg
			aria-hidden="true"
			class="mb-3 w-10 h-10 text-gray-400"
			fill="none"
			stroke="currentColor"
			viewBox="0 0 24 24"
			xmlns="http://www.w3.org/2000/svg"
			><path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
			/></svg
		>
		<p class="mb-2 text-sm text-gray-500 dark:text-gray-400">
			<span class="font-semibold">Click to upload</span> or drag and drop
		</p>
		<p class="text-xs text-gray-500 dark:text-gray-400">.ZIP files only</p>
	</Dropzone>
	<Progressbar class="my-6" bind:progress />
	<Button on:click={onUpload}>Upload</Button>
</div>
