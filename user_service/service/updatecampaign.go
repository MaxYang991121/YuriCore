package service

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type UpdateCampaignService interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateCampaignServiceImpl struct {
	UserID     uint32
	CampaignID uint8
}

func NewUpdateCampaignServiceImpl(UserID uint32, CampaignID uint8) UpdateCampaignService {
	return &updateCampaignServiceImpl{
		UserID:     UserID,
		CampaignID: CampaignID,
	}
}

func (u *updateCampaignServiceImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.UserID == 0 {
		return nil, errors.New("invalid userid")
	}

	return client.GetUserTableClient().UpdateCampaign(ctx, u.UserID, u.CampaignID)
}
