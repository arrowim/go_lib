package msg


import (
	cmd "arrowim/github.com/arrowim/message/command"
	"arrowim/github.com/arrowim/message/base"
)

func CreateMessage(cmdId cmd.Command, data []byte, key int) (error,base.ArrowMessage) {
	switch cmdId {


	case cmd.SEND_CUSTOM_CLIENT_MESSAGEREQUEST:
		msg := CreateSendCustomCliendMessageRequest()
		return msg.FromBytes(data, key), msg

	case cmd.SEND_CUSTOM_CLIENT_MESSAGERESPONSE:
		msg := CreateSendCustomClientMessageResponse()
		return msg.FromBytes(data, key), msg

	case cmd.SEND_CUSTOM_CLIENT_MESSAGEREQUEST2:
		msg := CreateSendCustomCliendMessageRequest()
		return msg.FromBytes(data, key), msg

	case cmd.SEND_CUSTOM_CLIENT_MESSAGERESPONSE2:
		msg := CreateSendCustomClientMessageResponse()
		return msg.FromBytes(data, key), msg

	case cmd.SEND_CUSTOM_SERVER_MESSAGEREQUEST:
		msg := CreateSendCustomServerMessageRequest()
		return msg.FromBytes(data, key), msg

	case cmd.SEND_CUSTOM_SERVER_MESSAGERESPONSE:
		msg := CreateSendCustomServerMessageResponse()
		return msg.FromBytes(data, key), msg

	default:
		return nil,nil
	}
}
