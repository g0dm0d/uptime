package user

import (
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/model"
	"github.com/g0dm0d/uptime/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResp struct {
	AccessToken string `json:"access_token"`
}

func (s *Service) Signin(ctx *req.Ctx) error {
	var r SignInReq

	err := ctx.ParseJSON(&r)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidJSON)
	}

	user, err := s.userStore.GetUserByUsername(r.Username)

	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.IncorrectLoginOrPassword)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.IncorrectLoginOrPassword)
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(model.NewUser(user))

	return ctx.JSON(SignInResp{
		AccessToken: accessToken,
	})
}
