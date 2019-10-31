package idcreator

import (
	dbm "arrowim/src/clcong/arrowim/db/dbmanager"
	"arrowim/src/clcong/arrowim/message"
	l4g "arrowim/clcong.com/utilslog4go"
	"code.google.com/p/log4go"
)

var logs *log4go.Logger = l4g.GetLogger()

var uid int32

func getUid() int32 {
	if uid == 0 {
		//logs.Debug("数据库中获取")
		um := new(dbm.UserManager)
		maxId, err := um.GetMaxUserId()
		if err != nil {
			logs.Info("GetMaxUserId:", err)
		}
		uid = maxId + 1
	}
	return uid
}

func Int32IdCreate(step int32) int32 {
	id := message.CreateIntID(getUid(), step)
	uid++
	return id.Id()
}
