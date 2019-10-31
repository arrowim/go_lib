package user

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
	"fmt"
	//	"fmt"
)

//func (self *SubDeviceToken) ToTokenString() string {
//	s := fmt.Sprintf("%d_%s", self.SubDeviceId, self.Token)
//	return s
//}

//func (self *SubDeviceToken) Equal(token string, subDeviceType int16) bool {
//	return self.SubDeviceId == subDeviceType &&
//		self.Token == token
//}

const (
	CURRENT_VERSION base.ProtocalVersion = 0x0201
)

type SecondLoginRequest2 struct {
	Header               *base.MessageHeader
	Version              base.ProtocalVersion `json:"protocolVersion"`
	UserId               int32                `json:"userId"`
	TimeStamp            int32                `json:"timeStamp"`
	Key                  string               `json:"key"`
	NeedReplaceOtherUser int16                `json:"needReplaceOtherUser"`
	LoginType            base.LoginType       `json:"loginType"`
	DeviceType           base.DeviceType      `json:"deviceType"`
	DevicePushType       base.DevicePushType  `json:"devicePushType"`
	UserToken            string               `json:"userToken"`
	AppId                int16                `json:"appId"`
	SubAppId             int16                `json:"subAppId"`
	LoginId              int64                `json:"loginId"`
}

func (self *SecondLoginRequest2) GetTokenString() string {
	return fmt.Sprintf("%d_%s", self.DevicePushType, self.UserToken)
}

func (self *SecondLoginRequest2) EquerToken(pushTypeId base.DevicePushType,
	token string) bool {
	return self.DevicePushType == pushTypeId &&
		self.UserToken == token
}

func (self *SecondLoginRequest2) IsNeedReplaceOtherUser() bool {
	return self.NeedReplaceOtherUser == 1
}

func (self *SecondLoginRequest2) SetNeedReplaceOtherUser(isNeed bool) {
	if isNeed {
		self.NeedReplaceOtherUser = 1
	} else {
		self.NeedReplaceOtherUser = 0
	}
}

func (self *SecondLoginRequest2) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SecondLoginRequest2) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateSecondLoginRequest2() *SecondLoginRequest2 {
	r := &SecondLoginRequest2{
		Header: base.CreateMessageHeader(cmd.SECOND_LOGIN_REQUEST2)}
	r.Version = CURRENT_VERSION
	r.Header.ServerHeader.MessagePriorty = base.PRIORTY_HIGH
	return r
}

func (self *SecondLoginRequest2) GetMessageHeader() *base.MessageHeader {
	return self.Header
}

func (self *SecondLoginRequest2) CreateResposne() *SecondLoginResponse2 {
	response := CreateSecondLoginResponse2()

	response.Header.SetHeader(self.Header)
	response.Header.CommandId = cmd.SECOND_LOGIN_RESPONSE2
	response.Header.ServerHeader.MessagePriorty = base.PRIORTY_HIGH

	return response

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SecondLoginRequest2) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt16(self.Version.GetValue())
	buf.WriteInt32(self.UserId)
	buf.WriteInt32(self.TimeStamp)
	buf.WriteString(self.Key)
	buf.WriteInt16(self.NeedReplaceOtherUser)
	buf.WriteInt16(self.AppId)
	buf.WriteInt16(self.SubAppId)

	buf.WriteInt16(int16(self.DeviceType))
	buf.WriteInt16(int16(self.DevicePushType))
	buf.WriteString(self.UserToken)
	if key == base.SERVER_MESSAGE {
		buf.WriteInt64(self.LoginId)
	}
	buf.WriteInt32(int32(self.LoginType))
	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SecondLoginRequest2) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}
	err, v := buf.ReadInt16()
	self.Version = base.ProtocalVersion(v)
	err, self.UserId = buf.ReadInt32()

	err, self.TimeStamp = buf.ReadInt32()
	err, self.Key = buf.ReadString()
	err, self.NeedReplaceOtherUser = buf.ReadInt16()
	err, self.AppId = buf.ReadInt16()
	err, self.SubAppId = buf.ReadInt16()
	err, t := buf.ReadInt16()
	self.DeviceType = base.DeviceType(t)
	err, t = buf.ReadInt16()
	self.DevicePushType = base.DevicePushType(t)
	err, self.UserToken = buf.ReadString()
	if key == base.SERVER_MESSAGE {
		err, self.LoginId = buf.ReadInt64()
	}

	err, t2 := buf.ReadFloat32()

	self.LoginType = base.LoginType(t2)


	//	fmt.Println("receive data:", b.Bytes2HexString(data))

	return err
}

func (self *SecondLoginRequest2) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
	switch targetVersion {
	default:
		return self
	}
}
