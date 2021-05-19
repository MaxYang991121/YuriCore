package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type UpdateOptionService interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateOptionServiceImpl struct {
	UserID uint32
	data   []byte
}

func NewUpdateOptionServiceImpl(UserID uint32, data []byte) UpdateOptionService {
	return &updateOptionServiceImpl{
		UserID: UserID,
		data:   data,
	}
}

func (u *updateOptionServiceImpl) Handle(ctx context.Context) (*user.UserInfo, error) {

	return client.GetUserTableClient().UpdateOption(ctx, u.UserID, u.data)
}
