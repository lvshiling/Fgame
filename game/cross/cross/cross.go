package cross

import (
	"fgame/fgame/common/message"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/runtimeutils"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

//TODO 限制场景人数
//跨服线程接口
type Cross interface {
	AddPlayer(player.Player)
	RemovePlayer(player.Player)
	Start()
	Stop()
	Post(msg message.Message)
}

//跨服结构
type cross struct {
	m      sync.Mutex
	stoped bool
	//玩家管理
	players *sync.Map
	// 定时器
	heartbeatTimer *time.Timer
	//消息
	msgs chan message.Message
}

func (c *cross) GetAllPlayers() *sync.Map {
	return c.players
}

func (c *cross) AddPlayer(p player.Player) {
	// c.m.Lock()
	// defer c.m.Unlock()
	log.WithFields(
		log.Fields{
			"id":   p.GetId(),
			"goId": runtimeutils.Goid(),
		}).Infoln("cross:玩家添加")
	c.players.Store(p.GetId(), p)
	// c.players[p.GetId()] = p
}

func (c *cross) RemovePlayer(p player.Player) {
	// c.m.Lock()
	// defer c.m.Unlock()
	log.WithFields(
		log.Fields{
			"id":   p.GetId(),
			"goId": runtimeutils.Goid(),
		}).Infoln("cross:玩家移除")
	c.players.Delete(p.GetId())
	// delete(c.players, p.GetId())
}

func (c *cross) Post(msg message.Message) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.stoped {
		log.WithFields(
			log.Fields{}).Infoln("cross:跨服协程已经停止")
		return
	}

	log.WithFields(
		log.Fields{
			"length": len(c.msgs),
		}).Infoln("cross:跨服协程post")
	c.msgs <- msg
}

const (
	heartbeatTime = 40 * time.Millisecond
	queueCapacity = 20000
)

//goroutine不安全 都会在scene runner 调用
//开始
func (c *cross) Start() {
	log.Info("cross:开启")
	c.heartbeatTimer.Reset(heartbeatTime)
Loop:
	for {
		select {
		case <-c.heartbeatTimer.C:
			c.tick()
			c.heartbeatTimer.Reset(heartbeatTime)
			break
		case m, ok := <-c.msgs:
			{
				if !ok {
					break Loop
				}
				err := global.GetGame().GetMessageHandler().HandleMessage(m)
				if err != nil {
					log.WithFields(
						log.Fields{
							"error": err,
						}).Error("cross:跨服处理消息,错误")
				}
			}
		}
	}
}

//消息处理
func (c *cross) tick() {
	// for _, p := range c.players {
	// 	p.Tick()
	// }
	c.players.Range(func(key interface{}, val interface{}) bool {
		p, ok := val.(player.Player)
		if !ok {
			return true
		}
		p.Tick()
		return true
	})
	c.heartbeat()

}

//数据更新
func (c *cross) heartbeat() error {
	c.players.Range(func(key interface{}, val interface{}) bool {
		p, ok := val.(player.Player)
		if !ok {
			return true
		}
		p.Heartbeat()
		return true
	})

	return nil
}

//goroutine不安全 都会在scene runner 调用
func (c *cross) Stop() {
	c.m.Lock()
	defer c.m.Unlock()
	if c.stoped {
		return
	}
	c.stoped = true

	c.beforeStop()

	//主动关闭
	close(c.msgs)
	log.Info("cross:关闭")
	return
}

//关闭前清理
func (c *cross) beforeStop() {
	c.clearPlayers()
}

//清理玩家
func (c *cross) clearPlayers() {
	// for _, p := range c.players {
	// }
}

func newCross() Cross {
	c := &cross{}
	c.players = &sync.Map{}
	c.msgs = make(chan message.Message, queueCapacity)
	c.heartbeatTimer = time.NewTimer(heartbeatTime)
	c.heartbeatTimer.Stop()
	return c
}
