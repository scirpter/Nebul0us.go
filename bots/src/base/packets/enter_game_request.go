package packets

import (
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/sigs"
	"neb/src/net"
)

// ## ?
// Used to send the client
// into a game.
type EnterGameRequest struct {
	enum      enums.PACKET_TYPE
	net       *specs.Net
	accountID int32
	worldName *string
}

func NewEnterGameRequestPacket(net *specs.Net, accountID int32, worldName *string) *EnterGameRequest {
	return &EnterGameRequest{enum: enums.PACKET_TYPE_ENTER_GAME_REQUEST, net: net, accountID: accountID, worldName: worldName}
}

func (p *EnterGameRequest) Write() *[]byte {
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteInt(p.net.RngToken1)
	dos.WriteInt(0xffffffff)
	dos.WriteUTF(*p.worldName)
	dos.WriteInt(uint32(p.accountID))
	dos.WriteByte(0xff)
	dos.WriteShort(sigs.APP_VERSION_SIG)

	return dos.GetData()
}
