package server

import (
	"fmt"
	"net/http"

	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/internal/server/socket"
	"github.com/g0dm0d/uptime/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	server *http.Server
	router chi.Router

	service *service.Service
	socket  *socket.Socket
}

type Config struct {
	Addr      string
	Port      int
	Service   *service.Service
	WebSocket *socket.Socket
}

func NewServer(config *Config) *Server {
	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprint(config.Addr, ":", config.Port),
			Handler: http.NotFoundHandler(),
		},
		router: chi.NewRouter(),

		service: config.Service,
		socket:  config.WebSocket,
	}
}

func (s *Server) SetupRouter() {
	s.setupCors()

	s.router.Route("/monitor", func(r chi.Router) {
		// r.Method("POST", "/add", req.NewHandler(s.service.Monitor.Add)) // Make mw.Auth
		r.Method("GET", "/getall", req.NewHandler(s.service.Monitor.GetAll))
		r.Method("GET", "/ws", req.NewHandler(s.socket.AddSubscriber))
		r.Method("GET", "/heartbeat/{monitor}", req.NewHandler(s.service.Monitor.GetHistory))
	})

	s.server.Handler = s.router
}

func (s *Server) RunServer() error {
	return s.server.ListenAndServe()
}

func (s *Server) setupCors() {
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
}
