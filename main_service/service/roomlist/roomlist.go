package roomlist

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

type RoomListService interface {
	Handler(ctx context.Context) error
}

type roomListServiceImpl struct {
	serverIndex  uint8
	channelIndex uint8
	client       net.Conn
}

func NewRoomListService(serverIndex, channelIndex uint8, client net.Conn) RoomListService {
	return &roomListServiceImpl{
		serverIndex:  serverIndex,
		channelIndex: channelIndex,
		client:       client,
	}
}

func (r *roomListServiceImpl) Handler(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, r.client)
	if u == nil {
		return errors.New("can't find user")
	}

	u_info, err := client.GetUserClient().GetUserInfo(ctx, u.UserID)
	if err != nil {
		return err
	}

	if u_info == nil || u_info.UserID == 0 {
		return errors.New("get null user for join lobby")
	}

	//设置用户所在频道
	if err := client.GetUserCacheClient().SetUserChannel(ctx, u.UserID, r.serverIndex, r.channelIndex); err != nil {
		return errors.New("set user channel failed")
	}

	// RoomList部分
	rooms, err := client.GetRoomClient().GetRoomList(ctx, r.serverIndex, r.channelIndex)
	if err != nil {
		return err
	}
	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeRoomList), out.BuildRoomList(rooms))
	packet.SendPacket(rst, u.CurrentConnection)

	// 找到所有用户
	users := []user.UserInfo{}
	userIDs := client.GetUserCacheClient().GetChannelUsers(ctx, r.serverIndex, r.channelIndex)
	for i := range userIDs {
		info, err := client.GetUserClient().GetUserInfo(ctx, userIDs[i])
		if err != nil {
			return err
		}
		users = append(users, *info)
	}

	// lobby
	lobbyreply := out.BuildLobbyReply(users)
	lobbyJoin := out.BuildLobbyJoin(u_info)
	for i := range userIDs {
		if userIDs[i] == u.UserID {
			rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeLobby), lobbyreply)
			packet.SendPacket(rst, u.CurrentConnection)
			continue
		}

		player := client.GetUserCacheClient().GetUserByID(ctx, userIDs[i])
		if player == nil {
			continue
		}
		rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeLobby), lobbyJoin)
		packet.SendPacket(rst, player.CurrentConnection)
	}

	return nil
}
