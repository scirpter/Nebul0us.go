package natives

import (
	"neb/src/net"
)

func RNATIVE_OBJ_DATA_RELATIVE(dis *net.DataInputStream, f1 float32, f2 float32) float32 {
	v0 := float32(dis.ReadByte2() & 0xff)
	v1 := float32(dis.ReadByte2() & 0xff)
	v2 := float32(dis.ReadByte2() & 0xff)

	return ((v0*65536.0)+(v1*256.0)+v2)*(f2-f1)/16777215.0 + f1
}

func RNATIVE_OBJ_DATA_RELATIVE_2(dis *net.DataInputStream, f float32) float32 {
	v0 := uint32(dis.ReadByte2() & 0xff)
	v1 := uint32(dis.ReadByte2() & 0xff)
	v2 := uint32(dis.ReadByte2() & 0xff)

	return (((f - 0.0) * float32(((v0<<16)+(v1<<8))+v2)) / 1.6777215e7) + 0.0
}

func RNATIVE_LE_SHORT_INT(dis *net.DataInputStream) uint32 {
	v0 := uint32(dis.ReadByte2() & 0xff)
	v1 := uint32(dis.ReadByte2() & 0xff)
	v2 := uint32(dis.ReadByte2() & 0xff)

	return (v0 << 16) + (v1 << 8) + v2
}

func RNATIVE_OBJ_ANGLE(dis *net.DataInputStream, f1 float32, f2 float32) float32 {
	v0 := float32(dis.ReadByte2() & 0xff)

	return v0*(f2-f1)/255.0 + f1
}

func RNATIVE_INTERPOLATE(dis *net.DataInputStream, f1, f2 float32) float32 {
	shortVal := dis.ReadShort() & 65535
	return float32(shortVal)*(f2-f1)/65535.0 + f1
}

func NATIVE_GET_CONTROL_ANGLE(f1 float32, f2 float32) uint16 {
	return uint16(((f1 - 0.0) * 65535.0) / (f2 - 0.0))
}

func RNATIVE_APPLY_SCALING_FACTOR(dis *net.DataInputStream, f1 float32) float32 {
	v0 := dis.ReadShort() & 65535
	return ((f1-0.0)*float32(v0))/65535.0 + 0.0
}
