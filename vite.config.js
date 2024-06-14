import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
    resolve: {
        alias: {
            $containers: '/src/containers',
            $components: '/src/components',
        }
    }
});
