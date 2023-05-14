package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/models"
	"neb/src/common"
	"neb/src/net"
)

// ## ?
// When a player rejoins,
// this is the packet received.
// Excludes self.
type JoinResult struct {
	enum            enums.PACKET_TYPE
	data            *[]byte
	net             *specs.Net
	statery         *specs.Statery
	eventDispatcher *event_handle.EventDispatcher
}

func NewJoinResultPacket(data *[]byte, eventDispatcher *event_handle.EventDispatcher, net *specs.Net, statery *specs.Statery) *JoinResult {
	return &JoinResult{enum: enums.PACKET_TYPE_JOIN_RESULT, data: data, net: net, statery: statery, eventDispatcher: eventDispatcher}
}

func (p *JoinResult) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	cr2Token2 := dis.ReadInt()

	readByte := dis.ReadByte2()
	if readByte != 0 {
		// aka this is "invalid"
		return
	}

	netID := common.NET_ID(dis.ReadByte2())
	dis.ReadShort()
	dis.ReadInt()
	dis.ReadByte2()
	playerName := dis.ReadUTF()

	// NOTE: skipped the rest as i don't need it rn

	p.net.World.PlayerMutx.Lock()
	player := p.net.World.Players[netID]
	if player == nil {
		player = models.NewPlayer(netID, &playerName, -1)
	}
	if p.statery.EmotionalSupportPlayerName != nil && *p.statery.EmotionalSupportPlayerName == *player.Name {
		p.statery.EmotionalSupportPlayerName = &playerName
	}
	if p.statery.PlasmaFarmingTargetPlayerName != nil && *p.statery.PlasmaFarmingTargetPlayerName == *player.Name {
		p.statery.PlasmaFarmingTargetPlayerName = &playerName
	}
	player.Name = &playerName
	player.Cr2Token2 = cr2Token2
	p.net.World.PlayerMutx.Unlock()

	eventData := event_handle.EventData{
		"player": player,
	}
	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnJoinResult, Data: eventData})
}
