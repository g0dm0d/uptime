import type { Tick } from "$lib/dto";
import type { Monitor } from "$lib/dto/monitor";
import { restController } from "./controller";

export type TicksInfoParams = {
  monitor_id: number;
  count: number;
  time_from: number;
}

export async function GetHeatbeat(params: TicksInfoParams): Promise<Tick[]> {
  return restController.call<Tick[]>({
    path: `/monitor/heartbeat/${params.monitor_id}?count=${params.count}&from=${params.time_from}`,
    method: "GET",
  })
}

export async function GetAll(): Promise<Monitor[]> {
  return restController.call<Monitor[]>({
    path: `/monitor/getall`,
    method: "GET",
  })
}

export type GetMonitorParams = {
  monitor_id: number;
}

export async function Get(params: GetMonitorParams): Promise<Monitor> {
  return restController.call<Monitor>({
    path: `/monitor/get/${params.monitor_id}`,
    method: "GET",
  })
}