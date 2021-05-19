package controller

import "context"

type UpdateController interface {
	Handle(ctx context.Context) error
}

type updateControllerImpl struct {
}

func NewUpdateController() UpdateController {
	return &updateControllerImpl{}
}

func (u *updateControllerImpl) Handle(ctx context.Context) error {
	return nil
}
