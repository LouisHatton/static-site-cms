import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

// If hot reload isn't working in wsl run:
// export CHOKIDAR_USEPOLLING=true
export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	server: {
		host: true,
		port: 5173,
		watch: {
			usePolling: true
		}
	}
});
