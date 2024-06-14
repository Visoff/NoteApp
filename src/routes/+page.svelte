<script>
    import Actiontab from "$containers/actiontab.svelte";
    import Fstab from "$containers/fstab.svelte";
    import Editor from "$containers/editor.svelte";
</script>

<div class="text-text h-full w-full max-w-full flex flex-row gap-[2px] bg-surface1">
    <div class="w-12 bg-base"><Actiontab /></div>
    <div class="flex-1 bg-base"><Fstab /></div>
    <div class="flex-[4] bg-base"><Editor /></div>
</div>
