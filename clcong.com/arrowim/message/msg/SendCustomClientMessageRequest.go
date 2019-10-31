package msg

import (
	"arrowim/clcong.com/arrowim/message/base"
	cmd "arrowim/clcong.com/arrowim/message/command"
	b "arrowim/clcong.com/utils/bytes"
)

//服务端发送给客户端自定义消息
type SendCustomClientMessageRequest struct {
	Header    *base.MessageHeader
	ServerId  int16  `json:"serverId"`//发送服务器ID
	TargetId  int32  `json:"targetId"`//接收消息者ID
	MessageId int64  `json:"messageId"`//消息Id

	MessageContent string `json:"messageContent"`//消息内容
	//	MessageContentStr string
	SendTime      int64                   `json:"sendTime"`//消息发送时间
	MessageFormat base.MessageFormatType  `json:"messageFormat"`//消息类型]
}

func CreateSendCustomCliendMessageRequest() *SendCustomClientMessageRequest {
	return &SendCustomClientMessageRequest{
		Header: base.CreateMessageHeader(cmd.SEND_CUSTOM_CLIENT_MESSAGEREQUEST)}
}

func (self *SendCustomClientMessageRequest) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SendCustomClientMessageRequest) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func (self *SendCustomClientMessageRequest) CreateSendCustomClientMessageResponse() *SendCustomClientMessageResponse {
	respone := CreateSendCustomClientMessageResponse()
	respone.Header.SetHeader(self.Header)
	respone.Header.CommandId = cmd.SEND_CUSTOM_CLIENT_MESSAGERESPONSE
	return respone
}

func (self *SendCustomClientMessageRequest) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt16(self.ServerId)
	buf.WriteInt32(self.TargetId)
	buf.WriteInt64(self.MessageId)
	buf.WriteString(self.MessageContent)
	buf.WriteInt64(self.SendTime)
	buf.WriteInt16(self.MessageFormat.GetValue())

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SendCustomClientMessageRequest) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.ServerId = buf.ReadInt16()
	err, self.TargetId = buf.ReadInt32()
	err, self.MessageId = buf.ReadInt64()
	err, self.MessageContent = buf.ReadString()
	err, self.SendTime = buf.ReadInt64()
	err, f := buf.ReadInt16()
	self.MessageFormat = base.MessageFormatType(f)
	return err

}

func (self *SendCustomClientMessageRequest) ToMessage(targetVersion base.ProtocalVersion) base.ArrowMessage {
	switch targetVersion {

	default:
		return self
	}
}
func (self *SendCustomClientMessageRequest) GetArrowMessageId() int64 {
	return self.MessageId
}

func (self *SendCustomClientMessageRequest) GetSourceId() int32 {
	return 0
}
