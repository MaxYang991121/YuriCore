package userTable

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
	. "github.com/KouKouChan/YuriCore/verbose"
)

type UserTableImpl struct {
	userNum     int64
	userNumLock sync.RWMutex
	table       sync.Map
	nameTable   sync.Map // 快表
	CurID       uint32
}

func NewUserTableImpl() *UserTableImpl {
	impl := &UserTableImpl{
		userNum:     0,
		userNumLock: sync.RWMutex{},
		table:       sync.Map{},
		nameTable:   sync.Map{},
	}
	if conf.Config.EnableDataBase == 1 {
		go func() {
			for {
				time.Sleep(60 * time.Second)
				DebugInfo(2, "Saving all users data")
				impl.table.Range(func(k, v interface{}) bool {
					client.GetDBClient().UpdateUser(context.TODO(), v.(*user.UserInfo))
					return true
				})
			}
		}()
	}
	return impl
}

func (u *UserTableImpl) GetUserByID(ctx context.Context, id uint32) *user.UserInfo {
	info, ok := u.table.Load(id)
	if ok {
		return info.(*user.UserInfo)
	}
	return nil
}

func (u *UserTableImpl) GetUserByUserName(ctx context.Context, username string) *user.UserInfo {
	id, ok := u.nameTable.Load(username)
	if ok {
		info, ok := u.table.Load(id)
		if ok {
			return info.(*user.UserInfo)
		}
	}
	return nil
}

func (u *UserTableImpl) GetNewUserID(ctx context.Context) uint32 {
	u.userNumLock.Lock()
	defer u.userNumLock.Unlock()
	if u.userNum >= math.MaxUint32 {
		return 0
	}

	u.CurID = (u.CurID % math.MaxUint32) + 1
	return u.CurID
}

func (u *UserTableImpl) UpdateUser(ctx context.Context, data *user.UserInfo) error {
	if data == nil || data.UserID == 0 || data.UserName == "" {
		return errors.New("userinfo is illegal !")
	}
	u.table.Store(data.UserID, data)
	u.nameTable.Store(data.UserName, data.UserID)

	return nil
}

func (u *UserTableImpl) DeleteUserByID(ctx context.Context, id uint32) error {
	info, loaded := u.table.LoadAndDelete(id)
	if loaded {
		if info != nil {
			u.table.Delete(info.(*user.UserInfo).UserID)
			u.userNumLock.Lock()
			u.userNum--
			u.userNumLock.Unlock()
		}
	}

	return nil
}

func (u *UserTableImpl) DeleteUserByName(ctx context.Context, username string) error {
	id, loaded := u.nameTable.LoadAndDelete(username)
	if loaded {
		u.table.Delete(id)

		u.userNumLock.Lock()
		u.userNum--
		u.userNumLock.Unlock()
	}

	return nil
}

func (u *UserTableImpl) AddUser(ctx context.Context, data *user.UserInfo) error {
	if data == nil || data.UserID == 0 || data.UserName == "" {
		return errors.New("userinfo is illegal !")
	}
	u.table.Store(data.UserID, data)
	u.nameTable.Store(data.UserName, data.UserID)

	u.userNumLock.Lock()
	u.userNum++
	u.userNumLock.Unlock()

	return nil
}

func (u *UserTableImpl) UpdateBag(ctx context.Context, UserID uint32, BagID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error) {
	if Slot > 3 {
		return nil, errors.New("Slot is too large")
	}

	var rst *user.UserInfo
	var err error
	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			if BagID < uint16(len(v.(*user.UserInfo).UserInventory.Loadouts)) {
				switch Slot {
				case 0:
					v.(*user.UserInfo).UserInventory.Loadouts[BagID].MainWeapon = ItemID
				case 1:
					v.(*user.UserInfo).UserInventory.Loadouts[BagID].SecondWeapon = ItemID
				case 2:
					v.(*user.UserInfo).UserInventory.Loadouts[BagID].Knife = ItemID
				case 3:
					v.(*user.UserInfo).UserInventory.Loadouts[BagID].Grenade = ItemID
				}
				rst = v.(*user.UserInfo)
				return false
			}
			err = errors.New("BagID is too large")
			return false
		}
		return true
	})

	return rst, err
}

func (u *UserTableImpl) UpdateBuymenu(ctx context.Context, UserID uint32, BuymenuID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error) {
	if Slot > 8 {
		return nil, errors.New("Slot is too large")
	}

	var rst *user.UserInfo
	var err error
	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			switch BuymenuID {
			case 0:
				v.(*user.UserInfo).UserInventory.BuyMenu.PistolsTR[Slot] = ItemID
			case 1:
				v.(*user.UserInfo).UserInventory.BuyMenu.ShotgunsTR[Slot] = ItemID
			case 2:
				v.(*user.UserInfo).UserInventory.BuyMenu.SmgsTR[Slot] = ItemID
			case 3:
				v.(*user.UserInfo).UserInventory.BuyMenu.RiflesTR[Slot] = ItemID
			case 4:
				v.(*user.UserInfo).UserInventory.BuyMenu.MachinegunsTR[Slot] = ItemID
			case 5:
				v.(*user.UserInfo).UserInventory.BuyMenu.EquipmentTR[Slot] = ItemID
			case 6:
				v.(*user.UserInfo).UserInventory.BuyMenu.ClassesTR[Slot] = ItemID
			case 14:
				v.(*user.UserInfo).UserInventory.BuyMenu.MeleesTR[Slot] = ItemID
			case 7:
				v.(*user.UserInfo).UserInventory.BuyMenu.PistolsCT[Slot] = ItemID
			case 8:
				v.(*user.UserInfo).UserInventory.BuyMenu.ShotgunsCT[Slot] = ItemID
			case 9:
				v.(*user.UserInfo).UserInventory.BuyMenu.SmgsCT[Slot] = ItemID
			case 10:
				v.(*user.UserInfo).UserInventory.BuyMenu.RiflesCT[Slot] = ItemID
			case 11:
				v.(*user.UserInfo).UserInventory.BuyMenu.MachinegunsCT[Slot] = ItemID
			case 12:
				v.(*user.UserInfo).UserInventory.BuyMenu.EquipmentCT[Slot] = ItemID
			case 13:
				v.(*user.UserInfo).UserInventory.BuyMenu.ClassesCT[Slot] = ItemID
			case 15:
				v.(*user.UserInfo).UserInventory.BuyMenu.MeleesCT[Slot] = ItemID
			}
			rst = v.(*user.UserInfo)
			return false
		}
		return true
	})

	return rst, err
}

func (u *UserTableImpl) UpdateCosmetics(ctx context.Context, UserID uint32, CosmeticsID uint8, cosmetics *user.UserCosmetics) (*user.UserInfo, error) {
	if cosmetics == nil {
		return nil, errors.New("cosmetics is null")
	}

	var rst *user.UserInfo
	var err error
	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			v.(*user.UserInfo).UserInventory.Cosmetics[CosmeticsID] = *cosmetics
			rst = v.(*user.UserInfo)
			return false
		}
		return true
	})

	return rst, err
}

func (u *UserTableImpl) UpdateCampaign(ctx context.Context, UserID uint32, CampaignID uint8) (*user.UserInfo, error) {
	var rst *user.UserInfo = nil
	var err error

	switch CampaignID {
	case 1, 2, 4, 8, 16, 32:
		u.table.Range(func(k, v interface{}) bool {
			if v != nil && v.(*user.UserInfo).UserID == UserID {
				if v.(*user.UserInfo).Campaign&CampaignID == 1 {
					err = errors.New("already completed campaign!")
				} else {
					v.(*user.UserInfo).Campaign |= CampaignID
					rst = v.(*user.UserInfo)
				}
				return false
			}
			return true
		})
		return rst, err
	default:
		return nil, errors.New("wrong campaign id")
	}

}
func (u *UserTableImpl) GetUserFriends(ctx context.Context, UserID uint32) ([]user.UserInfo, error) {
	rst := []user.UserInfo{}
	users := []string{}
	found := false
	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			users = v.(*user.UserInfo).Friends
			found = true
			return false
		}
		return true
	})

	if !found {
		return nil, errors.New("dest user not found")
	}
	if len(users) == 0 {
		return rst, nil
	}

	fastMap := map[string]bool{}
	for i := range users {
		fastMap[users[i]] = true
	}
	u.table.Range(func(k, v interface{}) bool {
		if v != nil && fastMap[v.(*user.UserInfo).UserName] {
			rst = append(rst, *(v.(*user.UserInfo)))
			return false
		}
		return true
	})

	return rst, nil
}

func (u *UserTableImpl) AddUserPoints(ctx context.Context, UserID uint32, num uint64) (uint64, error) {
	total := uint64(0)

	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			if math.MaxUint64-num < total {
				v.(*user.UserInfo).Points = math.MaxUint64
			} else {
				v.(*user.UserInfo).Points += num
			}
			total = v.(*user.UserInfo).Points
			return false
		}
		return true
	})

	return total, nil
}

func (u *UserTableImpl) AddUserCash(ctx context.Context, UserID uint32, num uint64) (uint64, error) {
	// TODO
	return 0, nil
}

func (u *UserTableImpl) PayPoints(ctx context.Context, UserID uint32, num uint64) (uint64, error) {
	total := uint64(0)
	var err error

	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			if num > v.(*user.UserInfo).Points {
				err = errors.New("user doesn't have enough points")
			} else {
				v.(*user.UserInfo).Points -= num
			}
			total = v.(*user.UserInfo).Points
			return false
		}
		return true
	})

	return total, err
}

func (u *UserTableImpl) PayCash(ctx context.Context, UserID uint32, num uint64) (uint64, error) {
	// TODO
	return 0, nil
}

func (u *UserTableImpl) AddFriend(ctx context.Context, UserID uint32, friend string) (*user.UserInfo, error) {
	var rst *user.UserInfo = nil
	var err error

	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			for i := range v.(*user.UserInfo).Friends {
				if v.(*user.UserInfo).Friends[i] == friend {
					err = errors.New("already friends")
					return false
				}
			}
			v.(*user.UserInfo).Friends = append(v.(*user.UserInfo).Friends, friend)
			rst = v.(*user.UserInfo)
			return false
		}
		return true
	})
	return rst, err
}

func (u *UserTableImpl) UpdateOption(ctx context.Context, UserID uint32, Options []byte) (*user.UserInfo, error) {
	var rst *user.UserInfo = nil
	var err error

	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			v.(*user.UserInfo).Options = Options
			rst = v.(*user.UserInfo)
			return false
		}
		return true
	})
	return rst, err

}

func (u *UserTableImpl) UpdateNickname(ctx context.Context, UserID uint32, nickname string) (*user.UserInfo, error) {
	var rst *user.UserInfo = nil
	var err error

	u.table.Range(func(k, v interface{}) bool {
		if v != nil && v.(*user.UserInfo).UserID == UserID {
			v.(*user.UserInfo).NickName = nickname
			rst = v.(*user.UserInfo)
			return false
		}
		return true
	})
	return rst, err

}
