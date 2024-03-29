<script lang="ts">
	import { Alert, Button, Dropzone, Modal, Spinner } from 'flowbite-svelte';
	import PageItem from './PageItem.svelte';
	import type { PageData } from './$types';
	import { invalidateAll } from '$app/navigation';

	export let data: PageData;
	let fileupload: FileList;
	let uploadSuccess: boolean | null = null;
	let uploading = false;
	let newFullSiteModal = false;

	const onUpload = async () => {
		uploading = true;
		let file = fileupload.item(fileupload.length - 1);
		if (file == null) return;
		if (file.type !== 'application/x-zip-compressed') return;

		var data = new FormData();
		data.append('file', file, file.name);
		try {
			let response = await fetch('/api/upload/full-site', {
				method: 'POST',
				body: data
			});
			if (response.status === 200) {
				uploadSuccess = true;
			} else {
				uploadSuccess = false;
			}
		} catch {
			uploadSuccess = false;
		} finally {
			uploading = false;
			newFullSiteModal = false;
			invalidateAll();
		}
	};
</script>

<div class="flex flex-row justify-between items-center mb-4">
	<div>
		<h1 class="text-2xl mb-2">Your Site</h1>
		<p>View and edit the pages for your site</p>
	</div>
	<div>
		<Button on:click={() => (newFullSiteModal = true)}>Upload New Site</Button>
	</div>
</div>

{#if uploadSuccess !== null}
	<Alert color={uploadSuccess ? 'green' : 'red'} class="mb-4">
		<span slot="icon"
			><svg
				aria-hidden="true"
				class="w-5 h-5"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					fill-rule="evenodd"
					d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
					clip-rule="evenodd"
				/></svg
			>
		</span>
		{uploadSuccess ? 'File was uploaded successfully' : 'There was an error uploading the file'}
	</Alert>
{/if}

<Modal title="Upload New Site" bind:open={newFullSiteModal} autoclose>
	<p>This will permanently override your current site and cannot be restored</p>
	<Dropzone bind:files={fileupload} id="dropzone">
		{#if !uploading}
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
		{:else}
			<div class="text-center"><Spinner /></div>
		{/if}
	</Dropzone>
	<p class="">
		File Uploaded: <span class="font-semibold font-mono"
			>{fileupload?.length != 0
				? fileupload?.item(fileupload?.length - 1)?.name ?? 'No file uploaded'
				: 'No file uploaded'}</span
		>
	</p>
	<div class="flex flex-row-reverse gap-x-2">
		<Button on:click={onUpload}>Upload</Button>
		<Button color="alternative">Close</Button>
	</div>
</Modal>

<div class="flex flex-col gap-y-6 mt-10">
	{#each data.siteStructure.routes as route}
		{#if route.files.includes('index.html')}
			<PageItem {route} />
		{/if}
	{/each}
</div>
