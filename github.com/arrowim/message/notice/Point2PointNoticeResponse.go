package notice

/**
服务器推送消息 响应
*/
import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type Point2PointNoticeResponse struct {
	Header       *base.MessageHeader
	RequestId    int64  `json:"messageId"`//对应的请求ID
	RequestTime  int64 `json:"requestTime"`
	SourceId     int32 `json:"sourceId"`
}

func (self *Point2PointNoticeResponse) GetArrowMessageId() int64 {
	return self.RequestId
}

func (self *Point2PointNoticeResponse) GetTargetId() int32 {
	return -1
}

func (self *Point2PointNoticeResponse) SetTargetId(targetId int32) {

}



func (self *Point2PointNoticeResponse) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *Point2PointNoticeResponse) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreatePoint2PointNoticeResponse() *Point2PointNoticeResponse {
	return &Point2PointNoticeResponse{
		Header: base.CreateMessageHeader(cmd.SEND_NOTIFYCATION_RESPONSE)}
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self Point2PointNoticeResponse) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt64(self.RequestId)
	buf.WriteInt64(self.RequestTime)
	buf.WriteInt32(self.SourceId)
	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *Point2PointNoticeResponse) FromBytes(data []byte, key int) error {

	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.RequestId = buf.ReadInt64()
	err, self.RequestTime = buf.ReadInt64()
	_, self.SourceId = buf.ReadInt32()
	//TODO为了兼容之前的版本，主动错误
	return nil
}

func (self *Point2PointNoticeResponse) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
	switch targetVersion {
	default:
		return self
	}
}
