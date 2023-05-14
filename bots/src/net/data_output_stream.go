package net

import "math"

type DataOutputStream struct {
	data *[]byte
}

func NewDataOutputStream() *DataOutputStream {
	data := make([]byte, 0)
	return &DataOutputStream{
		data: &data,
	}
}

func (dos *DataOutputStream) GetData() *[]byte {
	return dos.data
}

func (dos *DataOutputStream) WriteBool(value bool) {
	if value {
		*dos.data = append(*dos.data, 1)
	} else {
		*dos.data = append(*dos.data, 0)
	}
}

func (dos *DataOutputStream) WriteDouble(value float64) {
	dos.WriteLong(math.Float64bits(value))
}

func (dos *DataOutputStream) WriteFloat(value float32) {
	dos.WriteInt(math.Float32bits(value))
}

func (dos *DataOutputStream) WriteFully(buf []byte) {
	*dos.data = append(*dos.data, buf...)
}

func (dos *DataOutputStream) WriteByte(value byte) error {
	*dos.data = append(*dos.data, value)
	return nil
}

func (dos *DataOutputStream) WriteShort(value uint16) {
	*dos.data = append(*dos.data, byte(value>>8), byte(value))
}

func (dos *DataOutputStream) WriteInt(value uint32) {
	*dos.data = append(*dos.data, byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func (dos *DataOutputStream) WriteLong(value uint64) {
	*dos.data = append(*dos.data, byte(value>>56), byte(value>>48), byte(value>>40), byte(value>>32), byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func (dos *DataOutputStream) WriteUTF(value string) {
	dos.WriteShort(uint16(len(value)))
	dos.WriteFully([]byte(value))
}
