package service

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type UserInfoService interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type userInfoServiceImpl struct {
	userID uint32
}

func NewUserInfoService(id uint32) UserInfoService {
	return &userInfoServiceImpl{
		userID: id,
	}
}

func (u *userInfoServiceImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	info := client.GetUserTableClient().GetUserByID(ctx, u.userID)

	if info == nil {
		return nil, errors.New("can't get user")
	}

	return client.GetUserTableClient().GetUserByID(ctx, u.userID), nil
}
