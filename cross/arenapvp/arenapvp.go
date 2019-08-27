package arenapvp

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/cross/arenapvp/arenapvp"
	"fgame/fgame/cross/arenapvp/dao"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	"fgame/fgame/game/global"
	"time"

	_ "fgame/fgame/cross/arenapvp/battle/check_attack"
	_ "fgame/fgame/cross/arenapvp/cross_handler"
	_ "fgame/fgame/cross/arenapvp/event/listener"
	_ "fgame/fgame/cross/arenapvp/handler"
	_ "fgame/fgame/cross/arenapvp/login_handler"
	_ "fgame/fgame/cross/arenapvp/relive_handler"
	_ "fgame/fgame/cross/arenapvp/robot/ai"
)

const (
	arenapvpTaskTime = 3 * time.Second
)

type arenapvpModule struct {
	r runner.GoRunner
}

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

	m.r = runner.NewGoRunner("arenapvp", arenapvp.GetArenapvpService().Heartbeat, arenapvpTaskTime)
	return
}

func (m *arenapvpModule) Start() {
	arenapvp.GetArenapvpService().Start()
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
