package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/enums"
	"neb/src/net"
)

type BattleRoyaleListResult struct {
	enum            enums.PACKET_TYPE
	data            *[]byte
	eventDispatcher *event_handle.EventDispatcher
}

func NewBattleRoyaleListResultPacket(data *[]byte, eventDispatcher *event_handle.EventDispatcher) *BattleRoyaleListResult {
	return &BattleRoyaleListResult{enum: enums.PACKET_TYPE_BATTLE_ROYALE_LIST_RESULT, data: data, eventDispatcher: eventDispatcher}
}

func (p *BattleRoyaleListResult) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	dis.ReadByte2()
	BRRegistrantsCt := dis.ReadByte2()
	eventData := event_handle.EventData{
		"BRRegistrantsCt": uint8(BRRegistrantsCt),
	}
	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnBattleRoyaleListResult, Data: eventData})
}
