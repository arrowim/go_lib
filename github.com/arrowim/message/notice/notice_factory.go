package notice

import (
	cmd "arrowim/github.com/arrowim/message/command"
	"arrowim/github.com/arrowim/message/base"
)

func CreateMessage(cmdId cmd.Command, data []byte, key int) (error,base.ArrowMessage) {
	switch cmdId {


	//发送单条推送给客户端的通知
	case cmd.SEND_NOTIFYCATION_REQUEST:
		msg := CreatePoint2PointNoticeRequest()
		return msg.FromBytes(data, key), msg

		//发送单条推送给客户端的通知的响应
	case cmd.SEND_NOTIFYCATION_RESPONSE:
		msg := CreatePoint2PointNoticeResponse()
		return msg.FromBytes(data, key), msg

		//发送单条推送给客户端的通知
	case cmd.SEND_NOTIFYCATION_REQUEST2:
		msg := CreatePoint2PointNoticeRequest2()
		return msg.FromBytes(data, key), msg

		//发送单条推送给客户端的通知的响应
	case cmd.SEND_NOTIFYCATION_RESPONSE2:
		msg := CreatePoint2PointNoticeResponse2()
		return msg.FromBytes(data, key), msg

	default:
		return nil,nil
	}
}