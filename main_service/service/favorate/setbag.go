package favorate

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

type SetBagService interface {
	Handle(ctx context.Context) error
}

type setBagServiceImpl struct {
	p      *packet.PacketData
	client net.Conn
}

func NewSetBagService(p *packet.PacketData, client net.Conn) SetBagService {
	return &setBagServiceImpl{
		p:      p,
		client: client,
	}
}

func (s *setBagServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, s.client)
	if u == nil {
		return errors.New("can't find user")
	}

	subtype := utils.ReadUint8(s.p.Data, &s.p.CurOffset)

	switch subtype {
	case 0:

	case 1:

	case 2:

	default:
		if subtype < 10 || subtype >= 40 {
			return errors.New("wrong set bag type")
		}

		bag := subtype/10 - 1
		slot := subtype % 10

		if slot > 3 {
			return errors.New("wrong set bag slot")
		}

		item := utils.ReadUint16(s.p.Data, &s.p.CurOffset)

		info, err := client.GetUserClient().UpdateBag(ctx, u.UserID, uint16(bag), slot, item)
		if err != nil {
			return err
		}

		if info == nil || info.UserID == 0 {
			return errors.New("get null user for update bag")
		}

		err = client.GetUserCacheClient().FlushUserInventory(ctx, info.UserID, &info.UserInventory)
		if err != nil {
			return err
		}

	}

	return nil
}
