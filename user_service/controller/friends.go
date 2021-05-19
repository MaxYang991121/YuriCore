package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
)

type FriendsController interface {
	Handle(ctx context.Context) ([]user.UserInfo, error)
}

type friendsControllerImpl struct {
	userID uint32
}

func NewFriendsController(userID uint32) FriendsController {
	return &friendsControllerImpl{
		userID: userID,
	}
}

func (f *friendsControllerImpl) Handle(ctx context.Context) ([]user.UserInfo, error) {
	if f.userID == 0 {
		return nil, errors.New("wrong user info")
	}

	return service.NewFriendsService(f.userID).Handle(ctx)
}
