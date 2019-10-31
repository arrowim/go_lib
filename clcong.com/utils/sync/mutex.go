package sync

//	"runtime"

import (
	"sync"
)

type Mutex struct {
	lock sync.Mutex
	id   int32
}

func (self *Mutex) Lock() {
	//	f, file, line, _ := runtime.Caller(1)
	//	Logs.Debug("lock ", f, file, line)
	self.lock.Lock()
	//	Logs.Debug("lock ok", f, file, line)
}

func (self *Mutex) Unlock() {
	//	f, file, line, _ := runtime.Caller(1)
	//	Logs.Debug("un lock ", f, file, line)
	self.lock.Unlock()
}
