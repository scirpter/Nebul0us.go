package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/models"
	"neb/src/base/natives"
	"neb/src/common"
	"neb/src/net"
)

// ## ?
// Issued when
// the client has joined a game.
// It holds raw information
// like dot locations, ....
// This packet may be received
// multiple times per join depending
// on the amount of objects on the map.
type GameData struct {
	enum            enums.PACKET_TYPE
	net             *specs.Net
	nets            []*specs.Net
	eventDispatcher *event_handle.EventDispatcher
	data            *[]byte
}

func NewGameDataPacket(net *specs.Net, nets []*specs.Net, eventDispatcher *event_handle.EventDispatcher, data *[]byte) *GameData {
	return &GameData{enum: enums.PACKET_TYPE_GAME_DATA, net: net, nets: nets, eventDispatcher: eventDispatcher, data: data}
}

func (p *GameData) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	p.net.World.Token = dis.ReadInt()
	p.net.World.RawSize = dis.ReadFloat()

	// only information about *this* current stream of GAME_DATA
	playerCt := dis.ReadByte2()
	ejectionCt := dis.ReadByte2()
	dotCtOffset := common.NET_ID(dis.ReadShort())
	dotCt := dis.ReadShort()
	itemCtOffset := common.NET_ID(dis.ReadByte2())
	itemCt := dis.ReadByte2()

	for i := 0; i < int(playerCt); i++ {
		netID := common.NET_ID(dis.ReadByte2())
		dis.ReadShort()
		dis.ReadByte2() // eject range check
		dis.ReadInt()
		dis.ReadInt()
		dis.ReadByte2()
		dis.ReadShort()
		dis.ReadUTF()
		dis.ReadByte2() // hat range check
		dis.ReadByte2() // halo range check
		dis.ReadByte2()
		dis.ReadShort()
		dis.ReadUTF()
		dis.ReadInt()
		dis.ReadInt()
		dis.ReadByte2() // particle range check
		bArr1 := dis.ReadByte2()
		dis.ReadFully(int(bArr1))
		dis.ReadByte2() // cycle range check
		dis.ReadShort()
		natives.RNATIVE_INTERPOLATE(dis, 0.0, 60.0)
		dis.ReadInt()
		dis.ReadInt()
		dis.ReadByte2()
		playerName := dis.ReadUTF()
		dis.ReadByte2() // font range check
		bArr2 := dis.ReadByte2()
		var playerAccountID int32 = -1
		var playerLevel int16 = -1
		var playerClanName string = ""
		if bArr2 <= 16 {
			dis.ReadFully(int(bArr2))
			playerAccountID = int32(dis.ReadInt())
			playerLevel = int16(dis.ReadShort())
			playerClanName = dis.ReadUTF()
			bArr3 := dis.ReadByte2()
			dis.ReadFully(int(bArr3))
			dis.ReadByte2() // clan rank range check
			dis.ReadByte2() // click type range check
		} else {
			panic("INVALID ALIAS COLORS LENGTH!")
		}

		var player *models.Player
		for _, net := range p.nets {
			if playerName == "NULL" {
				continue
			}
			net.World.PlayerMutx.Lock()
			player = p.net.World.Players[netID]
			if player == nil {
				player = models.NewPlayer(netID, &playerName, playerAccountID)
			}
			player.Level = playerLevel
			player.ClanName = &playerClanName
			net.World.Players[netID] = player
			net.World.PlayerMutx.Unlock()
		}

		if p.net.Uniquifier == uint8(len(p.nets)-1) && player != nil {
			eventData := event_handle.EventData{
				"player":          player,
				"shortPacketType": "GD",
			}
			p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnPlayerJoin, Data: eventData})
		}
	}

	for i := 0; i < int(ejectionCt); i++ {
		netID := common.NET_ID(dis.ReadByte2())
		x := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, p.net.World.RawSize)
		y := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, p.net.World.RawSize)
		mass := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, 500000.0)
		ejection := models.NewEjection(netID, x, y, mass)
		p.net.World.EjectionMutx.Lock()
		p.net.World.Ejections[netID] = ejection
		p.net.World.EjectionMutx.Unlock()
	}

	for i := 0; i < int(dotCt); i++ {
		netID := common.NET_ID(i) + dotCtOffset
		x := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, p.net.World.RawSize)
		y := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, p.net.World.RawSize)
		dot := models.NewDot(netID, x, y)
		p.net.World.DotMutx.Lock()
		p.net.World.Dots[netID] = dot
		p.net.World.DotMutx.Unlock()
	}

	for i := 0; i < int(itemCt); i++ {
		netID := common.NET_ID(i) + itemCtOffset
		typeof := dis.ReadByte2()
		actualTypeof := enums.ITEM_TYPE(typeof)
		x := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, p.net.World.RawSize)
		y := natives.RNATIVE_OBJ_DATA_RELATIVE_2(dis, p.net.World.RawSize)
		item := models.NewItem(netID, actualTypeof, x, y)
		p.net.World.ItemMutx.Lock()
		p.net.World.Items[netID] = item
		p.net.World.ItemMutx.Unlock()
	}
}
