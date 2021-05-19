package controller

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/chat"
	"github.com/KouKouChan/YuriCore/utils"
	. "github.com/KouKouChan/YuriCore/verbose"
)

type ChatController interface {
	Handle(ctx context.Context) error
}

type chatControllerImpl struct {
	client      net.Conn
	messageType uint8
	message     string
	seq         *uint8
}

func GetChatController(client net.Conn, p *packet.PacketData) ChatController {
	impl := chatControllerImpl{}

	impl.client = client
	impl.messageType = utils.ReadUint8(p.Data, &p.CurOffset)
	impl.message = strings.TrimSpace(string(p.Data[p.CurOffset:])) // 读取剩下的所有消息，当作用户输入
	p.CurOffset = len(p.Data)                                      // offset 拉满，数据已经读完

	return &impl
}

func (c *chatControllerImpl) Handle(ctx context.Context) error {
	// 检查
	if len(c.message) == 0 {
		return fmt.Errorf("recived chat packet %+v with null message", c.messageType)
	}

	DebugPrintf(2, "got chat message %+v", c.message)

	switch c.messageType {
	case constant.MessageTypeChannel:
		return chat.GetChatInfra(c.message, c.client).ChannelHandler(ctx)
	case constant.MessageTypeRoom:
		return chat.GetChatInfra(c.message, c.client).RoomHandler(ctx)
	default:
		return fmt.Errorf("Unknown chat packet %+v", c.messageType)
	}
}
