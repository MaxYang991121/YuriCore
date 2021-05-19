package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
)

type UpdateOptionController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateOptionControllerImpl struct {
	userID uint32
	data   []byte
}

func NewUpdateOptionController(userID uint32, data []byte) UpdateOptionController {
	return &updateOptionControllerImpl{
		userID: userID,
		data:   data,
	}
}

func (u *updateOptionControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.userID == 0 {
		return nil, errors.New("invalid userid")
	}
	return service.NewUpdateOptionServiceImpl(u.userID, u.data).Handle(ctx)
}
