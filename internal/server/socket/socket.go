package socket

import (
	"context"
	"log"

	"github.com/g0dm0d/uptime/internal/server/req"
	"nhooyr.io/websocket"
)

type Socket struct {
	clients []*websocket.Conn
}

func New() *Socket {
	return &Socket{
		clients: []*websocket.Conn{},
	}
}

func (s *Socket) add(socket *websocket.Conn) {
	s.clients = append(s.clients, socket)
}

func (s *Socket) remove(socket *websocket.Conn) {
	for i, client := range s.clients {
		if client == socket {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
}

func (s *Socket) Emit(message string) {
	for _, client := range s.clients {
		err := client.Write(context.TODO(), websocket.MessageText, []byte(message))
		if err != nil {
			log.Println(err)
			s.remove(client)
		}
	}

}

func (s *Socket) AddSubscriber(ctx *req.Ctx) error {
	c, err := websocket.Accept(ctx.Writer, ctx.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return err
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")
	s.add(c)
	select {}
}
