package convert

import (
	"github.com/KouKouChan/YuriCore/main_service/gen-go/yuricore/room_service"
	"github.com/KouKouChan/YuriCore/main_service/gen-go/yuricore/user_service"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func ConvertToUserInfo(userInfo *user_service.UserInfo) *user.UserInfo {
	if userInfo == nil {
		return nil
	}
	user := &user.UserInfo{
		UserID:        uint32(userInfo.UserID),
		UserName:      userInfo.UserName,
		NickName:      userInfo.NickName,
		Password:      "",
		Level:         uint8(userInfo.Level),
		CurExp:        uint64(userInfo.CurrentEXP),
		MaxExp:        uint64(userInfo.MaxEXP),
		Points:        uint64(userInfo.Points),
		PlayedMatches: uint32(userInfo.PlayedMatches),
		Wins:          uint32(userInfo.Wins),
		Kills:         uint32(userInfo.Kills),
		Deaths:        uint32(userInfo.Deaths),
		UserInventory: ConvertToInventory(userInfo.UserInventory),
		Campaign:      uint8(userInfo.Campaign),
		NetInfo:       ConvertToUserNetInfo(userInfo.NetInfo),
		Friends:       userInfo.Friends,
		Rank:          uint32(userInfo.Rank),
		ChatTimes:     uint8(userInfo.ChatTimes),
		Options:       userInfo.Options,
	}
	return user
}

func ConvertToServiceUser(userInfo *user.UserInfo) *user_service.UserInfo {
	if userInfo == nil {
		return nil
	}
	user := &user_service.UserInfo{
		UserID:        int32(userInfo.UserID),
		UserName:      userInfo.UserName,
		NickName:      userInfo.NickName,
		Level:         int8(userInfo.Level),
		CurrentEXP:    int64(userInfo.CurExp),
		MaxEXP:        int64(userInfo.MaxExp),
		Points:        int64(userInfo.Points),
		PlayedMatches: int32(userInfo.PlayedMatches),
		Wins:          int32(userInfo.Wins),
		Kills:         int32(userInfo.Kills),
		Deaths:        int32(userInfo.Deaths),
		Campaign:      int8(userInfo.Campaign),
		UserInventory: ConvertToServiceInventory(userInfo.UserInventory),
		NetInfo:       ConvertToServiceNetInfo(&userInfo.NetInfo),
		Friends:       userInfo.Friends,
		Rank:          int32(userInfo.Rank),
		ChatTimes:     int8(userInfo.ChatTimes),
		Options:       userInfo.Options,
	}
	return user
}

func ConvertToInventory(userInventory *user_service.Inventory) user.Inventory {
	if userInventory == nil {
		return user.Inventory{}
	}
	items := []user.UserInventoryItem{}
	for i := range userInventory.Items {
		items = append(
			items,
			user.UserInventoryItem{
				Id:      uint16(userInventory.Items[i].ID),
				Count:   uint16(userInventory.Items[i].Count),
				Existed: uint8(userInventory.Items[i].Existed),
				Type:    uint8(userInventory.Items[i].Type),
				Time:    uint32(userInventory.Items[i].Time),
			})
	}

	rst := user.Inventory{
		Items:     items,
		BuyMenu:   ConvertToUserBuyMenu(userInventory.BuyMenu),
		Loadouts:  ConvertToUserLoadouts(userInventory.Loadouts),
		Cosmetics: ConvertToUserCosmetics(userInventory.Cosmetics),
	}
	return rst
}

func ConvertToServiceInventory(userInventory user.Inventory) *user_service.Inventory {
	items := []*user_service.UserInventoryItem{}
	for i := range userInventory.Items {
		items = append(
			items,
			&user_service.UserInventoryItem{
				ID:      int16(userInventory.Items[i].Id),
				Count:   int16(userInventory.Items[i].Count),
				Existed: int8(userInventory.Items[i].Existed),
				Type:    int8(userInventory.Items[i].Type),
				Time:    int32(userInventory.Items[i].Time),
			})
	}

	rst := &user_service.Inventory{
		Items:     items,
		BuyMenu:   ConvertToServiceBuyMenu(&userInventory.BuyMenu),
		Loadouts:  ConvertToServiceLoadouts(userInventory.Loadouts),
		Cosmetics: ConvertToServiceCosmetics(userInventory.Cosmetics),
	}
	return rst
}
func ConvertToUserBuyMenu(buymenu *user_service.UserBuyMenu) user.UserBuyMenu {
	return user.UserBuyMenu{
		PistolsTR:     utils.Convertint16sTo9Uint16s(buymenu.PistolsTR),
		ShotgunsTR:    utils.Convertint16sTo9Uint16s(buymenu.ShotgunsTR),
		SmgsTR:        utils.Convertint16sTo9Uint16s(buymenu.SmgsTR),
		RiflesTR:      utils.Convertint16sTo9Uint16s(buymenu.RiflesTR),
		MachinegunsTR: utils.Convertint16sTo9Uint16s(buymenu.MachinegunsTR),
		MeleesTR:      utils.Convertint16sTo9Uint16s(buymenu.MeleesTR),
		EquipmentTR:   utils.Convertint16sTo9Uint16s(buymenu.EquipmentTR),
		ClassesTR:     utils.Convertint16sTo9Uint16s(buymenu.ClassesTR),
		PistolsCT:     utils.Convertint16sTo9Uint16s(buymenu.PistolsCT),
		ShotgunsCT:    utils.Convertint16sTo9Uint16s(buymenu.ShotgunsCT),
		SmgsCT:        utils.Convertint16sTo9Uint16s(buymenu.SmgsCT),
		RiflesCT:      utils.Convertint16sTo9Uint16s(buymenu.RiflesCT),
		MachinegunsCT: utils.Convertint16sTo9Uint16s(buymenu.MachinegunsCT),
		MeleesCT:      utils.Convertint16sTo9Uint16s(buymenu.MeleesCT),
		EquipmentCT:   utils.Convertint16sTo9Uint16s(buymenu.EquipmentCT),
		ClassesCT:     utils.Convertint16sTo9Uint16s(buymenu.ClassesCT),
	}
}

func ConvertToServiceBuyMenu(buymenu *user.UserBuyMenu) *user_service.UserBuyMenu {
	return &user_service.UserBuyMenu{
		PistolsTR:     utils.Convert9Uint16sToInt16s(buymenu.PistolsTR),
		ShotgunsTR:    utils.Convert9Uint16sToInt16s(buymenu.ShotgunsTR),
		SmgsTR:        utils.Convert9Uint16sToInt16s(buymenu.SmgsTR),
		RiflesTR:      utils.Convert9Uint16sToInt16s(buymenu.RiflesTR),
		MachinegunsTR: utils.Convert9Uint16sToInt16s(buymenu.MachinegunsTR),
		MeleesTR:      utils.Convert9Uint16sToInt16s(buymenu.MeleesTR),
		EquipmentTR:   utils.Convert9Uint16sToInt16s(buymenu.EquipmentTR),
		ClassesTR:     utils.Convert9Uint16sToInt16s(buymenu.ClassesTR),
		PistolsCT:     utils.Convert9Uint16sToInt16s(buymenu.PistolsCT),
		ShotgunsCT:    utils.Convert9Uint16sToInt16s(buymenu.ShotgunsCT),
		SmgsCT:        utils.Convert9Uint16sToInt16s(buymenu.SmgsCT),
		RiflesCT:      utils.Convert9Uint16sToInt16s(buymenu.RiflesCT),
		MachinegunsCT: utils.Convert9Uint16sToInt16s(buymenu.MachinegunsCT),
		MeleesCT:      utils.Convert9Uint16sToInt16s(buymenu.MeleesCT),
		EquipmentCT:   utils.Convert9Uint16sToInt16s(buymenu.EquipmentCT),
		ClassesCT:     utils.Convert9Uint16sToInt16s(buymenu.ClassesCT),
	}
}

func ConvertToUserLoadouts(loadout []*user_service.UserLoadout) [3]user.UserLoadout {
	rst := [3]user.UserLoadout{}
	for i := range loadout {
		rst[i] = user.UserLoadout{
			MainWeapon:   uint16(loadout[i].MainWeapon),
			SecondWeapon: uint16(loadout[i].SecondWeapon),
			Knife:        uint16(loadout[i].Knife),
			Grenade:      uint16(loadout[i].Grenade),
		}
	}
	return rst
}

func ConvertToServiceLoadouts(loadout [3]user.UserLoadout) []*user_service.UserLoadout {
	rst := make([]*user_service.UserLoadout, 3)
	for i := range loadout {
		rst[i] = &user_service.UserLoadout{
			MainWeapon:   int16(loadout[i].MainWeapon),
			SecondWeapon: int16(loadout[i].SecondWeapon),
			Knife:        int16(loadout[i].Knife),
			Grenade:      int16(loadout[i].Grenade),
		}
	}
	return rst
}

func ConvertToUserCosmetics(cosmetics []*user_service.UserCosmetics) [5]user.UserCosmetics {
	rst := [5]user.UserCosmetics{}
	for i := range cosmetics {
		rst[i] = user.UserCosmetics{
			MainWeapon:     uint16(cosmetics[i].MainWeapon),
			SecondWeapon:   uint16(cosmetics[i].SecondWeapon),
			KnifeID:        uint16(cosmetics[i].KnifeID),
			GrenadeID:      uint16(cosmetics[i].GrenadeID),
			CosmeticsName:  cosmetics[i].CosmeticsName,
			MainBullet:     uint16(cosmetics[i].MainBullet),
			SecondBullet:   uint16(cosmetics[i].SecondBullet),
			FlashbangNum:   uint16(cosmetics[i].FlashbangNum),
			SmokeNum:       uint16(cosmetics[i].SmokeNum),
			DefuserNum:     uint16(cosmetics[i].DefuserNum),
			TelescopeNum:   uint16(cosmetics[i].TelescopeNum),
			BulletproofNum: uint16(cosmetics[i].BulletproofNum),
		}
	}
	return rst
}

func ConvertToServiceCosmetics(cosmetics [5]user.UserCosmetics) []*user_service.UserCosmetics {
	rst := make([]*user_service.UserCosmetics, 5)
	for i := range cosmetics {
		rst[i] = &user_service.UserCosmetics{
			MainWeapon:     int16(cosmetics[i].MainWeapon),
			SecondWeapon:   int16(cosmetics[i].SecondWeapon),
			KnifeID:        int16(cosmetics[i].KnifeID),
			GrenadeID:      int16(cosmetics[i].GrenadeID),
			CosmeticsName:  cosmetics[i].CosmeticsName,
			MainBullet:     int16(cosmetics[i].MainBullet),
			SecondBullet:   int16(cosmetics[i].SecondBullet),
			FlashbangNum:   int16(cosmetics[i].FlashbangNum),
			SmokeNum:       int16(cosmetics[i].SmokeNum),
			DefuserNum:     int16(cosmetics[i].DefuserNum),
			TelescopeNum:   int16(cosmetics[i].TelescopeNum),
			BulletproofNum: int16(cosmetics[i].BulletproofNum),
		}
	}
	return rst
}

func ConvertToUserNetInfo(info *user_service.UserNetInfo) user.UserNetInfo {
	return user.UserNetInfo{
		ExternalIpAddress:  uint32(info.ExternalIpAddress),
		ExternalClientPort: uint16(info.ExternalClientPort),
		ExternalServerPort: uint16(info.ExternalServerPort),
		LocalIpAddress:     uint32(info.LocalIpAddress),
		LocalClientPort:    uint16(info.LocalClientPort),
		LocalServerPort:    uint16(info.LocalServerPort),
	}
}

func ConvertToServiceNetInfo(info *user.UserNetInfo) *user_service.UserNetInfo {
	return &user_service.UserNetInfo{
		ExternalIpAddress:  int32(info.ExternalIpAddress),
		ExternalClientPort: int16(info.ExternalClientPort),
		ExternalServerPort: int16(info.ExternalServerPort),
		LocalIpAddress:     int32(info.LocalIpAddress),
		LocalClientPort:    int16(info.LocalClientPort),
		LocalServerPort:    int16(info.LocalServerPort),
	}
}

func ConvertToRoomInfo(roomInfo *room_service.ChannelRoom) *server.Room {
	if roomInfo == nil {
		return nil
	}
	return &server.Room{
		RoomId:              uint16(roomInfo.RoomID),
		RoomNumber:          uint8(roomInfo.RoomNumber),
		HostUserID:          uint32(roomInfo.HostUserID),
		HostUserName:        roomInfo.HostUserName,
		RoomName:            roomInfo.RoomName,
		PassWd:              roomInfo.PassWd,
		GameModeID:          uint8(roomInfo.GameModeID),
		MapID:               uint8(roomInfo.MapID),
		MaxPlayers:          uint8(roomInfo.MaxPlayers),
		WinLimit:            uint8(roomInfo.WinLimit),
		KillLimit:           uint16(roomInfo.KillLimit),
		Status:              uint8(roomInfo.Status),
		StartMoney:          uint16(roomInfo.StartMoney),
		TeamBalanceType:     uint8(roomInfo.TeamBalanceType),
		AreBotsEnabled:      uint8(roomInfo.AreBotsEnabled),
		BotDifficulty:       uint8(roomInfo.BotDifficulty),
		NumCtBots:           uint8(roomInfo.NumCtBots),
		NumTrBots:           uint8(roomInfo.NumTrBots),
		ChangeTeams:         uint8(roomInfo.ChangeTeams),
		RespawnTime:         uint8(roomInfo.RespawnTime),
		NextMapEnabled:      uint8(roomInfo.NextMapEnabled),
		Difficulty:          uint8(roomInfo.Difficulty),
		IsIngame:            uint8(roomInfo.IsIngame),
		Users:               utils.Convertint32sToUint32s(roomInfo.UserIDs),
		ParentChannelServer: uint8(roomInfo.ParentServer),
		ParentChannel:       uint8(roomInfo.ParentChannel),
		LevelLimit:          uint8(roomInfo.LevelLimit),
		GameTime:            uint8(roomInfo.GameTime),
		StopTime:            uint8(roomInfo.StopTime),
		BuyLimitTime:        uint8(roomInfo.BuyLimitTime),
		ShowName:            uint8(roomInfo.ShowName),
		ShowFlash:           uint8(roomInfo.ShowFlash),
		ViewAngle:           uint8(roomInfo.ViewAngle),
		EnableVoice:         uint8(roomInfo.EnableVoice),
		LimitDeaths:         uint8(roomInfo.LimitDeaths),
		GameTimePerRound:    uint8(roomInfo.GameTimePerRound),
		DisableEnhancement:  uint8(roomInfo.DisableEnhancement),
		BombCountdown:       uint8(roomInfo.BombCountdown),
		FriendHurt:          uint8(roomInfo.FriendHurt),
		BotBalance:          uint8(roomInfo.BotBalance),
		WeaponRestrictions:  uint8(roomInfo.WeaponRestrictions),
		HostagePunish:       uint8(roomInfo.HostagePunish),
		Unk:                 convertToRoomUnk(roomInfo.Unk),
	}
}

func ConvertToServiceRoomInfo(roomInfo *server.Room) *room_service.ChannelRoom {
	if roomInfo == nil {
		return nil
	}
	return &room_service.ChannelRoom{
		RoomID:             int16(roomInfo.RoomId),
		RoomNumber:         int8(roomInfo.RoomNumber),
		HostUserID:         int32(roomInfo.HostUserID),
		HostUserName:       roomInfo.HostUserName,
		RoomName:           roomInfo.RoomName,
		PassWd:             roomInfo.PassWd,
		GameModeID:         int8(roomInfo.GameModeID),
		MapID:              int8(roomInfo.MapID),
		MaxPlayers:         int8(roomInfo.MaxPlayers),
		WinLimit:           int8(roomInfo.WinLimit),
		KillLimit:          int16(roomInfo.KillLimit),
		Status:             int8(roomInfo.Status),
		StartMoney:         int16(roomInfo.StartMoney),
		TeamBalanceType:    int8(roomInfo.TeamBalanceType),
		AreBotsEnabled:     int8(roomInfo.AreBotsEnabled),
		BotDifficulty:      int8(roomInfo.BotDifficulty),
		NumCtBots:          int8(roomInfo.NumCtBots),
		NumTrBots:          int8(roomInfo.NumTrBots),
		ChangeTeams:        int8(roomInfo.ChangeTeams),
		RespawnTime:        int8(roomInfo.RespawnTime),
		NextMapEnabled:     int8(roomInfo.NextMapEnabled),
		Difficulty:         int8(roomInfo.Difficulty),
		IsIngame:           int8(roomInfo.IsIngame),
		UserIDs:            utils.ConvertUint32sToInt32s(roomInfo.Users),
		ParentServer:       int8(roomInfo.ParentChannelServer),
		ParentChannel:      int8(roomInfo.ParentChannel),
		LevelLimit:         int8(roomInfo.LevelLimit),
		GameTime:           int8(roomInfo.GameTime),
		StopTime:           int8(roomInfo.StopTime),
		BuyLimitTime:       int8(roomInfo.BuyLimitTime),
		ShowName:           int8(roomInfo.ShowName),
		ShowFlash:          int8(roomInfo.ShowFlash),
		ViewAngle:          int8(roomInfo.ViewAngle),
		EnableVoice:        int8(roomInfo.EnableVoice),
		LimitDeaths:        int8(roomInfo.LimitDeaths),
		GameTimePerRound:   int8(roomInfo.GameTimePerRound),
		DisableEnhancement: int8(roomInfo.DisableEnhancement),
		BombCountdown:      int8(roomInfo.BombCountdown),
		FriendHurt:         int8(roomInfo.FriendHurt),
		BotBalance:         int8(roomInfo.BotBalance),
		WeaponRestrictions: int8(roomInfo.WeaponRestrictions),
		HostagePunish:      int8(roomInfo.HostagePunish),
		Unk:                convertToServiceUnk(&roomInfo.Unk),
	}
}

func convertToRoomUnk(Unk *room_service.RoomSettingUnk) server.RoomSettingUnk {
	return server.RoomSettingUnk{
		Unk00: uint8(Unk.Unk00),
		Unk01: uint8(Unk.Unk01),
		Unk02: uint8(Unk.Unk02),
		Unk03: uint8(Unk.Unk03),
		Unk04: uint32(Unk.Unk04),
		Unk14: uint8(Unk.Unk14),
		Unk15: uint8(Unk.Unk15),
		Unk16: uint8(Unk.Unk16),
		Unk17: uint8(Unk.Unk17),
		Unk18: uint8(Unk.Unk18),
		Unk19: uint8(Unk.Unk19),
		Unk23: convertToRoomUnk23Struct(Unk.Unk23),
		Unk24: uint32(Unk.Unk24),
		Unk25: Unk.Unk25,
		Unk26: uint8(Unk.Unk26),
		Unk27: uint8(Unk.Unk27),
		Unk28: uint8(Unk.Unk28),
		Unk29: uint8(Unk.Unk29),
		Unk36: uint8(Unk.Unk36),
		Unk37: uint8(Unk.Unk37),
		Unk38: uint8(Unk.Unk38),
		Unk39: uint8(Unk.Unk39),
		Unk40: convertToRoomUnk40(Unk.Unk40),
		Unk41: uint8(Unk.Unk41),
		Unk42: uint8(Unk.Unk42),
	}
}

func convertToServiceUnk(Unk *server.RoomSettingUnk) *room_service.RoomSettingUnk {
	return &room_service.RoomSettingUnk{
		Unk00: int8(Unk.Unk00),
		Unk01: int8(Unk.Unk01),
		Unk02: int8(Unk.Unk02),
		Unk03: int8(Unk.Unk03),
		Unk04: int32(Unk.Unk04),
		Unk14: int8(Unk.Unk14),
		Unk15: int8(Unk.Unk15),
		Unk16: int8(Unk.Unk16),
		Unk17: int8(Unk.Unk17),
		Unk18: int8(Unk.Unk18),
		Unk19: int8(Unk.Unk19),
		Unk23: convertToServiceUnk23Struct(Unk.Unk23),
		Unk24: int32(Unk.Unk24),
		Unk25: Unk.Unk25,
		Unk26: int8(Unk.Unk26),
		Unk27: int8(Unk.Unk27),
		Unk28: int8(Unk.Unk28),
		Unk29: int8(Unk.Unk29),
		Unk36: int8(Unk.Unk36),
		Unk37: int8(Unk.Unk37),
		Unk38: int8(Unk.Unk38),
		Unk39: int8(Unk.Unk39),
		Unk40: convertToServiceUnk40(Unk.Unk40),
		Unk41: int8(Unk.Unk41),
		Unk42: int8(Unk.Unk42),
	}
}

func convertToRoomUnk23Struct(Unk23 []*room_service.Unk23Struct) []server.Unk23Struct {
	rst := make([]server.Unk23Struct, len(Unk23))
	for i := range Unk23 {
		rst[i] = server.Unk23Struct{
			Unk23_unk00: uint32(Unk23[i].Unk23Unk00),
			Unk23_unk01: uint32(Unk23[i].Unk23Unk01),
			Unk23_unk02: uint8(Unk23[i].Unk23Unk02),
			Unk23_unk03: uint8(Unk23[i].Unk23Unk03),
			Unk23_unk04: uint8(Unk23[i].Unk23Unk04),
			Unk23_unk05: uint8(Unk23[i].Unk23Unk05),
			Unk23_unk06: uint16(Unk23[i].Unk23Unk06),
			Unk23_unk07: uint8(Unk23[i].Unk23Unk07),
			Unk23_unk08: uint8(Unk23[i].Unk23Unk08),
		}
	}
	return rst
}

func convertToServiceUnk23Struct(Unk23 []server.Unk23Struct) []*room_service.Unk23Struct {
	rst := make([]*room_service.Unk23Struct, len(Unk23))
	for i := range Unk23 {
		rst[i] = &room_service.Unk23Struct{
			Unk23Unk00: int32(Unk23[i].Unk23_unk00),
			Unk23Unk01: int32(Unk23[i].Unk23_unk01),
			Unk23Unk02: int8(Unk23[i].Unk23_unk02),
			Unk23Unk03: int8(Unk23[i].Unk23_unk03),
			Unk23Unk04: int8(Unk23[i].Unk23_unk04),
			Unk23Unk05: int8(Unk23[i].Unk23_unk05),
			Unk23Unk06: int16(Unk23[i].Unk23_unk06),
			Unk23Unk07: int8(Unk23[i].Unk23_unk07),
			Unk23Unk08: int8(Unk23[i].Unk23_unk08),
		}
	}
	return rst
}

func convertToRoomUnk40(Unk40 []*room_service.Unk40Struct) []server.Unk40Struct {
	rst := make([]server.Unk40Struct, len(Unk40))
	for i := range Unk40 {
		rst[i] = server.Unk40Struct{
			Unk40_unk00: uint8(Unk40[i].Unk40Unk00),
			Unk40_unk01: uint8(Unk40[i].Unk40Unk01),
		}
	}
	return rst
}

func convertToServiceUnk40(Unk40 []server.Unk40Struct) []*room_service.Unk40Struct {
	rst := make([]*room_service.Unk40Struct, len(Unk40))
	for i := range Unk40 {
		rst[i] = &room_service.Unk40Struct{
			Unk40Unk00: int8(Unk40[i].Unk40_unk00),
			Unk40Unk01: int8(Unk40[i].Unk40_unk01),
		}
	}
	return rst
}
