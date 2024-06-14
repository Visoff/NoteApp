<script>
    /** @type {import('$lib/global.js').FSEntry|undefined} */
    export let tree = undefined;
    /** @type {string} */
    export let path = "";
    let open = false;
</script>

{#if tree}
    {#if tree.type == "file"}
        <button on:click={() => {alert(path+"/"+tree.name)}}>{tree.name}</button>
    {:else if tree.type == "dir"}
        <span>
            <button on:click={() => (open = !open)}>
                {open ? "-" : ">"} {tree.name}
            </button>
        </span>
        <ul class="pl-2">
            {#each tree.children as child}
                <li class="{open ? '' : 'hidden'}">
                    <svelte:self tree={child} path={path+"/"+tree.name} />
                </li>
            {/each}
        </ul>
    {/if}
{:else}
    {#if path == ""}
        <button>Open direcotry</button>
    {/if}
{/if}
