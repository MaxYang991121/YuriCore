package controller

import (
	"context"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/service/serverlist"
)

type ServerListController interface {
	Handle(ctx context.Context) error
}

type serverListControllerImpl struct {
	client net.Conn
}

func GetServerListController(client net.Conn) ServerListController {
	serverlist := serverListControllerImpl{}

	serverlist.client = client

	return &serverlist
}

func (s *serverListControllerImpl) Handle(ctx context.Context) error {

	return serverlist.NewServerListService(s.client).Handler(ctx)
}
