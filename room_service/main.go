package main

import (
	"fmt"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/room_service/start"
	"github.com/KouKouChan/YuriCore/utils"
)

func main() {
	ExePath, err := utils.GetExePath()
	if err != nil {
		panic(err)
	}

	start.Init(ExePath)
	initServer(ExePath, fmt.Sprintf("%s:%d", conf.Config.RoomAdress, conf.Config.RoomPort))
}
