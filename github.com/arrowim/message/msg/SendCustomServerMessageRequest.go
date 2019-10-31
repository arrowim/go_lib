package msg

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

const (
	MESSAGE_TYPE_NOTIFY = 0x0301
)

//客户端发送给服务端的自定义 透传消息
type SendCustomServerMessageRequest struct {
	Header         *base.MessageHeader
	SourceId       int32                   `json:"sourceId"`//发送消息者ID
	MessageId      int64                   `json:"messageId"`//消息ID
	CustomServerId int16                   `json:"customServerId"`//服务端Id
	MessageContent string                  `json:"messageContentId"`//消息内容
	SendTime       int64                   `json:"sendTime"`//发送消息时间
	MessageFormat  base.MessageFormatType  `json:"messageFormat"`//消息类型
}

func (self *SendCustomServerMessageRequest) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SendCustomServerMessageRequest) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateSendCustomServerMessageRequest() *SendCustomServerMessageRequest {
	return &SendCustomServerMessageRequest{
		Header: base.CreateMessageHeader(cmd.SEND_CUSTOM_SERVER_MESSAGEREQUEST)}
}

func (self *SendCustomServerMessageRequest) CreateResposne() *SendCustomServerMessageResponse {
	response := CreateSendCustomServerMessageResponse()

	response.Header.SetHeader(self.Header)
	response.Header.CommandId = cmd.SEND_CUSTOM_SERVER_MESSAGERESPONSE

	response.ClientMsgId = self.MessageId

	return response
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SendCustomServerMessageRequest) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt32(self.SourceId)
	buf.WriteInt64(self.MessageId)
	buf.WriteInt16(self.CustomServerId)
	buf.WriteString(self.MessageContent)
	buf.WriteInt64(self.SendTime)
	buf.WriteInt16(self.MessageFormat.GetValue())

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SendCustomServerMessageRequest) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.SourceId = buf.ReadInt32()
	err, self.MessageId = buf.ReadInt64()
	err, self.CustomServerId = buf.ReadInt16()
	err, self.MessageContent = buf.ReadString()
	err, self.SendTime = buf.ReadInt64()
	err, f := buf.ReadInt16()
	self.MessageFormat = base.MessageFormatType(f)
	return err

}

func (self *SendCustomServerMessageRequest) ToMessage(targetVersion base.ProtocalVersion) base.ArrowMessage {
	switch targetVersion {

	default:
		return self
	}
}

func (self *SendCustomServerMessageRequest) GetArrowMessageId() int64 {
	return self.MessageId
}

func (self *SendCustomServerMessageRequest) GetSourceId() int32 {
	return self.SourceId
}
