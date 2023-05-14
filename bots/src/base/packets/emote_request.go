package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// Used to let the client
// send an emote request,
// while playing of course.
type EmoteRequest struct {
	enum    enums.PACKET_TYPE
	net     *specs.Net
	emoteID uint8
}

func NewEmoteRequestPacket(net *specs.Net, emoteID uint8) *EmoteRequest {
	return &EmoteRequest{enum: enums.PACKET_TYPE_EMOTE_REQUEST, net: net, emoteID: emoteID}
}

func (p *EmoteRequest) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteByte(p.emoteID)
	dos.WriteInt(p.net.RngToken1)
	dos.WriteInt(0x00000000)

	return dos.GetData()
}
