package controller

import (
	"context"
	"fmt"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/host"
	"github.com/KouKouChan/YuriCore/utils"
)

type HostController interface {
	Handle(ctx context.Context) error
}

type hostControllerImpl struct {
	hostType uint8
	client   net.Conn
	packet   *packet.PacketData
}

func GetHostController(p *packet.PacketData, client net.Conn) HostController {
	roomlist := hostControllerImpl{}

	roomlist.hostType = utils.ReadUint8(p.Data, &p.CurOffset)
	roomlist.client = client
	roomlist.packet = p

	return &roomlist
}

func (h *hostControllerImpl) Handle(ctx context.Context) error {

	switch h.hostType {
	case 0:
		return host.NewStartService(h.client).Handle(ctx)
	case 1:
		return host.NewGameDataService(h.packet, h.client).Handle(ctx)
	case 5:
		return host.NewGameEndService(h.client).Handle(ctx)
	case 100:
		return host.NewItemUsingService(h.packet, h.client).Handle(ctx)
	case 101:
		return host.NewGetInventoryService(h.packet, h.client).Handle(ctx)
	default:
		return fmt.Errorf("Unknown host packet %+v", h.hostType)
	}
}
