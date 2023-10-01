<script lang="ts">
  import type { MonitorData } from "$lib/model";
  import { WsConnect, type Tick } from "$lib";
  import Cards from "$lib/components/cards.svelte";
  import { onDestroy, onMount } from "svelte";
  import { writable } from "svelte/store";

  export let data;
  let monitorsMap = writable(data.data);

  let ws: WebSocket;

  async function connectWebSocket() {
    ws = WsConnect("/monitor/ws");

    ws.onmessage = (event: MessageEvent) => {
      const message: Tick = JSON.parse(event.data);
      const currentData = $monitorsMap; // Access the current value of the store
      const updatedData = new Map(currentData);
      console.log(message.monitor_id);
      updatedData.get(message.monitor_id)?.ticks.shift();
      updatedData.get(message.monitor_id)?.ticks.push(message);
      monitorsMap.set(updatedData); // Update the store with the new data
    };

    ws.onerror = (error: Event) => {
      console.error("WebSocket error:", error);
    };

    ws.onclose = () => {
      console.log("WebSocket connection closed");
    };
  }

  function disconnectWebSocket() {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.close();
    }
  }

  onDestroy(() => {
    disconnectWebSocket();
  });

  onMount(async () => {
    await connectWebSocket();
  });
</script>

{#if $monitorsMap}
<div class="container h-full mx-auto flex justify-center">
  <div class="space-y-10">
    <Cards props={$monitorsMap}/> <!-- Use $monitorsMap to access the store's value -->
  </div>
</div>
{:else}
loading
{/if}
