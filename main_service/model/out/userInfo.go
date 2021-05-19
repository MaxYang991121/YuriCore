package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

//发送出去的包结构，其中一些未知，知道后会加入user里去
type OUTUserInfo struct {
	UserID uint32
	//flags                uint32 // should always be 0xFFFFFFFF for a full update
	Unk00         uint8
	UserName      string
	NickName      string
	Unk01         uint8
	Unk02         uint8
	Unk03         uint8
	Level         uint8
	Unk04         uint8
	CurExp        uint64
	MaxExp        uint64
	Points        uint64
	PlayedMatches uint32
	Wins          uint32
	Kills         uint32
	Deaths        uint32
	Headshots     uint32
	Unk05         uint32
	Regionname    string
	Unk06         uint16
	Unk07         uint16
	Unk08         uint16
	NetCafeName   string
	Unk09         uint32
	Unk10         uint8
	Unk11         uint32
	Unk12         uint32
	Unk13         string
	Unk14         uint8
	Unk15         uint8
	Unk16         uint8
	Unk17         uint8
	Rank          uint32
	Unk19         uint8
	Campaign      uint8
	Unk21         uint16
	Unk22         uint32
	Unk23         uint16
	Unk24         uint32
	Unk25         [128]uint8
	ChatTimes     uint8
	Unk28         uint32
}

func NewUserInfo(u *user.UserInfo) OUTUserInfo {
	return OUTUserInfo{
		UserID:        u.UserID,
		Unk00:         0x00,
		UserName:      u.UserName,
		NickName:      u.NickName,
		Unk01:         0x01,
		Unk02:         0x01,
		Unk03:         0x01,
		Level:         u.Level,
		Unk04:         0x00,
		CurExp:        u.CurExp,
		MaxExp:        u.MaxExp,
		Points:        u.Points,
		PlayedMatches: u.PlayedMatches,
		Wins:          u.Wins,
		Kills:         u.Kills,
		Deaths:        u.Deaths,
		Headshots:     0,
		Unk05:         1,
		Regionname:    "",
		Unk06:         0,
		Unk07:         1,
		Unk08:         1,
		NetCafeName:   "",
		Unk09:         0,
		Unk10:         0,
		Unk11:         0,
		Unk12:         0,
		Unk13:         "",
		Unk14:         0,
		Unk15:         0,
		Unk16:         0,
		Unk17:         0,
		Rank:          u.Rank,
		Unk19:         0,
		Campaign:      u.Campaign,
		Unk21:         0,
		Unk22:         0,
		Unk23:         0,
		Unk24:         0,
		Unk25: [128]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		ChatTimes: u.ChatTimes,
		Unk28:     0,
	}
}

func BuildUserInfo(info OUTUserInfo, needID bool, flags uint32) []byte {
	infobuf := make([]byte, 2048)
	offset := 0
	if needID {
		utils.WriteUint32(&infobuf, info.UserID, &offset)
	}
	utils.WriteUint32(&infobuf, flags, &offset)
	if flags&0x1 != 0 {
		utils.WriteUint8(&infobuf, info.Unk00, &offset)
	}
	if flags&0x2 != 0 {
		utils.WriteStringWithNull(&infobuf, []byte(info.UserName), &offset)
	}
	if flags&0x4 != 0 {
		utils.WriteStringWithNull(&infobuf, []byte(info.NickName), &offset)
		utils.WriteUint8(&infobuf, info.Unk01, &offset)
		utils.WriteUint8(&infobuf, info.Unk02, &offset)
		utils.WriteUint8(&infobuf, info.Unk03, &offset)
	}
	if flags&0x8 != 0 {
		utils.WriteUint8(&infobuf, info.Level, &offset)
	}
	if flags&0x10 != 0 {
		utils.WriteUint8(&infobuf, info.Unk04, &offset)
	}
	if flags&0x20 != 0 {
		utils.WriteUint64(&infobuf, info.CurExp, &offset)
	}
	if flags&0x40 != 0 {
		utils.WriteUint64(&infobuf, info.MaxExp, &offset)
	}
	if flags&0x80 != 0 {
		utils.WriteUint64(&infobuf, info.Points, &offset)
	}
	if flags&0x100 != 0 {
		utils.WriteUint32(&infobuf, info.PlayedMatches, &offset)
		utils.WriteUint32(&infobuf, info.Wins, &offset)
		utils.WriteUint32(&infobuf, info.Kills, &offset)
		utils.WriteUint32(&infobuf, info.Headshots, &offset)
		utils.WriteUint32(&infobuf, info.Deaths, &offset)
		utils.WriteUint32(&infobuf, info.Unk05, &offset)
	}
	if flags&0x200 != 0 {
		utils.WriteStringWithNull(&infobuf, []byte(info.Regionname), &offset)
		utils.WriteUint16(&infobuf, info.Unk06, &offset)
		utils.WriteUint16(&infobuf, info.Unk07, &offset)
		utils.WriteUint16(&infobuf, info.Unk08, &offset)
		utils.WriteStringWithNull(&infobuf, []byte(info.NetCafeName), &offset)
	}
	if flags&0x400 != 0 {
		utils.WriteUint32(&infobuf, info.Unk09, &offset)
	}
	if flags&0x800 != 0 {
		utils.WriteUint8(&infobuf, info.Unk10, &offset)
	}
	if flags&0x1000 != 0 {
		utils.WriteUint32(&infobuf, info.Unk11, &offset)
		utils.WriteUint32(&infobuf, info.Unk12, &offset)
		utils.WriteStringWithNull(&infobuf, []byte(info.Unk13), &offset)
		utils.WriteUint8(&infobuf, info.Unk14, &offset)
		utils.WriteUint8(&infobuf, info.Unk15, &offset)
		utils.WriteUint8(&infobuf, info.Unk16, &offset)
	}
	if flags&0x2000 != 0 {
		utils.WriteUint8(&infobuf, info.Unk17, &offset)
	}
	if flags&0x4000 != 0 {
		utils.WriteUint32(&infobuf, info.Rank, &offset)
	}
	if flags&0x8000 != 0 {
		utils.WriteUint8(&infobuf, info.Unk19, &offset)
		utils.WriteUint8(&infobuf, info.Campaign, &offset)
	}
	if flags&0x10000 != 0 {
		utils.WriteUint16(&infobuf, info.Unk21, &offset)
	}
	if flags&0x20000 != 0 {
		utils.WriteUint32(&infobuf, info.Unk22, &offset)
	}
	if flags&0x40000 != 0 {
		utils.WriteUint16(&infobuf, info.Unk23, &offset)
		utils.WriteUint32(&infobuf, info.Unk24, &offset)
	}
	if flags&0x80000 != 0 {
		for i := range info.Unk25 {
			utils.WriteUint8(&infobuf, info.Unk25[i], &offset)
		}
	}
	if flags&0x200000 != 0 {
		utils.WriteUint8(&infobuf, info.ChatTimes, &offset)
		utils.WriteUint32(&infobuf, info.Unk28, &offset)
	}
	return infobuf[:offset]
}
