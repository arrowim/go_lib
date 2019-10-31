package user

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type UserLoginWithOtherDeviceRequest struct {
	Header       *base.MessageHeader
	RequestId    int64 `json:"messageId"`
	UserId       int32 `json:"userId"`
	LoginId      int64 `json:"loginId"`
}


func CreateUserLoginWithOtherDeviceRequest() *UserLoginWithOtherDeviceRequest {
	return &UserLoginWithOtherDeviceRequest{
		Header: base.CreateMessageHeader(
			cmd.USER_LOGIN_WITH_OTHER_DEVICE_REQUEST)}

}

func (self *UserLoginWithOtherDeviceRequest) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *UserLoginWithOtherDeviceRequest) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

func (self UserLoginWithOtherDeviceRequest) CreateResposne() *UserLoginWithOtherDeviceResponse {
	response := CreateUserLoginWithOtherDeviceResponse()

	response.Header.SetHeader(self.Header)
	response.Header.CommandId = cmd.USER_LOGIN_WITH_OTHER_DEVICE_RESPONSE

	return response
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self UserLoginWithOtherDeviceRequest) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt64(self.RequestId)
	buf.WriteInt32(self.UserId)
	buf.WriteInt64(self.LoginId)
	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *UserLoginWithOtherDeviceRequest) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)

	err, self.RequestId = buf.ReadInt64()
	err, self.UserId = buf.ReadInt32()
	err, self.LoginId = buf.ReadInt64()
	return err
}

func (self *UserLoginWithOtherDeviceRequest) ToMessage (targetVersion base.ProtocalVersion) (base.ArrowMessage){
	switch targetVersion{

	default:
		return self
	}
}
