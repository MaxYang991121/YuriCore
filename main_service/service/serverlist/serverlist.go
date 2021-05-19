package serverlist

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

type ServerListService interface {
	Handler(ctx context.Context) error
}

type serverListServiceImpl struct {
	client net.Conn
}

func NewServerListService(client net.Conn) ServerListService {
	return &serverListServiceImpl{
		client: client,
	}
}

func (s *serverListServiceImpl) Handler(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, s.client)
	if u == nil {
		return errors.New("can't find user")
	}

	// Server
	servers, err := client.GetRoomClient().GetServiceList(ctx)
	if err != nil {
		return err
	}

	//设置用户所在频道
	srvID := u.CurrentServerIndex
	chlID := u.CurrentChannelIndex
	if err := client.GetUserCacheClient().SetUserQuitChannel(ctx, u.UserID); err != nil {
		return errors.New("set quit channel failed")
	}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeServerList), out.BuildServerList(servers))
	packet.SendPacket(rst, u.CurrentConnection)

	if srvID == 0 || chlID == 0 {
		return nil
	}
	// 找到所有用户
	userIDs := client.GetUserCacheClient().GetChannelUsers(ctx, srvID, chlID)

	// lobby
	lobbyleave := out.BuildLobbyLeave(u.UserID)
	for i := range userIDs {
		player := client.GetUserCacheClient().GetUserByID(ctx, userIDs[i])
		if player == nil {
			continue
		}
		rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeLobby), lobbyleave)
		packet.SendPacket(rst, player.CurrentConnection)
	}
	return nil
}
