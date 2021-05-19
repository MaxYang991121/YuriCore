package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
)

type UpdateBagController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateBagControllerImpl struct {
	UserID uint32
	BagID  uint16
	Slot   uint8
	ItemID uint16
}

func NewUpdateBagController(UserID uint32, BagID uint16, Slot uint8, ItemID uint16) UpdateBagController {
	return &updateBagControllerImpl{
		UserID: UserID,
		BagID:  BagID,
		Slot:   Slot,
		ItemID: ItemID,
	}
}

func (u *updateBagControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.UserID == 0 {
		return nil, errors.New("wrong userid")
	}

	return service.NewUpdateBagServiceImpl(u.UserID, u.BagID, u.Slot, u.ItemID).Handle(ctx)
}

type UpdateBuymenuController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateBuymenuControllerImpl struct {
	UserID    uint32
	BuymenuID uint16
	Slot      uint8
	ItemID    uint16
}

func NewUpdateBuymenuController(UserID uint32, BuymenuID uint16, Slot uint8, ItemID uint16) UpdateBuymenuController {
	return &updateBuymenuControllerImpl{
		UserID:    UserID,
		BuymenuID: BuymenuID,
		Slot:      Slot,
		ItemID:    ItemID,
	}
}

func (u *updateBuymenuControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.UserID == 0 {
		return nil, errors.New("wrong userid")
	}

	return service.NewUpdateBuymenuServiceImpl(u.UserID, u.BuymenuID, u.Slot, u.ItemID).Handle(ctx)
}

type UpdateCosmeticsController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateCosmeticsControllerImpl struct {
	UserID      uint32
	CosmeticsID uint8
	cosmetics   *user.UserCosmetics
}

func NewUpdateCosmeticsController(UserID uint32, CosmeticsID uint8, cosmetics *user.UserCosmetics) UpdateCosmeticsController {
	return &updateCosmeticsControllerImpl{
		cosmetics:   cosmetics,
		UserID:      UserID,
		CosmeticsID: CosmeticsID,
	}
}

func (u *updateCosmeticsControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.UserID == 0 {
		return nil, errors.New("wrong userid")
	}

	return service.NewUpdateCosmeticsServiceImpl(u.UserID, u.CosmeticsID, u.cosmetics).Handle(ctx)
}
