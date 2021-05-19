package start

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
	"github.com/KouKouChan/YuriCore/room_service/infrastructure/room_table"
)

func initRoomTable() {
	client.InitRoomTableClient(room_table.NewRoomTable())

	client.GetRoomTableClient().AddServer(
		context.TODO(),
		server.Server{
			ServerIndex:  0,
			ServerName:   "普通服务器[1/2]",
			ServerStatus: 1,
			ServerType:   0,
			Channels:     []server.Channel{},
		},
	)
	client.GetRoomTableClient().AddChannel(
		context.TODO(),
		1,
		server.Channel{
			ChannelIndex:  0,
			ChannelName:   "普通频道[1/2]",
			ChannelStatus: 0,
			ChannelType:   0,
			Rooms:         []server.Room{},
		},
	)
	client.GetRoomTableClient().AddChannel(
		context.TODO(),
		1,
		server.Channel{
			ChannelIndex:  0,
			ChannelName:   "普通频道[2/2]",
			ChannelStatus: 0,
			ChannelType:   0,
			Rooms:         []server.Room{},
		},
	)

	client.GetRoomTableClient().AddServer(
		context.TODO(),
		server.Server{
			ServerIndex:  0,
			ServerName:   "普通服务器[2/2]",
			ServerStatus: 1,
			ServerType:   0,
			Channels:     []server.Channel{},
		},
	)
	client.GetRoomTableClient().AddChannel(
		context.TODO(),
		2,
		server.Channel{
			ChannelIndex:  0,
			ChannelName:   "高手频道[1/1]",
			ChannelStatus: 0,
			ChannelType:   0,
			Rooms:         []server.Room{},
		},
	)

}
