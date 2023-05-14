package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/sigs"
	"neb/src/net"
)

// ## ?
// This packet must be sent periodically,
// otherwise the server will disconnect
// the client after few seconds.
type KeepAlive struct {
	enum enums.PACKET_TYPE
	net  *specs.Net
}

func NewKeepAlivePacket(net *specs.Net) *KeepAlive {
	return &KeepAlive{enum: enums.PACKET_TYPE_KEEP_ALIVE, net: net}
}

func (p *KeepAlive) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteInt(p.net.Cr2Token2)
	dos.WriteInt(sigs.KEEP_ALIVE_SIG)
	dos.WriteInt(p.net.RngToken1)

	return dos.GetData()
}
