<script>
    import { signal_to_writable, signal_setter } from "$lib/wasm";
    import { onMount } from "svelte";

    let content = signal_to_writable("text_field_content");
    let keydown = signal_setter("keypress");
    /** @type {HTMLElement} */
    let editor;
    onMount(() => {
        window.onkeydown = (event) => {
            if (keydown == undefined) {
                keydown = signal_setter("keypress");
                return
            }
            event.preventDefault();
            keydown(event);
        }
    })
    $: {
        if (editor != undefined) {
            editor.innerHTML = $content;
        }
    }
</script>

<div bind:this={editor}>
</div>
