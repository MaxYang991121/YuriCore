package user_cache

import (
	"context"
	"errors"
	"net"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	. "github.com/KouKouChan/YuriCore/verbose"
)

type UserCacheImpl struct {
	cache sync.Map
}

func NewUserCacheImpl() *UserCacheImpl {
	return &UserCacheImpl{
		cache: sync.Map{},
	}
}

func (u *UserCacheImpl) GetUserByID(ctx context.Context, id uint32) *user.UserCache {
	info, ok := u.cache.Load(id)
	if ok {
		return info.(*user.UserCache)
	}
	DebugPrintf(2, "Call UserCache get user id=%+v failed", id)
	return nil
}

func (u *UserCacheImpl) GetUserByUserName(ctx context.Context, username string) *user.UserCache {
	var info *user.UserCache = nil
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserName == username {
			info = v.(*user.UserCache)
			DebugPrintf(2, "Call UserCache get user=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
	return info
}

func (u *UserCacheImpl) GetUserByConnection(ctx context.Context, client net.Conn) *user.UserCache {
	var info *user.UserCache = nil
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).CurrentConnection == client {
			info = v.(*user.UserCache)
			DebugPrintf(2, "Call UserCache get user=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
	return info
}

func (u *UserCacheImpl) SetUser(ctx context.Context, data *user.UserCache) error {
	if data == nil || data.UserID == 0 || data.UserName == "" {
		return errors.New("userinfo is illegal !")
	}
	u.cache.Store(data.UserID, data)
	DebugPrintf(2, "Call UserCache Set user=%+v", data)
	return nil
}

func (u *UserCacheImpl) DeleteUserByID(ctx context.Context, id uint32) {
	u.cache.Delete(id)
	DebugPrintf(2, "Call UserCache Deleted user id=%+v", id)
}

func (u *UserCacheImpl) DeleteUserByName(ctx context.Context, username string) {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserName == username {
			u.cache.Delete(k)
			DebugPrintf(2, "Call UserCache Deleted user=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
}

func (u *UserCacheImpl) DeleteUserByConnection(ctx context.Context, client net.Conn) {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).CurrentConnection == client {
			u.cache.Delete(k)
			DebugPrintf(2, "Call UserCache Deleted user=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
}

func (u *UserCacheImpl) SetUserChannel(ctx context.Context, userID uint32, serverID, channelID uint8) error {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentServerIndex = serverID
			v.(*user.UserCache).CurrentChannelIndex = channelID
			DebugPrintf(2, "Call UserCache set user channel=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) SetUserQuitChannel(ctx context.Context, userID uint32) error {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentServerIndex = 0
			v.(*user.UserCache).CurrentChannelIndex = 0
			DebugPrintf(2, "Call UserCache quit user channel=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) SetUserRoom(ctx context.Context, userID uint32, roomID uint16, team uint8) error {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentRoomId = roomID
			v.(*user.UserCache).CurrentTeam = team
			DebugPrintf(2, "Call SetUserRoom room=%+v", roomID)
			return false
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) QuitUserRoom(ctx context.Context, userID uint32) error {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentRoomId = 0
			v.(*user.UserCache).CurrentTeam = 0
			v.(*user.UserCache).Currentstatus = constant.UserNotReady
			DebugPrintf(2, "Call UserCache quit user room=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) FlushUserUDP(ctx context.Context, userID uint32, portId uint16, localPort uint16, externalPort uint16, externalIPAddress, localIpAddress uint32) (uint16, error) {
	rst := 0xff
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			switch portId {
			case constant.UDPTypeClient:
				v.(*user.UserCache).NetInfo.LocalClientPort = localPort
				v.(*user.UserCache).NetInfo.ExternalClientPort = externalPort
				rst = 0
			case constant.UDPTypeServer:
				v.(*user.UserCache).NetInfo.LocalServerPort = localPort
				v.(*user.UserCache).NetInfo.ExternalServerPort = externalPort
				rst = 1
			default:
				return false
			}
			v.(*user.UserCache).NetInfo.ExternalIpAddress = externalIPAddress
			v.(*user.UserCache).NetInfo.LocalIpAddress = localIpAddress
			return false
		}
		return true
	})
	return uint16(rst), nil
}

func (u *UserCacheImpl) SetUserStatus(ctx context.Context, userID uint32, status uint8) error {
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).Currentstatus = status
			DebugPrintf(2, "Call SetUserStatus user status=%+v", status)
			return false
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) FlushUserInventory(ctx context.Context, userID uint32, inventory *user.Inventory) error {
	if inventory == nil {
		return errors.New("null inventory")
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).UserInventory = *inventory
			DebugPrintf(2, "Call FlushUserInventory user inventory=%+v", v.(*user.UserCache))
			return false
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) GetChannelUsers(ctx context.Context, serverID, channelID uint8) []uint32 {
	rst := []uint32{}
	if serverID == 0 || channelID == 0 {
		return rst
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).CurrentServerIndex == serverID && v.(*user.UserCache).CurrentChannelIndex == channelID {
			rst = append(rst, v.(*user.UserCache).UserID)
		}
		return true
	})
	DebugPrintf(2, "Call GetChannelUsers users=%+v", rst)
	return rst
}

func (u *UserCacheImpl) SetUserIngame(ctx context.Context, userID uint32, ingame bool) error {
	if userID == 0 {
		return errors.New("wrong userid")
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentIsIngame = ingame
			if ingame {
				v.(*user.UserCache).Currentstatus = constant.UserIngame
			} else {
				v.(*user.UserCache).Currentstatus = constant.UserNotReady
			}
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) ResetKillNum(ctx context.Context, userID uint32) error {
	if userID == 0 {
		return errors.New("wrong userid")
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentKillNum = 0
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) ResetDeadNum(ctx context.Context, userID uint32) error {
	if userID == 0 {
		return errors.New("wrong userid")
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentDeathNum = 0
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) ResetAssistNum(ctx context.Context, userID uint32) error {
	if userID == 0 {
		return errors.New("wrong userid")
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentAssistNum = 0
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) GetChannelNoRoomUsers(ctx context.Context, serverID, channelID uint8) []uint32 {
	rst := []uint32{}
	if serverID == 0 || channelID == 0 {
		return rst
	}

	u.cache.Range(func(k, v interface{}) bool {
		if v != nil &&
			v.(*user.UserCache).CurrentServerIndex == serverID &&
			v.(*user.UserCache).CurrentChannelIndex == channelID &&
			v.(*user.UserCache).CurrentRoomId == 0 {
			rst = append(rst, v.(*user.UserCache).UserID) // no danger
		}
		return true
	})
	DebugPrintf(2, "Call GetChannelNoRoomUsers users=%+v", rst)
	return rst
}

func (u *UserCacheImpl) FlushUserRoomData(ctx context.Context, userID uint32, data []byte) error {
	if userID == 0 {
		return errors.New("wrong userid")
	}
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).CurrentRoomData = data
		}
		return true
	})
	return nil
}

func (u *UserCacheImpl) SetNickname(ctx context.Context, userID uint32, nickname string) error {
	if userID == 0 {
		return errors.New("wrong userid")
	}
	u.cache.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserCache).UserID == userID {
			v.(*user.UserCache).NickName = nickname
		}
		return true
	})
	return nil
}
