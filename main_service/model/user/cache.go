package user

import (
	"net"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/constant"
)

type UserCache struct {
	UserID   uint32
	UserName string
	NickName string

	UserInventory Inventory
	NetInfo       UserNetInfo
	//连接
	CurrentConnection net.Conn
	CurrentSequence   *uint8
	//频道房间信息
	CurrentServerIndex  uint8
	CurrentChannelIndex uint8
	CurrentRoomId       uint16
	CurrentTeam         uint8
	Currentstatus       uint8
	CurrentIsIngame     bool
	CurrentKillNum      uint16
	CurrentDeathNum     uint16
	CurrentAssistNum    uint16
	CurrentRoomData     []byte

	SequenceLocker sync.Mutex
}

func (u *UserCache) IsUserReady() bool {
	return u.Currentstatus != constant.UserNotReady
}

func (u *UserCache) GetNextSeq() uint8 {
	u.SequenceLocker.Lock()
	defer u.SequenceLocker.Unlock()

	if *u.CurrentSequence >= constant.MAXSEQUENCE {
		*u.CurrentSequence = 0
		return 0
	}
	(*u.CurrentSequence)++
	return *u.CurrentSequence
}
