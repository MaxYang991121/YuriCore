package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/user_service/service"
	"github.com/KouKouChan/YuriCore/utils"
)

type RegisterController interface {
	Handle(ctx context.Context) error
}

type registerControllerImpl struct {
	username string
	nickname string
	password string
}

func NewRegisterController(username, nickname, password string) RegisterController {
	return &registerControllerImpl{
		username: username,
		nickname: nickname,
		password: password,
	}
}

func (u *registerControllerImpl) Handle(ctx context.Context) error {
	if u.username == "" ||
		u.username == " " ||
		u.password == "" ||
		!utils.PasswordFilter([]byte(u.password)) ||
		len(u.nickname) > constant.MaxLen_UserName ||
		len(u.username) > constant.MaxLen_UserName ||
		len(u.password) > constant.MaxLen_Password {
		return errors.New("wrong register info")
	}

	return service.NewRegisterServiceImpl(u.username, u.nickname, u.password).Register(ctx)
}
