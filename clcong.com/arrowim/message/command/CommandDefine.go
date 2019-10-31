package command

/***
各种命令编号 全在这里
*/

type Command int16

func (self Command) GetValue() int16 {
	return int16(self)
}

const (
	//////////////////用户消息
	//SECOND_LOGIN_REQUEST                  Command = 0x0101 //第二次登录请求
	//SECOND_LOGIN_RESPONSE                 Command = 0x0102 //第二次登录请求响应
	GET_USER_INFO_REQUEST                 Command = 0x0103 //得到用户信息请求
	GET_USER_INFO_RESPONSE                Command = 0x0104 //得到用户信息响应
	LOG_OUT_REQUEST                       Command = 0x0105 //登出请求
	LOG_OUT_RESPONSE                      Command = 0x0106 //登出响应
	MODIFY_USER_INFO_REQUEST              Command = 0x0107 //	修改用户信息的通知
	MODIFY_USER_INFO_RESPONSE             Command = 0x0108 //修改用户信息的响应
	MODIFY_PASSWORD_REQUEST               Command = 0x0111 //更改密码请求
	MODIFY_PASSWORD_RESPONSE              Command = 0x0112 //更改密码响应
	DELETE_USER_REQUEST                   Command = 0x0113 //删除用户
	DELETE_USER_RESPONSE                  Command = 0x0114 //删除用户响应
	USER_LOGIN_WITH_OTHER_DEVICE_REQUEST  Command = 0x0115 //用户从其他设备登录通知
	USER_LOGIN_WITH_OTHER_DEVICE_RESPONSE Command = 0x0116 //用户从其他设备登录通知响应
	MODIFY_PASSWORD_NOTIFY_REQUEST        Command = 0x0117 //修改密码通知
	MODIFY_PASSWORD_NOTIFY_RESPONSE       Command = 0x0118 //修改密码通知响应
	SECOND_LOGIN_REQUEST2                 Command = 0x0119 //第二次登录请求
	SECOND_LOGIN_RESPONSE2                Command = 0x0120 //第二次登录请求响应
	USER_OFFLINE_REQUEST                  Command = 0x0121 //用户下线请求
	USER_OFFLINE_RESPONSE                 Command = 0x0122 //用户下线响应
	LOAD_USER_OFFLINE_MESSAGE_REQUEST     Command = 0x0123
	SEND_OWN_GROUP_AND_FRIEND_REQUEST     Command = 0x0151 //服务端发送用户的好友与群到客户端
	SEND_OWN_GROUP_AND_FRIEND_RESPONSE    Command = 0x0152 //登录成功响应，

	////////////////////好友
	ADD_FRIENDS_REQUEST                                    Command = 0x0201 //添加好友请求
	ADD_FRIENDS_RESPONSE                                   Command = 0x0202 //添加好友响应
	DELETE_FRIENDS_REQUEST                                 Command = 0x0203 //删除好友
	DELETE_FRIENDS_RESPONSE                                Command = 0x0204 //删除好友响应
	RESULT_OF_ADD_FRIEND_REQUEST                           Command = 0x0205 //添加好友,被邀请者返回结果请求
	RESULT_OF_ADD_FRIEND_RESPONSE                          Command = 0x0206 //添加好友,被邀请者返回结果响应
	GET_FRIEND_IDS_REQUEST                                 Command = 0x0207 //获取好友列表
	GET_FRIEND_IDS_RESPONSE                                Command = 0x0208 //获取好友列表响应
	SEND_FRIEND_IDS_TO_CLIENT_REQUEST                      Command = 0x0209 //发送好友列表到客户端的通知
	SEND_FRIEND_IDS_TO_CLIENT_RESPONSE                     Command = 0X0210 //发送好友列表到客户端的通知的响应
	GET_FRIEND_DETAILS_REQUEST                             Command = 0x0211 //获取好友列表(新增好友备注名)
	GET_FRIEND_DETAILS_RESPONSE                            Command = 0x0212 //获取好友列表响应(新增好友备注名)
	SEND_FRIEND_DETAILS_TO_CLIENT_REQUEST                  Command = 0x0213 //发送好友列表到客户端的通知(新增好友备注名)
	SEND_FRIEND_DETAILS_TO_CLIENT_RESPONSE                 Command = 0X0214 //发送好友列表到客户端的通知的响应(新增好友备注名)
	SEND_MODIFY_FRIEND_REMARKNAME_TO_SELF_CLIENTS_REQUEST  Command = 0X0215 //发送修改好友备注名到自己的设备的通知
	SEND_MODIFY_FRIEND_REMARKNAME_TO_SELF_CLIENTS_RESPONSE Command = 0X0216 //发送修改好友备注名到自己的设备的响应
	SEND_SHIELD_FRIEND_MESSAGE_TO_SELF_CLIENTS_REQUEST     Command = 0X0217 //发送屏蔽或者取消屏蔽好友信息的通知给自己的其他设备的通知
	SEND_SHIELD_FRIEND_MESSAGE_TO_SELF_CLIENTS_RESPONSE    Command = 0X0218 //发送屏蔽或者取消屏蔽好友信息的通知给自己的其他设备的通知的响应

	///////////////////群
	CREATE_GROUP_REQUEST                                   Command = 0x0301 //建立群请求命令
	CREATE_GROUP_RESPONSE                                  Command = 0x0302 //建立群请求响应
	DELETE_GROUP_REQUEST                                   Command = 0x0303 //删除群请求命令
	DELETE_GROUP_RESPONSE                                  Command = 0x0304 //删除群请求响应
	ADD_GROUP_MEMBER_REQUEST                               Command = 0x0305 //添加群成员命令
	ADD_GROUP_MEMBER_RESPONSE                              Command = 0x0306 //添加群成员响应
	DELETE_GROUP_MEMBER_REQUEST                            Command = 0x0307 //删除群成员命令
	DELETE_GROUP_MEMBER_RESPONSE                           Command = 0x0308 //删除群成员响应
	GET_GROUP_INFO_REQUEST                                 Command = 0x0309 //得到群信息请求
	GET_GROUP_INFO_RESPOSNE                                Command = 0x0310 //得到群信息响应
	ALLOW_OR_DISALLOW_INVITED_TO_GROUP_REQUEST             Command = 0x0311 //邀请群成员,被邀请者返回结果请求
	ALLOW_OR_DISALLOW_INVITED_TO_GROUP_RESPONSE            Command = 0x0312 //邀请群成员,被邀请者返回结果响应
	ALLOW_OR_DISALLOW_USER_ADDED_TO_GROUP_REQUEST          Command = 0x0313 //申请入群,管理员返回结果请求
	ALLOW_OR_DISALLOW_USER_ADDED_TO_GROUP_RESPONSE         Command = 0x0314 //申请入群,管理员返回结果响应
	MODIFY_GROUP_INFO_REQUEST                              Command = 0x0315 //修改群资料的请求
	MODIFY_GROUP_INFO_RESPONSE                             Command = 0x0316 //修改群资料的响应
	SEND_GROUP_INFO_TO_CLIENT_REQUEST                      Command = 0x0317 //发送群资料给客户端的请求
	SEND_GROUP_INFO_TO_CLIENT_RESPONSE                     Command = 0x0318 //发送群资料给客户端的响应
	GET_GROUP_IDS_REQUEST                                  Command = 0x0319 //获得群列表信息
	GET_GROUP_IDS_RESPONSE                                 Command = 0x0320 //获得群列表信息响应
	GET_GROUP_MEMBER_IDS_REQUEST                           Command = 0x0321 //获取群成员列表信息请求
	GET_GROUP_MEMBER_IDS_RESPONSE                          Command = 0x0322 //获取群成员列表响应
	SEND_GROUP_IDS_TO_CLIENT_REQUEST                       Command = 0X0323 //发送群ID列表给客户端的通知
	SEND_GROUP_IDS_TO_CLIENT_RESPONSE                      Command = 0X0324 //发送群ID列表给客户端的通知的响应
	SEND_GROUP_MEMBER_IDS_TO_CLIENT_REQUEST                Command = 0x0325 //发送群成员ID列表给客户端的通知
	SEND_GROUP_MEMBER_IDS_TO_CLIENT_RESPONSE               Command = 0x0326 //发送群成员ID列表给客户端的通知的响应
	GET_GROUP_MEMBER_DETAILS_REQUEST                       Command = 0x0327 //获取群成员列表信息请求（新增群成员备注）
	GET_GROUP_MEMBER_DETAILS_RESPONSE                      Command = 0x0328 //获取群成员列表响应（新增群成员备注）
	SEND_GROUP_MEMBER_DETAILS_TO_CLIENT_REQUEST            Command = 0x0329 //发送群成员资料列表给客户端的通知（新增群成员备注）
	SEND_GROUP_MEMBER_DETAILS_TO_CLIENT_RESPONSE           Command = 0x0330 //发送群成员资料列表给客户端的通知的响应（新增群成员备注）
	MODIFY_OWN_GROUP_LIST_REQUEST                          Command = 0x0331 //修改用户群列表的通知
	MODIFY_OWN_GROUP_LIST_RESPONSE                         Command = 0x0332 //修改用户群列表的响应
	SEND_MODIFY_GROUP_MEMBER_REMARKNAME_TO_CLIENT_REQUEST  Command = 0X0333 //发送修改群成员备注的信息到群成员客户端的通知
	SEND_MODIFY_GROUP_MEMBER_REMARKNAME_TO_CLIENT_RESPONSE Command = 0X0334 //发送修改群成员备注的信息到群成员客户端的响应
	SEND_GROUP_DETAILS_TO_CLIENT_REQUEST                   Command = 0X0335 //发送群列表（新增是否群，是否屏蔽等参数）到客户端的通知
	SEND_GROUP_DETAILS_TO_CLIENT_RESPONSE                  Command = 0X0336 //发送群列表（新增是否群，是否屏蔽等参数）到客户端通知的响应
	GET_GROUP_DETAILS_REQUEST                              Command = 0X0337 //获取群列表信息请求（新增是否群，是否屏蔽等参数）
	GET_GROUP_DETAILS_RESPONSE                             Command = 0X0338 //获取群列表信息响应（新增是否群，是否屏蔽等参数）
	SNED_SHIELD_GROUP_MESSAGE_TO_SELF_CLIENTS_REQUEST      Command = 0X0339 //发送屏蔽或者取消屏蔽群信息的通知给自己的其他设备的通知
	SNED_SHIELD_GROUP_MESSAGE_TO_SELF_CLIENTS_RESPONSE     Command = 0X0340 //发送屏蔽或者取消屏蔽群信息的通知给自己的其他设备的通知的响应
	GET_GROUP_BY_ID_REQUEST                                Command = 0x0341 //根据群ID获取群信息（新增是否群，是否屏蔽等参数）的请求
	GET_GROUP_BY_ID_RESPONSE                               Command = 0x0342 //根据群ID获取群信息（新增是否群，是否屏蔽等参数）的响应
	SEND_SINGLE_GROUP_DETAIL_TO_CLIENT_REQUEST             Command = 0x0343 //发送单个群资料到客户端的通知
	SEND_SINGLE_GROUP_DETAIL_TO_CLIENT_RESPONSE            Command = 0x0344 //发送单个群资料到客户端的通知的响应
	SEND_GROUP_PERMISSION_UPDATE_TO_CLIENT_REQUEST         Command = 0x0345 //发送群管理权限变更通知到客户端的请求
	SEND_GROUP_PERMISSION_UPDATE_TO_CLIENT_RESPONSE        Command = 0x0346 //发送群管理权限变更通知到客户端的响应
	SEND_GROUP_IS_NOT_EXISTS_TO_CLIENT_REQUEST             Command = 0x0347 //群不存在的时候，发给客户端的请求
	SEND_GROUP_IS_NOT_EXISTS_TO_CLIENT_RESPONSE            Command = 0x0348 //群不存在的时候，发给客户端的响应
	SEND_GROUP_CREATOR_IS_UPDATE_REQUEST                   Command = 0x0349 // 群主修改后，发送给客户端的请求
	SEND_GROUP_CREATOR_IS_UPDATE_RESPONSE                  Command = 0x0350 // 群主修改后，发送给客户端的返回
	ALLOW_OR_DISALLOW_INVITED_TO_GROUP_REQUEST2            Command = 0x0351 // 邀请群成员,被邀请者返回结果请求
	ALLOW_OR_DISALLOW_INVITED_TO_GROUP_RESPONSE2           Command = 0x0352 // 邀请群成员,被邀请者返回结果响应
	PUBLISH_PROCLAMATION_REQUEST                           Command = 0x0353 //群公告通知
	PUBLISH_PROCLAMATION_REPONSE                           Command = 0x0354 //群公告响应
	//ADD_GROUP_MEMBER_REQUEST2                              Command = 0x0353 //添加群成员命令
	//ADD_GROUP_MEMBER_RESPONSE2                             Command = 0x0354 //添加群成员响应

	//////////////消息
	SEND_MESSAGE_REQUEST                Command = 0x0401 //发送消息请求
	SEND_MESSAGE_RESPONSE               Command = 0x0402 //发送消息请求响应
	SEND_GROUP_MESSAGE_REQUEST          Command = 0x0403 //发送群消息请求
	SEND_GROUP_MESSAGE_RESPONSE         Command = 0x0404 //发送群消息响应
	SEND_CUSTOM_SERVER_MESSAGEREQUEST   Command = 0x405  //发送自定义消息到服务器
	SEND_CUSTOM_SERVER_MESSAGERESPONSE  Command = 0x406  //发送自定义消息到服务器的响应
	SEND_CUSTOM_CLIENT_MESSAGEREQUEST   Command = 0x407  //服务器发送到客户端的消息
	SEND_CUSTOM_CLIENT_MESSAGERESPONSE  Command = 0x408  //服务器发送到客户端消息的响应
	SEND_CUSTOM_CLIENT_MESSAGEREQUEST2  Command = 0x409  //服务器发送到客户端的消息
	SEND_CUSTOM_CLIENT_MESSAGERESPONSE2 Command = 0x410  //服务器发送到客户端消息的响应
	SINGLE_MESSAGE_IS_READ_REQUEST      Command = 0x0413 //单聊消息已读请求
	SINGLE_MESSAGE_IS_READ_RESPONSE     Command = 0x0414 //单聊消息已读响应
	GROUP_MESSAGE_IS_READ_REQUEST       Command = 0x0415 //群聊消息已读请求
	GROUP_MESSAGE_IS_READ_RESPONSE      Command = 0x0416 //群聊消息已读响应
	SEND_INPUT_STATE_REQUEST            Command = 0x0417 // 发送状态请求
	SEND_INPUT_STATE_RESPONSE           Command = 0x0418 // 发送状态响应
	SEND_GROUP_MESSAGE_REQUEST2         Command = 0x0419 // 新发送群消息请求
	SEND_GROUP_MESSAGE_RESPONSE2        Command = 0x0420 // 新发送群消息响应
	///////////////基础命令
	KEEP_ALIVE_REQUEST              Command = 0x4001 //心跳连接请求
	KEEP_ALIVE_RESPONSE             Command = 0x4002 //心跳连接响应
	MESSAGE_DELIVER_REPORT_REQUEST  Command = 0x4003 //消息状态报告给app
	MESSAGE_DELIVER_REPORT_RESPONSE Command = 0x4004 //消息状态报告响应

	ERROR_MESSAGE_REQUEST  Command = 0x4F01 //
	ERROR_MESSAGE_RESPONSE Command = 0x4F02 //

	///////////////自定义通知
	SEND_NOTIFYCATION_REQUEST   Command = 0x0601 //推送消息请求
	SEND_NOTIFYCATION_RESPONSE  Command = 0x0602 //推送消息请求响应
	SEND_NOTIFYCATION_REQUEST2  Command = 0x0603 //推送群消息请求
	SEND_NOTIFYCATION_RESPONSE2 Command = 0x0604 //推送群消息响应

	//////////////消息回执
	REPORT_REQUEST  Command = 0x0701 //发送消息回执到
	REPORT_RESPONSE Command = 0x0702 //发送消息回执响应

)
