package notice

/**
服务器推送消息 响应
*/
import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type Point2PointNoticeResponse2 struct {
	Header       *base.MessageHeader
	RequestId    int64  `json:"messageId"`//对应的请求ID
	RequestTime  int64 `json:"requestTime"`
}



func (self *Point2PointNoticeResponse2) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *Point2PointNoticeResponse2) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreatePoint2PointNoticeResponse2() *Point2PointNoticeResponse2 {
	return &Point2PointNoticeResponse2{
		Header: base.CreateMessageHeader(cmd.SEND_NOTIFYCATION_RESPONSE2)}
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self Point2PointNoticeResponse2) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt64(self.RequestId)
	buf.WriteInt64(self.RequestTime)

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *Point2PointNoticeResponse2) FromBytes(data []byte, key int) error {

	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.RequestId = buf.ReadInt64()
	err, self.RequestTime = buf.ReadInt64()

	//TODO为了兼容之前的版本，主动错误
	return nil
}

func (self *Point2PointNoticeResponse2) ToMessage (targetVersion base.ProtocalVersion) (base.ArrowMessage){
	switch targetVersion{
	default:
		return self
	}
}