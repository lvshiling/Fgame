package arena

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/arena/arena"
	"fgame/fgame/game/arena/dao"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/global"
	"time"

	_ "fgame/fgame/game/arena/cross_handler"
	_ "fgame/fgame/game/arena/cross_loginflow"
	_ "fgame/fgame/game/arena/drop_handler"
	_ "fgame/fgame/game/arena/event/listener"
	_ "fgame/fgame/game/arena/handler"
	_ "fgame/fgame/game/arena/player"
)

type arenaModule struct {
	r runner.GoRunner
}

const (
	taskTime = 5 * time.Second
)

func (m *arenaModule) InitTemplate() (err error) {
	err = arenatemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *arenaModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}

	err = arena.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("arena", arena.GetArenaService().Heartbeat, taskTime)
	return
}

func (m *arenaModule) Start() {
	arena.GetArenaService().Star()
	m.r.Start()
}

func (m *arenaModule) Stop() {
	m.r.Stop()
}

func (m *arenaModule) String() string {
	return "arena"
}

var (
	m = &arenaModule{}
)

func init() {
	module.Register(m)
}
