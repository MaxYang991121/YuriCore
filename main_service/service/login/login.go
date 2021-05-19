package login

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/main_service/service/disconnect"
	"github.com/KouKouChan/YuriCore/utils"
)

type LoginService interface {
	Handle(ctx context.Context) error
}

type loginServiceImpl struct {
	username string
	password string
	client   net.Conn
	seq      *uint8
}

func GetLoginServiceImpl(username, password string, client net.Conn, seq *uint8) LoginService {
	return &loginServiceImpl{
		username: username,
		password: password,
		client:   client,
		seq:      seq,
	}
}

func (l *loginServiceImpl) Handle(ctx context.Context) error {
	userinfo, code := client.GetUserClient().Login(ctx, l.username, l.password)
	switch code {
	case constant.Login_RPC_ERROR, constant.Login_NULL_Resp, constant.Login_DB_ERROR:
		out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_ServerFailed)
		return fmt.Errorf("login failed code=%+v", code)
	case constant.Login_Already:
		out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_Already)
		dest_u := client.GetUserCacheClient().GetUserByUserName(ctx, l.username)
		defer dest_u.CurrentConnection.Close()
		if dest_u == nil {
			out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_Already_Failed_NotFound)
			return fmt.Errorf("login failed code=%+v , try find cache failed", code)
		}

		if err := disconnect.NewDisconnectService(dest_u.CurrentConnection).Handle(ctx); err != nil {
			out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_Already_Failed)
			return fmt.Errorf("login failed code=%+v , try kick user failed err=%+v", code, err)
		}
		return fmt.Errorf("login failed code=%+v", code)
	case constant.Login_Wrong_Info, constant.Login_Wrong_PASSWORD:
		out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_Wrong)
		return fmt.Errorf("login failed code=%+v", code)
	case constant.Login_Not_Registed:
		out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_Not_Registed)
		return fmt.Errorf("login failed code=%+v", code)
	}

	if userinfo.UserID == 0 {
		return fmt.Errorf("null userid for login=%+v %+v", l.username, l.password)
	}

	// UserReply部分
	rst := utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeReply), out.BuildLoginReply())
	packet.SendPacket(rst, l.client)

	if userinfo.NickName == "" {
		rst := utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeCharacter))
		packet.SendPacket(rst, l.client)

		// 加进Cache
		cache := &user.UserCache{
			UserID:            userinfo.UserID,
			UserName:          userinfo.UserName,
			NickName:          userinfo.NickName,
			UserInventory:     userinfo.UserInventory,
			CurrentConnection: l.client,
			CurrentSequence:   l.seq,
			SequenceLocker:    sync.Mutex{},
		}

		return client.GetUserCacheClient().SetUser(ctx, cache)
	}

	// UserStart部分
	rst = utils.BytesCombine(packet.BuildHeader(
		packet.GetNextSeq(l.seq), constant.PacketTypeUserStart),
		out.BuildUserStart(userinfo.UserID, []byte(userinfo.UserName), []byte(userinfo.NickName)))
	packet.SendPacket(rst, l.client)

	// UserInfo部分
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeUserInfo), out.BuildUserInfo(out.NewUserInfo(userinfo), true, 0xFFFFFFFF))
	packet.SendPacket(rst, l.client)

	// ServerList部分
	servers, err := client.GetRoomClient().GetServiceList(ctx)
	if err != nil {
		return err
	}
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeServerList), out.BuildServerList(servers))
	packet.SendPacket(rst, l.client)

	// Inventory部分
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeInventory_Create), out.BuildUserInventory(&userinfo.UserInventory))
	packet.SendPacket(rst, l.client)

	// buymenu
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeFavorite), out.BuildUserBuymenu(&userinfo.UserInventory.BuyMenu))
	packet.SendPacket(rst, l.client)

	// bag
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeFavorite), out.BuildUserBag(userinfo.UserInventory.Loadouts))
	packet.SendPacket(rst, l.client)

	// cosmetics
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeFavorite), out.BuildUserCosmetics(userinfo.UserInventory.Cosmetics))
	packet.SendPacket(rst, l.client)

	// option
	rst = utils.BytesCombine(packet.BuildHeader(packet.GetNextSeq(l.seq), constant.PacketTypeOption), out.BuildUserOptions(userinfo))
	packet.SendPacket(rst, l.client)

	// 加进Cache
	cache := &user.UserCache{
		UserID:            userinfo.UserID,
		UserName:          userinfo.UserName,
		NickName:          userinfo.NickName,
		UserInventory:     userinfo.UserInventory,
		CurrentConnection: l.client,
		CurrentSequence:   l.seq,
		SequenceLocker:    sync.Mutex{},
	}

	return client.GetUserCacheClient().SetUser(ctx, cache)
}
