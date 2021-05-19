package controller

import (
	"context"
	"fmt"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/option"
	"github.com/KouKouChan/YuriCore/utils"
)

type OptionController interface {
	Handle(ctx context.Context) error
}

type optionControllerImpl struct {
	client net.Conn
	p      *packet.PacketData
	Type   uint8
}

func GetOptionController(p *packet.PacketData, client net.Conn) OptionController {
	impl := optionControllerImpl{}

	impl.client = client
	impl.Type = utils.ReadUint8(p.Data, &p.CurOffset)
	impl.p = p

	return &impl
}

func (o *optionControllerImpl) Handle(ctx context.Context) error {

	switch o.Type {
	case 0:
		return option.NewOptionService(o.p, o.client).Handle(ctx)
	default:
		return fmt.Errorf("Unknown option packet %+v", o.Type)
	}
}
