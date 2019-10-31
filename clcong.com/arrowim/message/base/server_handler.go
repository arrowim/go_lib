package base


type ServerMessageHeader struct {
	////////////////////////服务器端使用
	UserId    int32 `json:"userId"`//请求者
	GatewayId int16 //请求者网关ID

	AppId          int16
	SubAppId       int16
	DeviceType     DeviceType
	MessagePriorty MESSAGE_PRIORTY
	LoginId        int64
}


func (self *ServerMessageHeader) SetHeader(h *ServerMessageHeader) {
	if h == nil {
		return
	}
	self.UserId = h.UserId
	self.GatewayId = h.GatewayId

	self.AppId = h.AppId
	self.SubAppId = h.SubAppId
	self.DeviceType = h.DeviceType
	self.LoginId = h.LoginId
	self.MessagePriorty = h.MessagePriorty
}


func (self *ServerMessageHeader) SetData (userId int32, appId, subAppId, gatewayId int16,
	deviceType DeviceType, loginId int64){
		self.UserId = userId
		self.AppId = appId
		self.SubAppId = subAppId
		self.GatewayId = gatewayId
		self.DeviceType = deviceType
		self.LoginId = loginId
}