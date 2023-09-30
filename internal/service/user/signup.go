package user

import (
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/model"
	"github.com/g0dm0d/uptime/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type SignUpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Service) Signup(ctx *req.Ctx) error {
	var r SignUpReq

	err := ctx.ParseJSON(&r)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidJSON)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	id, err := s.userStore.CreateUser(store.CreateUserOpts{
		Username: r.Username,
		Password: string(passwordHash),
	})
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.UserAlreadyExists)
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(model.User{
		Username: r.Username,
		ID:       id,
	})

	return ctx.JSON(SignInResp{
		AccessToken: accessToken,
	})
}
