package shareboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/shareboss/shareboss"
	sharebosstemplate "fgame/fgame/game/shareboss/template"
	"time"
)

import (
	//注册
	_ "fgame/fgame/game/shareboss/boss_handler"
	_ "fgame/fgame/game/shareboss/check_enter"
	_ "fgame/fgame/game/shareboss/handler"
)

//跨服世界boss
type shareBossModule struct {
	r runner.GoRunner
}

func (m *shareBossModule) InitTemplate() (err error) {
	err = sharebosstemplate.Init()
	if err != nil {
		return
	}

	return
}

const (
	shareBossTimer = time.Second * 5
)

func (m *shareBossModule) Init() (err error) {

	err = shareboss.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("shareboss", shareboss.GetShareBossService().Heartbeat, shareBossTimer)
	return
}

func (m *shareBossModule) Start() {
	shareboss.GetShareBossService().Start()

	m.r.Start()
	return
}

func (m *shareBossModule) Stop() {
	m.r.Stop()
}

func (m *shareBossModule) String() string {
	return "shareboss"
}

var (
	m = &shareBossModule{}
)

func init() {
	module.Register(m)
}
