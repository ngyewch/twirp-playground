import {defineConfig} from 'vite';
import {svelte} from '@sveltejs/vite-plugin-svelte';

export default defineConfig({
    plugins: [
        svelte(),
    ],
    build: {
        sourcemap: true,
    },
    server: {
        proxy: {
            '^/twirp/.*': {
                target: 'http://localhost:8080',
            },
        },
    },
});
