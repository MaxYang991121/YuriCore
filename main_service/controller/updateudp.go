package controller

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/main_service/service/udp"
)

type UpdateUDPController interface {
	Handle(ctx context.Context) (uint16, error)
}

type updateUDPControllerImpl struct {
	portId            uint16
	localPort         uint16
	externalPort      uint16
	externalIPAddress uint32
	localIpAddress    uint32

	user *user.UserCache
}

func GetUpdateUDPController(portId uint16, localPort uint16, externalPort uint16, externalIPAddress, localIpAddress uint32, user *user.UserCache) UpdateUDPController {
	update := updateUDPControllerImpl{
		portId:            portId,
		localPort:         localPort,
		externalPort:      externalPort,
		externalIPAddress: externalIPAddress,
		localIpAddress:    localIpAddress,
		user:              user,
	}

	return &update
}

func (r *updateUDPControllerImpl) Handle(ctx context.Context) (uint16, error) {

	return udp.GetUpdateUDPService(r.portId, r.localPort, r.externalPort, r.externalIPAddress, r.localIpAddress, r.user).Handle(ctx)
}
