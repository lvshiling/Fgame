package heartbeat

import (
	"time"
)

//心跳任务
type HeartbeatTask interface {
	Run()
	//间隔时间
	ElapseTime() time.Duration
}

//定时心跳任务
type heartbeatTimerTask struct {
	*timer
	HeartbeatTask
}

func (h *heartbeatTimerTask) Run() {
	if h.timer.IsTimeUp() {
		h.HeartbeatTask.Run()
		h.timer.Reset(h.HeartbeatTask.ElapseTime())
	}
}

//TODO heartbeatTimerTask复用
func createHeartbeatTimerTask(task HeartbeatTask) *heartbeatTimerTask {
	hbtt := &heartbeatTimerTask{}
	hbtt.timer = createTimer(task.ElapseTime())
	hbtt.HeartbeatTask = task
	return hbtt
}

//心跳
type HeartbeatTaskRunner interface {
	//心跳
	Heartbeat()
	//添加任务
	AddTask(task HeartbeatTask)
	//清楚所有
	Clear()
}

type heartbeatTaskRunner struct {
	tasks []*heartbeatTimerTask
}

func (htr *heartbeatTaskRunner) Heartbeat() {
	for _, task := range htr.tasks {
		task.Run()
	}
}

func (htr *heartbeatTaskRunner) AddTask(task HeartbeatTask) {
	hbtt := createHeartbeatTimerTask(task)
	htr.tasks = append(htr.tasks, hbtt)
}

func (htr *heartbeatTaskRunner) Clear() {
	htr.tasks = nil
}

func NewHeartbeatTaskRunner() HeartbeatTaskRunner {
	htr := &heartbeatTaskRunner{}
	return htr
}
