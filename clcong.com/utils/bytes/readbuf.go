package bytes

import (
	bytes "bytes"
	"encoding/binary"
	"errors"
	//	"fmt"
	"strconv"
)

type ReadBuf struct {
	index int
	data  *[]byte
	err   error
}

func (self *ReadBuf) GetIndex() int {
	return self.index
}

func CreateReadBuf(data *[]byte, isNeedReadLengthFirst bool) *ReadBuf {
	r := ReadBuf{}
	r.data = data

	if isNeedReadLengthFirst {
		r.ReadInt16()
	}

	return &r
}

func (self *ReadBuf) ReadInt8() (error, int8) {
	if self.err != nil {
		return self.err, 0
	}
	self.err = checkDataLength(*self.data, self.index, 1)

	if self.err != nil {
		return self.err, 0
	}

	tmp := (*self.data)[self.index : self.index+1]
	self.index += 1

	b_buf := bytes.NewBuffer(tmp)

	var result int8
	binary.Read(b_buf, binary.BigEndian, &result)
	return nil, result
}

func (self *ReadBuf) ReadInt16() (error, int16) {
	if self.err != nil {
		return self.err, 0
	}
	self.err = checkDataLength(*self.data, self.index, 2)

	if self.err != nil {
		return self.err, 0
	}

	tmp := (*self.data)[self.index : self.index+2]
	self.index += 2
	b_buf := bytes.NewBuffer(tmp)

	var result int16
	binary.Read(b_buf, binary.BigEndian, &result)
	return nil, result
}

func (self *ReadBuf) ReadInt32() (error, int32) {
	if self.err != nil {
		return self.err, 0
	}
	self.err = checkDataLength(*self.data, self.index, 4)

	if self.err != nil {
		return self.err, 0
	}

	tmp := (*self.data)[self.index : self.index+4]
	self.index += 4
	b_buf := bytes.NewBuffer(tmp)

	var result int32
	binary.Read(b_buf, binary.BigEndian, &result)
	return nil, result
}

func (self *ReadBuf) ReadInt64() (error, int64) {
	if self.err != nil {
		return self.err, 0
	}
	self.err = checkDataLength(*self.data, self.index, 8)

	if self.err != nil {
		return self.err, 0
	}

	tmp := (*self.data)[self.index : self.index+8]
	self.index += 8
	b_buf := bytes.NewBuffer(tmp)

	var result int64
	binary.Read(b_buf, binary.BigEndian, &result)
	return nil, result
}

func (self *ReadBuf) ReadBytes() (error, []byte) {
	//	fmt.Println(fmt.Sprintf("11111 len:%d,index:%d\n"), len(*self.data), self.GetIndex())
	if self.err != nil {
		return self.err, nil
	}
	var length int16
	self.err, length = self.ReadInt16()

	if self.err != nil {
		return self.err, nil
	}

	if length == 0 {
		return nil, []byte{}
	}

	if length < 0 || length >= int16(0x8000-1) {
		return errors.New("length error:" + strconv.Itoa(int(length))), []byte{}
	}

	self.err = checkDataLength(*self.data, self.index, int(length))
	if self.err != nil {
		return self.err, nil
	}

	d := (*self.data)[self.index : self.index+int(length)]
	self.index += int(length)

	return nil, d
}

func (self *ReadBuf) ReadString() (error, string) {

	err, d := self.ReadBytes()

	return err, string(d)

}

func (self *ReadBuf) ReadFloat32() (error, float32) {
	err := checkDataLength(*self.data, self.index, 4)

	if err != nil {
		return err, 0
	}

	tmp := (*self.data)[self.index : self.index+4]
	self.index += 4
	b_buf := bytes.NewBuffer(tmp)

	var result float32
	binary.Read(b_buf, binary.BigEndian, &result)

	return nil, result
}
