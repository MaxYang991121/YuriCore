package controller

import (
	"context"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/service/disconnect"
)

type DisconnectController interface {
	Handle(ctx context.Context) error
}

type disconnectControllerImpl struct {
	client net.Conn
}

func GetDisconnectController(client net.Conn) DisconnectController {
	impl := disconnectControllerImpl{}

	impl.client = client

	return &impl
}

func (d *disconnectControllerImpl) Handle(ctx context.Context) error {
	return disconnect.NewDisconnectService(d.client).Handle(ctx)
}
