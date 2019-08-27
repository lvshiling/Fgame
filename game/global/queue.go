package global

import (
	"fgame/fgame/common/message"
	"time"

	log "github.com/Sirupsen/logrus"
)

//消息队列
type MessageQueue struct {
	t       *time.Timer
	maxTime time.Duration
	msgs    chan message.Message
}

func (mq *MessageQueue) Tick() {
	mq.t.Reset(mq.maxTime)
Loop:
	for {
		select {
		case <-mq.t.C:
			break Loop
		case m, ok := <-mq.msgs:

			if !ok {
				break Loop
			}

			err := GetGame().GetMessageHandler().HandleMessage(m)
			if err != nil {
				log.WithFields(
					log.Fields{
						"error": err,
					}).Error("message:处理消息,错误")
				// mq.pl.Session().Close(true)
				break Loop
			}
			break
		default:
			break Loop
		}

	}
	mq.t.Stop()
}

func (ms *MessageQueue) Post(msg message.Message) {
	ms.msgs <- msg
}

func NewMessageQueue(capacity int32, maxTime time.Duration) *MessageQueue {
	ms := &MessageQueue{
		msgs:    make(chan message.Message, capacity),
		maxTime: maxTime,
	}
	ms.t = time.NewTimer(maxTime)
	ms.t.Stop()
	return ms
}
