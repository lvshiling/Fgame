package job

import (
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
)

type RunnerState int

const (
	RunnerStateInit = iota
	RunnerStateRunning
	RunnerStateStop
)

type IJobRunner interface {
	Start() error
	Stop()
	GetState() RunnerState
}

type jobRunner struct {
	job   IJob
	done  chan struct{}
	state RunnerState
}

func (m *jobRunner) Start() error {
	go func() {
		m.tick()
		m.state = RunnerStateRunning
	}()
	log.WithFields(log.Fields{
		"Id": m.job.GetId(),
	}).Debug("作业开始")
	return nil
}

func (m *jobRunner) Stop() {
	m.state = RunnerStateStop
	m.done <- struct{}{}
}

func (m *jobRunner) GetState() RunnerState {
	return m.state
}

func (m *jobRunner) tick() {
loop:
	for {
		select {
		case <-time.After(time.Second * time.Duration(m.job.GetTickSecond())):
			{
				err := m.runJob()
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
						"jobId": m.job.GetId(),
					}).Error("作业运行异常")
				}
			}
		case <-m.done:
			{
				break loop
			}
		}
	}
}

func (m *jobRunner) runJob() error {
	defer func() {
		if rerr := recover(); rerr != nil {
			stackBuffer := make([]byte, 10240)
			tempPos := runtime.Stack(stackBuffer, true)
			log.WithFields(log.Fields{
				"error": rerr,
				"stack": string(stackBuffer[:tempPos]),
				"jobId": m.job.GetId(),
			}).Error("运行作业期间出现异常，已停止")
			m.Stop()
		}
	}()
	return m.job.Run()
}

func NewJobRunner(job IJob) IJobRunner {
	rst := &jobRunner{}
	rst.job = job
	rst.done = make(chan struct{})
	return rst
}
