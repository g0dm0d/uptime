<script lang="ts">
  import type { PopupSettings } from '@skeletonlabs/skeleton';
  import { popup } from '@skeletonlabs/skeleton';
  import Popup from './popup.svelte';
    import type { Tick } from '$lib/dto';

  export let index: number;
  export let item: Tick;
  const popupHover: PopupSettings = {
	  event: 'hover',
	  target: `popupHover-${item.monitor_id}-${index}`,
	  placement: 'top'
  };
</script>

<button class="[&>*]:pointer-events-none" use:popup={popupHover}>
  <Popup dataPopup={`popupHover-${item.monitor_id}-${index}`} msg={item.message} ping={item.ping} date={item.date} />
  <div class="{item.success === 1 ? 'bg-success-400' : 'bg-error-400'} h-6 px-1 rounded-full mr-2"></div>
</button>