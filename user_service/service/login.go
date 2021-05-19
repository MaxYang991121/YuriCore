package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"sync"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type LoginService interface {
	Login(ctx context.Context) (*user.UserInfo, int8)
}

type loginServiceImpl struct {
	username  string
	password  string
	loginLock sync.Mutex
}

func NewLoginServiceImpl(username, password string) *loginServiceImpl {
	return &loginServiceImpl{
		username:  username,
		password:  password,
		loginLock: sync.Mutex{},
	}
}

func (l *loginServiceImpl) Login(ctx context.Context) (*user.UserInfo, int8) {
	l.loginLock.Lock()
	defer l.loginLock.Unlock()
	var user *user.UserInfo
	if conf.Config.EnableDataBase != 1 {
		// 创建用户
		user = NewRegisterServiceImpl(l.username, l.username, l.password).newUser_Default()
	} else {
		// 读取用户数据
		db_user, err := client.GetDBClient().GetUser(ctx, l.username)
		if err != nil {
			return nil, constant.Login_DB_ERROR
		}
		if db_user == nil || db_user.UserName == "" || db_user.Password == "" {
			return nil, constant.Login_Not_Registed
		}
		// 校验用户数据
		if !checkPassword(l.username, l.password, db_user.Password) {
			return nil, constant.Login_Wrong_PASSWORD
		}
		user = db_user
	}

	// 是否已经登录
	check_u := client.GetUserTableClient().GetUserByUserName(ctx, user.UserName)
	if check_u != nil && check_u.UserID != 0 {
		return nil, constant.Login_Already
	}
	// 获取用户ID
	user.UserID = client.GetUserTableClient().GetNewUserID(ctx)
	user.ChatTimes = 0xff
	// TODO fixme
	user.UserInventory.Items = full_inventory.Items
	// 存入table
	client.GetUserTableClient().AddUser(ctx, user)
	return user, constant.Login_Success
}

func checkPassword(username, password, DBpassword string) bool {
	str := fmt.Sprintf("%x", md5.Sum([]byte(username+password)))
	for i := 0; i < 32; i++ {
		if str[i] != DBpassword[i] {
			return false
		}
	}
	return true
}
