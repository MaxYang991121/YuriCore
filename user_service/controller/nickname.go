package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/service"
)

type NickNameController interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type nicknameControllerImpl struct {
	userID   uint32
	nickname string
}

func NewNickNameController(userID uint32, nickname string) NickNameController {
	return &nicknameControllerImpl{
		userID:   userID,
		nickname: nickname,
	}
}

func (n *nicknameControllerImpl) Handle(ctx context.Context) (*user.UserInfo, error) {
	if n.nickname == "" ||
		n.userID == 0 ||
		len(n.nickname) > constant.MaxLen_UserName {
		return nil, errors.New("wrong userID or nickname")
	}

	return service.NewNickNameServiceImpl(n.userID, n.nickname).SetNickName(ctx)
}
