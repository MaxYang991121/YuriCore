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

type LeaveRoomService interface {
	Handle(ctx context.Context) error
	UserRoomEnd(ctx context.Context, u *user.UserCache) error
}

type leaveRoomServiceImpl struct {
	client net.Conn
}

func NewLeaveRoomService(client net.Conn) LeaveRoomService {
	return &leaveRoomServiceImpl{
		client: client,
	}
}

func (l *leaveRoomServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, l.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	// 获取用户房间
	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}

	if u.CurrentIsIngame {
		// 如果是房主
		if u.UserID == room.HostUserID && len(room.Users) > 1 {
			// 找到新房主
			var new_Host *user.UserCache = nil
			for i := range room.Users {
				if room.Users[i] == u.UserID {
					continue
				}
				guest_u := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
				if guest_u == nil || guest_u.CurrentRoomId == 0 || !guest_u.CurrentIsIngame {
					continue
				}
				new_Host = guest_u
			}
			if new_Host != nil {
				room, err = client.GetRoomClient().SetRoomHost(ctx, new_Host.UserID, new_Host.NickName, new_Host.CurrentRoomId)
				if err != nil {
					return err
				}
				sethost := out.BuildSetHost(room.HostUserID, 1)
				roominfo := out.BuildUpdateChannelRoom(room, true, 0xFFFFFFFF)
				for i := range room.Users {
					player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
					if player == nil {
						continue
					}
					rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), sethost)
					packet.SendPacket(rst, player.CurrentConnection)
					rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoomList), roominfo)
					packet.SendPacket(rst, player.CurrentConnection)
				}
			}
		}

		// 发送hoststop
		rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeHost), out.BuildHostStop())
		packet.SendPacket(rst, u.CurrentConnection)

		// 发送结果
		rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeRoom), out.BuildRoomResult(room))
		packet.SendPacket(rst, u.CurrentConnection)

		// 设置用户
		if err := client.GetUserCacheClient().SetUserIngame(ctx, u.UserID, false); err != nil {
			return errors.New("set user no ingame failed")
		}
		if err := client.GetUserCacheClient().SetUserStatus(ctx, u.UserID, constant.UserNotReady); err != nil {
			return errors.New("set user not ready failed")
		}

		status := out.BuildUserReadyStatus(u.UserID, u.Currentstatus)
		for i := range room.Users {
			player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
			if player == nil {
				continue
			}
			rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), status)
			packet.SendPacket(rst, player.CurrentConnection)
		}
	} else {
		// 退出用户房间
		room, err = client.GetRoomClient().LeaveRoom(ctx, u.UserID, u.CurrentRoomId)
		if err != nil {
			return err
		}

		// 是否没有剩余用户了
		if len(room.Users) != 0 {
			// 选举新房主
			set_host := false
			if u.UserID == room.HostUserID {
				new_host := client.GetUserCacheClient().GetUserByID(ctx, room.Users[0])
				if new_host == nil || new_host.CurrentRoomId == 0 {
					return errors.New("can't find new host!")
				}
				room, err = client.GetRoomClient().SetRoomHost(ctx, new_host.UserID, new_host.NickName, new_host.CurrentRoomId)
				if err != nil {
					return err
				}
				set_host = true
			}

			// 发送离开消息
			leave := out.BuildLeaveRoom(u.UserID)
			sethost := out.BuildSetHost(room.HostUserID, 1)
			roominfo := out.BuildUpdateChannelRoom(room, true, 0xFFFFFFFF)
			for i := range room.Users {
				player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
				if player == nil {
					continue
				}
				rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), leave)
				packet.SendPacket(rst, player.CurrentConnection)
				if set_host {
					rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), sethost)
					packet.SendPacket(rst, player.CurrentConnection)
					rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoomList), roominfo)
					packet.SendPacket(rst, player.CurrentConnection)
				}
			}
		} else {
			err = client.GetRoomClient().DelRoom(ctx, room.RoomId)
			if err != nil {
				return err
			}
			// 发送房间删除消息
			deleteroom := out.BuildDeleteChannelRoom(room.RoomId)
			userIDs := client.GetUserCacheClient().GetChannelNoRoomUsers(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)
			for i := range userIDs {
				if userIDs[i] == u.UserID {
					continue
				}

				player := client.GetUserCacheClient().GetUserByID(ctx, userIDs[i])
				if player == nil {
					continue
				}
				rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoomList), deleteroom)
				packet.SendPacket(rst, player.CurrentConnection)
			}
		}

		// 设置用户
		if err := client.GetUserCacheClient().QuitUserRoom(ctx, u.UserID); err != nil {
			return errors.New("set user room failed")
		}

		// 发送房间列表
		if u.CurrentServerIndex != 0 && u.CurrentChannelIndex != 0 {
			rooms, err := client.GetRoomClient().GetRoomList(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)
			if err != nil {
				return err
			}
			rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeRoomList), out.BuildRoomList(rooms))
			packet.SendPacket(rst, u.CurrentConnection)

			// 找到所有用户
			users := []user.UserInfo{}
			userIDs := client.GetUserCacheClient().GetChannelUsers(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)
			for i := range userIDs {
				info, err := client.GetUserClient().GetUserInfo(ctx, userIDs[i])
				if err != nil {
					return err
				}
				users = append(users, *info)
			}

			// lobby
			lobbyreply := out.BuildLobbyReply(users)
			rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeLobby), lobbyreply)
			packet.SendPacket(rst, u.CurrentConnection)
		}

	}

	// 判断游戏内玩家人数
	ingame_num := 0
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}

		if player.CurrentIsIngame {
			ingame_num++
		}
	}
	if ingame_num != 0 {
		return nil
	}

	room, err = client.GetRoomClient().EndGame(ctx, room.HostUserID, room.RoomId)
	if err != nil {
		return err
	}

	setting := utils.BytesCombine([]byte{constant.OUTUpdateSettings}, out.BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}

		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), setting)
		packet.SendPacket(rst, player.CurrentConnection)
	}
	return nil
}

func (l *leaveRoomServiceImpl) UserRoomEnd(ctx context.Context, u *user.UserCache) error {
	// 获取用户房间
	room, err := client.GetRoomClient().LeaveRoom(ctx, u.UserID, u.CurrentRoomId)
	if err != nil {
		return err
	}

	// 是否没有剩余用户了
	if len(room.Users) == 0 {
		err = client.GetRoomClient().DelRoom(ctx, room.RoomId)
		if err != nil {
			return err
		}
		// 发送房间删除消息
		deleteroom := out.BuildDeleteChannelRoom(room.RoomId)
		userIDs := client.GetUserCacheClient().GetChannelNoRoomUsers(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)
		for i := range userIDs {
			if userIDs[i] == u.UserID {
				continue
			}

			player := client.GetUserCacheClient().GetUserByID(ctx, userIDs[i])
			if player == nil {
				continue
			}
			rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoomList), deleteroom)
			packet.SendPacket(rst, player.CurrentConnection)
		}
		return nil
	}

	// 选举新房主
	set_host := false
	if u.UserID == room.HostUserID {
		new_host := client.GetUserCacheClient().GetUserByID(ctx, room.Users[0])
		if new_host == nil || new_host.CurrentRoomId == 0 {
			return errors.New("can't find new host!")
		}
		room, err = client.GetRoomClient().SetRoomHost(ctx, new_host.UserID, new_host.NickName, new_host.CurrentRoomId)
		if err != nil {
			return err
		}
		set_host = true
	}

	// 发送离开消息
	leave := out.BuildLeaveRoom(u.UserID)
	sethost := out.BuildSetHost(room.HostUserID, 1)
	roominfo := out.BuildUpdateChannelRoom(room, true, 0xFFFFFFFF)
	ingame_num := 0
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}

		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), leave)
		packet.SendPacket(rst, player.CurrentConnection)

		if set_host {
			rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), sethost)
			packet.SendPacket(rst, player.CurrentConnection)
			rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoomList), roominfo)
			packet.SendPacket(rst, player.CurrentConnection)
		}

		if player.CurrentIsIngame {
			ingame_num++
		}
	}
	// 如果还有人在游戏内
	if ingame_num != 0 {
		return nil
	}

	// 所有人都不在游戏内了,但是还有在房间内的
	room, err = client.GetRoomClient().EndGame(ctx, room.HostUserID, room.RoomId)
	if err != nil {
		return err
	}

	setting := utils.BytesCombine([]byte{constant.OUTUpdateSettings}, out.BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}

		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), setting)
		packet.SendPacket(rst, player.CurrentConnection)
	}
	return nil
}
