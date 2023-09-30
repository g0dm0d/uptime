package req

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Ctx) error

type Handler struct {
	handlerFunc HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Ctx{
		Writer:  w,
		Request: r,
	}
	if err := h.handlerFunc(ctx); err != nil {
		log.Println(err)
	}
}
func NewHandler(f HandlerFunc) Handler {
	return Handler{
		handlerFunc: f,
	}
}
