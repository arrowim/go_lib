package message

/**
消息创建工厂，通过 commandId 来创建各种request 与 response
*/
import (
	"arrowim/github.com/arrowim/message/base"
	cmd "arrowim/github.com/arrowim/message/command"
	msgpkg "arrowim/github.com/arrowim/message/msg"
	noticepkg "arrowim/github.com/arrowim/message/notice"
	"arrowim/github.com/arrowim/message/user"
	b "arrowim/github.com/arrowim/utils/bytes"
)

//获取命令编号
func getCommandId(data []byte) cmd.Command {
	if len(data) == 0 {
		return 0
	}
	_, v16 := b.Bytes2Int16(data, 4)
	return cmd.Command(v16)
}

func CreateBaseMessage(cmdId cmd.Command, data []byte, key int) (error, base.ArrowMessage) {
	switch cmdId {
	case cmd.KEEP_ALIVE_REQUEST:
		msg := CreateKeepAliveRequest()
		return msg.FromBytes(data, key), msg

	case cmd.KEEP_ALIVE_RESPONSE:
		msg := CreateKeepAliveResponse()
		return msg.FromBytes(data, key), msg

	case cmd.LOG_OUT_REQUEST:
		//		logs.Debug("messagefactory-->create logoutRequest")
		msg := user.CreateLogOutRequest()
		return msg.FromBytes(data, key), msg

	case cmd.LOG_OUT_RESPONSE:
		//		logs.Debug("messagefactory-->create logoutResponse")
		msg := user.CreateLogOutResponse()
		return msg.FromBytes(data, key), msg

		//	case cmd.READ_MESSAGE_RESPONSE:
		//		logs.Debug("messagefactory-->create readMessageResponse")
		//		msg := msgpkg.CreateReadMessageResponse(sequnceId)
		//		msg.FromBytes(data, key)
		//		return msg

	default:
		return nil, nil
	}
}

//创建消息体
func CreateMessageFromBytes(data []byte, key int) (error, base.ArrowMessage) {
	commandId := getCommandId(data)
	b := cmd.Command(commandId.GetValue() / 0x100 * 0x100)
	var err error
	var m base.ArrowMessage
	switch b {
	case cmd.USER_MESSAGE:
		err, m = user.CreateMessage(commandId, data, key)

	case cmd.MESSAGE_MESSAGE_BASE:
		err, m = msgpkg.CreateMessage(commandId, data, key)

	case cmd.BASE_MESSAGE:
		err, m = CreateBaseMessage(commandId, data, key)

	case cmd.PUSH_MESSAGE:
		err, m = noticepkg.CreateMessage(commandId, data, key)

	default:
		//		logs.Info("commandId is ---->" + strconv.Itoa(int(commandId)))
		//		logs.Info("data--------->:", data)
		msg := CreateErrorMessageRequest()
		msg.Data = data
		err, m = nil, msg
	}

	if m == nil {
		msg := CreateErrorMessageRequest()
		msg.Data = data
		err, m = nil, msg
		return err, msg
	}
	return err, m
}
