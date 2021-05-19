package room_service

import (
	"context"
	"errors"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/gen-go/yuricore/room_service"
	"github.com/KouKouChan/YuriCore/main_service/model/convert"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	. "github.com/KouKouChan/YuriCore/verbose"
	"github.com/apache/thrift/lib/go/thrift"
)

type RoomServiceImpl struct {
	client *room_service.RoomServiceClient
	lock   sync.Mutex
}

func NewRoomServiceImpl(addr string) *RoomServiceImpl {
	var transport thrift.TTransport
	var err error
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err = thrift.NewTSocket(addr)

	if err != nil {
		panic(err)
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		panic(err)
	}
	if err := transport.Open(); err != nil {
		panic(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return &RoomServiceImpl{
		client: room_service.NewRoomServiceClient(thrift.NewTStandardClient(iprot, oprot)),
	}
}

func (r *RoomServiceImpl) GetServiceList(ctx context.Context) ([]server.Server, error) {
	req := room_service.NewServerListRequest()

	DebugPrintf(2, "Call GetServiceList req=%+v", req)
	r.lock.Lock()
	resp, err := r.client.ServerList(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call GetServiceList resp=%+v", resp)

	if resp == nil || resp.Servers == nil {
		return nil, errors.New("Call GetServiceList got null resp")
	}

	rst := []server.Server{}

	for i := range resp.Servers {
		tmp_server := server.Server{
			ServerIndex:  uint8(resp.Servers[i].ServerIndex),
			ServerName:   resp.Servers[i].ServerName,
			ServerStatus: uint8(resp.Servers[i].ServerStatus),
			ServerType:   uint8(resp.Servers[i].ServerType),
			Channels:     []server.Channel{},
		}
		for j := range resp.Servers[i].Channels {
			tmp_channel := server.Channel{
				ChannelIndex:  uint8(resp.Servers[i].Channels[j].ChannelIndex),
				ChannelName:   resp.Servers[i].Channels[j].ChannelName,
				ChannelStatus: uint8(resp.Servers[i].Channels[j].ChannelStatus),
				ChannelType:   uint8(resp.Servers[i].Channels[j].ChannelType),
			}
			tmp_server.Channels = append(tmp_server.Channels, tmp_channel)
		}
		rst = append(rst, tmp_server)
	}

	return rst, nil
}

func (r *RoomServiceImpl) GetRoomList(ctx context.Context, ServerIndex, ChannelIndex uint8) ([]server.Room, error) {
	req := room_service.NewRoomListRequest()
	req.ServerIndex = int8(ServerIndex)
	req.ChannelIndex = int8(ChannelIndex)

	DebugPrintf(2, "Call GetRoomList req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.RoomList(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call GetRoomList resp=%+v", resp)

	if resp == nil || resp.Rooms == nil {
		return nil, errors.New("Call GetRoomList got null resp")
	}

	rst := []server.Room{}

	for i := range resp.Rooms {
		rst = append(rst, *convert.ConvertToRoomInfo(resp.Rooms[i]))
	}

	return rst, nil
}

func (r *RoomServiceImpl) NewRoom(ctx context.Context, room *server.Room) (*server.Room, error) {
	if room == nil {
		return nil, errors.New("wrong room info")
	}

	users := make([]int32, len(room.Users))
	for i := range room.Users {
		users[i] = int32(room.Users[i])
	}

	req := room_service.NewNewRoomRequest_()
	req.RoomInfo = convert.ConvertToServiceRoomInfo(room)

	DebugPrintf(2, "Call NewRoom req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.NewRoom_(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call NewRoom resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call NewRoom got null resp")
	}

	users_uint32 := make([]uint32, len(resp.RoomInfo.UserIDs))
	for i := range resp.RoomInfo.UserIDs {
		users_uint32[i] = uint32(resp.RoomInfo.UserIDs[i])
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) UpdateRoom(ctx context.Context, room *server.Room) (*server.Room, error) {
	if room == nil {
		return nil, errors.New("wrong room info")
	}

	users := make([]int32, len(room.Users))
	for i := range room.Users {
		users[i] = int32(room.Users[i])
	}

	req := room_service.NewUpdateRoomRequest()
	req.RoomInfo = convert.ConvertToServiceRoomInfo(room)

	DebugPrintf(2, "Call UpdateRoom req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.UpdateRoom(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateRoom resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call UpdateRoom got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) JoinRoom(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error) {
	req := room_service.NewJoinRoomRequest()
	req.RoomID = int16(roomID)
	req.Player = int32(userID)

	DebugPrintf(2, "Call JoinRoom req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.JoinRoom(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call JoinRoom resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call JoinRoom got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) LeaveRoom(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error) {
	req := room_service.NewLeaveRoomRequest()
	req.RoomID = int16(roomID)
	req.Player = int32(userID)

	DebugPrintf(2, "Call LeaveRoom req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.LeaveRoom(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call LeaveRoom resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call LeaveRoom got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) StartGame(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error) {
	req := room_service.NewStartGameRequest()
	req.RoomID = int16(roomID)
	req.HostID = int32(userID)

	DebugPrintf(2, "Call StartGame req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.StartGame(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call StartGame resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call StartGame got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) GetRoomInfo(ctx context.Context, roomID uint16) (*server.Room, error) {
	req := room_service.NewGetRoomInfoRequest()
	req.RoomID = int16(roomID)

	DebugPrintf(2, "Call GetRoomInfo req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.GetRoomInfo(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call GetRoomInfo resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call GetRoomInfo got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) UpdateRoomSafe(ctx context.Context, room *server.Room) (*server.Room, error) {
	if room == nil {
		return nil, errors.New("wrong room info")
	}

	users := make([]int32, len(room.Users))
	for i := range room.Users {
		users[i] = int32(room.Users[i])
	}

	req := room_service.NewUpdateRoomSafeRequest()
	req.RoomInfo = convert.ConvertToServiceRoomInfo(room)

	DebugPrintf(2, "Call UpdateRoomSafe req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.UpdateRoomSafe(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateRoomSafe resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call UpdateRoomSafe got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) SetRoomHost(ctx context.Context, userID uint32, name string, roomID uint16) (*server.Room, error) {
	if userID == 0 || roomID == 0 {
		return nil, errors.New("wrong roomid or userid")
	}

	req := room_service.NewSetRoomHostRequest()
	req.Name = name
	req.RoomID = int16(roomID)
	req.UserID = int32(userID)

	DebugPrintf(2, "Call SetRoomHost req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.SetRoomHost(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call SetRoomHost resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call SetRoomHost got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}

func (r *RoomServiceImpl) DelRoom(ctx context.Context, roomID uint16) error {
	if roomID == 0 {
		return errors.New("wrong roomid")
	}

	req := room_service.NewDelRoomRequest()
	req.RoomID = int16(roomID)

	DebugPrintf(2, "Call DelRoom req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.DelRoom(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return err
	}
	DebugPrintf(2, "Call DelRoom resp=%+v", resp)

	if resp == nil ||
		!resp.Success {
		return errors.New("Call DelRoom failed")
	}

	return nil
}
func (r *RoomServiceImpl) EndGame(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error) {
	req := room_service.NewEndGameRequest()
	req.RoomID = int16(roomID)
	req.HostID = int32(userID)

	DebugPrintf(2, "Call EndGame req=%+v", req)

	r.lock.Lock()
	resp, err := r.client.EndGame(ctx, req)
	r.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call EndGame resp=%+v", resp)

	if resp == nil ||
		resp.RoomInfo == nil ||
		resp.RoomInfo.RoomID == 0 ||
		resp.RoomInfo.RoomNumber == 0 {
		return nil, errors.New("Call EndGame got null resp")
	}

	rst := convert.ConvertToRoomInfo(resp.RoomInfo)

	return rst, nil
}
