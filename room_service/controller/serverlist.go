package controller

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type ServerListController interface {
	Handle(ctx context.Context) ([]server.Server, error)
}

type serverListControllerImpl struct {
}

func NewServerListController() ServerListController {
	return &serverListControllerImpl{}
}

func (u *serverListControllerImpl) Handle(ctx context.Context) ([]server.Server, error) {

	return service.NewServerListService().Handle(ctx)
}
