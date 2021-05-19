package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
)

type UpdateCampaignController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateCampaignControllerImpl struct {
	UserID     uint32
	campaignID uint8
}

func NewUpdateCampaignController(UserID uint32, campaignID uint8) UpdateCampaignController {
	return &updateCampaignControllerImpl{
		UserID:     UserID,
		campaignID: campaignID,
	}
}

func (u *updateCampaignControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if u.UserID == 0 {
		return nil, errors.New("wrong userid")
	}

	return service.NewUpdateCampaignServiceImpl(u.UserID, u.campaignID).Handle(ctx)
}
