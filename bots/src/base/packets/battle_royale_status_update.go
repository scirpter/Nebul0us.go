package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

type BattleRoyaleStatusUpdate struct {
	enum            enums.PACKET_TYPE
	data            *[]byte
	statery         *specs.Statery
	eventDispatcher *event_handle.EventDispatcher
}

func NewBattleRoyaleStatusUpdatePacket(data *[]byte, statery *specs.Statery, eventDispatcher *event_handle.EventDispatcher) *BattleRoyaleStatusUpdate {
	return &BattleRoyaleStatusUpdate{enum: enums.PACKET_TYPE_BATTLE_ROYALE_STATUS_UPDATE, data: data, statery: statery, eventDispatcher: eventDispatcher}
}

func (p *BattleRoyaleStatusUpdate) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	dis.ReadByte2()
	BRRegistrantsCt := dis.ReadByte2()

	if BRRegistrantsCt == 16 {
		// this block is actually never properly reached..
		p.statery.BRRegistrantsCt = 0
		// p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnBattleRoyaleStart})
		return
	}

	p.statery.BRRegistrantsCt = BRRegistrantsCt
}
