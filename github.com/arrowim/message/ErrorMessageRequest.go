package message

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	"arrowim/github.com/arrowim/utils/bytes"
)

type ErrorMessageRequest struct {
	Header *base.MessageHeader
	Data   []byte  `json:"data"`
}


func (self *ErrorMessageRequest) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *ErrorMessageRequest) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateErrorMessageRequest() *ErrorMessageRequest {
	return &ErrorMessageRequest{
		Header: base.CreateMessageHeader(cmd.ERROR_MESSAGE_REQUEST)}
}

func (self ErrorMessageRequest) CreateResposne() *ErrorMessageResponse {
	return CreateErrorMessageResponse()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self ErrorMessageRequest) ToBytes(key int) []byte {
	buf := bytes.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)
	buf.WriteBytes(self.Data)
	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *ErrorMessageRequest) FromBytes(data []byte, key int) error {
	buf := bytes.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	err, self.Data = buf.ReadBytes()
	return err
}

func (self *ErrorMessageRequest) ToMessage (targetVersion base.ProtocalVersion) (base.ArrowMessage){
	switch targetVersion{

	default:
		return self
	}
}