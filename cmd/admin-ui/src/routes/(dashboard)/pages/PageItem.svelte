<script lang="ts">
	import type { SiteRoute } from '$lib/pages/managePage';
	import { Button, Tooltip } from 'flowbite-svelte';

	export let route: SiteRoute;
	let online = true;

	function onEdit() {
		console.log(`edit page ${route.path}`);
		return;
	}
</script>

<div
	class="bg-slate-100  py-4 pl-4 pr-8 border border-transparent rounded-md hover:cursor-pointer transition-all duration-200 hover:border-blue-600 ease-in-out"
>
	<div class="flex flex-row justify-between items-center">
		<div class="flex items-center gap-x-4">
			<div>
				{#if online}
					<Tooltip triggeredBy="[id='svg-online']">Page is online</Tooltip>
					<svg
						id="svg-online"
						clip-rule="evenodd"
						fill-rule="evenodd"
						stroke-linejoin="round"
						stroke-miterlimit="2"
						viewBox="0 0 24 24"
						class="fill-current text-green-600 w-6"
						xmlns="http://www.w3.org/2000/svg"
						><path
							d="m11.998 2.005c5.517 0 9.997 4.48 9.997 9.997 0 5.518-4.48 9.998-9.997 9.998-5.518 0-9.998-4.48-9.998-9.998 0-5.517 4.48-9.997 9.998-9.997zm0 1.5c-4.69 0-8.498 3.807-8.498 8.497s3.808 8.498 8.498 8.498 8.497-3.808 8.497-8.498-3.807-8.497-8.497-8.497zm-5.049 8.886 3.851 3.43c.142.128.321.19.499.19.202 0 .405-.081.552-.242l5.953-6.509c.131-.143.196-.323.196-.502 0-.41-.331-.747-.748-.747-.204 0-.405.082-.554.243l-5.453 5.962-3.298-2.938c-.144-.127-.321-.19-.499-.19-.415 0-.748.335-.748.746 0 .205.084.409.249.557z"
							fill-rule="nonzero"
						/></svg
					>
				{:else}
					<Tooltip triggeredBy="[id='svg-offline']">Page is offline</Tooltip>
					<svg
						id="svg-offline"
						clip-rule="evenodd"
						fill-rule="evenodd"
						stroke-linejoin="round"
						stroke-miterlimit="2"
						viewBox="0 0 24 24"
						class="fill-current text-red-600 w-6"
						xmlns="http://www.w3.org/2000/svg"
						><path
							d="m12.002 2.005c5.518 0 9.998 4.48 9.998 9.997 0 5.518-4.48 9.998-9.998 9.998-5.517 0-9.997-4.48-9.997-9.998 0-5.517 4.48-9.997 9.997-9.997zm0 1.5c-4.69 0-8.497 3.807-8.497 8.497s3.807 8.498 8.497 8.498 8.498-3.808 8.498-8.498-3.808-8.497-8.498-8.497zm0 7.425 2.717-2.718c.146-.146.339-.219.531-.219.404 0 .75.325.75.75 0 .193-.073.384-.219.531l-2.717 2.717 2.727 2.728c.147.147.22.339.22.531 0 .427-.349.75-.75.75-.192 0-.384-.073-.53-.219l-2.729-2.728-2.728 2.728c-.146.146-.338.219-.53.219-.401 0-.751-.323-.751-.75 0-.192.073-.384.22-.531l2.728-2.728-2.722-2.722c-.146-.147-.219-.338-.219-.531 0-.425.346-.749.75-.749.192 0 .385.073.531.219z"
							fill-rule="nonzero"
						/></svg
					>
				{/if}
			</div>
			<div>
				<h2 class="text-lg ">{route.path}</h2>
				<p class="text-sm mt-2">Last edited: 18/03/2023</p>
			</div>
		</div>

		<div><Button on:click={onEdit}>Edit</Button></div>
	</div>
</div>
