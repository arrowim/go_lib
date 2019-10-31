package job

import (
	"sync/atomic"
)

type Runnable interface {
	Run()
}

type Job struct {
	status int32
}

func CreateJob() *Job {
	return &Job{}
}

func (self *Job) Stop() {
	atomic.StoreInt32(&self.status, 1)
}

func (self *Job) RunOnce(job Runnable) {
	go func() {
		job.Run()
	}()
}

func (self *Job) Start(jobNum int, job Runnable) {
	for i := 0; i < jobNum; i++ {
		go func() {
			for {
				if atomic.LoadInt32(&self.status) == 1 {
					return
				}
				job.Run()
			}
		}()
	}
}
