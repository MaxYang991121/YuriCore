package favorate

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

type SetCosmeticsService interface {
	Handle(ctx context.Context) error
}

type setCosmeticsService struct {
	CosmeticsID uint8
	cosmetics   *user.UserCosmetics
	client      net.Conn
}

func NewSetCosmeticsService(p *packet.PacketData, client net.Conn) SetCosmeticsService {
	return &setCosmeticsService{
		CosmeticsID: utils.ReadUint8(p.Data, &p.CurOffset),
		cosmetics: &user.UserCosmetics{
			CosmeticsName:  utils.ReadStringToNULL(p.Data, &p.CurOffset),
			MainWeapon:     utils.ReadUint16(p.Data, &p.CurOffset),
			MainBullet:     utils.ReadUint16(p.Data, &p.CurOffset),
			SecondWeapon:   utils.ReadUint16(p.Data, &p.CurOffset),
			SecondBullet:   utils.ReadUint16(p.Data, &p.CurOffset),
			FlashbangNum:   utils.ReadUint16(p.Data, &p.CurOffset),
			GrenadeID:      utils.ReadUint16(p.Data, &p.CurOffset),
			SmokeNum:       utils.ReadUint16(p.Data, &p.CurOffset),
			DefuserNum:     utils.ReadUint16(p.Data, &p.CurOffset),
			TelescopeNum:   utils.ReadUint16(p.Data, &p.CurOffset),
			BulletproofNum: utils.ReadUint16(p.Data, &p.CurOffset),
			KnifeID:        utils.ReadUint16(p.Data, &p.CurOffset),
		},
		client: client,
	}
}

func (s *setCosmeticsService) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, s.client)
	if u == nil {
		return errors.New("can't find user")
	}

	_, err := client.GetUserClient().UpdateCosmetics(ctx, u.UserID, s.CosmeticsID, s.cosmetics)
	return err
}
