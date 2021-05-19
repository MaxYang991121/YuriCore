package service

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/client"
)

type NickNameService interface {
	SetNickName(ctx context.Context) (*user.UserInfo, error)
}

type nicknameServiceImpl struct {
	userID   uint32
	nickname string
}

func NewNickNameServiceImpl(userID uint32, nickname string) *nicknameServiceImpl {
	return &nicknameServiceImpl{
		userID:   userID,
		nickname: nickname,
	}
}

func (n *nicknameServiceImpl) SetNickName(ctx context.Context) (*user.UserInfo, error) {
	// 检查是否重复
	db_user, err := client.GetDBClient().GetUserByNickName(ctx, n.nickname)
	if err != nil &&
		err.Error() != "mongo: no documents in result" {
		return nil, err
	}
	if db_user != nil {
		return nil, errors.New("nickname existed!")
	}

	return client.GetUserTableClient().UpdateNickname(ctx, n.userID, n.nickname)
}
