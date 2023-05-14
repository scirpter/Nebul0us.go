package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/natives"
	"neb/src/common"
	"neb/src/net"
)

// ## ?
// Issued every ~1s while
// the client is in playing state.
type TopScores struct {
	enum enums.PACKET_TYPE
	net  *specs.Net
	data *[]byte
}

func NewTopScoresPacket(net *specs.Net, data *[]byte) *TopScores {
	return &TopScores{enum: enums.PACKET_TYPE_TOP_SCORES, net: net, data: data}
}

func (p *TopScores) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	name := dis.ReadUTF()
	p.net.World.Name = &name
	p.net.World.Token = dis.ReadInt()
	p.net.World.TimeLeft = dis.ReadShort()
	p.net.World.GameMode = enums.GAME_MODE(dis.ReadByte2())
	p.net.World.WorldSize = enums.WORLD_SIZE((dis.ReadByte2() & 12) >> 2)
	p.net.World.MaxPlayers = dis.ReadByte2()
	p.net.World.PlayerCt = dis.ReadByte2()

	for i := 0; i < int(p.net.World.PlayerCt); i++ {
		dis.ReadShort()

		fluctNetID := common.NET_ID(dis.ReadByte2())
		var netID common.NET_ID = 0
		if p.net.World.GameMode == enums.GAME_MODE_BATTLE_ROYALE {
			netID = fluctNetID
		} else {
			netID = fluctNetID & 31
		}
		_ = netID
	}

	for i := 0; i < 5; i++ {
		dis.ReadByte2()
		natives.RNATIVE_LE_SHORT_INT(dis)
	}

	// NOTE: skipped the rest, don't need this now

	allData := dis.GetData()
	p.net.World.SpectatorCt = (*allData)[len(*allData)-13] / 2
}
