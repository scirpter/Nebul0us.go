package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

type BattleRoyaleAction struct {
	enum       enums.PACKET_TYPE
	net        *specs.Net
	doRegister bool
}

func NewBattleRoyaleActionPacket(net *specs.Net, doRegister bool) *BattleRoyaleAction {
	return &BattleRoyaleAction{enum: enums.PACKET_TYPE_BATTLE_ROYALE_ACTION, net: net, doRegister: doRegister}
}

func (p *BattleRoyaleAction) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteInt(p.net.Cr2Token2)

	if p.doRegister {
		dos.WriteByte(0x01)
	} else {
		dos.WriteByte(0x02)
	}

	return dos.GetData()
}
