package onearena

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/onearena/dao"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/onearena/template"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/onearena/event/listener"
	_ "fgame/fgame/game/onearena/handler"
	_ "fgame/fgame/game/onearena/player"
)

type oneArenaModule struct {
	r runner.GoRunner
}

func (m *oneArenaModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

const (
	//TODO 暂时不修改,是否太频繁
	oneArenaTimer = time.Second
)

func (m *oneArenaModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = onearena.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("onearena", onearena.GetOneArenaService().Heartbeat, oneArenaTimer)
	return
}

func (m *oneArenaModule) Start() {
	m.r.Start()

}

func (m *oneArenaModule) Stop() {

}

func (m *oneArenaModule) String() string {
	return "onearena"
}

var (
	m = &oneArenaModule{}
)

func init() {
	module.Register(m)
}
