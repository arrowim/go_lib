package user

import (
	cmd "arrowim/github.com/arrowim/message/command"
	"arrowim/github.com/arrowim/message/base"
)

func CreateMessage(cmdId cmd.Command, data []byte, key int) (error, base.ArrowMessage) {
	switch cmdId {

	//case cmd.SECOND_LOGIN_REQUEST:
	//	msg := CreateSecondLoginRequest2()
	//	err := msg.FromBytes(data, key)
	//	return err, msg.ToMessage(base.CURRENT_VERSION)
	//
	//case cmd.SECOND_LOGIN_RESPONSE:
	//	msg := CreateSecondLoginResponse2()
	//	return msg.FromBytes(data, key), msg

	case cmd.SECOND_LOGIN_REQUEST2:
		msg := CreateSecondLoginRequest2()
		return msg.FromBytes(data, key), msg

	case cmd.SECOND_LOGIN_RESPONSE2:
		msg := CreateSecondLoginResponse2()
		return msg.FromBytes(data, key), msg

	case cmd.MODIFY_USER_INFO_REQUEST:
		//		logs.Debug("messagefactory-->create modifyUserInfoRequest")
		msg := CreateModifyUserInfoRequest()
		return msg.FromBytes(data, key), msg

	case cmd.MODIFY_USER_INFO_RESPONSE:
		//		logs.Debug("messagefactory-->create modifyUserInfoResponse")
		msg := CreateModifyUserInfoResponse()
		return msg.FromBytes(data, key), msg



		//获取用户信息的请求
	case cmd.GET_USER_INFO_REQUEST:
		//		logs.Debug("messagefactory-->create getUserInfoRequest")
		msg := CreateGetUserInfoRequest()
		return msg.FromBytes(data, key), msg

		//获取用户信息的响应
	case cmd.GET_USER_INFO_RESPONSE:
		//		logs.Debug("messagefactory-->create getUserInfoResponse")
		msg := CreateGetUserInfoResponse()
		return msg.FromBytes(data, key), msg



	case cmd.USER_LOGIN_WITH_OTHER_DEVICE_REQUEST:
		msg := CreateUserLoginWithOtherDeviceRequest()
		return msg.FromBytes(data, key), msg

	case cmd.USER_LOGIN_WITH_OTHER_DEVICE_RESPONSE:
		msg := CreateUserLoginWithOtherDeviceResponse()
		return msg.FromBytes(data, key), msg

	case cmd.USER_OFFLINE_REQUEST:
		msg := CreateUserOfflineRequest()
		return msg.FromBytes(data, key), msg

		//删除用户
	case cmd.DELETE_USER_REQUEST:
		msg := CreateDeleteUserRequest()
		return msg.FromBytes(data, key), msg

	case cmd.LOAD_USER_OFFLINE_MESSAGE_REQUEST:
		msg := CreateLoadUserOfflineRequest()
		return msg.FromBytes(data, key),msg

	case cmd.SEND_OWN_GROUP_AND_FRIEND_RESPONSE:
		msg := CreateSendOwnGroupAndFriendResponse()
		return msg.FromBytes(data, key),msg

	case cmd.SEND_OWN_GROUP_AND_FRIEND_REQUEST:
		msg := CreateSendOwnGroupAndFriendRequest()
		return msg.FromBytes(data, key),msg

	default:
		return nil, nil
	}
}
