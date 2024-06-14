<script>
    /** @typedef {string} Tab */
    /** @type {{name:Tab, icon: string}[]} */
    let tabs = [{name:"Files", icon: "#"}];
    /** @type {Record<Tab, {icon: string, popover: string, action: () => void}[]>} */
    let actions = {"Files": []};
    /** @type {Number} */
    export let current_tab = 0;
</script>

<ul>
    {#each tabs as tab, i}
        <li>
            <button on:click={() => (current_tab = i)}>
                <img src={tab.icon} alt={tab.name} />
                <span>{tab.name}</span>
            </button>
        </li>
    {/each}
    {#each actions[tabs[current_tab].name] as action}
        <li>
            <button on:click={action.action}>
                <img src={action.icon} alt={action.popover} />
                <span>{action.popover}</span>
            </button>
        </li>
    {/each}
</ul>
