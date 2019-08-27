package godsiege

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/godsiege/godsiege"
	"fgame/fgame/game/godsiege/template"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"time"

	//注册管理器
	_ "fgame/fgame/game/godsiege/activity_handler"
	_ "fgame/fgame/game/godsiege/cross_handler"
	_ "fgame/fgame/game/godsiege/cross_loginflow"
	_ "fgame/fgame/game/godsiege/event/listener"
	_ "fgame/fgame/game/godsiege/found_handler"
	_ "fgame/fgame/game/godsiege/handler"

	// 通用
	_ "fgame/fgame/game/godsiege/event/listener/common"
	_ "fgame/fgame/game/godsiege/guaji"
	_ "fgame/fgame/game/godsiege/relive_handler"
)

//神兽攻城
type godSiegeModule struct {
	r runner.GoRunner
}

const (
	taskTime = 3 * time.Second
)

func (m *godSiegeModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *godSiegeModule) Init() (err error) {
	err = godsiege.Init(godsiegetypes.GodSiegeServerTypeLocal)
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("godsiege", godsiege.GetGodSiegeService().Heartbeat, taskTime)
	return
}

func (m *godSiegeModule) Start() {
	m.r.Start()
}
func (m *godSiegeModule) Stop() {
	m.r.Stop()
}

func (m *godSiegeModule) String() string {
	return "godsiege"
}

var (
	m = &godSiegeModule{}
)

func init() {
	module.Register(m)
}
