package message

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	"arrowim/github.com/arrowim/utils/bytes"
)

type ErrorMessageResponse struct {
	Header *base.MessageHeader
}

func (self *ErrorMessageResponse) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *ErrorMessageResponse) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateErrorMessageResponse() *ErrorMessageResponse {
	return &ErrorMessageResponse{
		Header: base.CreateMessageHeader(cmd.ERROR_MESSAGE_RESPONSE)}
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self ErrorMessageResponse) ToBytes(key int) []byte {
	buf := bytes.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *ErrorMessageResponse) FromBytes(data []byte, key int) error {
	buf := bytes.CreateReadBuf(&data, true)
	return self.Header.FromBytes(buf, key)
}
