package timecheckutil

import (
	list2 "arrowim/github.com/arrowim/utils/list"
	sync2 "arrowim/github.com/arrowim/utils/sync"
	timeutils2 "arrowim/github.com/arrowim/utils/timeutils"
	"time"
)

type TimeChecker interface {
	OnTimeOut()
}

type TimeCheckNode struct {
	DelineTime int32
	Checker    map[TimeChecker]bool
}

func (self *TimeCheckNode) Compare(comparator list2.Comparator) int {
	if v, ok := comparator.(*TimeCheckNode); ok {
		if v.DelineTime > self.DelineTime {
			return -1
		} else if v.DelineTime < self.DelineTime {
			return 1
		} else {
			return 0
		}
	}

	return -1
}

func CreateTimeCheckNode(delineTime int32) *TimeCheckNode {
	s := &TimeCheckNode{}
	s.DelineTime = delineTime
	s.Checker = make(map[TimeChecker]bool)
	return s
}

type TimeCheck struct {
	data list2.SortList
	lock sync2.Mutex
}

func createTimeCheck() *TimeCheck {
	v := TimeCheck{}
	return &v
}

var inst *TimeCheck

func GetTimeChecker() *TimeCheck {
	if inst == nil {
		inst = createTimeCheck()
		go inst.start()
	}

	return inst
}

func (self *TimeCheck) GetTimeoutNode() []TimeChecker {
	r := []TimeChecker{}

	self.lock.Lock()
	defer self.lock.Unlock()

	s := self.data.Front()

	if v, ok := s.(*TimeCheckNode); ok {
		if v.DelineTime <= timeutils2.GetTimeStamp() {
			for k, _ := range v.Checker {
				r = append(r, k)
			}

			self.data.PopBack()
		}
	}

	return r
}

func (self *TimeCheck) work() {
	r := self.GetTimeoutNode()

	for _, v := range r {
		v.OnTimeOut()
	}
}

func (self *TimeCheck) start() {

	for {

		self.work()

		time.Sleep(1 * time.Second)
	}
}

//func (self *TimeCheck) RemoveTimeChecker(id int64) {
//	self.lock.Lock()
//	self.lock.Unlock()

//	if t, ok := self.dataWitId[id]; ok {
//		if t2, ok2 := self.datas[t]; ok2 {
//			delete(t2.Nodes, id)
//		}
//	}

//	delete(self.dataWitId, id)
//}

func (self *TimeCheck) PutTimeChecker(delineTime int32, d TimeChecker) {
	s := CreateTimeCheckNode(delineTime)

	r := self.data.Find(s)

	if r != nil {
		if v, ok := r.(*TimeCheckNode); ok {
			v.Checker[d] = true
		}
	} else {
		s.Checker[d] = true
		self.data.Add(s)
	}
}
