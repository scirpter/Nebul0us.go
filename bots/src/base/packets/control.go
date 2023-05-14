package packets

import (
	"math"
	"math/rand"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/base/natives"
	"neb/src/net"
)

// ## ?
// Used to send client movement
// instructions to the server.
type Control struct {
	enum        enums.PACKET_TYPE
	net         *specs.Net
	controlData *specs.ControlData

	Split bool
	Eject bool
	Drop  bool
}

func NewControlPacket(net *specs.Net, controlData *specs.ControlData) *Control {
	return &Control{
		enum:        enums.PACKET_TYPE_CONTROL,
		net:         net,
		controlData: controlData,
		Split:       false,
		Eject:       false,
		Drop:        false,
	}
}

func (p *Control) Write() *[]byte {
	var pi2 float32 = 2 * math.Pi

	angle := natives.NATIVE_GET_CONTROL_ANGLE(p.controlData.Angle, pi2)
	speed := p.controlData.Speed * 0xff
	dos := net.NewDataOutputStream()

	dos.WriteByte(byte(p.enum))
	dos.WriteInt(p.net.Cr2Token1)
	dos.WriteShort(angle)
	dos.WriteByte(speed)
	dos.WriteByte(p.controlData.ControlTick)

	p.controlData.ControlTick = uint8(uint16((p.controlData.ControlTick + 1)) % 256)
	var button byte = 0

	if p.Split {
		button |= 1
		p.controlData.SplitCt--
	}
	if p.Eject {
		button |= 2
		p.controlData.EjectCt--
	}
	if p.Drop {
		button |= 8
		p.controlData.DropCt--
	}

	dos.WriteByte(button)

	var idkID uint8 = 0xff

	arr := []byte{0x02, 0x0c, 0x00, 0x07, 0xff, 0x09, 0x03, 0x01, 0x05, 0x04}
	idkID = uint8(arr[rand.Intn(len(arr))])

	// NOTE: differs in multi-blob
	dos.WriteByte(idkID) // first time in world: ff
	dos.WriteInt(p.net.RngToken1)
	dos.WriteByte(0x63)

	return dos.GetData()
}
