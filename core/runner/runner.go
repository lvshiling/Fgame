package runner

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type Task interface {
	Heartbeat()
}

type Runner interface {
	Start() error
	Stop()
	Hearbeat()
	AddTask(t Task)
}

type runner struct {
	refreshTime    time.Duration
	heartbeatTimer *time.Timer
	stopChan       chan struct{}
	tasks          []Task
}

func (r *runner) Start() error {
	go func() {
	Loop:
		for {
			select {
			case <-r.heartbeatTimer.C:
				{
					r.Hearbeat()
					r.heartbeatTimer.Reset(r.refreshTime)
				}
			case <-r.stopChan:
				break Loop
			}
		}
		log.Infoln("定时器结束")
	}()
	log.Infoln("定时器启动")
	return nil
}

func (r *runner) Stop() {
	r.heartbeatTimer.Stop()
	r.stopChan <- struct{}{}
}

func (r *runner) AddTask(t Task) {
	r.tasks = append(r.tasks, t)
}

func (r *runner) Hearbeat() {
	for _, task := range r.tasks {
		task.Heartbeat()
	}
}

func NewRunner(refreshTime time.Duration) Runner {
	r := &runner{}
	r.stopChan = make(chan struct{})
	r.heartbeatTimer = time.NewTimer(refreshTime)
	r.refreshTime = refreshTime
	return r
}
