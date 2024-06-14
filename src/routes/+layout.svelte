<script>
    import "../app.css";
    import "../wasm/wasm_exec.js";

    import { init } from "$lib/wasm.js";
    if (typeof window !== "undefined") {
        init();
    }
</script>

<slot />
