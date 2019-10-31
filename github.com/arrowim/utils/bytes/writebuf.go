package bytes

import (
	b "bytes"
	"encoding/binary"
)

type WriteBuf struct {
	buf                     *b.Buffer
	isNeedIncludeDataLength bool
}

//func CreateWriteBuf()

func CreateWriteBuf(isNeedIncludeDataLength bool) *WriteBuf {
	w := WriteBuf{}
	w.buf = b.NewBuffer([]byte{})

	if isNeedIncludeDataLength {
		w.WriteInt32(0)

	}
	w.isNeedIncludeDataLength = isNeedIncludeDataLength

	return &w
}

func (self *WriteBuf) WriteInt8(data int8) {
	binary.Write(self.buf, binary.BigEndian, data)
}

func (self *WriteBuf) WriteInt16(data int16) {
	binary.Write(self.buf, binary.BigEndian, data)
}

func (self *WriteBuf) WriteInt32(data int32) {
	binary.Write(self.buf, binary.BigEndian, data)
}

func (self *WriteBuf) WriteInt64(data int64) {
	binary.Write(self.buf, binary.BigEndian, data)
}

func (self *WriteBuf) WriteString(data string) {
	self.WriteInt16(int16(len(data)))

	if len(data) > 0 {
		self.buf.Write([]byte(data))
	}

}

func (self *WriteBuf) WriteBytes(data []byte) {
	self.WriteInt16(int16(len(data)))

	if len(data) > 0 {
		self.buf.Write(data)
	}

}

func (self *WriteBuf) Bytes() []byte {
	d := self.buf.Bytes()

	if !self.isNeedIncludeDataLength {
		return self.buf.Bytes()
	}
	l := Int32ToBytes(int32(len(d)))
	d[0] = l[0]
	d[1] = l[1]
	d[2] = l[2]
	d[3] = l[3]
	return d
}

func (self *WriteBuf) WriteFloat32(data float32) {
	binary.Write(self.buf, binary.BigEndian, data)
}
