package disconnect

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/room"
	"github.com/KouKouChan/YuriCore/utils"
)

type DisconnectService interface {
	Handle(ctx context.Context) error
}

type disconnectServiceImpl struct {
	client net.Conn
}

func NewDisconnectService(client net.Conn) DisconnectService {
	return &disconnectServiceImpl{
		client: client,
	}
}

func (d *disconnectServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, d.client)
	if u == nil {
		return errors.New("can't find user")
	}
	// 是否在频道内
	if u.CurrentServerIndex != 0 && u.CurrentChannelIndex != 0 {
		// 找到所有用户
		userIDs := client.GetUserCacheClient().GetChannelUsers(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)

		// lobby
		lobbyleave := out.BuildLobbyLeave(u.UserID)
		for i := range userIDs {
			player := client.GetUserCacheClient().GetUserByID(ctx, userIDs[i])
			if player == nil {
				continue
			}
			rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeLobby), lobbyleave)
			packet.SendPacket(rst, player.CurrentConnection)
		}
	}

	// 是否在房间内
	if u.CurrentRoomId != 0 {
		room.NewLeaveRoomService(d.client).UserRoomEnd(ctx, u)
	}

	// 删除用户
	uid := u.UserID
	client.GetUserCacheClient().DeleteUserByID(ctx, uid)

	// 等待用户数据刷新
	time.Sleep(time.Second / 4)
	_, err := client.GetUserClient().UserDown(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}
