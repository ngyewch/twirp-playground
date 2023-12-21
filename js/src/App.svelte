<script lang="ts">
    import {onMount} from 'svelte';
    import {client} from 'twirpscript';
    import * as TestService from './service.pb';

    onMount(() => {
        client.baseURL = `${window.location.protocol}//${window.location.host}`;
    });

    let a: number | undefined = undefined;
    let b: number | undefined = undefined;
    let a1: number | undefined;
    let b1: number | undefined;
    let value: number | undefined;

    function submit() {
        if ((a == undefined) || (b == undefined)) {
            return;
        }
        a1 = undefined;
        b1 = undefined;
        value = undefined;
        TestService.Add({a: a, b: b})
            .then(res => {
                a1 = a;
                b1 = b;
                value = res.value;
            })
            .catch(error => {
                console.log('[ERROR]', error);
            });
    }
</script>

<main class="container">
    <form>
        <div class="grid">
            <input type="number" id="a" name="a" placeholder="a" required bind:value={a}>
            <input type="number" id="b" name="b" placeholder="b" required bind:value={b}>
            <button type="button" on:click={() => submit()}>Add</button>
        </div>
        <div class="grid">
            {#if a1 && b1 && value}
                {a1} + {b1} = {value}
            {/if}
        </div>
    </form>
</main>
