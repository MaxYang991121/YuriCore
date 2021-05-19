package room

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

type CloseResultService interface {
	Handle(ctx context.Context) error
}

type closeResultServiceImpl struct {
	client net.Conn
}

func NewCloseResultService(client net.Conn) CloseResultService {
	return &closeResultServiceImpl{
		client: client,
	}
}

func (c *closeResultServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeHost), out.BuildCloseResultWindow())
	packet.SendPacket(rst, u.CurrentConnection)

	return nil
}
