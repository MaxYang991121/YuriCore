package favorate

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

type SetBuymenuService interface {
	Handle(ctx context.Context) error
}

type setBuymenuService struct {
	buymenuID uint8
	sloutID   uint8
	ItemID    uint16
	client    net.Conn
}

func NewSetBuymenuService(p *packet.PacketData, client net.Conn) SetBuymenuService {
	return &setBuymenuService{
		buymenuID: utils.ReadUint8(p.Data, &p.CurOffset),
		sloutID:   utils.ReadUint8(p.Data, &p.CurOffset),
		ItemID:    utils.ReadUint16(p.Data, &p.CurOffset),
		client:    client,
	}
}

func (s *setBuymenuService) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, s.client)
	if u == nil {
		return errors.New("can't find user")
	}

	_, err := client.GetUserClient().UpdateBuymenu(ctx, u.UserID, uint16(s.buymenuID), s.sloutID, s.ItemID)
	return err
}
