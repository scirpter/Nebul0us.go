package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// When the bot dies or a round ends,
// this packet is received.
// This is more used to determine whether
// the bot has died.
type SessionStats struct {
	enum            enums.PACKET_TYPE
	data            *[]byte
	eventDispatcher *event_handle.EventDispatcher
}

func NewSessionStatsPacket(data *[]byte, eventDispatcher *event_handle.EventDispatcher) *SessionStats {
	return &SessionStats{enum: enums.PACKET_TYPE_SESSION_STATS, data: data, eventDispatcher: eventDispatcher}
}

func (p *SessionStats) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	cr2Token2 := dis.ReadInt()

	eventData := event_handle.EventData{
		"cr2Token2": cr2Token2,
	}

	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnSessionStats, Data: eventData})
}
