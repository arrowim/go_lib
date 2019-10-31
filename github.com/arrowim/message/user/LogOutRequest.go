package user

import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	b "arrowim/github.com/arrowim/utils/bytes"
)

type LogOutRequest struct {
	Header     *base.MessageHeader
	AppId      int16 `json:"appId"`
	SubAppId   int16 `json:"subAppId"`
	DeviceType int16 `json:"deviceType"`
}


func CreateLogOutRequest() *LogOutRequest {
	return &LogOutRequest{
		Header: base.CreateMessageHeader(cmd.LOG_OUT_REQUEST)}
}

func (self LogOutRequest) CreateResposne() *LogOutResponse {
	response := CreateLogOutResponse()

	response.Header.SetHeader(self.Header)
	response.Header.CommandId = cmd.LOG_OUT_RESPONSE

	return response
}

func (self *LogOutRequest) GetHeader() *base.MessageHeader {
	return self.Header
}

func (self *LogOutRequest) SetHeader(header *base.MessageHeader) {
	self.Header.SetHeader(header)
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self LogOutRequest) ToBytes(key int) []byte {
	buf := b.CreateWriteBuf(true)
	self.Header.ToBytes(buf, key)

	buf.WriteInt16(self.AppId)
	buf.WriteInt16(self.SubAppId)
	buf.WriteInt16(self.DeviceType)

	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (self *LogOutRequest) FromBytes(data []byte, key int) error {
	buf := b.CreateReadBuf(&data, true)
	err := self.Header.FromBytes(buf, key)
	if err != nil {
		return err
	}

	err, self.AppId = buf.ReadInt16()
	err, self.SubAppId = buf.ReadInt16()
	err, self.DeviceType = buf.ReadInt16()

	return err
}

func (self *LogOutRequest) ToMessage(targetVersion base.ProtocalVersion) (base.ArrowMessage) {
	switch targetVersion {
	default:
		return self
	}
}
