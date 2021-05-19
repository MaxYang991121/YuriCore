package room

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

type CreateRoomService interface {
	Handle(ctx context.Context) error
}

type createRoomServiceImpl struct {
	roomName   string
	GameModeID uint8
	maxPlayers uint8
	unk01      uint8
	WinLimit   uint8
	unk03      uint8
	KillLimit  uint16
	TimeLimit  uint8
	unk05      uint8
	client     net.Conn
}

func NewCreateRoomService(p *packet.PacketData, client net.Conn) CreateRoomService {
	return &createRoomServiceImpl{
		roomName:   string(utils.ReadStringToNULL(p.Data, &p.CurOffset)),
		GameModeID: utils.ReadUint8(p.Data, &p.CurOffset),
		maxPlayers: utils.ReadUint8(p.Data, &p.CurOffset),
		unk01:      utils.ReadUint8(p.Data, &p.CurOffset),
		WinLimit:   utils.ReadUint8(p.Data, &p.CurOffset),
		unk03:      utils.ReadUint8(p.Data, &p.CurOffset),
		KillLimit:  utils.ReadUint16(p.Data, &p.CurOffset),
		TimeLimit:  utils.ReadUint8(p.Data, &p.CurOffset),
		unk05:      utils.ReadUint8(p.Data, &p.CurOffset),
		client:     client,
	}
}

func (c *createRoomServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil {
		return errors.New("can't find user")
	}

	// 判断用户房间
	if u.CurrentRoomId > 0 {
		NewLeaveRoomService(c.client).UserRoomEnd(ctx, u)
		client.GetUserCacheClient().QuitUserRoom(ctx, u.UserID)
		return errors.New("user try to create room but already in another room")

	}

	room, err := client.GetRoomClient().NewRoom(
		ctx,
		&server.Room{
			RoomId:              0,
			RoomNumber:          0,
			HostUserID:          u.UserID,
			HostUserName:        u.NickName,
			CanSpec:             0,
			IsVipRoom:           0,
			VipRoomLevel:        0,
			RoomName:            c.roomName,
			PassWd:              "",
			GameModeID:          c.GameModeID,
			MapID:               1,
			MaxPlayers:          c.maxPlayers,
			WinLimit:            c.WinLimit,
			KillLimit:           c.KillLimit,
			WeaponRestrictions:  0,
			Status:              constant.StatusWaiting,
			TeamBalanceType:     0,
			AreBotsEnabled:      0,
			BotDifficulty:       0,
			NumCtBots:           0,
			NumTrBots:           0,
			StartMoney:          7500,
			ChangeTeams:         1,
			RespawnTime:         3,
			NextMapEnabled:      0,
			Difficulty:          0,
			IsIngame:            0,
			Users:               []uint32{u.UserID},
			ParentChannelServer: u.CurrentServerIndex,
			ParentChannel:       u.CurrentChannelIndex,
			GameTime:            c.TimeLimit,
			GameTimePerRound:    2,
			EnableVoice:         1,
			ShowFlash:           1,
			LevelLimit:          2,
			BuyLimitTime:        90,
		},
	)
	if err != nil {
		return err
	}

	// 获取用户信息
	info, err := client.GetUserClient().GetUserInfo(ctx, u.UserID)
	if err != nil {
		return err
	}

	if info == nil || info.UserID == 0 {
		return errors.New("get null user for create room")
	}

	// 设置用户房间ID
	if err := client.GetUserCacheClient().SetUserRoom(ctx, u.UserID, room.RoomId, constant.UserForceCounterTerrorist); err != nil {
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

	infos := []*user.UserInfo{info}
	caches := []*user.UserCache{u}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeRoom), out.BuildCreateRoom(infos, caches, room))
	packet.SendPacket(rst, u.CurrentConnection)

	// 发送新房间
	addroom := out.BuildAddChannelRoom(room, true, 0xFFFFFFFF)
	userIDs := client.GetUserCacheClient().GetChannelNoRoomUsers(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)
	for i := range userIDs {
		if userIDs[i] == u.UserID {
			continue
		}

		player := client.GetUserCacheClient().GetUserByID(ctx, userIDs[i])
		if player == nil {
			continue
		}
		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoomList), addroom)
		packet.SendPacket(rst, player.CurrentConnection)
	}
	return nil
}
