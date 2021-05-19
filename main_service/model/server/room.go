package server

type (
	Unk40Struct struct {
		Unk40_unk00 uint8
		Unk40_unk01 uint8
	}
	Unk23Struct struct {
		Unk23_unk00 uint32
		Unk23_unk01 uint32
		Unk23_unk02 uint8
		Unk23_unk03 uint8
		Unk23_unk04 uint8
		Unk23_unk05 uint8
		Unk23_unk06 uint16
		Unk23_unk07 uint8
		Unk23_unk08 uint8
	}
	RoomSettingUnk struct {
		Unk00 uint8
		Unk01 uint8
		Unk02 uint8
		Unk03 uint8
		Unk04 uint32
		Unk14 uint8
		Unk15 uint8
		Unk16 uint8
		Unk17 uint8
		Unk18 uint8
		Unk19 uint8
		Unk23 []Unk23Struct
		Unk24 uint32
		Unk25 string
		Unk26 uint8
		Unk27 uint8
		Unk28 uint8
		Unk29 uint8
		Unk36 uint8
		Unk37 uint8
		Unk38 uint8
		Unk39 uint8
		Unk40 []Unk40Struct
		Unk41 uint8
		Unk42 uint8
	}
	Room struct {
		RoomId       uint16
		RoomNumber   uint8
		HostUserID   uint32
		HostUserName string
		CanSpec      uint8
		IsVipRoom    uint8
		VipRoomLevel uint8

		//设置
		RoomName           string
		PassWd             string
		GameModeID         uint8
		MapID              uint8
		MaxPlayers         uint8
		WinLimit           uint8
		KillLimit          uint16
		LevelLimit         uint8
		GameTime           uint8
		GameTimePerRound   uint8
		WeaponRestrictions uint8
		Status             uint8
		HostagePunish      uint8
		StopTime           uint8
		BuyLimitTime       uint8
		ShowName           uint8
		ShowFlash          uint8
		ViewAngle          uint8
		EnableVoice        uint8
		LimitDeaths        uint8
		TeamBalanceType    uint8
		AreBotsEnabled     uint8
		BotDifficulty      uint8
		NumCtBots          uint8
		NumTrBots          uint8
		BotBalance         uint8
		StartMoney         uint16
		ChangeTeams        uint8
		RespawnTime        uint8
		NextMapEnabled     uint8
		Difficulty         uint8
		IsIngame           uint8
		ForceCamera        uint8
		DisableEnhancement uint8
		BombCountdown      uint8
		FriendHurt         uint8

		CountingDown        bool
		Countdown           uint8
		Users               []uint32
		ParentChannelServer uint8
		ParentChannel       uint8
		CtScore             uint8
		TrScore             uint8
		CtKillNum           uint32
		TrKillNum           uint32
		WinnerTeam          uint8
		Cache               []byte
		PageNum             uint8

		Unk RoomSettingUnk
	}
)
