package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/user_service/service"
)

type DownController interface {
	Handle(ctx context.Context) error
}

type downControllerImpl struct {
	userID uint32
}

func NewDownController(userID uint32) DownController {
	return &downControllerImpl{
		userID: userID,
	}
}

func (d *downControllerImpl) Handle(ctx context.Context) error {
	if d.userID == 0 {
		return errors.New("wrong user info")
	}

	return service.NewDownService(d.userID).Handle(ctx)
}
