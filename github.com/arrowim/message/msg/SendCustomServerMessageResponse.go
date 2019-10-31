package msg

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type SendCustomServerMessageResponse struct {
	Header             *base.MessageHeader
	Result             int32   `json:"result"`//响应结果
	ServerMessageId    int64   `json:"serverMessageId"`//后台生成的消息的ID
	SendTime           int64   `json:"sendTime"`//消息发送的时间
	ClientMsgId        int64   `json:"clientMsgId"`//客户端所传信息ID
}

func (self *SendCustomServerMessageResponse) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SendCustomServerMessageResponse) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateSendCustomServerMessageResponse() *SendCustomServerMessageResponse {
	return &SendCustomServerMessageResponse{
		Header: base.CreateMessageHeader(cmd.SEND_CUSTOM_SERVER_MESSAGERESPONSE)}
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self SendCustomServerMessageResponse) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt32(self.Result)
	buf.WriteInt64(self.ServerMessageId)
	buf.WriteInt64(self.SendTime)
	buf.WriteInt64(self.ClientMsgId)

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SendCustomServerMessageResponse) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.Result = buf.ReadInt32()
	err, self.ServerMessageId = buf.ReadInt64()
	err, self.SendTime = buf.ReadInt64()
	err, self.ClientMsgId = buf.ReadInt64()

	return err

}


func (self *SendCustomServerMessageResponse) ToMessage (targetVersion base.ProtocalVersion) (base.ArrowMessage){
	switch targetVersion{

	default:
		return self
	}
}
