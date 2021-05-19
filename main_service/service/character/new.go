package character

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

type NewService interface {
	Handle(ctx context.Context) error
}

type newServiceImpl struct {
	nickname string
	client   net.Conn
}

func GetNewServiceImpl(nickname string, client net.Conn) NewService {
	return &newServiceImpl{
		nickname: nickname,
		client:   client,
	}
}

func (n *newServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, n.client)
	if u == nil {
		return errors.New("can't find user")
	}

	if u.NickName != "" {
		out.OnSendMessage(u.GetNextSeq(), u.CurrentConnection, constant.MessageDialogBox, constant.CSO_NewCharacter_Failed)
		return errors.New("user already has nickname")
	}

	userinfo, err := client.GetUserClient().UpdateNickname(ctx, u.UserID, n.nickname)
	if err != nil {
		out.OnSendMessage(u.GetNextSeq(), u.CurrentConnection, constant.MessageDialogBox, constant.CSO_NewCharacter_Failed)
		return errors.New("change nickname failed!")
	}

	// UserReply部分
	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeReply), out.BuildNicknameReply())
	packet.SendPacket(rst, u.CurrentConnection)

	// UserStart部分
	rst = utils.BytesCombine(packet.BuildHeader(
		u.GetNextSeq(), constant.PacketTypeUserStart),
		out.BuildUserStart(userinfo.UserID, []byte(userinfo.UserName), []byte(userinfo.NickName)))
	packet.SendPacket(rst, u.CurrentConnection)

	// UserInfo部分
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeUserInfo), out.BuildUserInfo(out.NewUserInfo(userinfo), true, 0xFFFFFFFF))
	packet.SendPacket(rst, u.CurrentConnection)

	// ServerList部分
	servers, err := client.GetRoomClient().GetServiceList(ctx)
	if err != nil {
		return err
	}
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeServerList), out.BuildServerList(servers))
	packet.SendPacket(rst, u.CurrentConnection)

	// Inventory部分
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeInventory_Create), out.BuildUserInventory(&userinfo.UserInventory))
	packet.SendPacket(rst, u.CurrentConnection)

	// buymenu
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeFavorite), out.BuildUserBuymenu(&userinfo.UserInventory.BuyMenu))
	packet.SendPacket(rst, u.CurrentConnection)

	// bag
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeFavorite), out.BuildUserBag(userinfo.UserInventory.Loadouts))
	packet.SendPacket(rst, u.CurrentConnection)

	// cosmetics
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeFavorite), out.BuildUserCosmetics(userinfo.UserInventory.Cosmetics))
	packet.SendPacket(rst, u.CurrentConnection)

	// option
	rst = utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeOption), out.BuildUserOptions(userinfo))
	packet.SendPacket(rst, u.CurrentConnection)

	return client.GetUserCacheClient().SetNickname(ctx, u.UserID, userinfo.NickName)
}
