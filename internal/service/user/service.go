package user

import (
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/pkg/jwtmanager"
)

type User interface {
	Signup(ctx *req.Ctx) error
	Signin(ctx *req.Ctx) error
}

type Service struct {
	userStore  store.UserStore
	jwtManager jwtmanager.Tool
}

func New(userStore store.UserStore, jwtManager jwtmanager.Tool) *Service {
	return &Service{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}
