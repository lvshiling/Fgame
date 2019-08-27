package player

import (
	"fgame/fgame/common/message"
	"fgame/fgame/game/global"
	"time"

	log "github.com/Sirupsen/logrus"
)

//消息队列
type MessageQueue struct {
	msgs     chan message.Message
	t        *time.Timer
	maxTime  time.Duration
	pl       *Player
	pause    bool
	pauseSig chan struct{}
}

func (mq *MessageQueue) Tick() {
	mq.t.Reset(mq.maxTime)
	// mq.pauseSig = make(chan struct{}, 1)
	mq.pause = false
Loop:
	for !mq.pause {
		select {
		case <-mq.t.C:
			break Loop
		case m, ok := <-mq.msgs:
			if !ok {
				break Loop
			}
			err := global.GetGame().GetMessageHandler().HandleMessage(m)
			if err != nil {
				log.WithFields(
					log.Fields{
						"playerId": mq.pl.GetId(),
						"error":    err,
					}).Error("player:玩家处理消息,错误")
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

func (mq *MessageQueue) Pause() {

	log.WithFields(
		log.Fields{
			"playerId": mq.pl.GetId(),
		}).Info("player:玩家处理消息暂停前")
	mq.pause = true
	log.WithFields(
		log.Fields{
			"playerId": mq.pl.GetId(),
		}).Info("player:玩家处理消息暂停后")
}

func NewMessageQueue(pl *Player, capacity int32, maxTime time.Duration) *MessageQueue {
	ms := &MessageQueue{
		msgs:     make(chan message.Message, capacity),
		pauseSig: make(chan struct{}, 1),
		maxTime:  maxTime,
		pl:       pl,
	}
	ms.t = time.NewTimer(maxTime)
	ms.t.Stop()
	return ms
}
