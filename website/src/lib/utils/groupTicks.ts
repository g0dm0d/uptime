import type { Tick } from "$lib/dto";
import { formatDate } from "./date";

export function groupTicks(ticks: Tick[]) {
  const groupedTicks: Record<string, { totalPing: number; count: number }> = {};

  for (const tick of ticks) {
    const date = new Date(tick.date);
    const hour = date.getHours();

    const groupKey = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()} ${hour}:00:00`;

    if (!groupedTicks[groupKey]) {
      groupedTicks[groupKey] = { totalPing: 0, count: 0 };
    }

    groupedTicks[groupKey].totalPing += tick.ping;
    groupedTicks[groupKey].count++;
  }

  const result: Record<string, number> = {};
  for (const groupKey in groupedTicks) {
    const group = groupedTicks[groupKey];
    const averagePing = group.totalPing / group.count;
    result[groupKey] = averagePing;
  }

  return result;
}