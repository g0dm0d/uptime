export type Tick = {
    monitor_id: number;
    success: number;
    ping: number;
    message: string;
    date: string;
}

export type Monitor = {
    id: number,
    hostname: string,
    interval: number,
    protocol: string,
    address: string,
    port: number,
    tags: string[],
}