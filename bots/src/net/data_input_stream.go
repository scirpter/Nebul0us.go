package net

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"neb/src/utils/cnsl"
	"time"
)

type DataInputStream struct {
	data *[]byte
	pos  int
}

func NewDataInputStream(data *[]byte) *DataInputStream {
	return &DataInputStream{data: data, pos: 0}
}

// ensures that out of bounds does not crash, but instead
// permanently warns about it to fix
func (dis *DataInputStream) EnsurePos() {
	if dis.pos > 8100 {
		cnsl.Error("packet reader offset is out of bounds. please contact the bot developer.")
		time.Sleep(5 * time.Second)
	}
}

func (dis *DataInputStream) GetData() *[]byte {
	return dis.data
}

func (dis *DataInputStream) SpaceLeft() int {
	return len(*dis.data) - dis.pos
}

func (dis *DataInputStream) SkipBytes(n int) {
	dis.pos += n
	dis.EnsurePos()
}

func (dis *DataInputStream) ToHexString() string {
	data := *dis.data
	return hex.EncodeToString(data)
}

func (dis *DataInputStream) ReadBool() bool {
	data := *dis.data
	value := data[dis.pos] != 0
	dis.pos++
	dis.EnsurePos()
	return value
}

func (dis *DataInputStream) ReadByte2() byte {
	data := *dis.data
	value := data[dis.pos]
	dis.pos++
	dis.EnsurePos()
	return value
}

func (dis *DataInputStream) ReadShort() uint16 {
	data := *dis.data
	value := binary.BigEndian.Uint16(data[dis.pos : dis.pos+2])
	dis.pos += 2
	dis.EnsurePos()
	return value
}

func (dis *DataInputStream) ReadInt() uint32 {
	data := *dis.data
	value := binary.BigEndian.Uint32(data[dis.pos : dis.pos+4])
	dis.pos += 4
	dis.EnsurePos()
	return value
}

func (dis *DataInputStream) ReadLong() uint64 {
	data := *dis.data
	value := binary.BigEndian.Uint64(data[dis.pos : dis.pos+8])
	dis.pos += 8
	dis.EnsurePos()
	return value
}

func (dis *DataInputStream) ReadFloat() float32 {
	data := *dis.data
	bits := binary.BigEndian.Uint32(data[dis.pos : dis.pos+4])
	dis.pos += 4
	dis.EnsurePos()
	return math.Float32frombits(bits)
}

func (dis *DataInputStream) ReadDouble() float64 {
	data := *dis.data
	bits := binary.BigEndian.Uint64(data[dis.pos : dis.pos+8])
	dis.pos += 8
	dis.EnsurePos()
	return math.Float64frombits(bits)
}

func (dis *DataInputStream) ReadFully(len int) ([]byte, error) {
	data := *dis.data
	value := data[dis.pos : dis.pos+len]
	dis.pos += len
	dis.EnsurePos()
	return value, nil
}

func (dis *DataInputStream) ReadUTF() string {
	data := *dis.data
	length := dis.ReadShort()
	ret := string(data[dis.pos : dis.pos+int(length)])
	dis.pos += int(length)
	dis.EnsurePos()
	return ret
}
