package controller

import (
	"context"
	"fmt"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/playerinfo"
	"github.com/KouKouChan/YuriCore/utils"
)

type PlayerInfoController interface {
	Handle(ctx context.Context) error
}

type playerInfoControllerImpl struct {
	playerType uint8
	client     net.Conn
	packet     *packet.PacketData
}

func GetPlayerInfoController(p *packet.PacketData, client net.Conn) PlayerInfoController {
	player := playerInfoControllerImpl{}

	player.playerType = utils.ReadUint8(p.Data, &p.CurOffset)
	player.client = client
	player.packet = p

	return &player
}

func (p *playerInfoControllerImpl) Handle(ctx context.Context) error {

	switch p.playerType {
	case 5: // campaign
		return playerinfo.NewUpdateCampaignService(p.packet, p.client).Handle(ctx)

	default:
		return fmt.Errorf("Unknown playerInfo packet %+v", p.playerType)
	}
}
