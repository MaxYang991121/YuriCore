package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildRoomList(rooms []server.Room) []byte {
	buf := make([]byte, 3)
	offset := 0
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint16(&buf, uint16(len(rooms)), &offset)

	for i := range rooms {
		buf = utils.BytesCombine(buf, BuildRoomChannelInfo(&rooms[i], true, 0xFFFFFFFF))
	}

	return buf
}

func BuildUpdateChannelRoom(room *server.Room, needID bool, flag uint32) []byte {
	buf := make([]byte, 1)
	offset := 0
	utils.WriteUint8(&buf, 3, &offset)
	buf = utils.BytesCombine(buf, BuildRoomChannelInfo(room, needID, flag))
	return buf
}

func BuildDeleteChannelRoom(roomID uint16) []byte {
	buf := make([]byte, 4)
	offset := 0
	utils.WriteUint8(&buf, 2, &offset)
	utils.WriteUint16(&buf, roomID, &offset)
	return buf[:offset]
}

func BuildAddChannelRoom(room *server.Room, needID bool, flag uint32) []byte {
	buf := make([]byte, 1)
	offset := 0
	utils.WriteUint8(&buf, 1, &offset)
	buf = utils.BytesCombine(buf, BuildRoomChannelInfo(room, needID, flag))
	return buf
}

func BuildRoomChannelInfo(room *server.Room, needID bool, flag uint32) []byte {
	buf := make([]byte, 512)
	offset := 0

	if needID {
		utils.WriteUint16(&buf, room.RoomId, &offset)
	}

	utils.WriteUint32(&buf, flag, &offset)

	if flag&0x1 != 0 {
		utils.WriteStringWithNull(&buf, []byte(room.RoomName), &offset)
	}

	if flag&0x2 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x4 != 0 {
		if room.PassWd != "" {
			utils.WriteUint8(&buf, 1, &offset)
		} else {
			utils.WriteUint8(&buf, 0, &offset)
		}
	}

	if flag&0x8 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x10 != 0 {
		utils.WriteUint8(&buf, room.GameModeID, &offset)
	}

	if flag&0x20 != 0 {
		utils.WriteUint8(&buf, room.MapID, &offset)
	}

	if flag&0x40 != 0 {
		utils.WriteUint8(&buf, uint8(len(room.Users)), &offset)
	}

	if flag&0x80 != 0 {
		utils.WriteUint8(&buf, room.MaxPlayers, &offset)
	}

	if flag&0x100 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x200 != 0 {
		utils.WriteUint32(&buf, room.HostUserID, &offset)
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x400 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x800 != 0 {
		utils.WriteUint32(&buf, 3, &offset)
		utils.WriteUint16(&buf, 0, &offset)
		utils.WriteUint32(&buf, 3, &offset)
		utils.WriteUint16(&buf, 0, &offset)
		utils.WriteUint8(&buf, 2, &offset)
	}

	if flag&0x1000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x2000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x4000 != 0 {
		utils.WriteUint8(&buf, uint8(len(room.Users)), &offset)
	}

	if flag&0x8000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x10000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x20000 != 0 {
		utils.WriteUint16(&buf, 0, &offset)
	}

	if flag&0x40000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x80000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x100000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x200000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x400000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x800000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	if flag&0x1000000 != 0 {
		utils.WriteUint8(&buf, 0, &offset)
	}

	return buf[:offset]
}
