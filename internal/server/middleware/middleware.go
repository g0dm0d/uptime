package middleware

import (
	"github.com/g0dm0d/uptime/internal/service"
	"github.com/g0dm0d/uptime/pkg/jwtmanager"
)

type Middleware struct {
	service    *service.Service
	jwtManager *jwtmanager.Tool
}

func New(s *service.Service, j *jwtmanager.Tool) *Middleware {
	return &Middleware{
		service:    s,
		jwtManager: j,
	}
}
