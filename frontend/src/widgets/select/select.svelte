<script lang="ts">
  let { value = $bindable(null), options = [] } = $props<{
    value?: string;
    "bind:value"?: string;
    options?: string[];
  }>();
  let active = $state(false);
</script>

<div class="relative bg-blue-500 w-fit">
  <div class="w-full" onclick={() => (active = !active)}>{value}</div>
  <div class="absolute bg-blue-500">
    {#if active}
      {#each options as el}
        <div>{el}</div>
      {/each}
    {/if}
  </div>
</div>
