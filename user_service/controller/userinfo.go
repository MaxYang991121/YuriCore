package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
)

type UserInfoController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type userInfoControllerImpl struct {
	userID uint32
}

func NewUserInfoController(id uint32) UserInfoController {
	return &userInfoControllerImpl{
		userID: id,
	}
}

func (u *userInfoControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.userID == 0 {
		return nil, errors.New("null userid")
	}

	return service.NewUserInfoService(u.userID).Handle(ctx)
}
