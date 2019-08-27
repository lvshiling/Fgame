package arena

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/cross/arena/arena"
	arenatemplate "fgame/fgame/game/arena/template"
	"time"

	_ "fgame/fgame/cross/arena/battle"
	_ "fgame/fgame/cross/arena/cross_handler"
	_ "fgame/fgame/cross/arena/event/listener"
	_ "fgame/fgame/cross/arena/handler"
	_ "fgame/fgame/cross/arena/login_handler"
	_ "fgame/fgame/cross/arena/relive_handler"
	_ "fgame/fgame/cross/arena/robot/ai"
)

type arenaModule struct {
	r runner.GoRunner
}

func (m *arenaModule) InitTemplate() (err error) {
	err = arenatemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *arenaModule) Init() (err error) {
	err = arenatemplate.Init()
	if err != nil {
		return
	}

	err = arena.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("arena", arena.GetArenaService().Heartbeat, time.Second)
	return
}

func (m *arenaModule) Start() {
	arena.GetArenaService().Start()

	m.r.Start()

}

func (m *arenaModule) Stop() {
	m.r.Stop()
	arena.GetArenaService().Stop()
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
