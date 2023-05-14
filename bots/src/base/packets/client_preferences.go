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
// Issued when
// a new player has joined the lobby, excluding self.
type ClientPreferences struct {
	enum            enums.PACKET_TYPE
	net             *specs.Net
	data            *[]byte
	eventDispatcher *event_handle.EventDispatcher
}

func NewClientPreferencesPacket(net *specs.Net, data *[]byte, eventDispatcher *event_handle.EventDispatcher) *ClientPreferences {
	return &ClientPreferences{enum: enums.PACKET_TYPE_CLIENT_PREFERENCES, net: net, data: data, eventDispatcher: eventDispatcher}
}

func (p *ClientPreferences) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	netID := common.NET_ID(dis.ReadInt())
	dis.ReadShort()

	playerName := dis.ReadUTF()
	playerID := dis.ReadInt()

	p.net.World.PlayerMutx.Lock()
	player := p.net.World.Players[netID]
	if player == nil {
		player = models.NewPlayer(netID, &playerName, int32(playerID))
	}
	p.net.World.Players[netID] = player
	p.net.World.PlayerMutx.Unlock()

	// NOTE: skipped the rest, don't need this now
	// 		 also, the fuck is up with all those nests?
	//          ==>> https://imgur.com/a/ZI2O94B <<==

	eventData := event_handle.EventData{
		"player":          player,
		"shortPacketType": "CP",
	}
	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnPlayerJoin, Data: eventData})
}
