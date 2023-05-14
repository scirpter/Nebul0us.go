package packets

import (
	"math/rand"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/sigs"
	"neb/src/common"
	"neb/src/net"
)

// ## ?
// Used to establish a session
// between the client and the server.
// Ref: Game's home screen.
type ConnectRequest3 struct {
	enum                     enums.PACKET_TYPE
	net                      *specs.Net
	blobData                 *specs.BlobData
	gameMode                 enums.GAME_MODE
	doConnectInPrivateSearch bool
	isMayhemTicked           bool
}

func NewConnectRequest3Packet(net *specs.Net, blobData *specs.BlobData, gameMode enums.GAME_MODE, doConnectInPrivateSearch bool, isMayhemTicked bool) *ConnectRequest3 {
	return &ConnectRequest3{enum: enums.PACKET_TYPE_CONNECT_REQUEST_3, net: net, blobData: blobData, gameMode: gameMode, doConnectInPrivateSearch: doConnectInPrivateSearch, isMayhemTicked: isMayhemTicked}
}

func (p *ConnectRequest3) Write() *[]byte {
	randomLong := rand.Uint64()
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(0x00000000)
	dos.WriteLong(randomLong)
	dos.WriteShort(sigs.APP_VERSION_SIG)
	dos.WriteInt(p.net.RngToken1)
	dos.WriteByte(byte(p.gameMode))
	dos.WriteBool(p.doConnectInPrivateSearch) // should actually be WriteByte
	dos.WriteInt(0xffffffff)
	dos.WriteUTF(*p.blobData.Token)
	dos.WriteByte(byte(p.blobData.ProfileVisibility))
	dos.WriteBool(p.isMayhemTicked)
	dos.WriteShort(uint16(p.blobData.Skin))
	dos.WriteByte(0xFF)
	dos.WriteUTF(*p.blobData.Name)
	dos.WriteInt(0x00000000)
	dos.WriteByte(byte(len(*p.blobData.Name)))

	for i := 0; i < len(*p.blobData.Name); i++ {
		dos.WriteByte(0xff)
	}

	dos.WriteByte(0xff)
	dos.WriteInt(p.net.RngToken2)
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

	for i := 0; i < 5; i++ {
		dos.WriteByte(0x77)
	}

	dos.WriteByte(byte(p.blobData.ColorCycle))
	dos.WriteShort(0x0000) // source code said: uint16(p.client.blobData.skin)), maybe secondary skin?
	dos.WriteShort(0x0000)
	dos.WriteInt(0x00000000)
	dos.WriteLong(common.GetCommonRequestStamp())

	secured := RNGEncrypt(dos.GetData(), int64(randomLong))

	return secured
}
