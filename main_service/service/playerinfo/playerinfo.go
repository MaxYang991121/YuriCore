package playerinfo

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

type UpdateCampaignService interface {
	Handle(ctx context.Context) error
}

type updateCampaignServiceImpl struct {
	client     net.Conn
	subtype    uint8
	campaignID uint8
}

func NewUpdateCampaignService(p *packet.PacketData, client net.Conn) UpdateCampaignService {
	return &updateCampaignServiceImpl{
		client:     client,
		subtype:    utils.ReadUint8(p.Data, &p.CurOffset),
		campaignID: utils.ReadUint8(p.Data, &p.CurOffset),
	}
}

func (u *updateCampaignServiceImpl) Handle(ctx context.Context) error {
	user := client.GetUserCacheClient().GetUserByConnection(ctx, u.client)
	if user == nil || user.UserID == 0 {
		return errors.New("can't find user")
	}

	if u.subtype != 1 {
		return errors.New("not update campaign request")
	}

	switch u.campaignID {
	case 1, 2, 4, 8, 16, 32:
		info, err := client.GetUserClient().UpdateCampaign(ctx, user.UserID, u.campaignID)
		if err != nil {
			return err
		}

		// UserInfo部分
		rst := utils.BytesCombine(packet.BuildHeader(user.GetNextSeq(), constant.PacketTypeUserInfo), out.BuildUserInfo(out.NewUserInfo(info), true, 0xFFFFFFFF))
		packet.SendPacket(rst, user.CurrentConnection)

		return nil
	default:
		return errors.New("Unknown canpaign id")
	}

}
