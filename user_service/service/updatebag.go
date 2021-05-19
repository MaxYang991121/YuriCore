package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type UpdateBagService interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateBagServiceImpl struct {
	UserID uint32
	BagID  uint16
	Slot   uint8
	ItemID uint16
}

func NewUpdateBagServiceImpl(UserID uint32, BagID uint16, Slot uint8, ItemID uint16) UpdateBagService {
	return &updateBagServiceImpl{
		UserID: UserID,
		BagID:  BagID,
		Slot:   Slot,
		ItemID: ItemID,
	}
}

func (u *updateBagServiceImpl) Handle(ctx context.Context) (*user.UserInfo, error) {

	return client.GetUserTableClient().UpdateBag(ctx, u.UserID, u.BagID, u.Slot, u.ItemID)
}
