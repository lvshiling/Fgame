package marry

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/marry/dao"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/template"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/marry/check_enter"
	_ "fgame/fgame/game/marry/event/listener"
	_ "fgame/fgame/game/marry/handler"
	_ "fgame/fgame/game/marry/npc"
	_ "fgame/fgame/game/marry/player"
)

type marryModule struct {
	r runner.GoRunner
}

func (m *marryModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

const (
	marryTimer = 2 * time.Second
)

func (m *marryModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = marry.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("marry", marry.GetMarryService().Heartbeat, marryTimer)
	return
}

func (m *marryModule) Start() {
	marry.GetMarryService().CreateMarrySceneData()
	m.r.Start()
}

func (m *marryModule) Stop() {

}

func (m *marryModule) String() string {
	return "marry"
}

var (
	m = &marryModule{}
)

func init() {
	module.Register(m)
}
