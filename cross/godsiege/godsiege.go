package godsiege

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"time"

	//注册管理器
	_ "fgame/fgame/cross/godsiege/cross_handler"
	_ "fgame/fgame/cross/godsiege/event/listener"
	_ "fgame/fgame/cross/godsiege/handler"
	_ "fgame/fgame/cross/godsiege/login_handler"

	_ "fgame/fgame/game/godsiege/event/listener/common"
	_ "fgame/fgame/game/godsiege/found_handler"
	_ "fgame/fgame/game/godsiege/guaji"
	_ "fgame/fgame/game/godsiege/relive_handler"
)

//神兽攻城
type godSiegeModule struct {
	gr runner.GoRunner
}

const (
	taskTime = 3 * time.Second
)

func (m *godSiegeModule) InitTemplate() (err error) {
	err = godsiegetemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *godSiegeModule) Init() (err error) {
	err = godsiege.Init(godsiegetypes.GodSiegeServerTypeCross)
	if err != nil {
		return
	}

	m.gr = runner.NewGoRunner("godsiege", godsiege.GetGodSiegeService().Heartbeat, taskTime)
	return
}

func (m *godSiegeModule) Start() {
	m.gr.Start()

}
func (m *godSiegeModule) Stop() {
	m.gr.Stop()
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
