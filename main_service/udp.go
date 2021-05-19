package main

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/controller"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
	. "github.com/KouKouChan/YuriCore/verbose"
)

func initUDP() {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", conf.Config.MainPort))
	if err != nil {
		fmt.Println("Init udp addr error !\n")
		panic(err)
	}
	holepunchserver, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Init udp socket error !\n")
		panic(err)
	}
	defer holepunchserver.Close()
	StartHolePunchServer(strconv.Itoa(int(conf.Config.MainPort)), holepunchserver)
}

func StartHolePunchServer(port string, server *net.UDPConn) {
	defer server.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("UDP server suffered a fault !")
			fmt.Println("error:", err)
			fmt.Println("Fault end!")
		}
	}()
	fmt.Println("Server UDPholepunch is running at", "[AnyAdapter]:"+port)
	for {
		data := make([]byte, 1024)
		n, ClientAddr, err := server.ReadFromUDP(data)
		if err != nil {
			DebugInfo(2, "UDP read error from", ClientAddr.String())
			continue
		}
		go RecvHolePunchMessage(data[:n], n, ClientAddr, server)
	}
}

//RecvHolePunchMessage 处理收到的包
func RecvHolePunchMessage(data []byte, len int, UDPClient *net.UDPAddr, server *net.UDPConn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recv UDP packet suffered a fault !")
			fmt.Println(err)
			fmt.Println("Fault end!")
			return
		}
	}()
	//分析数据包
	var p packet.InUDPmsg
	if !p.PraseUDPpacket(data, len) {
		DebugInfo(2, "UDP had a illegal packet from", UDPClient.String())
		return
	}
	if p.IsHeartbeat() {
		return
	}
	cliadr := UDPClient.IP.To4().String()
	externalIPAddress, err := utils.IPToUint32(cliadr)
	if err != nil {
		DebugInfo(2, "Error : Prasing externalIpAddress error !")
		return
	}
	//找到对应玩家
	uPtr := client.GetUserCacheClient().GetUserByID(context.TODO(), p.UserId)
	if uPtr == nil ||
		uPtr.UserID <= 0 {
		return
	}
	//更新netinfo
	index, err := controller.GetUpdateUDPController(p.PortId, p.Port, uint16(UDPClient.Port), externalIPAddress, p.IpAddress, uPtr).Handle(context.TODO())
	if index == 0xFF || err != nil {
		return
	}

	//发送返回数据
	rst := out.BuildUDPHolepunch(index)
	server.WriteToUDP(rst, UDPClient)
}
