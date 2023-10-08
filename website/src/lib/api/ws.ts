export function WsConnect(path: string): WebSocket {
  const url = new URL(import.meta.env.VITE_API_URL);
  url.pathname += path;
  url.protocol = "ws"

  let ws = new WebSocket(url);
  
  ws.onerror = (error: Event) => {
    console.error("WebSocket error:", error);
  };

  ws.onclose = () => {
    console.log("WebSocket connection closed");
  };
  return ws;
}