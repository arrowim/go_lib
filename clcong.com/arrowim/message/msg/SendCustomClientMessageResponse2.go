package msg

import (
	"arrowim/clcong.com/arrowim/message/base"
	cmd "arrowim/clcong.com/arrowim/message/command"
	b "arrowim/clcong.com/utils/bytes"
)

type SendCustomClientMessageResponse2 struct {
	Header       *base.MessageHeader
	Result       int32 `json:"result"`
	MsgId        int64 `json:"msgId"`
	SendTime     int64 `json:"sendTime"`
}


func (self *SendCustomClientMessageResponse2) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SendCustomClientMessageResponse2) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateSendCustomClientMessageResponse2() *SendCustomClientMessageResponse2 {
	return &SendCustomClientMessageResponse2{
		Header: base.CreateMessageHeader(cmd.SEND_CUSTOM_CLIENT_MESSAGERESPONSE2)}
}

func (self *SendCustomClientMessageResponse2) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt32(self.Result)
	buf.WriteInt64(self.MsgId)
	buf.WriteInt64(self.SendTime)

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SendCustomClientMessageResponse2) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.Result = buf.ReadInt32()
	err, self.MsgId = buf.ReadInt64()
	err, self.SendTime = buf.ReadInt64()

	return err

}

func (self *SendCustomClientMessageResponse2) ToMessage (targetVersion base.ProtocalVersion) (base.ArrowMessage){
	switch targetVersion{

	default:
		return self
	}
}