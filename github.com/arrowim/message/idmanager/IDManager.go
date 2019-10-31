package idmanager

import (
	"arrowim/github.com/arrowim/utils/timeutils"
	"fmt"
	//	"fmt"
	"sync"
	"time"
)

type MessageIdManager struct {
	currentId      int64
	serverId       int64
	mutext         sync.Mutex
	currentSecond  int64
	serverType     int64
	lastTimeSecond int64
}

const (
	GATEWAY_SERVER  = 1
	IM_SERVER       = 2
	GROUP_SERVER    = 3
	CUSTOMER_SERVER = 4
)

const (
	startTime     = 1517329836 / 8
	secondLen     = 30
	serverTypeLen = 3
	serverIdLen   = 7
	idLen         = 13
)

var inst *MessageIdManager

func CreateMessageIdManager(serverType int32, serverNum int32) *MessageIdManager {
	v := MessageIdManager{}
	v.serverType = int64(serverType)
	v.serverId = int64(serverNum)
	inst = &v
	return &v
}

func (self *MessageIdManager) assembleId() int64 {
	data := self.lastTimeSecond - startTime;
	//fmt.Println("data1::", data)

	data = data << (53 - secondLen)

	//fmt.Println("data2::", data, secondLen)

	data = data + (self.serverType << (serverIdLen + idLen))
	//fmt.Println("data3::", data, serverIdLen+idLen, self.serverType<<(serverIdLen+idLen))

	data = data + (self.serverId << idLen)
	//fmt.Println("data4::", data, self.serverId<<idLen, self.serverId<<idLen)

	data += self.currentId;
	//fmt.Println("data5::", data)
	return data
}

func (self *MessageIdManager) CreateId() int64 {

	t := timeutils.GetUnix13NowTime() / 1000 / 8

	self.mutext.Lock()
	defer self.mutext.Unlock()

	if (t <= self.lastTimeSecond) {
		self.currentId ++

		if (self.currentId >= 8192) {
			time.Sleep(100 * time.Second)
		}

		return self.assembleId()
	}

	self.currentId = 1

	if (self.lastTimeSecond >= t) {
		self.lastTimeSecond ++
	} else {
		self.lastTimeSecond = t
	}

	return self.assembleId()
}

func CreateMessageId() int64 {
	return inst.CreateId()
}

func getSecond(id int64) int64 {
	second := id >> (53 - secondLen)
	return second;
}

func getServerType(id int64) int64 {
	second := (id >> (53 - secondLen)) << (53 - secondLen)

	return (id - second) >> (serverIdLen + idLen);
}

func getServerId(id int64) int64 {
	second := getSecond(id)
	serverType := getServerType(id)
	serverId := id - second<<(53-secondLen) - (serverType << (53 - secondLen - serverTypeLen))

	serverId = serverId >> (idLen)

	return serverId
}

func getId(id int64) int64 {
	return (id << (64 - idLen)) >> (64 - idLen)

}

func getIDDetal(id int64) {
	//second := id >> (53 - secondLen)
	//fmt.Println("aaaaaaaa:", second<<(53-secondLen))
	//serverType := (id - (second << (53 - secondLen))) >> (serverIdLen + idLen)
	//serverId := (id - (second << (53 - secondLen))) >> (serverIdLen + idLen)
	//sId := (id << (secondLen + serverTypeLen + serverIdLen)) >> (53 - idLen)

	fmt.Println(fmt.Sprintf("id:%d\nsecond:%d\nserverType:%d\nserverId:%d\nid:%d\n", id,
		getSecond(id), getServerType(id), getServerId(id), getId(id)))
}
