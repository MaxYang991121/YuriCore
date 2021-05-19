package room_table

import (
	"context"
	"errors"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
)

type RoomTable struct {
	serverlist []server.Server
	idMap      sync.Map
	lock       sync.RWMutex
}

func NewRoomTable() *RoomTable {
	return &RoomTable{
		serverlist: []server.Server{},
		idMap:      sync.Map{},
		lock:       sync.RWMutex{},
	}
}

func (r *RoomTable) AddRoom(ctx context.Context, data *server.Room) (*server.Room, error) {
	if data == nil || data.ParentChannelServer == 0 || data.ParentChannel == 0 {
		return nil, errors.New("wrong room info")
	}

	newRoomNumber := uint8(0)
	newRoomID := uint16(0)

	r.lock.Lock()
	defer r.lock.Unlock()

	// 分配ID
	for i := 1; i <= server.MAXSERVERROOM; i++ {
		if _, ok := r.idMap.Load(uint16(i)); !ok {
			newRoomID = uint16(i)
			break
		}
	}
	if newRoomID == 0 {
		return nil, errors.New("reach server room limit")
	}

	// 查找空位
	for i := range r.serverlist {
		if r.serverlist[i].ServerIndex == data.ParentChannelServer {
			for j := range r.serverlist[i].Channels {
				if r.serverlist[i].Channels[j].ChannelIndex == data.ParentChannel {
					if len(r.serverlist[i].Channels[j].Rooms) >= server.MAXCHANNELROOM {
						return nil, errors.New("reach channel room limit")
					}
					numberMap := map[uint8]bool{}
					// 分配Number
					for k := range r.serverlist[i].Channels[j].Rooms {
						numberMap[r.serverlist[i].Channels[j].Rooms[k].RoomNumber] = true
					}
					for tmp_number := 1; tmp_number <= server.MAXCHANNELROOM; tmp_number++ {
						if !numberMap[uint8(tmp_number)] {
							newRoomNumber = uint8(tmp_number)
							break
						}
					}
					// 设置ID
					data.RoomId = newRoomID
					data.RoomNumber = newRoomNumber
					r.idMap.Store(newRoomID, "true")
					// 添加房间
					r.serverlist[i].Channels[j].Rooms = append(r.serverlist[i].Channels[j].Rooms, *data)

					return data, nil
				}
			}
		}
	}
	return nil, errors.New("dest server or channel not found")
}

func (r *RoomTable) UpdateRoom(ctx context.Context, data *server.Room) error {
	if data == nil || data.RoomId == 0 || data.RoomNumber == 0 {
		return errors.New("wrong room info")
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		if r.serverlist[i].ServerIndex == data.ParentChannelServer {
			for j := range r.serverlist[i].Channels {
				if r.serverlist[i].Channels[j].ChannelIndex == data.ParentChannel {
					for k := range r.serverlist[i].Channels[j].Rooms {
						if r.serverlist[i].Channels[j].Rooms[k].RoomId == data.RoomId {
							r.serverlist[i].Channels[j].Rooms[k] = *data
							return nil
						}
					}
					return errors.New("dest room not found in channel")
				}
			}
			return errors.New("dest room not found in server")
		}
	}

	return errors.New("dest room not found")
}

func (r *RoomTable) DeleteRoom(ctx context.Context, roomid uint16) error {
	if roomid == 0 {
		return errors.New("wrong room id")
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomid {
					r.serverlist[i].Channels[j].Rooms = append(
						r.serverlist[i].Channels[j].Rooms[:k],
						r.serverlist[i].Channels[j].Rooms[k+1:]...,
					)
					r.idMap.Delete(roomid)
					return nil
				}
			}
		}
	}

	return errors.New("dest room not found")
}

func (r *RoomTable) GetRoomList(ctx context.Context, serverID, ChannelID uint8) ([]server.Room, error) {
	rst := []server.Room{}
	r.lock.RLock()
	defer r.lock.RUnlock()

	for i := range r.serverlist {
		if r.serverlist[i].ServerIndex == serverID {
			for j := range r.serverlist[i].Channels {
				if r.serverlist[i].Channels[j].ChannelIndex == ChannelID {
					for k := range r.serverlist[i].Channels[j].Rooms {
						rst = append(
							rst,
							r.serverlist[i].Channels[j].Rooms[k],
						)
					}
					return rst, nil
				}
			}
			return rst, errors.New("dest channel not found in server")
		}
	}
	return rst, errors.New("dest channel not found")
}

func (r *RoomTable) GetServerList(ctx context.Context) ([]server.Server, error) {
	rst := []server.Server{}
	// 服务器固定，不需要lock

	for i := range r.serverlist {
		rst = append(rst, r.serverlist[i])
	}

	return rst, nil
}

func (r *RoomTable) AddServer(ctx context.Context, srv server.Server) {
	r.lock.Lock()
	defer r.lock.Unlock()

	srv.ServerIndex = uint8(len(r.serverlist) + 1)
	r.serverlist = append(r.serverlist, srv)
}

func (r *RoomTable) AddChannel(ctx context.Context, serverID uint8, chl server.Channel) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		if r.serverlist[i].ServerIndex == serverID {
			chl.ChannelIndex = uint8(len(r.serverlist[i].Channels) + 1)
			r.serverlist[i].Channels = append(r.serverlist[i].Channels, chl)
		}
	}
}

func (r *RoomTable) GetRoomInfo(ctx context.Context, roomID uint16) (*server.Room, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomID {
					rst := server.Room{}
					rst = r.serverlist[i].Channels[j].Rooms[k]
					return &rst, nil
				}
			}
		}
	}
	return nil, errors.New("dest room not found")
}

func (r *RoomTable) UpdateRoomSafe(ctx context.Context, data *server.Room) error {
	if data == nil || data.RoomId == 0 || data.RoomNumber == 0 {
		return errors.New("wrong room info")
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		if r.serverlist[i].ServerIndex == data.ParentChannelServer {
			for j := range r.serverlist[i].Channels {
				if r.serverlist[i].Channels[j].ChannelIndex == data.ParentChannel {
					for k := range r.serverlist[i].Channels[j].Rooms {
						if r.serverlist[i].Channels[j].Rooms[k].RoomId == data.RoomId {
							data.Users = r.serverlist[i].Channels[j].Rooms[k].Users // 保存用户
							r.serverlist[i].Channels[j].Rooms[k] = *data
							return nil
						}
					}
					return errors.New("dest room not found in channel")
				}
			}
			return errors.New("dest room not found in server")
		}
	}

	return errors.New("dest room not found")
}

func (r *RoomTable) AddUser(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomID {
					if len(r.serverlist[i].Channels[j].Rooms[k].Users) >= int(r.serverlist[i].Channels[j].Rooms[k].MaxPlayers) {
						return nil, errors.New("the room has too many users")
					}

					for idx := range r.serverlist[i].Channels[j].Rooms[k].Users {
						if r.serverlist[i].Channels[j].Rooms[k].Users[idx] == userID {
							rst := server.Room{}
							rst = r.serverlist[i].Channels[j].Rooms[k]
							return &rst, nil
						}
					}

					r.serverlist[i].Channels[j].Rooms[k].Users = append(r.serverlist[i].Channels[j].Rooms[k].Users, userID)
					rst := server.Room{}
					rst = r.serverlist[i].Channels[j].Rooms[k]
					return &rst, nil
				}
			}
		}
	}
	return nil, errors.New("dest room not found")
}

func (r *RoomTable) LeaveUser(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomID {
					if len(r.serverlist[i].Channels[j].Rooms[k].Users) == 0 {
						return nil, errors.New("the room has no users")
					}
					for idx := range r.serverlist[i].Channels[j].Rooms[k].Users {
						if r.serverlist[i].Channels[j].Rooms[k].Users[idx] == userID {
							r.serverlist[i].Channels[j].Rooms[k].Users = append(
								r.serverlist[i].Channels[j].Rooms[k].Users[:idx],
								r.serverlist[i].Channels[j].Rooms[k].Users[idx+1:]...)
							rst := server.Room{}
							rst = r.serverlist[i].Channels[j].Rooms[k]
							return &rst, nil
						}
					}
					return nil, errors.New("user not found")
				}
			}
		}
	}
	return nil, errors.New("dest room not found")
}

func (r *RoomTable) SetUserHost(ctx context.Context, roomID uint16, userID uint32, name string) (*server.Room, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomID {
					r.serverlist[i].Channels[j].Rooms[k].HostUserID = userID
					r.serverlist[i].Channels[j].Rooms[k].HostUserName = name
					rst := server.Room{}
					rst = r.serverlist[i].Channels[j].Rooms[k]
					return &rst, nil
				}
			}
		}
	}
	return nil, errors.New("dest room not found")
}

func (r *RoomTable) StartGame(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomID &&
					r.serverlist[i].Channels[j].Rooms[k].HostUserID == userID {
					r.serverlist[i].Channels[j].Rooms[k].Status = constant.StatusIngame
					rst := server.Room{}
					rst = r.serverlist[i].Channels[j].Rooms[k]
					return &rst, nil
				}
			}
		}
	}
	return nil, errors.New("dest room not found")
}

func (r *RoomTable) EndGame(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i := range r.serverlist {
		for j := range r.serverlist[i].Channels {
			for k := range r.serverlist[i].Channels[j].Rooms {
				if r.serverlist[i].Channels[j].Rooms[k].RoomId == roomID &&
					r.serverlist[i].Channels[j].Rooms[k].HostUserID == userID {
					r.serverlist[i].Channels[j].Rooms[k].Status = constant.StatusWaiting
					rst := server.Room{}
					rst = r.serverlist[i].Channels[j].Rooms[k]
					return &rst, nil
				}
			}
		}
	}
	return nil, errors.New("dest room not found")
}
