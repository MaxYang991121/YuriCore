package service

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type UpdateCosmeticsService interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateCosmeticsServiceImpl struct {
	UserID      uint32
	CosmeticsID uint8
	cosmetics   *user.UserCosmetics
}

func NewUpdateCosmeticsServiceImpl(UserID uint32, CosmeticsID uint8, cosmetics *user.UserCosmetics) UpdateCosmeticsService {
	return &updateCosmeticsServiceImpl{
		cosmetics:   cosmetics,
		UserID:      UserID,
		CosmeticsID: CosmeticsID,
	}
}

func (u *updateCosmeticsServiceImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.UserID == 0 {
		return nil, errors.New("invalid userid")
	}

	return client.GetUserTableClient().UpdateCosmetics(ctx, u.UserID, u.CosmeticsID, u.cosmetics)
}
