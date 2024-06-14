<script>
    /** @type {{icon: string, popup: string, action: () => void}[]} */
    const actions = [
    ]
</script>

<ul>
    {#each actions as action}
        <li>
            <button on:click={action.action}>
                <img src={action.icon} alt={action.popup} />
                <span>{action.popup}</span>
            </button>
        </li>
    {/each}
</ul>
