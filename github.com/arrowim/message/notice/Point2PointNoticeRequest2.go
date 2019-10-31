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

type Point2PointNoticeRequest2 struct {
	Header        *base.MessageHeader
	RequestId     int64   `json:"messageId"`     //请求ID
	SourceId      int32   `json:"sourceId"`      //发送通知者ID
	TargetIds     []int32 `json:"targetIds"`     //接收通知者ID
	NoticeType    int16   `json:"noticeType"`    //发送通知类型
	NoticeContent string  `json:"noticeContent"` //发送通知内容
	NoticeTime    int64   `json:"noticeTime"`    //推送时间
	SourceName    string  `json:"sourceName"`    //发送者名称
	SourceIcon    string  `json:"sourceIcon"`
	ShowContent   string  `json:"showContent"` //离线推送的显示内容
}

func (self *Point2PointNoticeRequest2) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *Point2PointNoticeRequest2) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreatePoint2PointNoticeRequest2() *Point2PointNoticeRequest2 {
	return &Point2PointNoticeRequest2{
		Header: base.CreateMessageHeader(cmd.SEND_NOTIFYCATION_REQUEST2)}
}

func (self Point2PointNoticeRequest2) CreateResposne() *Point2PointNoticeResponse2 {
	r := CreatePoint2PointNoticeResponse2()
	r.Header.ServerHeader.SetHeader(self.Header.ServerHeader)
	r.RequestId = self.RequestId
	r.RequestTime = timeutils.GetUnix13NowTime()
	return r
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self Point2PointNoticeRequest2) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt64(self.RequestId)
	buf.WriteInt32(self.SourceId)
	buf.WriteInt16(int16(len(self.TargetIds)))
	for _, v := range self.TargetIds {
		buf.WriteInt32(v)
	}

	buf.WriteInt16(self.NoticeType)
	buf.WriteString(self.NoticeContent)
	buf.WriteInt64(self.NoticeTime)

	buf.WriteString(self.SourceName)
	buf.WriteString(self.SourceIcon)

	return buf.Bytes()

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *Point2PointNoticeRequest2) FromBytes(data []byte, key int) error {

	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.RequestId = buf.ReadInt64()
	err, self.SourceId = buf.ReadInt32()
	err, l := buf.ReadInt16()
	for i := int16(0); i < l; i++ {
		_, id := buf.ReadInt32()
		self.TargetIds = append(self.TargetIds, id)
	}

	err, self.NoticeType = buf.ReadInt16()
	err, self.NoticeContent = buf.ReadString()
	err, self.NoticeTime = buf.ReadInt64()

	err, self.SourceName = buf.ReadString()
	err, self.SourceIcon = buf.ReadString()

	return err

}

func (self *Point2PointNoticeRequest2) createPoint2PointNoticeRequest() *Point2PointNoticeRequest {
	s := CreatePoint2PointNoticeRequest()
	s.Header.ServerHeader.SetHeader(self.Header.ServerHeader)
	s.RequestId = self.RequestId
	s.SourceName = self.SourceName
	s.SourceIcon = self.SourceIcon
	s.TargetId = self.TargetIds[0]
	s.SourceId = self.SourceId
	s.NoticeType = self.NoticeType
	s.NoticeContent = self.NoticeContent
	s.NoticeTime = self.NoticeTime
	s.ShowContent = self.ShowContent
	return s
}

func (self *Point2PointNoticeRequest2) CreatePoint2PointNoticeRequest(targetId int32) *Point2PointNoticeRequest {
	s := CreatePoint2PointNoticeRequest()
	s.Header.ServerHeader.SetHeader(self.Header.ServerHeader)
	s.RequestId = self.RequestId
	s.SourceName = self.SourceName
	s.SourceIcon = self.SourceIcon
	s.TargetId = targetId
	s.SourceId = self.SourceId
	s.NoticeType = self.NoticeType
	s.NoticeContent = self.NoticeContent
	s.NoticeTime = self.NoticeTime
	s.ShowContent = self.ShowContent
	return s
}

func (self *Point2PointNoticeRequest2) ToMessage(targetVersion base.ProtocalVersion) base.ArrowMessage {
	switch targetVersion {
	case 0:
		if len(self.TargetIds) < 0 {
			return nil
		}
		return self.createPoint2PointNoticeRequest()
	default:
		return self
	}
}
func (self *Point2PointNoticeRequest2) GetArrowMessageId() int64 {
	return self.RequestId
}

func (self *Point2PointNoticeRequest2) GetSourceId() int32 {
	return self.SourceId
}
