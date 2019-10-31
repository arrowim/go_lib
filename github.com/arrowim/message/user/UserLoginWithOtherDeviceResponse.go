package user

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type UserLoginWithOtherDeviceResponse struct {
	Header       *base.MessageHeader
	RequestId    int64 `json:"messageId"`
}


func (self *UserLoginWithOtherDeviceResponse) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *UserLoginWithOtherDeviceResponse) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func CreateUserLoginWithOtherDeviceResponse() *UserLoginWithOtherDeviceResponse {
	return &UserLoginWithOtherDeviceResponse{
		Header: base.CreateMessageHeader(cmd.USER_LOGIN_WITH_OTHER_DEVICE_RESPONSE)}

}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self UserLoginWithOtherDeviceResponse) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt64(self.RequestId)

	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *UserLoginWithOtherDeviceResponse) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.RequestId = buf.ReadInt64()

	return err
}



func (self *UserLoginWithOtherDeviceResponse) ToMessage (targetVersion base.ProtocalVersion) (base.ArrowMessage){
	switch targetVersion{

	default:
		return self
	}
}