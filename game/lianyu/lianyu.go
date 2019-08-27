package lianyu

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/lianyu/lianyu"
	"fgame/fgame/game/lianyu/template"
	"time"

	//注册管理器
	_ "fgame/fgame/game/lianyu/activity_handler"
	_ "fgame/fgame/game/lianyu/cross_handler"
	_ "fgame/fgame/game/lianyu/cross_loginflow"
	_ "fgame/fgame/game/lianyu/event/listener"
	_ "fgame/fgame/game/lianyu/handler"

	// 通用
	_ "fgame/fgame/game/lianyu/event/listener/common"
	_ "fgame/fgame/game/lianyu/guaji"
	_ "fgame/fgame/game/lianyu/relive_handler"
)

//无间炼狱
type lianYuModule struct {
	gr runner.GoRunner
}

func (m *lianYuModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return

}

func (m *lianYuModule) Init() (err error) {
	err = lianyu.Init(activitytypes.ActivityTypeLocalLianYu)
	if err != nil {
		return
	}

	m.gr = runner.NewGoRunner("lianyu", lianyu.GetLianYuService().Heartbeat, 3*time.Second)
	return
}

func (m *lianYuModule) Start() {
	m.gr.Start()
}
func (m *lianYuModule) Stop() {
	m.gr.Stop()
}

func (m *lianYuModule) String() string {
	return "lianyu"
}

var (
	m = &lianYuModule{}
)

func init() {
	module.Register(m)
}
