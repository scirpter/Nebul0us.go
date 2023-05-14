package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// Used to read or write
// clan chat messages.
type ClanChatMessage struct {
	enum enums.PACKET_TYPE
	data *[]byte

	eventDispatcher *event_handle.EventDispatcher
}

func NewClanChatMessagePacket(data *[]byte, eventDispatcher *event_handle.EventDispatcher) *ClanChatMessage {
	return &ClanChatMessage{enum: enums.PACKET_TYPE_CLAN_CHAT_MESSAGE, data: data, eventDispatcher: eventDispatcher}
}

func (p *ClanChatMessage) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	dis.ReadInt()
	playerName := dis.ReadUTF()
	message := dis.ReadUTF()
	readByte := dis.ReadByte2() // ? always writes 0? (ref c6500n1.writeByte(0);)
	enumClanRankLength := int8(enums.CLAN_RANK_INITIATE + 1)
	if readByte <= 0 || readByte >= byte(enumClanRankLength) {
		readByte = 0 // => invalid clan rank
	}
	playerClanRank := enums.CLAN_RANK(readByte)
	playerAccountID := dis.ReadInt()

	if dis.SpaceLeft() > 0 {
		dis.ReadInt()
	}
	if dis.SpaceLeft() > 0 {
		dis.ReadBool()
	}

	eventData := event_handle.EventData{
		"playerName":      &playerName,
		"message":         &message,
		"playerClanRank":  playerClanRank,
		"playerAccountID": int32(playerAccountID),
	}

	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnClanChatMessage, Data: eventData})
}
