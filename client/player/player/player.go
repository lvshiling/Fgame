package player

import (
	commonpbutil "fgame/fgame/client/common/pbutil"
	"fgame/fgame/client/scene/pbutil"
	"fgame/fgame/client/session"
	"fgame/fgame/core/fsm"
	"fmt"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

//玩家
type Player struct {
	//用户id
	id            int64
	userName      string
	state         fsm.State
	session       session.Session
	managers      map[PlayerDataKey]PlayerDataManager
	s             Strategy
	wg            sync.WaitGroup
	pingTimer     *time.Timer
	lastPingTime  int64
	hearbeatTimer *time.Ticker
}

func (p *Player) Id() int64 {
	return p.id
}

func (p *Player) UserName() string {
	return p.userName
}

func (p *Player) CurrentState() fsm.State {
	return p.state
}

func (p *Player) OnEnter(state fsm.State) {
	p.state = state
}

func (p *Player) OnExit(state fsm.State) {

}

func (p *Player) GetStrategy() Strategy {
	return p.s
}

func (p *Player) GetManager(pdk PlayerDataKey) PlayerDataManager {
	pdm, ok := p.managers[pdk]
	if !ok {
		return nil
	}
	return pdm
}

func (p *Player) SendMessage(msg proto.Message) {
	p.session.Send(msg)
}

//初始化
func (p *Player) init() {
	p.state = PlayerStateInit
	for k, factory := range playerDataFactoryMap {
		p.managers[k] = factory.Create(p)
	}
}

func (p *Player) Auth() bool {
	flag := stateMachine.Trigger(p, EventPlayerAuth)
	if !flag {
		return false
	}
	return true
}

func (p *Player) SelectRole() bool {
	flag := stateMachine.Trigger(p, EventPlayerSelectJob)
	if !flag {
		return false
	}
	return true
}

func (p *Player) StartStrategy(s Strategy) {
	p.StartHeartbeat()
	p.StartPing()
	go func() {
		s.Run()
	}()
}

func (p *Player) StartHeartbeat() {
	ti := time.Second * 30
	p.hearbeatTimer = time.NewTicker(ti)
	go func() {
		for {
			select {
			case <-p.hearbeatTimer.C:
				//发送事件
				csHeartBeat := commonpbutil.BuildCSHeartBeat()
				p.SendMessage(csHeartBeat)
			}
		}
	}()
}

func (p *Player) StartPing() {
	ti := time.Second * 2
	p.pingTimer = time.NewTimer(ti)
	go func() {
		for {
			select {
			case <-p.pingTimer.C:
				//发送ping
				now := time.Now().UnixNano() / int64(time.Millisecond)
				p.lastPingTime = now
				//发送事件
				csPing := pbutil.BuildPing()
				p.SendMessage(csPing)
			}
		}
	}()
}

const (
	maxElapse = 100
)

func (p *Player) Ping() {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	elapse := now - p.lastPingTime
	if elapse >= maxElapse {
		log.WithFields(
			log.Fields{
				"playerId": p.GetPlayerId(),
				"ping":     elapse,
			}).Warn("player:ping超时")
	} else {
		log.WithFields(
			log.Fields{
				"playerId": p.GetPlayerId(),
				"ping":     elapse,
			}).Info("player:ping")
	}
	ti := time.Second * 2
	p.pingTimer.Reset(ti)
}

func (p *Player) Game() bool {
	flag := stateMachine.Trigger(p, EventPlayerGame)
	if !flag {
		return false
	}
	log.WithFields(
		log.Fields{
			"playerId": p.id,
		}).Info("player:玩家进入场景")

	return true
}

func (p *Player) OnError(code int32) {
	p.s.OnError(code)
}

func (p *Player) ping() {

}

func (p *Player) Stop() error {
	return nil
}

func (p *Player) Close() {
	p.session.Close()
}

func (p *Player) getPlayerBasicManager() *PlayerBasicManager {
	m, ok := p.managers[PlayerDataKeyBasic]
	if !ok {
		return nil
	}
	tm, ok := m.(*PlayerBasicManager)
	if !ok {
		return nil
	}
	return tm
}

func (p *Player) GetPlayerId() int64 {
	m := p.getPlayerBasicManager()
	if m == nil {
		return int64(0)
	}
	return m.GetPlayerId()
}

func NewPlayer(id int64, session session.Session) *Player {
	p := &Player{
		id:       id,
		userName: fmt.Sprintf("guest_%d", id),
		session:  session,
		managers: make(map[PlayerDataKey]PlayerDataManager),
	}
	p.init()
	return p
}
