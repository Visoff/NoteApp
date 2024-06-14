<script>
    import { fs } from "$lib/global.js";
    import Fstree from "$components/fstree.svelte";
    import Fsactions from "$components/fsactions.svelte";
</script>

<div class="w-full h-full max-w-full flex flex-col">
    <div class="h-10 bg-crust border-b-2 border-surface1"><Fsactions /></div>
    <div class="flex-1">
        {#if $fs != undefined && $fs.tree != undefined}
            <Fstree tree={$fs.tree} />
        {:else}
            <div class="w-full flex flex-col align-center">
                <button>Open Directory</button>
                <button>Open Remote Vault</button>
            </div>
        {/if}
    </div>
    <div class="h-10 flex flex-row justify-between">
        <div></div>
        <span>{$fs ? $fs.path : ""}</span>
    </div>
</div>
