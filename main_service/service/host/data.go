package host

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

type GameDataService interface {
	Handle(ctx context.Context) error
}

type gameDataServiceImpl struct {
	data   []byte
	client net.Conn
}

func NewGameDataService(p *packet.PacketData, client net.Conn) GameDataService {
	len := utils.ReadUint16(p.Data, &p.CurOffset)
	return &gameDataServiceImpl{
		data:   utils.ReadString(p.Data, &p.CurOffset, int(len)),
		client: client,
	}
}

func (g *gameDataServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, g.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	if u.CurrentRoomId == 0 {
		return errors.New("user send game data but is not host")
	}

	return client.GetUserCacheClient().FlushUserRoomData(ctx, u.UserID, g.data)
}
