package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/gen-go/yuricore/room_service"
	"github.com/KouKouChan/YuriCore/main_service/model/convert"
	"github.com/KouKouChan/YuriCore/room_service/controller"
	. "github.com/KouKouChan/YuriCore/verbose"
	"github.com/apache/thrift/lib/go/thrift"
)

type RoomServiceImpl struct{}

func (r *RoomServiceImpl) ServerList(ctx context.Context, req *room_service.ServerListRequest) (*room_service.ServerListResponse, error) {
	DebugPrintf(2, "ServerList called req=%+v", req)
	resp := room_service.NewServerListResponse()
	defer DebugPrintf(2, "Get ServerList resp=%+v", resp)

	servers, err := controller.NewServerListController().Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call ServerList err=%+v", err)
		return resp, err
	}

	for i := range servers {
		tmp_server := room_service.NewServer()
		tmp_server.ServerIndex = int8(servers[i].ServerIndex)
		tmp_server.ServerName = servers[i].ServerName
		tmp_server.ServerStatus = int8(servers[i].ServerStatus)
		tmp_server.ServerType = int8(servers[i].ServerType)
		for j := range servers[i].Channels {
			tmp_channel := room_service.NewChannel()
			tmp_channel.ChannelIndex = int8(servers[i].Channels[j].ChannelIndex)
			tmp_channel.ChannelName = servers[i].Channels[j].ChannelName
			tmp_channel.ChannelStatus = int8(servers[i].Channels[j].ChannelStatus)
			tmp_channel.ChannelType = int8(servers[i].Channels[j].ChannelType)
			tmp_server.Channels = append(tmp_server.Channels, tmp_channel)
		}
		resp.Servers = append(resp.Servers, tmp_server)
	}

	return resp, nil
}

func (r *RoomServiceImpl) RoomList(ctx context.Context, req *room_service.RoomListRequest) (*room_service.RoomListResponse, error) {
	DebugPrintf(2, "RoomList called req=%+v", req)
	resp := room_service.NewRoomListResponse()
	defer DebugPrintf(2, "Get RoomList resp=%+v", resp)

	rooms, err := controller.NewRoomListController(uint8(req.ServerIndex), uint8(req.ChannelIndex)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call RoomList err=%+v", err)
		return resp, err
	}

	for i := range rooms {
		resp.Rooms = append(resp.Rooms, convert.ConvertToServiceRoomInfo(&rooms[i]))
	}

	return resp, nil
}
func (r *RoomServiceImpl) NewRoom_(ctx context.Context, req *room_service.NewRoomRequest_) (*room_service.NewRoomResponse_, error) {
	DebugPrintf(2, "NewRoom called req=%+v", req)
	resp := room_service.NewNewRoomResponse_()
	resp.RoomInfo = nil
	defer DebugPrintf(2, "Get NewRoom_ resp=%+v", resp)

	if req.RoomInfo == nil {
		return resp, errors.New("null roominfo")
	}

	info, err := controller.NewNewRoomController(*convert.ConvertToRoomInfo(req.RoomInfo)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call NewRoom err=%+v", err)
		return resp, err

	}

	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)

	return resp, nil
}

func (r *RoomServiceImpl) UpdateRoom(ctx context.Context, req *room_service.UpdateRoomRequest) (*room_service.UpdateRoomResponse, error) {
	DebugPrintf(2, "UpdateRoom called req=%+v", req)
	resp := room_service.NewUpdateRoomResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get UpdateRoom resp=%+v", resp)

	if req.RoomInfo == nil {
		return resp, errors.New("null roominfo")
	}

	info, err := controller.NewUpdateRoomController(convert.ConvertToRoomInfo(req.RoomInfo), false).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call UpdateRoom err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func (r *RoomServiceImpl) JoinRoom(ctx context.Context, req *room_service.JoinRoomRequest) (*room_service.JoinRoomResponse, error) {
	DebugPrintf(2, "JoinRoom called req=%+v", req)
	resp := room_service.NewJoinRoomResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get JoinRoom resp=%+v", resp)

	info, err := controller.NewJoinRoomInfoController(uint16(req.RoomID), uint32(req.Player)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call JoinRoom err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func (r *RoomServiceImpl) LeaveRoom(ctx context.Context, req *room_service.LeaveRoomRequest) (*room_service.LeaveRoomResponse, error) {
	DebugPrintf(2, "LeaveRoom called req=%+v", req)
	resp := room_service.NewLeaveRoomResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get LeaveRoom resp=%+v", resp)

	info, err := controller.NewLeaveRoomInfoController(uint16(req.RoomID), uint32(req.Player)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call LeaveRoom err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func (r *RoomServiceImpl) StartGame(ctx context.Context, req *room_service.StartGameRequest) (*room_service.StartGameResponse, error) {
	DebugPrintf(2, "StartGame called req=%+v", req)
	resp := room_service.NewStartGameResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get StartGame resp=%+v", resp)

	info, err := controller.NewStartGameController(uint16(req.RoomID), uint32(req.HostID)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call StartGame err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func (r *RoomServiceImpl) StartCountdown(ctx context.Context, req *room_service.StartCountdownRequest) (*room_service.StartCountdownResponse, error) {
	return nil, nil
}

func (r *RoomServiceImpl) ToggleReady(ctx context.Context, req *room_service.ToggleReadyRequest) (*room_service.ToggleReadyResponse, error) {
	return nil, nil
}

func (r *RoomServiceImpl) GetRoomInfo(ctx context.Context, req *room_service.GetRoomInfoRequest) (*room_service.GetRoomInfoResponse, error) {
	DebugPrintf(2, "GetRoomInfo called req=%+v", req)
	resp := room_service.NewGetRoomInfoResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get GetRoomInfo resp=%+v", resp)

	if req.RoomID == 0 {
		return resp, errors.New("null roomid")
	}

	info, err := controller.NewGetRoomInfoController(uint16(req.RoomID)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call GetRoomInfo err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)

	return resp, nil
}

func (r *RoomServiceImpl) UpdateRoomSafe(ctx context.Context, req *room_service.UpdateRoomSafeRequest) (*room_service.UpdateRoomSafeResponse, error) {
	DebugPrintf(2, "UpdateRoomSafe called req=%+v", req)
	resp := room_service.NewUpdateRoomSafeResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get UpdateRoomSafe resp=%+v", resp)

	if req.RoomInfo == nil {
		return resp, errors.New("null roominfo")
	}

	info, err := controller.NewUpdateRoomController(convert.ConvertToRoomInfo(req.RoomInfo), true).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call UpdateRoomSafe err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func (r *RoomServiceImpl) DelRoom(ctx context.Context, req *room_service.DelRoomRequest) (*room_service.DelRoomResponse, error) {
	DebugPrintf(2, "DelRoom called req=%+v", req)
	resp := room_service.NewDelRoomResponse()
	resp.Success = false
	defer DebugPrintf(2, "Get DelRoom resp=%+v", resp)

	err := controller.NewDelRoomController(uint16(req.RoomID)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call UpdateRoomSafe err=%+v", err)
		return resp, err

	}

	resp.Success = true
	return resp, nil
}

func (r *RoomServiceImpl) SetRoomHost(ctx context.Context, req *room_service.SetRoomHostRequest) (*room_service.SetRoomHostResponse, error) {
	DebugPrintf(2, "SetRoomHost called req=%+v", req)
	resp := room_service.NewSetRoomHostResponse()
	resp.Success = false
	resp.RoomInfo = nil
	defer DebugPrintf(2, "Get SetRoomHost resp=%+v", resp)

	info, err := controller.NewSetHostController(uint16(req.RoomID), uint32(req.UserID), req.Name).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call SetRoomHost err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func (r *RoomServiceImpl) EndGame(ctx context.Context, req *room_service.EndGameRequest) (*room_service.EndGameResponse, error) {
	DebugPrintf(2, "EndGame called req=%+v", req)
	resp := room_service.NewEndGameResponse()
	resp.RoomInfo = nil
	resp.Success = false
	defer DebugPrintf(2, "Get EndGame resp=%+v", resp)

	info, err := controller.NewEndGameController(uint16(req.RoomID), uint32(req.HostID)).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Call EndGame err=%+v", err)
		resp.RoomInfo = nil
		return resp, err

	}

	resp.Success = true
	resp.RoomInfo = convert.ConvertToServiceRoomInfo(info)
	return resp, nil
}

func initServer(path string, addr string) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		panic(err)
	}

	handler := &RoomServiceImpl{}
	processor := room_service.NewRoomServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("room service is running at", fmt.Sprintf("%s:%d", conf.Config.RoomAdress, conf.Config.RoomPort))
	server.Serve()
}
