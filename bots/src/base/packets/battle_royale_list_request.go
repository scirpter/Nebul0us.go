package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

type BattleRoyaleListRequest struct {
	enum enums.PACKET_TYPE
	net  *specs.Net
}

func NewBattleRoyaleListRequestPacket(net *specs.Net) *BattleRoyaleListRequest {
	return &BattleRoyaleListRequest{enum: enums.PACKET_TYPE_BATTLE_ROYALE_LIST_REQUEST, net: net}
}

func (p *BattleRoyaleListRequest) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteInt(p.net.RngToken1)

	return dos.GetData()
}
