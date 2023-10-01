import type { Monitor, Tick } from "$lib/dto";

export type MonitorData = {
    monitor: Monitor;
    ticks: Tick[];
}