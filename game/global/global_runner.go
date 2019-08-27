package global

import (
	"fgame/fgame/common/message"
	gamesession "fgame/fgame/game/session"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	defaultCapacity = 10000
	heartbeatTime   = time.Second
)

//全局业务
//TODO 加锁
type GlobalRunner struct {
	msgs           chan message.Message
	heartbeatTimer *time.Timer
	mh             message.Handler
	stopChan       chan struct{}
	globalUpdater  *GlobalUpdater
}

func (gp *GlobalRunner) Heartbeat() (err error) {
	log.WithFields(
		log.Fields{}).Debug("全局业务心跳")
	gp.globalUpdater.Update()
	return
}

func (gp *GlobalRunner) Start() {
	go func() {
	Loop:
		for {
			select {
			case <-gp.heartbeatTimer.C:
				{
					err := gp.Heartbeat()
					if err != nil {
						log.WithFields(
							log.Fields{
								"error": err,
							}).Error("全局业务心跳,错误")
					}
					gp.heartbeatTimer.Reset(heartbeatTime)
				}
			case msg, ok := <-gp.msgs:
				{
					if !ok {
						break Loop
					}

					//TODO 运行消息
					err := gp.mh.HandleMessage(msg)
					if err != nil {
						log.WithFields(
							log.Fields{
								"error": err,
							}).Error("全局业务消息处理,错误")
						switch tmsg := msg.(type) {
						case message.SessionMessage:
							tmsg.Session().Close()
							break
						case message.ScheduleMessage:
							{
								s := gamesession.SessionInContext(tmsg.Context())
								if s != nil {
									s.Close(true)
								}
							}
							break
						}
					}
				}
			}
		}
		log.Infoln("全局业务结束")
		gp.stopChan <- struct{}{}
	}()
	log.Infoln("全局业务启动")

}

func (gp *GlobalRunner) Post(msg message.Message) {
	if !GetGame().Open() {
		return
	}
	log.WithFields(
		log.Fields{
			"length": len(gp.msgs),
		}).Infoln("global:全局post")
	gp.msgs <- msg
}

func (gp *GlobalRunner) GetUpdater() *GlobalUpdater {
	return gp.globalUpdater
}

func (gp *GlobalRunner) Stop() {

	gp.heartbeatTimer.Stop()
	close(gp.msgs)
	<-gp.stopChan
}

func NewGlobalRunner(mh message.Handler) *GlobalRunner {
	gr := &GlobalRunner{}
	gr.mh = mh
	gr.msgs = make(chan message.Message, defaultCapacity)
	gr.stopChan = make(chan struct{})
	gr.heartbeatTimer = time.NewTimer(heartbeatTime)
	gr.globalUpdater = NewGlobalUpdater()
	return gr
}
