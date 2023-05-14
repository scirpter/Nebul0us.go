package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// Read or write a message to
// the public chat.
type GameChatMessage struct {
	enum            enums.PACKET_TYPE
	net             *specs.Net
	blobData        *specs.BlobData
	data            *[]byte
	eventDispatcher *event_handle.EventDispatcher
	message         *string
}

func NewGameChatMessagePacket(net *specs.Net, blobData *specs.BlobData, data *[]byte, eventDispatcher *event_handle.EventDispatcher, message *string) *GameChatMessage {
	return &GameChatMessage{enum: enums.PACKET_TYPE_GAME_CHAT_MESSAGE, net: net, blobData: blobData, data: data, eventDispatcher: eventDispatcher, message: message}
}

func (p *GameChatMessage) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	dis.ReadInt()
	playerName := dis.ReadUTF()

	if playerName == *p.blobData.Name {
		return
	}

	message := dis.ReadUTF()
	playerAccountID := int32(dis.ReadInt())
	dis.ReadBool()
	dis.ReadLong()

	eventData := event_handle.EventData{
		"playerName":      &playerName,
		"message":         &message,
		"playerAccountID": playerAccountID,
	}
	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnGameChatMessage, Data: eventData})
}

func (p *GameChatMessage) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteUTF(*p.blobData.Name)
	dos.WriteUTF(*p.message)
	dos.WriteInt(0xffffffff)
	dos.WriteLong(0x0000000000000000)
	dos.WriteShort(uint16(len(*p.blobData.Name)))

	for i := 0; i < len(*p.blobData.Name); i++ {
		dos.WriteByte(0xff)
	}

	dos.WriteShort(0x0000)
	dos.WriteInt(p.net.RngToken1)

	return dos.GetData()
}
