package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// If a client session is running,
// send this to disconnect.
// Requires the client to reconnect.
type Disconnect struct {
	enum enums.PACKET_TYPE
	net  *specs.Net
}

func NewDisconnectPacket(net *specs.Net) *Disconnect {
	return &Disconnect{enum: enums.PACKET_TYPE_DISCONNECT, net: net}
}

func (p *Disconnect) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteInt(p.net.Cr2Token2)
	dos.WriteInt(p.net.RngToken1)

	return dos.GetData()
}
