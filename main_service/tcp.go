package main

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/controller"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	. "github.com/KouKouChan/YuriCore/verbose"
)

func initTCP() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("TCP server suffered a fault !")
			fmt.Println("error:", err)
			fmt.Println("Fault end!")
		}
	}()

	server, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Config.MainPort))
	if err != nil {
		fmt.Println("Init tcp socket error !\n")
		panic(err)
	}
	defer server.Close()

	fmt.Println("Server is running at", "[AnyAdapter]:"+strconv.Itoa(int(conf.Config.MainPort)))
	for {
		client, err := server.Accept()
		if err != nil {
			DebugInfo(2, "Server accept data error !\n")
			continue
		}
		DebugInfo(2, "Server accept a new connection request at", client.RemoteAddr().String())
		client.Write([]byte("~SERVERCONNECTED\n"))
		go RecvMessage(client)
	}
}

//RecvMessage 循环处理收到的包
func RecvMessage(client net.Conn) {
	var dataPacket packet.PacketData
	var seq uint8
	packetController := controller.NewPacketController()

	defer client.Close() //关闭con
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Client", client.RemoteAddr().String(), "suffered a fault !")
			fmt.Println(err)
			fmt.Println("dump data", dataPacket.Data, "offset:", dataPacket.CurOffset)
			fmt.Println("Fault end!")
			controller.GetDisconnectController(client).Handle(context.TODO())
		}
	}()
	for {
		//读取4字节数据包头部
		headBytes, err := packetController.ReadHead(client)
		if err != nil {
			goto close
		}
		header := packetController.GetHeadPacket(headBytes)
		//读取数据部分
		databytes, err := packetController.ReadData(client, header.Length)
		if err != nil {
			goto close
		}
		dataPacket = packetController.GetDataPacket(header, databytes)

		DebugPrintf(2, "data packet %+v from %s", dataPacket, client.RemoteAddr().String())
		//执行功能
		switch dataPacket.Id {
		case constant.PacketTypeLogin:
			err = controller.GetLoginController(client, &dataPacket, &seq).Handle(context.TODO())
		case constant.PacketTypeRegister:
			err = controller.GetRegisterController(client, &dataPacket, &seq).Handle(context.TODO())
		case constant.PacketTypeNewCharacter:
			err = controller.GetNewCharacter(client, &dataPacket, &seq).Handle(context.TODO())
		case constant.PacketTypeChat:
			err = controller.GetChatController(client, &dataPacket).Handle(context.TODO())
		case constant.PacketTypeRequestChannels:
			err = controller.GetServerListController(client).Handle(context.TODO())
		case constant.PacketTypeRequestRoomList:
			err = controller.GetRoomListController(&dataPacket, client).Handle(context.TODO())
		case constant.PacketTypeRoom:
			err = controller.GetRoomController(&dataPacket, client).Handle(context.TODO())
		case constant.PacketTypeVersion:
			err = controller.GetVersionController(&dataPacket, client).Handle()
		case constant.PacketTypeFavorite:
			err = controller.GetFavorateController(&dataPacket, client).Handle(context.TODO())
		case constant.PacketTypePlayerInfo:
			err = controller.GetPlayerInfoController(&dataPacket, client).Handle(context.TODO())
		case constant.PacketTypeHost:
			err = controller.GetHostController(&dataPacket, client).Handle(context.TODO())
		case constant.PacketTypeOption:
			err = controller.GetOptionController(&dataPacket, client).Handle(context.TODO())
		default:
			DebugInfo(2, "Unknown packet", dataPacket.Id, "from", client.RemoteAddr().String())
		}
		if err != nil {
			DebugInfo(2, "handle packet ID=", dataPacket.Id, "failed ! error=", err)
		}
	}

close:
	controller.GetDisconnectController(client).Handle(context.TODO())
	DebugInfo(1, "client", client.RemoteAddr().String(), "closed the connection")
}
