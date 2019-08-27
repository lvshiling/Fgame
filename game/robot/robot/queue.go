package robot

import (
	"fgame/fgame/common/message"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	"time"
)

//消息队列
type MessageQueue struct {
	msgs    chan message.Message
	t       *time.Timer
	maxTime time.Duration
	pl      scene.RobotPlayer
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
			err := global.GetGame().GetMessageHandler().HandleMessage(m)
			if err != nil {
				mq.pl.Close(nil)
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

func NewMessageQueue(pl scene.RobotPlayer, capacity int32, maxTime time.Duration) *MessageQueue {
	ms := &MessageQueue{
		msgs:    make(chan message.Message, capacity),
		maxTime: maxTime,
		pl:      pl,
	}

	ms.t = time.NewTimer(maxTime)
	ms.t.Stop()
	return ms
}
