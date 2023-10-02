<script lang="ts">
	import '../app.postcss';
	import { AppShell, AppBar, Toast, initializeStores, getToastStore, Modal, Drawer } from '@skeletonlabs/skeleton';
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });
	initializeStores()
	export const toastStore = getToastStore();
</script>

<!-- <Toast/>
<Modal />
<Drawer /> -->

<AppShell>
	<svelte:fragment slot="header">
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">Uptime monitor</strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<a
					class="btn btn-sm variant-ringed-primary"
					href="https://github.com/g0dm0d/uptime"
					target="_blank"
					rel="noreferrer"
				>
					GitHub
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<slot />
</AppShell>
