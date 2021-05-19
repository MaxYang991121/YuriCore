package controller

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
	"github.com/KouKouChan/YuriCore/utils"
)

type LoginController interface {
	Handle(ctx context.Context) (*user.UserInfo, int8)
}

type loginControllerImpl struct {
	username string
	password string
}

func NewLoginController(username, password string) LoginController {
	return &loginControllerImpl{
		username: username,
		password: password,
	}
}

func (u *loginControllerImpl) Handle(ctx context.Context) (*user.UserInfo, int8) {
	if u.username == "" ||
		u.username == " " ||
		u.password == "" ||
		!utils.PasswordFilter([]byte(u.password)) ||
		len(u.username) > constant.MaxLen_UserName ||
		len(u.password) > constant.MaxLen_Password {
		return nil, constant.Login_Wrong_Info
	}

	return service.NewLoginServiceImpl(u.username, u.password).Login(ctx)
}
