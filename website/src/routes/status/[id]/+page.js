// @ts-ignore
import { Get, GetHeatbeat } from '$lib';


// @ts-ignore
export const load = ({ params }) => {
    // @ts-ignore
    const hearbeatData = async(id) => {
        // @ts-ignore
        let result = await GetHeatbeat({ monitor_id: id, count: 86400, time_from: 0 });
        return result
    }

    // @ts-ignore
    const monitorData = async(id) => {
        // @ts-ignore
        let result = await Get({ monitor_id: id});
        return result
    }

    return {
        heartbeat: hearbeatData(params.id),
        monitor: monitorData(params.id)
    }
}