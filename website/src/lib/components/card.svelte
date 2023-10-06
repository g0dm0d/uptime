<script lang="ts">
  import type { MonitorData } from "$lib/model";
    import Infocard from "./infocard.svelte";
  import Status from "./status.svelte";

  export let monitorData: MonitorData;

  let ticks = monitorData.ticks.slice()
</script>

{#key monitorData}
<div class="card bg-surface-800 overflow-hidden flex flex-col lg:flex-row p-4 rounded-xl m-5 w-full">
  <div>
    <div class="flex items-center flex-row mt-2">
      <div class="w-16 items-center text-center font-bold text-black h-6 rounded-full mr-2
                  {ticks[ticks.length-1].success === 1 ? 'bg-success-400' : 'bg-error-400'}">
          {#if ticks[ticks.length-1].success === 1}
          <p class="text-surface-900">online</p>
          {:else}
          <p class="text-surface-900">offline</p>
          {/if}
          <footer class="flex items-center">
            <p class="text-surface-300 text-sm mt-2">{monitorData.monitor.protocol}://{monitorData.monitor.address}</p>
            {#if monitorData.monitor.port}
            <span>
              <p class="text-surface-300 text-sm mt-2">:{monitorData.monitor.port}</p>
            </span>
            {/if}
          </footer>
      </div>
      <h3 class="text-white">
        {monitorData.monitor.hostname}
      </h3>
    </div>
  </div>
  <div class="ml-auto">
    <div class="mt-2 flex flex-row">
      {#each ticks as item, index}
        <Status item={item} index={index}/>
      {/each}

    </div>
    <footer class="justify-end flex items-center">
      <p class="text-surface-300 text-sm mt-2">Last 24 hours</p>
    </footer>
  </div>
</div>
{/key}
