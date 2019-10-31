package user
//
//import (
//	"arrowim/github.com/arrowim/message/base"
//	cmd "arrowim/github.com/arrowim/message/command"
//	"arrowim/github.com/arrowim/message/user/loginresult"
//	b "arrowim/github.com/arrowim/utils/bytes"
//)
//
//type SecondLoginResponse struct {
//	Header  *base.MessageHeader
//	Result  loginresult.LoginResult  `json:"result"`//响应结果
//	LoginId int32 `json:"loginId"`
//
//	Token string `json:"token"`
//}
//
//func (self *SecondLoginResponse) GetHeader() *base.MessageHeader {
//	return self.Header
//}
//
//func (self *SecondLoginResponse) SetHeader(header *base.MessageHeader) {
//	self.Header.SetHeader(header)
//}
//
//func CreateSecondLoginResponse() *SecondLoginResponse {
//	r := &SecondLoginResponse{
//		Header: base.CreateMessageHeader(cmd.SECOND_LOGIN_RESPONSE)}
//	r.Header.ServerHeader.MessagePriorty=base.PRIORTY_HIGH
//	return r
//}
//
////参数key：1标识发给客户端用，2标识服务器内部使用
//func (self *SecondLoginResponse) ToBytes(key int) []byte {
//	buf := b.CreateWriteBuf(true)
//	self.Header.ToBytes(buf, key)
//
//	buf.WriteInt32(self.Result.GetValue())
//	buf.WriteInt32(self.LoginId)
//	buf.WriteString(self.Token)
//
//	return buf.Bytes()
//}
//
////参数key：1标识发给客户端用，2标识服务器内部使用
//func (self *SecondLoginResponse) FromBytes(data []byte, key int) error {
//	buf := b.CreateReadBuf(&data, true)
//	err := self.Header.FromBytes(buf, key)
//	if err != nil {
//		return err
//	}
//
//	err, r := buf.ReadInt32()
//	self.Result = loginresult.LoginResult(r)
//	err, self.LoginId = buf.ReadInt32()
//	err, self.Token = buf.ReadString()
//
//	return err
//}
//
//func (self *SecondLoginResponse) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
//	switch targetVersion {
//	default:
//		return self
//	}
//}