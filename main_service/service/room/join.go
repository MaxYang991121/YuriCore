package room

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

type JoinRoomService interface {
	Handle(ctx context.Context) error
}

type joinRoomServiceImpl struct {
	roomID   uint16
	password string
	client   net.Conn
}

func NewJoinRoomService(p *packet.PacketData, client net.Conn) JoinRoomService {
	return &joinRoomServiceImpl{
		roomID:   utils.ReadUint16(p.Data, &p.CurOffset),
		password: utils.ReadStringToNULL(p.Data, &p.CurOffset),
		client:   client,
	}
}

func (j *joinRoomServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, j.client)
	if u == nil {
		return errors.New("can't find user")
	}

	// 房间
	if j.roomID == 0 {
		return errors.New("the room which user want to join is illegal")

	}

	// 判断用户房间
	if u.CurrentRoomId > 0 {
		NewLeaveRoomService(j.client).UserRoomEnd(ctx, u)
		client.GetUserCacheClient().QuitUserRoom(ctx, u.UserID)
		return errors.New("user try to join room but already in another room")

	}

	room, err := client.GetRoomClient().JoinRoom(ctx, u.UserID, j.roomID)
	if err != nil {
		return err
	}

	// TODO检索房间状态、密码、加入的team等
	if room.PassWd != "" && room.PassWd != j.password {
		out.OnSendMessage(u.GetNextSeq(), u.CurrentConnection, constant.MessageDialogBox, constant.CSO_Warning_ROOM_JOIN_FAILED_INVALID_PASSWD)
		return nil
	}

	// 获取用户信息
	info, err := client.GetUserClient().GetUserInfo(ctx, u.UserID)
	if err != nil {
		return err
	}

	if info == nil || info.UserID == 0 {
		return errors.New("get null user for join room")
	}

	// 设置用户房间ID
	if err := client.GetUserCacheClient().SetUserRoom(ctx, u.UserID, room.RoomId, constant.UserForceTerrorist); err != nil {
		return errors.New("set user room failed")
	}
	if err := client.GetUserCacheClient().SetUserStatus(ctx, u.UserID, constant.UserNotReady); err != nil {
		return errors.New("set user room failed")
	}
	if err := client.GetUserCacheClient().SetUserIngame(ctx, u.UserID, false); err != nil {
		return errors.New("set user not in game failed")
	}
	if err := client.GetUserCacheClient().FlushUserRoomData(ctx, u.UserID, []byte{}); err != nil {
		return errors.New("clear user data failed")
	}

	info.NetInfo = u.NetInfo

	infos := []*user.UserInfo{}
	caches := []*user.UserCache{}

	uplayjoin := out.BuildPlayerJoin(u, info)
	ustatus := out.BuildUserReadyStatus(u.UserID, u.Currentstatus)
	uteam := out.BuildChangTeam(u.UserID, u.CurrentTeam)
	for i := 0; i < len(room.Users); i++ {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}
		caches = append(caches, player)

		player_info, err := client.GetUserClient().GetUserInfo(ctx, room.Users[i])
		if err != nil {
			return err
		}
		if player_info == nil {
			continue
		}

		infos = append(infos, player_info)

		if player.UserID != u.UserID {
			rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), uplayjoin)
			packet.SendPacket(rst, player.CurrentConnection)
			rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), ustatus)
			packet.SendPacket(rst, player.CurrentConnection)
			rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), uteam)
			packet.SendPacket(rst, player.CurrentConnection)
		}
	}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeRoom), out.BuildCreateRoom(infos, caches, room))
	packet.SendPacket(rst, u.CurrentConnection)

	return nil
}
