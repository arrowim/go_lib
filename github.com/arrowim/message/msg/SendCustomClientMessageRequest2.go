package msg

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

//服务端发送给客户端自定义消息
type SendCustomClientMessageRequest2 struct {
	Header    *base.MessageHeader
	ServerId  int16    `json:"serverId"`//发送服务器ID
	TargetIds []int32  `json:"targetIds"`//接收消息者ID
	MessageId int64    `json:"messageId"`//消息Id

	MessageContent string `json:"messageContent"`//消息内容
	//	MessageContentStr string
	SendTime      int64                   `json:"sendTime"`//消息发送时间
	MessageFormat base.MessageFormatType  `json:"messageFormat"`//消息类型
}

func CreateSendCustomCliendMessageRequest2() *SendCustomClientMessageRequest2 {
	return &SendCustomClientMessageRequest2{
		Header: base.CreateMessageHeader(cmd.SEND_CUSTOM_CLIENT_MESSAGEREQUEST2)}
}

func (self *SendCustomClientMessageRequest2) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SendCustomClientMessageRequest2) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func (self *SendCustomClientMessageRequest2) CreateSendCustomClientMessageResponse() *SendCustomClientMessageResponse {
	respone := CreateSendCustomClientMessageResponse()
	respone.Header.SetHeader(self.Header)
	respone.Header.CommandId = cmd.SEND_CUSTOM_CLIENT_MESSAGERESPONSE
	return respone
}

func (self *SendCustomClientMessageRequest2) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt16(self.ServerId)
	buf.WriteInt16(int16(len(self.TargetIds)))

	for _, v := range self.TargetIds {
		buf.WriteInt32(v)
	}

	buf.WriteInt64(self.MessageId)
	buf.WriteString(self.MessageContent)
	buf.WriteInt64(self.SendTime)
	buf.WriteInt16(self.MessageFormat.GetValue())

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SendCustomClientMessageRequest2) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.ServerId = buf.ReadInt16()
	err, l := buf.ReadInt16()

	for i := int16(0); i < l; i++ {
		_, i := buf.ReadInt32()
		self.TargetIds = append(self.TargetIds, i)
	}

	err, self.MessageId = buf.ReadInt64()
	err, self.MessageContent = buf.ReadString()
	err, self.SendTime = buf.ReadInt64()
	err, f := buf.ReadInt16()
	self.MessageFormat = base.MessageFormatType(f)
	return err

}

func (self *SendCustomClientMessageRequest2) createSendCustomClientMessageRequest(
	userId int32) *SendCustomClientMessageRequest {
	s := CreateSendCustomCliendMessageRequest()

	s.ServerId = self.ServerId
	s.TargetId = userId
	s.MessageId = self.MessageId
	s.MessageFormat = self.MessageFormat
	s.MessageContent = self.MessageContent
	s.SendTime = self.SendTime

	return s
}

func (self *SendCustomClientMessageRequest2) ToMessage(targetVersion base.ProtocalVersion) base.ArrowMessage {
	switch targetVersion {
	case 0:
		return self.createSendCustomClientMessageRequest(
			self.TargetIds[0])
	default:
		return self
	}
}

func (self *SendCustomClientMessageRequest2) GetArrowMessageId() int64 {
	return self.MessageId
}

func (self *SendCustomClientMessageRequest2) GetSourceId() int32 {
	return 0
}
