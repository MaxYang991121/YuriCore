package start

import (
	"fmt"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/infrastructure/room_service"
)

func initRoomService() {
	client.InitRoomClient(room_service.NewRoomServiceImpl(fmt.Sprintf("%s:%d", conf.Config.RoomAdress, conf.Config.RoomPort)))
}
