package arenapvp

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/dao"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	"fgame/fgame/game/global"
	"time"

	_ "fgame/fgame/game/arenapvp/activity_handler"
	_ "fgame/fgame/game/arenapvp/cross_handler"
	_ "fgame/fgame/game/arenapvp/cross_loginflow"
	_ "fgame/fgame/game/arenapvp/drop_handler"
	_ "fgame/fgame/game/arenapvp/event/listener"
	_ "fgame/fgame/game/arenapvp/handler"
	_ "fgame/fgame/game/arenapvp/player"
	_ "fgame/fgame/game/arenapvp/types/scene"
	_ "fgame/fgame/game/arenapvp/use"
)

type arenapvpModule struct {
	r runner.GoRunner
}

const (
	taskTime = 5 * time.Second
)

func (m *arenapvpModule) InitTemplate() (err error) {
	err = arenapvptemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *arenapvpModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}

	err = arenapvp.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("arenapvp", arenapvp.GetArenapvpService().Heartbeat, taskTime)
	return
}

func (m *arenapvpModule) Start() {
	arenapvp.GetArenapvpService().Star()
	m.r.Start()
}

func (m *arenapvpModule) Stop() {
	m.r.Stop()
}

func (m *arenapvpModule) String() string {
	return "arenapvp"
}

var (
	m = &arenapvpModule{}
)

func init() {
	module.Register(m)
}
