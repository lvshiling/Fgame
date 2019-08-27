package bagua

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/bagua/bagua"
	"fgame/fgame/game/bagua/dao"
	"fgame/fgame/game/bagua/template"
	"fgame/fgame/game/global"
	"time"

	//注册管理器
	_ "fgame/fgame/game/bagua/event/listener"
	_ "fgame/fgame/game/bagua/guaji"
	_ "fgame/fgame/game/bagua/handler"
	_ "fgame/fgame/game/bagua/player"
)

type baGuaModule struct {
	r runner.GoRunner
}

func (m *baGuaModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *baGuaModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = bagua.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("bagua", bagua.GetBaGuaService().Heartbeat, 5*time.Second)
	return
}

func (m *baGuaModule) Start() {
	m.r.Start()

}

func (m *baGuaModule) Stop() {
	m.r.Stop()
}

func (m *baGuaModule) String() string {
	return "bagua"
}

var (
	m = &baGuaModule{}
)

func init() {
	module.Register(m)
}
