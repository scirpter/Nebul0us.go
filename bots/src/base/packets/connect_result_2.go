package packets

import (
	event_handle "neb/src/base/client/events"
	"neb/src/base/client/specs"
	"neb/src/base/enums"
	"neb/src/net"
)

// ## ?
// Received when trying to establish
// a connection to the server.
// This holds important information like
// `cr2_token1` and `cr2_token2`.
// Those tokens are used to identify the
// client server-side.
type ConnectResult2 struct {
	enum            enums.PACKET_TYPE
	net             *specs.Net
	data            *[]byte
	eventDispatcher *event_handle.EventDispatcher
}

func NewConnectResult2Packet(net *specs.Net, data *[]byte, eventDispatcher *event_handle.EventDispatcher) *ConnectResult2 {
	return &ConnectResult2{enum: enums.PACKET_TYPE_CONNECT_RESULT_2, net: net, data: data, eventDispatcher: eventDispatcher}
}

func (p *ConnectResult2) Parse() {
	dis := net.NewDataInputStream(p.data)

	dis.ReadByte2() // packet type
	p.net.RngToken1 = dis.ReadInt()
	dis.ReadByte2() // 0
	p.net.Cr2Token1 = dis.ReadInt()
	p.net.Cr2Token2 = dis.ReadInt()
	p.net.World.Token = dis.ReadInt()
	dis.ReadFloat() // ?
	dis.ReadByte2() // ?

	p.eventDispatcher.Dispatch(event_handle.Event{Type: event_handle.OnConnected})
}
