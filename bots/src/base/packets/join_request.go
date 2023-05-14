package packets

import (
	"encoding/hex"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// When we have joined a game and
// want the bots to start playing.
type JoinRequest struct {
	enum     enums.PACKET_TYPE
	net      *specs.Net
	blobData *specs.BlobData
}

func NewJoinRequestPacket(net *specs.Net, blobData *specs.BlobData) *JoinRequest {
	return &JoinRequest{enum: enums.PACKET_TYPE_JOIN_REQUEST, net: net, blobData: blobData}
}

func (p *JoinRequest) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteShort(uint16(p.blobData.Skin))
	dos.WriteUTF(*p.blobData.Name)
	dos.WriteShort(0xff00)
	dos.WriteInt(uint32(len(*p.blobData.Name)))
	dos.WriteShort(0xffff)

	for i := 0; i < len(*p.blobData.Name); i++ {
		dos.WriteByte(0xff)
	}

	raw0, _ := hex.DecodeString("e1d452")
	dos.WriteFully(raw0)

	dos.WriteUTF("")
	dos.WriteByte(byte(p.blobData.Hat))
	dos.WriteInt(0x00000000)
	dos.WriteByte(byte(p.blobData.Halo))
	dos.WriteByte(0xff)
	dos.WriteUTF("")
	dos.WriteInt(0x00000000)
	dos.WriteInt(0x00000000)
	dos.WriteByte(byte(p.blobData.Particle))
	dos.WriteByte(byte(p.blobData.NameFont))
	dos.WriteByte(0x05)
	dos.WriteByte(byte(p.blobData.ColorCycle))
	dos.WriteShort(0x0000)
	dos.WriteInt(0x00000000)
	dos.WriteShort(0x0000)
	dos.WriteInt(0x00000000)
	dos.WriteInt(p.net.RngToken1)

	raw1, _ := hex.DecodeString("7777777777")
	dos.WriteFully(raw1)

	return dos.GetData()
}
