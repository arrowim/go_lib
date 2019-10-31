package notice

/**
服务器推送消息请求
*/
import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
	"arrowim/github.com/arrowim/utils/timeutils"
)

type Point2PointNoticeRequest struct {
	Header        *base.MessageHeader
	RequestId     int64  `json:"messageId"`     //请求ID
	SourceId      int32  `json:"sourceId"`      //发送通知者ID
	TargetId      int32  `json:"targetId"`      //接收通知者ID
	NoticeType    int16  `json:"noticeType"`    //发送通知类型
	NoticeContent string `json:"noticeContent"` //发送通知内容
	NoticeTime    int64  `json:"noticeTime"`    //推送时间
	SourceName    string `json:"sourceName"`    //发送者名称
	SourceIcon    string `json:"sourceIcon"`
	ShowContent   string `json:"showContent"` //离线推送的显示内容
}

func (self *Point2PointNoticeRequest) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *Point2PointNoticeRequest) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreatePoint2PointNoticeRequest() *Point2PointNoticeRequest {
	return &Point2PointNoticeRequest{
		Header: base.CreateMessageHeader(cmd.SEND_NOTIFYCATION_REQUEST)}
}

func CreatePoint2PointNoticeRequest_nopoint() Point2PointNoticeRequest {
	return Point2PointNoticeRequest{
		Header: base.CreateMessageHeader(cmd.SEND_NOTIFYCATION_REQUEST)}
}

func (self Point2PointNoticeRequest) CreateResposne() *Point2PointNoticeResponse {
	r := CreatePoint2PointNoticeResponse()
	r.Header.ServerHeader.SetHeader(self.Header.ServerHeader)
	r.RequestId = self.RequestId
	r.RequestTime = timeutils.GetUnix13NowTime()
	return r
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self Point2PointNoticeRequest) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt64(self.RequestId)
	buf.WriteInt32(self.SourceId)
	buf.WriteInt32(self.TargetId)
	buf.WriteInt16(self.NoticeType)
	buf.WriteString(self.NoticeContent)
	buf.WriteInt64(self.NoticeTime)

	buf.WriteString(self.SourceName)
	buf.WriteString(self.SourceIcon)

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *Point2PointNoticeRequest) FromBytes(data []byte, key int) error {

	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.RequestId = buf.ReadInt64()
	err, self.SourceId = buf.ReadInt32()
	err, self.TargetId = buf.ReadInt32()
	err, self.NoticeType = buf.ReadInt16()
	err, self.NoticeContent = buf.ReadString()
	err, self.NoticeTime = buf.ReadInt64()

	err, self.SourceName = buf.ReadString()
	err, self.SourceIcon = buf.ReadString()
	return err

}

func (self *Point2PointNoticeRequest) ToMessage(targetVersion base.ProtocalVersion) base.ArrowMessage {
	switch targetVersion {
	default:
		return self
	}
}
func (self *Point2PointNoticeRequest) GetArrowMessageId() int64 {
	return self.RequestId
}

func (self *Point2PointNoticeRequest) GetSourceId() int32 {
	return self.SourceId
}
