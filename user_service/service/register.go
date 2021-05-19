package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"sync"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type RegisterService interface {
	Register(ctx context.Context) error
	newUser_Default() *user.UserInfo
}

type registerServiceImpl struct {
	username     string
	nickname     string
	password     string
	registerLock sync.Mutex
}

var (
	full_inventory user.Inventory
)

func init() {
	full_inventory = create_fullinventory()
}

func NewRegisterServiceImpl(username, nickname, password string) RegisterService {
	return &registerServiceImpl{
		username:     username,
		nickname:     nickname,
		password:     password,
		registerLock: sync.Mutex{},
	}
}

func (r *registerServiceImpl) Register(ctx context.Context) error {
	if conf.Config.EnableDataBase != 1 {
		return nil
	}
	r.registerLock.Lock()
	defer r.registerLock.Unlock()
	// 检查是否有重叠
	db_user, err := client.GetDBClient().GetUser(ctx, r.username)
	if err != nil &&
		err.Error() != "mongo: no documents in result" {
		return err
	}
	if db_user != nil {
		return errors.New("already registed!")
	}
	if r.nickname != "" {
		db_user, err = client.GetDBClient().GetUserByNickName(ctx, r.nickname)
		if err != nil &&
			err.Error() != "mongo: no documents in result" {
			return err
		}
		if db_user != nil {
			return errors.New("already registed!")
		}
	}
	// 新建用户
	user := r.newUser_Default()
	// 保存用户
	return client.GetDBClient().UpdateUser(ctx, user)
}

func create_fullinventory() user.Inventory {
	inventory := user.Inventory{
		BuyMenu: user.UserBuyMenu{
			PistolsTR:     [9]uint16{3, 6, 2, 4, 1},
			PistolsCT:     [9]uint16{3, 6, 2, 4, 5},
			ShotgunsTR:    [9]uint16{7, 8, 38},
			ShotgunsCT:    [9]uint16{7, 8},
			SmgsTR:        [9]uint16{10, 12, 13, 11, 37},
			SmgsCT:        [9]uint16{12, 13, 11, 36, 37},
			RiflesTR:      [9]uint16{23, 21, 17, 19, 14, 22, 34, 39, 114},
			RiflesCT:      [9]uint16{20, 17, 15, 16, 18, 14, 33, 35, 113},
			MachinegunsTR: [9]uint16{24, 32},
			MachinegunsCT: [9]uint16{24, 32},
			MeleesTR:      [9]uint16{161},
			MeleesCT:      [9]uint16{161},
			EquipmentTR:   [9]uint16{27, 28, 30, 31, 26, 25},
			EquipmentCT:   [9]uint16{27, 28, 30, 31, 26, 25, 29},
			ClassesTR:     [9]uint16{41, 42, 43, 40, 44, 46},
			ClassesCT:     [9]uint16{49, 50, 52, 51, 53, 56},
		},
		Loadouts: [3]user.UserLoadout{
			{428, 270, 363, 394},
			{344, 407, 405, 394},
			{466, 440, 420, 313},
		},
		Cosmetics: [5]user.UserCosmetics{},
	}

	for i := 1; i <= 57; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      72,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 113; i <= 114; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      138,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 148; i <= 149; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 150; i <= 155; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 157; i <= 161; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 163; i <= 164; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 169; i <= 172; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 174; i <= 176; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      187,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      197,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 200; i <= 203; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 211; i <= 216; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 220; i <= 222; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      233,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      237,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 242; i <= 249; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      251,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 253; i <= 255; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	// 加血
	for i := 258; i <= 259; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   100,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}
	// 战斗复活
	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      260,
		Count:   30,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 261; i <= 262; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 267; i <= 268; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 270; i <= 271; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 279; i <= 283; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	// 生命补液
	for i := 285; i <= 288; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	// 回合重置
	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      300,
		Count:   5,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 304; i <= 322; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 327; i <= 330; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	// 火力守卫
	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      334,
		Count:   100,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      337,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      339,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 341; i <= 347; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 350; i <= 365; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      367,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 370; i <= 372; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 376; i <= 377; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 379; i <= 380; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 382; i <= 385; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 388; i <= 395; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 397; i <= 400; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 405; i <= 407; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 413; i <= 435; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      437,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 440; i <= 454; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	for i := 456; i <= 460; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}

	inventory.Items = append(inventory.Items, user.UserInventoryItem{
		Id:      463,
		Count:   1,
		Existed: 1,
		Type:    1,
		Time:    0,
	})

	for i := 466; i <= 469; i++ {
		inventory.Items = append(inventory.Items, user.UserInventoryItem{
			Id:      uint16(i),
			Count:   1,
			Existed: 1,
			Type:    1,
			Time:    0,
		})
	}
	return inventory
}

func (r *registerServiceImpl) newUser_Default() *user.UserInfo {
	userinfo := &user.UserInfo{
		UserID:        0,
		UserName:      r.username,
		NickName:      r.nickname,
		Password:      fmt.Sprintf("%x", md5.Sum([]byte(r.username+r.password))),
		Level:         1,
		CurExp:        1,
		MaxExp:        1000,
		Points:        1000,
		PlayedMatches: 0,
		Wins:          0,
		Kills:         0,
		Deaths:        0,
		UserInventory: full_inventory,
		NetInfo:       user.UserNetInfo{},
		Rank:          1,
		ChatTimes:     0xff,
	}
	return userinfo
}
