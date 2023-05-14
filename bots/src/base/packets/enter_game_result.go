package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/natives"
	"neb/src/net"
)

// ## ?
// Received when we request to join
// a game.
type EnterGameResult struct {
	enum            enums.PACKET_TYPE
	data            *[]byte
	net             *specs.Net
	eventDispatcher *event_handle.EventDispatcher
}

func NewEnterGameResultPacket(data *[]byte, net *specs.Net, eventDispatcher *event_handle.EventDispatcher) *EnterGameResult {
	return &EnterGameResult{enum: enums.PACKET_TYPE_ENTER_GAME_RESULT, net: net, data: data, eventDispatcher: eventDispatcher}
}

func (p *EnterGameResult) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	result := enums.JOIN_RESULT(dis.ReadByte2())

	if result >= 34 {
		result = enums.JOIN_RESULT_UNKNOWN_ERROR
	}

	var roomName string
	var gameMode enums.GAME_MODE

	if result == enums.JOIN_RESULT_SUCCESS {
		readByte2 := dis.ReadByte2()
		e := readByte2
		if e >= 2 {
			e = 1
		}

		g := dis.ReadByte2()
		if g >= 4 {
			g = 2
		}

		gameMode = enums.GAME_MODE(dis.ReadByte2())
		p.net.World.GameMode = gameMode
		dis.ReadInt()
		dis.ReadInt()
		dis.ReadInt()
		roomName = dis.ReadUTF()
		p.net.World.Name = &roomName
		dis.ReadByte2()
		dis.ReadBool()
		dis.ReadLong()
		dis.ReadByte2()
		dis.ReadInt()
		dis.ReadBool()
		dis.ReadBool()
		dis.ReadBool()
		dis.ReadInt()
		dis.ReadBool()
		dis.ReadByte2()
		dis.ReadBool()

		roomSize := enums.WORLD_SIZE(dis.ReadByte2())
		if roomSize >= 4 {
			roomSize = enums.WORLD_SIZE_TINY
		}
		bArr := make([]byte, dis.ReadByte2())
		dis.ReadFully(len(bArr))
		dis.ReadBool()
	}
	dis.ReadLong()

	if result == enums.JOIN_RESULT_SUCCESS {
		dis.ReadByte2()
		dis.ReadByte2()
		natives.RNATIVE_LE_SHORT_INT(dis)
		dis.ReadByte2()
		dis.ReadInt()
		dis.ReadBool()
	}

	eventData := event_handle.EventData{
		"result":   result,
		"roomName": &roomName,
		"gameMode": gameMode,
	}

	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnEnterGameResult, Data: eventData})
}
