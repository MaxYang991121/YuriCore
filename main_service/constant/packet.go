package constant

import "math"

const (
	HeaderLen   = 4
	MINSEQUENCE = 0
	MAXSEQUENCE = math.MaxUint8

	PacketTypeVersion          = 0
	PacketTypeReply            = 1
	PacketTypeNewCharacter     = 2
	PacketTypeLogin            = 3
	PacketTypeServerList       = 5
	PacketTypeCharacter        = 6
	PacketTypeRequestRoomList  = 7
	PacketTypeRequestChannels  = 10
	PacketTypeRoom             = 65
	PacketTypeChat             = 67
	PacketTypeHost             = 68
	PacketTypePlayerInfo       = 69
	PacketTypeUdp              = 70
	PacketTypeShop             = 72
	PacketTypeBan              = 74
	PacketTypeOption           = 76
	PacketTypeFavorite         = 77
	PacketTypeUseItem          = 78
	PacketTypeQuickJoin        = 80
	PacketTypeReport           = 83
	PacketTypeSignature        = 85
	PacketTypeQuickStart       = 86
	PacketTypeAutomatch        = 88
	PacketTypeFriend           = 89
	PacketTypeUnlock           = 90
	PacketTypeMail             = 91
	PacketTypeGZ               = 95
	PacketTypeAchievement      = 96
	PacketTypeSupply           = 102
	PacketTypeDisassemble      = 104
	PacketTypeConfigInfo       = 106
	PacketTypeUserStart        = 150
	PacketTypeRoomList         = 151
	PacketTypeInventory_Add    = 152
	PacketTypeLobby            = 153
	PacketTypeInventory_Create = 154
	PacketTypeUserInfo         = 157
	PacketTypeRegister         = 163

	UdpPacketSignature = 87
	UDPTypeClient      = 0
	UDPTypeServer      = 256
	UDPTypeSourceTV    = 512
)
