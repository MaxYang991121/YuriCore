package controller

import (
	"context"
	"fmt"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/favorate"
	"github.com/KouKouChan/YuriCore/utils"
)

type FavorateController interface {
	Handle(ctx context.Context) error
}

type favorateControllerImpl struct {
	favorateType uint8
	client       net.Conn
	packet       *packet.PacketData
}

func GetFavorateController(p *packet.PacketData, client net.Conn) FavorateController {
	favorate := favorateControllerImpl{}

	favorate.favorateType = utils.ReadUint8(p.Data, &p.CurOffset)
	favorate.client = client
	favorate.packet = p

	return &favorate
}

func (r *favorateControllerImpl) Handle(ctx context.Context) error {

	switch r.favorateType {
	case 0: // buymenu
		return favorate.NewSetBuymenuService(r.packet, r.client).Handle(ctx)
	case 1: // Cosmetics
		return favorate.NewSetCosmeticsService(r.packet, r.client).Handle(ctx)
	case 2: // bag
		return favorate.NewSetBagService(r.packet, r.client).Handle(ctx)
	default:
		return fmt.Errorf("Unknown favorate packet %+v", r.favorateType)
	}
}
