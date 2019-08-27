package runner

import (
	"runtime/debug"
	"time"

	log "github.com/Sirupsen/logrus"
)

type GoRunner interface {
	Start()
	Stop()
}

type goRunner struct {
	name           string
	refreshTime    time.Duration
	heartbeatTimer *time.Timer
	stopChan       chan struct{}
	action         func()
}

func (r *goRunner) Start() {
	go func() {
	Loop:
		for {
			select {
			case <-r.heartbeatTimer.C:
				{
					r.heartbeat()
					r.heartbeatTimer.Reset(r.refreshTime)
				}
			case <-r.stopChan:
				break Loop
			}
		}
		log.WithFields(
			log.Fields{
				"name": r.name,
			}).Info("定时器,结束")
	}()
	log.WithFields(
		log.Fields{
			"name":       r.name,
			"elapseTime": float64(r.refreshTime) / float64(time.Second),
		}).Info("定时器,开始")
	return
}

func (r *goRunner) heartbeat() {
	defer func() {
		//恢复
		if rerr := recover(); rerr != nil {
			debug.PrintStack()
			log.WithFields(
				log.Fields{
					"name":  r.name,
					"error": rerr,
					"stack": string(debug.Stack()),
				}).Error("定时器,异常")
		}

	}()
	r.action()
}

func (r *goRunner) Stop() {
	r.heartbeatTimer.Stop()
	r.stopChan <- struct{}{}
}

func NewGoRunner(name string, action func(), refreshTime time.Duration) GoRunner {
	r := &goRunner{}
	r.name = name
	r.stopChan = make(chan struct{})
	r.heartbeatTimer = time.NewTimer(refreshTime)
	r.refreshTime = refreshTime
	r.action = action
	return r
}
