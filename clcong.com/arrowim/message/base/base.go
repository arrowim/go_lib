package base

import (
	"arrowim/clcong.com/arrowim/message/command"
	"arrowim/clcong.com/arrowim/message/idmanager"
	b "arrowim/clcong.com/utils/bytes"
	//	"fmt"
	//	"fmt"
)

type MESSAGE_PRIORTY int16

func (self *MESSAGE_PRIORTY) GetPriortyString() string {
	s := "_priorty_"
	switch *self {
	case PRIORTY_HIGH:
		s += "high"
	case PRIORTY_MIDDLE:
		s += "middle"
	case PRIORTY_LOW:
		s += "low"
	default:
		s += "default"
	}

	return s

}

func CreateAllPriortys() []MESSAGE_PRIORTY {
	d := []MESSAGE_PRIORTY{PRIORTY_HIGH, PRIORTY_MIDDLE, PRIORTY_LOW}
	return d
}

const (
	CLIENT_MESSAGE = 1
	SAVE_MESSAGE   = 2
	SERVER_MESSAGE = 6

	PRIORTY_HIGH   MESSAGE_PRIORTY = 1
	PRIORTY_MIDDLE MESSAGE_PRIORTY = 0
	PRIORTY_LOW    MESSAGE_PRIORTY = -1
)

type GROUP_TYPE int32;

const (
	GROUP_TYPE_GROUP   GROUP_TYPE = 2
	GROUP_TYPE_DISCUSS GROUP_TYPE = 1
	GROUP_TYPE_ALL     GROUP_TYPE = 3
)

const (
	USER_OFF_ONLINE = -100 //用户不在线
)

type MessageHeader struct {
	MessageLength int32
	CommandId     command.Command `json:"cmd"`

	ServerHeader *ServerMessageHeader
}


func (self *MessageHeader) SetHeader(header *MessageHeader) {
	self.CommandId = header.CommandId
	self.MessageLength = header.MessageLength

	if header.ServerHeader != nil {
		if self.ServerHeader == nil {
			self.ServerHeader = &ServerMessageHeader{}
		}
		self.ServerHeader.SetHeader(header.ServerHeader)
	}
}

func GetHeaderLength(key int) int {
	if key == SERVER_MESSAGE {
		return 14
	} else {
		return 6
	}
}

var ids *idmanager.MessageIdManager

func getIdManager() *idmanager.MessageIdManager {
	if ids == nil {
		ids = idmanager.CreateMessageIdManager(0, 0)

	}

	return ids
}

// offlineType 处理离线消息的方式
func CreateMessageHeader(commandId command.Command) *MessageHeader {
	s := &MessageHeader{
		CommandId: commandId}
	s.ServerHeader = &ServerMessageHeader{}
	s.ServerHeader.LoginId = getIdManager().CreateId()
	//	s.ServerHeader.IS_pushOffLine = PUSH_OFF_LINE
	return s
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (m *MessageHeader) ToBytes(buf *b.WriteBuf, key int) []byte {

	buf.WriteInt16(m.CommandId.GetValue())

	if key == SERVER_MESSAGE {
		h := m.ServerHeader
		buf.WriteInt32(h.UserId)
		buf.WriteInt16(h.GatewayId)

		buf.WriteInt16(h.AppId)
		buf.WriteInt16(h.SubAppId)
		buf.WriteInt16(int16(h.DeviceType))
		buf.WriteInt64(h.LoginId)
		buf.WriteInt16(int16(h.MessagePriorty))
	}
	return buf.Bytes()
}

//参数key：1标识发给客户端用，2标识服务器内部使用
func (m *MessageHeader) FromBytes(buf *b.ReadBuf, key int) error {
	var err error
	err, m.MessageLength = buf.ReadInt32()
	if err != nil {
		return err
	}

	if key == SERVER_MESSAGE {
		err, m.ServerHeader.UserId = buf.ReadInt32()
		err, m.ServerHeader.GatewayId = buf.ReadInt16()

		err, m.ServerHeader.AppId = buf.ReadInt16()
		err, m.ServerHeader.SubAppId = buf.ReadInt16()
		_, t := buf.ReadInt16()
		m.ServerHeader.DeviceType = DeviceType(t)
		err, m.ServerHeader.LoginId = buf.ReadInt64()
		var d int16

		err, d = buf.ReadInt16()
		m.ServerHeader.MessagePriorty = MESSAGE_PRIORTY(d)

	}

	return err
}

type ArrowMessage interface {
	ToBytes(key int) []byte

	GetHeader() *MessageHeader

	SetHeader(header *MessageHeader)

	ToMessage(targetVersion ProtocalVersion) ArrowMessage
}

type ArrowRequest interface {
	ArrowMessage

	GetArrowMessageId() int64

	GetSourceId() int32
}

type TargetType int

const (
	TARGET_TYPE_GROUP TargetType = 2;
	TARGET_TYPE_USER  TargetType = 1;
)

type LoginType int

const (
	LOGIN_TYPE_NORMAL      LoginType = 0
	LOGIN_TYPE_HTTP_NORMAL LoginType = 20
	LOGIN_TYPE_HTTP_TOKEN  LoginType = 21
	LOGIN_TYPE_AD          LoginType = 30
)

const CURRENT_VERSION ProtocalVersion = 0x0101

type ProtocalVersion int16

func (self ProtocalVersion) GetValue() int16 {
	return int16(self)
}
