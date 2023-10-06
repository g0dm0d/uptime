<script lang="ts">
  import type { Monitor, Tick } from '$lib';
  import Infocard from '$lib/components/infocard.svelte';
  import { groupTicks } from '$lib/utils/groupTicks.js';

  export let data;

  let heartbeat: Tick[] = data.heartbeat.reverse()
  let monitor: Monitor = data.monitor

  function calculateAveragePing(ticks: Tick[]): number {
    let totalPing = 0;

    for (const tick of ticks) {
      totalPing += tick.ping;
    }

    return totalPing / heartbeat.length;
  }

  function calculateAverageUptime(ticks: Tick[]): number {
    let totalUptime = 0;

    for (const tick of ticks) {
      totalUptime += tick.success;
    }

    return (ticks.length - totalUptime) * 100;
  }
    
  const avg_ping = calculateAveragePing(heartbeat);
  const avg_uptime = (heartbeat.filter((tick) => tick.success === 1).length / heartbeat.length) * 100;
</script>


<div class="container mx-auto block justify-center card bg-surface-800 overflow-hidden lg:flex-row p-4 rounded-xl m-5 w-full">
  <div class="grid grid-cols-3 content-around">
    <div class="flex justify-center">
      <div>
        <h1>AVG PING</h1>
        <p>{avg_ping.toFixed(2)}ms</p>
      </div>
    </div>
    <div class="flex justify-center">
      <div>
        <h1 class="text-xl">{monitor.address}</h1>
      </div>
    </div>
    <div class="flex justify-center">
      <div>
        <h1>UPTIME</h1>
        <p>{avg_uptime.toFixed(2)}%</p>
      </div>
    </div>
  </div>
  <Infocard ticks={heartbeat}/>
</div>