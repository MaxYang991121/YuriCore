package udp

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
)

type UpdateUDPService interface {
	Handle(ctx context.Context) (uint16, error)
}

type updateUDPServiceImpl struct {
	portId            uint16
	localPort         uint16
	externalPort      uint16
	externalIPAddress uint32
	localIpAddress    uint32

	user *user.UserCache
}

func GetUpdateUDPService(portId uint16, localPort uint16, externalPort uint16, externalIPAddress, localIpAddress uint32, user *user.UserCache) UpdateUDPService {
	update := updateUDPServiceImpl{
		portId:            portId,
		localPort:         localPort,
		externalPort:      externalPort,
		externalIPAddress: externalIPAddress,
		localIpAddress:    localIpAddress,
		user:              user,
	}

	return &update
}

func (r *updateUDPServiceImpl) Handle(ctx context.Context) (uint16, error) {
	return client.GetUserCacheClient().FlushUserUDP(ctx, r.user.UserID, r.portId, r.localPort, r.externalPort, r.externalIPAddress, r.localIpAddress)
}
