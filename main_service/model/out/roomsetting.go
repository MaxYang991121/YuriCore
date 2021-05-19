package out

import (
	"unsafe"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildRoomSetting(room *server.Room, flags uint64) []byte {
	buf := make([]byte, 4096)
	offset := 0

	utils.WriteUint64(&buf, flags, &offset)
	lowFlag := *(*uint32)(unsafe.Pointer(&flags))
	flags = flags >> 32
	highFlag := *(*uint32)(unsafe.Pointer(&flags))

	if lowFlag&0x1 != 0 {
		utils.WriteStringWithNull(&buf, []byte(room.RoomName), &offset)
	}

	if lowFlag&0x2 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk00, &offset)
	}

	if lowFlag&0x4 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk01, &offset)
		utils.WriteUint8(&buf, room.Unk.Unk02, &offset)
		utils.WriteUint8(&buf, room.Unk.Unk03, &offset)
		utils.WriteUint32(&buf, room.Unk.Unk04, &offset)
	}

	if lowFlag&0x8 != 0 {
		utils.WriteStringWithNull(&buf, []byte(room.PassWd), &offset)
	}

	if lowFlag&0x10 != 0 { // 等级限制
		utils.WriteUint8(&buf, room.LevelLimit, &offset)
	}

	if lowFlag&0x20 != 0 {
		utils.WriteUint8(&buf, room.ForceCamera, &offset)
	}

	if lowFlag&0x40 != 0 { // 模式
		utils.WriteUint8(&buf, room.GameModeID, &offset)
	}

	if lowFlag&0x80 != 0 { // 地图
		utils.WriteUint8(&buf, room.MapID, &offset)
	}

	if lowFlag&0x100 != 0 {
		utils.WriteUint8(&buf, room.MaxPlayers, &offset)
	}

	if lowFlag&0x200 != 0 {
		utils.WriteUint8(&buf, room.WinLimit, &offset)
	}

	if lowFlag&0x400 != 0 {
		utils.WriteUint16(&buf, room.KillLimit, &offset)
	}

	if lowFlag&0x800 != 0 { // 游戏时间
		utils.WriteUint8(&buf, room.GameTime, &offset)
	}

	if lowFlag&0x1000 != 0 {
		utils.WriteUint8(&buf, room.GameTimePerRound, &offset)
	}

	if lowFlag&0x2000 != 0 { // 武器限制
		utils.WriteUint8(&buf, room.WeaponRestrictions, &offset)
	}

	if lowFlag&0x4000 != 0 { // 击杀人质惩罚
		utils.WriteUint8(&buf, room.HostagePunish, &offset)
	}

	if lowFlag&0x8000 != 0 { // 禁止时间
		utils.WriteUint8(&buf, room.StopTime, &offset)
	}

	if lowFlag&0x10000 != 0 { // 购买时间限制
		utils.WriteUint8(&buf, room.BuyLimitTime, &offset)
	}

	if lowFlag&0x20000 != 0 { //显示昵称
		utils.WriteUint8(&buf, room.ShowName, &offset)
	}

	if lowFlag&0x40000 != 0 {
		utils.WriteUint8(&buf, room.TeamBalanceType, &offset)
	}

	if lowFlag&0x80000 != 0 {
		utils.WriteUint8(&buf, room.BombCountdown, &offset)
	}

	if lowFlag&0x100000 != 0 {
		utils.WriteUint8(&buf, room.FriendHurt, &offset)
	}

	if lowFlag&0x200000 != 0 { // 允许战术手电
		utils.WriteUint8(&buf, room.ShowFlash, &offset)
	}

	if lowFlag&0x400000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk14, &offset)
	}

	if lowFlag&0x800000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk15, &offset)
	}

	if lowFlag&0x1000000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk16, &offset)
	}

	if lowFlag&0x2000000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk17, &offset)
	}

	if lowFlag&0x4000000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk18, &offset)
	}

	if lowFlag&0x8000000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk19, &offset)
	}

	if lowFlag&0x10000000 != 0 { // 视角
		utils.WriteUint8(&buf, room.ViewAngle, &offset)
	}

	if lowFlag&0x20000000 != 0 { // 允许语音聊天
		utils.WriteUint8(&buf, room.EnableVoice, &offset)
	}

	if lowFlag&0x40000000 != 0 {
		utils.WriteUint8(&buf, room.Status, &offset)
	}

	if lowFlag&0x80000000 != 0 {
		utils.WriteUint8(&buf, uint8(len(room.Unk.Unk23)), &offset)
		for i := range room.Unk.Unk23 {
			utils.WriteUint32(&buf, room.Unk.Unk23[i].Unk23_unk00, &offset)
			utils.WriteUint32(&buf, room.Unk.Unk23[i].Unk23_unk01, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk23[i].Unk23_unk02, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk23[i].Unk23_unk03, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk23[i].Unk23_unk04, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk23[i].Unk23_unk05, &offset)
			utils.WriteUint16(&buf, room.Unk.Unk23[i].Unk23_unk06, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk23[i].Unk23_unk07, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk23[i].Unk23_unk08, &offset)
		}
	}

	if highFlag&0x1 != 0 {
		utils.WriteUint32(&buf, room.Unk.Unk24, &offset)
		utils.WriteStringWithNull(&buf, []byte(room.Unk.Unk25), &offset)
		utils.WriteUint8(&buf, room.Unk.Unk26, &offset)
		utils.WriteUint8(&buf, room.Unk.Unk27, &offset)
		utils.WriteUint8(&buf, room.Unk.Unk28, &offset)

	}

	if highFlag&0x2 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk29, &offset)
	}

	if highFlag&0x4 != 0 { //
		utils.WriteUint8(&buf, room.BotDifficulty, &offset)
		utils.WriteUint8(&buf, room.NumCtBots, &offset)
		utils.WriteUint8(&buf, room.NumTrBots, &offset)
		utils.WriteUint8(&buf, room.BotBalance, &offset)
	}

	if highFlag&0x8 != 0 { // 限制死亡>杀敌
		utils.WriteUint8(&buf, room.LimitDeaths, &offset)
	}

	if highFlag&0x10 != 0 { // 金钱
		utils.WriteUint16(&buf, room.StartMoney, &offset)
	}

	if highFlag&0x20 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk36, &offset)
	}

	if highFlag&0x40 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk37, &offset)
	}

	if highFlag&0x80 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk38, &offset)
	}

	if highFlag&0x100 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk39, &offset)
	}

	if highFlag&0x200 != 0 {
		utils.WriteUint8(&buf, uint8(len(room.Unk.Unk40)), &offset)
		for i := range room.Unk.Unk40 {
			utils.WriteUint8(&buf, room.Unk.Unk40[i].Unk40_unk00, &offset)
			utils.WriteUint8(&buf, room.Unk.Unk40[i].Unk40_unk01, &offset)
		}
	}

	if highFlag&0x400 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk41, &offset)
	}

	if highFlag&0x800 != 0 {
		utils.WriteUint8(&buf, room.DisableEnhancement, &offset)
	}

	if highFlag&0x1000 != 0 {
		utils.WriteUint8(&buf, room.Unk.Unk42, &offset)
	}

	return buf[:offset]
}
