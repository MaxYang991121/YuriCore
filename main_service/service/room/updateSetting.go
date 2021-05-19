package room

import (
	"context"
	"errors"
	"net"
	"unsafe"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/utils"
)

type UpdateRoomService interface {
	Handle(ctx context.Context) error
}

type updateRoomServiceImpl struct {
	p      *packet.PacketData
	client net.Conn
}

func NewUpdateRoomService(p *packet.PacketData, client net.Conn) UpdateRoomService {
	return &updateRoomServiceImpl{
		p:      p,
		client: client,
	}
}

func (u *updateRoomServiceImpl) Handle(ctx context.Context) error {
	user := client.GetUserCacheClient().GetUserByConnection(ctx, u.client)
	if user == nil {
		return errors.New("can't find user")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, user.CurrentRoomId)
	if err != nil {
		return err
	}

	if room == nil || room.RoomId == 0 || room.HostUserID != user.UserID {
		return errors.New("got null resp or user is not host")
	}

	u.buildRoomSetting(room, u.p)

	room, err = client.GetRoomClient().UpdateRoomSafe(ctx, room)
	if err != nil {
		return err
	}

	if room == nil || room.RoomId == 0 || room.HostUserID != user.UserID {
		return errors.New("got null resp or user is not host")
	}

	// 给所有玩家发送
	setting := utils.BytesCombine([]byte{constant.OUTUpdateSettings}, out.BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))

	for i := range room.Users {
		dest_player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if dest_player == nil {
			continue
		}
		rst := utils.BytesCombine(packet.BuildHeader(dest_player.GetNextSeq(), constant.PacketTypeRoom), setting)
		packet.SendPacket(rst, dest_player.CurrentConnection)
	}

	return nil
}

func (u *updateRoomServiceImpl) buildRoomSetting(room *server.Room, p *packet.PacketData) {

	flags := utils.ReadUint64(p.Data, &p.CurOffset)
	lowFlag := *(*uint32)(unsafe.Pointer(&flags))
	tmp_flag := flags >> 32
	highFlag := *(*uint32)(unsafe.Pointer(&tmp_flag))

	if lowFlag&0x1 != 0 {
		room.RoomName = utils.ReadStringToNULL(p.Data, &p.CurOffset)
	}

	if lowFlag&0x2 != 0 {
		room.Unk.Unk00 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x4 != 0 {
		room.Unk.Unk01 = utils.ReadUint8(p.Data, &p.CurOffset)
		room.Unk.Unk02 = utils.ReadUint8(p.Data, &p.CurOffset)
		room.Unk.Unk03 = utils.ReadUint8(p.Data, &p.CurOffset)
		room.Unk.Unk04 = utils.ReadUint32(p.Data, &p.CurOffset)
	}

	if lowFlag&0x8 != 0 {
		room.PassWd = utils.ReadStringToNULL(p.Data, &p.CurOffset)
	}

	if lowFlag&0x10 != 0 { // 等级限制
		room.LevelLimit = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x20 != 0 {
		room.ForceCamera = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x40 != 0 { // 模式
		room.GameModeID = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x80 != 0 { // 地图
		room.MapID = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x100 != 0 {
		room.MaxPlayers = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x200 != 0 {
		room.WinLimit = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x400 != 0 {
		room.KillLimit = utils.ReadUint16(p.Data, &p.CurOffset)
	}

	if lowFlag&0x800 != 0 { // 游戏时间
		room.GameTime = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x1000 != 0 { // 单局游戏时间
		room.GameTimePerRound = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x2000 != 0 { // 武器限制
		room.WeaponRestrictions = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x4000 != 0 { // 击杀人质惩罚
		room.HostagePunish = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x8000 != 0 { // 禁止时间
		room.StopTime = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x10000 != 0 { // 购买时间限制
		room.BuyLimitTime = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x20000 != 0 { //显示昵称
		room.ShowName = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x40000 != 0 { // 人数平衡
		room.TeamBalanceType = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x80000 != 0 { // 炸弹读秒
		room.BombCountdown = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x100000 != 0 { // 友军伤害
		room.FriendHurt = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x200000 != 0 { // 允许战术手电
		room.ShowFlash = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x400000 != 0 {
		room.Unk.Unk14 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x800000 != 0 {
		room.Unk.Unk15 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x1000000 != 0 {
		room.Unk.Unk16 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x2000000 != 0 {
		room.Unk.Unk17 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x4000000 != 0 {
		room.Unk.Unk18 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x8000000 != 0 {
		room.Unk.Unk19 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x10000000 != 0 { // 视角
		room.ViewAngle = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x20000000 != 0 { // 允许语音聊天
		room.EnableVoice = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x40000000 != 0 {
		room.Status = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if lowFlag&0x80000000 != 0 {
		len_Unk23 := utils.ReadUint8(p.Data, &p.CurOffset)
		tmp_Unk23 := []server.Unk23Struct{}
		for i := 0; i < int(len_Unk23); i++ {
			tmp_Unk23 = append(
				tmp_Unk23,
				server.Unk23Struct{
					Unk23_unk00: utils.ReadUint32(p.Data, &p.CurOffset),
					Unk23_unk01: utils.ReadUint32(p.Data, &p.CurOffset),
					Unk23_unk02: utils.ReadUint8(p.Data, &p.CurOffset),
					Unk23_unk03: utils.ReadUint8(p.Data, &p.CurOffset),
					Unk23_unk04: utils.ReadUint8(p.Data, &p.CurOffset),
					Unk23_unk05: utils.ReadUint8(p.Data, &p.CurOffset),
					Unk23_unk06: utils.ReadUint16(p.Data, &p.CurOffset),
					Unk23_unk07: utils.ReadUint8(p.Data, &p.CurOffset),
					Unk23_unk08: utils.ReadUint8(p.Data, &p.CurOffset),
				},
			)
		}
		room.Unk.Unk23 = tmp_Unk23
	}

	if highFlag&0x1 != 0 {
		room.Unk.Unk24 = utils.ReadUint32(p.Data, &p.CurOffset)
		room.Unk.Unk25 = utils.ReadStringToNULL(p.Data, &p.CurOffset)
		room.Unk.Unk26 = utils.ReadUint8(p.Data, &p.CurOffset)
		room.Unk.Unk27 = utils.ReadUint8(p.Data, &p.CurOffset)
		room.Unk.Unk28 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x2 != 0 {
		room.Unk.Unk29 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x4 != 0 {
		room.BotDifficulty = utils.ReadUint8(p.Data, &p.CurOffset)
		room.NumCtBots = utils.ReadUint8(p.Data, &p.CurOffset)
		room.NumTrBots = utils.ReadUint8(p.Data, &p.CurOffset)
		room.BotBalance = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x8 != 0 { // 限制死亡>杀敌
		room.LimitDeaths = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x10 != 0 {
		room.StartMoney = utils.ReadUint16(p.Data, &p.CurOffset)
	}

	if highFlag&0x20 != 0 {
		room.Unk.Unk36 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x40 != 0 {
		room.Unk.Unk37 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x80 != 0 {
		room.Unk.Unk38 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x100 != 0 {
		room.Unk.Unk39 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x200 != 0 {
		len_Unk40 := utils.ReadUint8(p.Data, &p.CurOffset)
		tmp_Unk40 := []server.Unk40Struct{}
		for i := 0; i < int(len_Unk40); i++ {
			tmp_Unk40 = append(
				tmp_Unk40,
				server.Unk40Struct{
					Unk40_unk00: utils.ReadUint8(p.Data, &p.CurOffset),
					Unk40_unk01: utils.ReadUint8(p.Data, &p.CurOffset),
				},
			)
		}
		room.Unk.Unk40 = tmp_Unk40
	}

	if highFlag&0x400 != 0 {
		room.Unk.Unk41 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x800 != 0 { // 取消强化
		room.DisableEnhancement = utils.ReadUint8(p.Data, &p.CurOffset)
	}

	if highFlag&0x1000 != 0 { // 取消强化
		room.Unk.Unk42 = utils.ReadUint8(p.Data, &p.CurOffset)
	}

}
