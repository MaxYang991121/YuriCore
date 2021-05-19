package service

import (
	"context"
)

type UpdateService interface {
	Update(ctx context.Context) error
}

type updateServiceImpl struct {
}

func NewUpdateServiceImpl() *updateServiceImpl {
	return &updateServiceImpl{}
}

func (u *updateServiceImpl) Update(ctx context.Context) error {
	return nil
}
