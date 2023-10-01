<script lang="ts">
    import { formatDate } from "$lib/utils";

    export let msg: string;
    export let ping: number;
    export let date: string;
    
    export let dataPopup: string;
</script>


<div class="card p-1 shadow-xl z-50" data-popup={dataPopup}>
    <div class="text-left">
        <p>{ping} ms - {msg} - {formatDate(date)}</p>
    </div>
    <div class="arrow bg-surface-100-800-token" />
</div>