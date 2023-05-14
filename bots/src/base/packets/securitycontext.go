package packets

import (
	"neb/src/net"
	"neb/src/utils/javarng"
)

type PacketSecurityContext struct {
	xlength  int
	xrandint int32
}

func NewPacketSecurityContext(xlength int, xrandint int32) *PacketSecurityContext {
	return &PacketSecurityContext{
		xlength:  xlength,
		xrandint: xrandint,
	}
}

func RNGEncrypt(buf *[]byte, nextLong int64) *[]byte {
	length := len(*buf) - 13
	arrayList := make([]*PacketSecurityContext, 0)
	random := javarng.NewJavaRandom(nextLong)
	for {
		length--
		if length <= 0 {
			break
		}
		arrayList = append(arrayList, NewPacketSecurityContext(length, random.NextInt32(int32(length+1))))
	}
	for _, k2Var := range arrayList {
		i11 := k2Var.xlength + 13
		b10 := (*buf)[i11]
		i12 := k2Var.xrandint + 13
		(*buf)[i11] = (*buf)[i12]
		(*buf)[i12] = b10
	}
	return buf
}

func RNGDecrypt(buf *[]byte) *[]byte {
	dis := net.NewDataInputStream(buf)
	dis.SkipBytes(5)

	length := len(*buf) - 13
	arrayList := make([]*PacketSecurityContext, 0)
	random := javarng.NewJavaRandom(int64(dis.ReadLong()))
	for {
		length--
		if length <= 0 {
			break
		}
		arrayList = append(arrayList, NewPacketSecurityContext(length, random.NextInt32(int32(length+1))))
	}
	for i := len(arrayList) - 1; i >= 0; i-- {
		k2Var := arrayList[i]
		i11 := k2Var.xlength + 13
		b10 := (*buf)[i11]
		i12 := k2Var.xrandint + 13
		(*buf)[i11] = (*buf)[i12]
		(*buf)[i12] = b10
	}
	return buf
}
