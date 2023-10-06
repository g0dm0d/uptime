<script lang="ts">
  import type { Tick } from '$lib/dto';
  import { formatDate, groupTicks } from '$lib/utils';
    
  import Chart from 'chart.js/auto';
  import { onMount } from 'svelte';

  export let ticks: Tick[];
  let charts: HTMLCanvasElement;

  function createLabels(ticks: Tick[]) {
    const labels = [];
    let prevDate = null;

    for (const tick of ticks) {
      const formattedDate = formatDate(tick.date);

      if (formattedDate !== prevDate) {
        labels.push(formattedDate);
      }

      prevDate = formattedDate;
    }

    return labels;
  }

  onMount(async() => {
  let ctx = charts.getContext('2d')
  if (ctx) {
    const data = {
      labels: Object.keys(groupTicks(ticks)),
      datasets: [{
        label: 'Ping history',
        data: Object.values(groupTicks(ticks)),
        fill: true,
        borderColor: '#a6d189',
        tension: 0.1
      }]
    };

    const chart = new Chart(ctx, {
      type: 'line',
      data: data,
    });
  }});
</script>

{#if ticks}
<canvas bind:this={charts} width="500" height="200"/>
{/if}
