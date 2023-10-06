<script lang="ts">
  import type { MonitorData } from '$lib/model';
  import Card from './card.svelte';

  export let props: Map<number, MonitorData>;
</script>

{#key props}
{#each props.entries() as [key, value]}
  <a href="/status/{key}">
  <Card monitorData={value}/>
  </a>
{/each}
{/key}