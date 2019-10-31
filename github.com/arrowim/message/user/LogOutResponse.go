package user

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type LogOutResponse struct {
	Header *base.MessageHeader
}


func (self *LogOutResponse) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *LogOutResponse) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateLogOutResponse() *LogOutResponse {
	return &LogOutResponse{
		Header: base.CreateMessageHeader(cmd.LOG_OUT_RESPONSE)}
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self LogOutResponse) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *LogOutResponse) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	return err
}


func (self *LogOutResponse) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
	switch targetVersion {
	default:
		return self
	}
}