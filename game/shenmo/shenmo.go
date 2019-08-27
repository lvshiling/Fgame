package shenmo

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	activitytypes "fgame/fgame/game/activity/types"

	"fgame/fgame/game/global"
	"fgame/fgame/game/shenmo/dao"
	"fgame/fgame/game/shenmo/shenmo"
	"fgame/fgame/game/shenmo/template"
	"time"

	//注册管理器
	_ "fgame/fgame/game/shenmo/activity_handler"
	_ "fgame/fgame/game/shenmo/cross_handler"
	_ "fgame/fgame/game/shenmo/cross_loginflow"
	_ "fgame/fgame/game/shenmo/drop_handler"
	_ "fgame/fgame/game/shenmo/event/listener"
	_ "fgame/fgame/game/shenmo/found_handler"
	_ "fgame/fgame/game/shenmo/handler"

	//
	_ "fgame/fgame/game/shenmo/event/listener/common"
	_ "fgame/fgame/game/shenmo/guaji"
	_ "fgame/fgame/game/shenmo/relive_handler"
)

//神魔
type shenMoModule struct {
	gr runner.GoRunner
}

func (m *shenMoModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *shenMoModule) Init() (err error) {

	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = shenmo.Init(activitytypes.ActivityTypeLocalShenMoWar)
	if err != nil {
		return
	}
	m.gr = runner.NewGoRunner("shenmo", shenmo.GetShenMoService().Heartbeat, time.Second)
	return
}

func (m *shenMoModule) Start() {
	shenmo.GetShenMoService().Star()
	m.gr.Start()
}

func (m *shenMoModule) Stop() {
	m.gr.Stop()
}

func (m *shenMoModule) String() string {
	return "shenmo"
}

var (
	m = &shenMoModule{}
)

func init() {
	module.Register(m)
}
