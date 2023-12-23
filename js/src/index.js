import App from './App.svelte';
import {client} from 'twirpscript';

import '@picocss/pico/css/pico.css';

client.baseURL = `${window.location.protocol}//${window.location.host}`;
client.on('error', (context, error) => {
    console.log('ERROR', error);
});

let app = new App({
    target: document.body,
});

export default app;
