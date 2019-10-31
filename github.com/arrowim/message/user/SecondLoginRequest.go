package user
//
//import (
//	"arrowim/github.com/arrowim/message/base"
//	cmd "arrowim/github.com/arrowim/message/command"
//	b "arrowim/github.com/arrowim/utils/bytes"
//	"fmt"
//)
//
//type SecondLoginRequest struct {
//	Header               *base.MessageHeader
//	UserId               int32 `json:"userId"`
//	TimeStamp            int32 `json:"timeStamp"`
//	Key                  string `json:"key"`
//	NeedReplaceOtherUser int16 `json:"needReplaceOtherUser"`
//	DeviceType           base.DeviceType `json:"deviceType"`
//	DevicePushType       base.DevicePushType `json:"devicePushType"`
//	UserToken            string `json:"userToken"`
//	AppId                int16 `json:"appId"`
//	SubAppId             int16 `json:"subAppId"`
//}
//
//
//func (self *SecondLoginRequest) GetHeader() *base.MessageHeader {
//	return self.Header
//}
//
//func (self *SecondLoginRequest) SetHeader(header *base.MessageHeader) {
//	self.Header.SetHeader(header)
//}
//
//func CreateSecondLoginRequest() *SecondLoginRequest {
//	r := &SecondLoginRequest{
//		Header: base.CreateMessageHeader(cmd.SECOND_LOGIN_REQUEST)}
//	r.Header.ServerHeader.MessagePriorty=base.PRIORTY_HIGH
//
//	return r
//}
//
//func (self *SecondLoginRequest) GetMessageHeader() *base.MessageHeader {
//	return self.Header
//}
//
//func (self *SecondLoginRequest) CreateResposne() *SecondLoginResponse {
//	response := CreateSecondLoginResponse()
//
//	response.Header.SetHeader(self.Header)
//	response.Header.CommandId = cmd.SECOND_LOGIN_RESPONSE
//	response.Header.ServerHeader.MessagePriorty=base.PRIORTY_HIGH
//	return response
//
//}
//
////参数key：1标识发给客户端用，2标识服务器内部使用
//func (self *SecondLoginRequest) ToBytes(key int) []byte {
//	buf := b.CreateWriteBuf(true)
//	self.Header.ToBytes(buf, key)
//
//	buf.WriteInt32(self.UserId)
//	buf.WriteInt32(self.TimeStamp)
//	buf.WriteString(self.Key)
//	buf.WriteInt16(self.NeedReplaceOtherUser)
//	buf.WriteInt16(self.AppId)
//	buf.WriteInt16(self.SubAppId)
//	buf.WriteInt16(int16(self.DeviceType))
//	buf.WriteString(self.UserToken)
//	fmt.Println("token", self.UserToken)
//
//	return buf.Bytes()
//}
//
////参数key：1标识发给客户端用，2标识服务器内部使用
//func (self *SecondLoginRequest) FromBytes(data []byte, key int) error {
//	buf := b.CreateReadBuf(&data, true)
//	err := self.Header.FromBytes(buf, key)
//	if err != nil {
//		return err
//	}
//
//	//	fmt.Println("data:", b.Bytes2HexString(data))
//
//	err, self.UserId = buf.ReadInt32()
//
//	err, self.TimeStamp = buf.ReadInt32()
//	err, self.Key = buf.ReadString()
//	err, self.NeedReplaceOtherUser = buf.ReadInt16()
//	err, self.AppId = buf.ReadInt16()
//	err, self.SubAppId = buf.ReadInt16()
//	err, t := buf.ReadInt16()
//	self.DeviceType = base.DeviceType(t)
//	err, self.UserToken = buf.ReadString()
//
//	if len(self.UserToken) > 50 {
//		self.DevicePushType = base.DEVICE_PUSH_IOS
//	} else {
//		self.DevicePushType = base.DEVICE_PUSH_ANDROID
//	}
//
//	return err
//}
//
//func (self *SecondLoginRequest) createSecondLoginRequest2() *SecondLoginRequest2 {
//	s := CreateSecondLoginRequest2()
//	s.AppId = self.AppId
//	s.UserId = self.UserId
//	s.TimeStamp = self.TimeStamp
//	s.Key = self.Key
//	s.NeedReplaceOtherUser = self.NeedReplaceOtherUser
//	s.AppId = self.AppId
//	s.SubAppId = self.SubAppId
//	s.DeviceType = self.DeviceType
//	s.DevicePushType = self.DevicePushType
//	s.UserToken = self.UserToken
//	s.Version = 0
//	return s
//}
//
//
//func (self *SecondLoginRequest) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
//	switch targetVersion {
//	case 0:
//		return self
//	default:
//		return self.createSecondLoginRequest2()
//	}
//}