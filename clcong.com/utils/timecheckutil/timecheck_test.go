package timecheckutil

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

type TimeCheckObject struct {
	ss int64
}

var t1 int = 1

var lock sync.Mutex

func (self TimeCheckObject) OnTimeOut() {

	//	if self.ss%500 == 1 {
	fmt.Println("on time out", self.ss)
	//	}

}

func (self TimeCheckObject) GetId() int64 {
	return self.ss

}

func Test_timeCheck(t *testing.T) {
	fmt.Println("add checker")
	checker := GetTimeChecker()

	fmt.Println("add checker2")
	for i := 1; i < 10; i++ {
		s := TimeCheckObject{int64(i)}

		//		if i%500 == 1 {
		fmt.Println("add checker " + strconv.Itoa(i))
		//		}

		checker.PutAfterTimeChecker(1, s.ss, s)

		//		time.Sleep(1 * time.Second)
	}

	checker.RemoveTimeChecker(1)
	checker.RemoveTimeChecker(3)
	checker.RemoveTimeChecker(6)

	time.Sleep(100 * time.Second)
}
