package user

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	"arrowim/github.com/arrowim/message/user/loginresult"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type SecondLoginResponse2 struct {
	Header   *base.MessageHeader
	Result   loginresult.LoginResult `json:"result"`//响应结果
	UserId   int32 `json:"userId"`
	LoginId  int64 `json:"loginId"`
	Token    string `json:"token"`
	Friends  []int32 `json:"friends"`
	Groups   []int32 `json:"groups"`
	UserName string `json:"userName"`
	UserIcon string `json:"userIcon"`
}

func (self *SecondLoginResponse2) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *SecondLoginResponse2) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateSecondLoginResponse2() *SecondLoginResponse2 {
	r := &SecondLoginResponse2{
		Header: base.CreateMessageHeader(cmd.SECOND_LOGIN_RESPONSE2)}
	r.Header.ServerHeader.MessagePriorty = base.PRIORTY_HIGH

	return r
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SecondLoginResponse2) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt32(self.Result.GetValue())
	buf.WriteInt32(self.UserId)
	buf.WriteString(self.Token)

	if key == base.SERVER_MESSAGE {
		buf.WriteInt64(self.LoginId)
	}

	buf.WriteInt16(int16(len(self.Friends)));
	for _, v := range self.Friends {
		buf.WriteInt32(v)
	}

	buf.WriteInt16(int16(len(self.Groups)))
	for _, v := range self.Groups {
		buf.WriteInt32(v)
	}

	buf.WriteString(self.UserName)
	buf.WriteString(self.UserIcon)

	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *SecondLoginResponse2) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, r := buf.ReadInt32()
	self.Result = loginresult.LoginResult(r)
	err, self.UserId = buf.ReadInt32()
	err, self.Token = buf.ReadString()
	if key == base.SERVER_MESSAGE {
		err, self.LoginId = buf.ReadInt64()
	}

	err, l := buf.ReadInt16()
	for i := int16(0); i < l; i ++ {
		err, v := buf.ReadInt32()
		if (err == nil) {
			self.Friends = append(self.Friends, v)
		}
	}

	err, l = buf.ReadInt16()
	for i := int16(0); i < l; i ++ {
		err, v := buf.ReadInt32()
		if (err == nil) {
			self.Groups = append(self.Groups, v)
		}
	}

	err, self.UserName = buf.ReadString()
	err, self.UserIcon = buf.ReadString()
	return err
}
//
//func (self *SecondLoginResponse2) createSecondLoginResponse() *SecondLoginResponse {
//	r := CreateSecondLoginResponse()
//	r.Result = self.Result
//	r.LoginId = self.UserId
//	r.Token = self.Token
//	return r
//}

func (self *SecondLoginResponse2) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
	switch targetVersion {
	//case 0:
	//	return self.createSecondLoginResponse()
	default:
		return self
	}
}
