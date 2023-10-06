// Import necessary functions/modules
// Import your data fetching functions (e.g., GetAll and GetHeatbeat)

import { GetAll, GetHeatbeat } from '$lib';

// Initialize an empty monitorsMap
let monitorsMap = new Map();

/** @type {import('./$types').PageLoad} */
export const load = ({ params }) => {
    const monitorsData = async() => {
        const res = await GetAll();
        if (res) {
            let monitors = res;
    
            for (const monitor of monitors) {
                let result = await GetHeatbeat({ monitor_id: monitor.id, count: 20, time_from: 0 });
    
                if (result) {
                    let ticks = result;
                    monitorsMap.set(monitor.id, { monitor: monitor, ticks: ticks.reverse() });
                }
            }
        }
        return monitorsMap
    }

    return {
        data: monitorsData()
    }
}
