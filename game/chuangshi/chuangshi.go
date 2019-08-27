package chuangshi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/chuangshi/chuangshi"
	"fgame/fgame/game/chuangshi/dao"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	"fgame/fgame/game/global"
	"time"

	//注册管理器
	// _ "fgame/fgame/game/chuangshi/activity_handler"
	// _ "fgame/fgame/game/chuangshi/cross_handler"
	// _ "fgame/fgame/game/chuangshi/cross_loginflow"
	_ "fgame/fgame/game/chuangshi/event/listener"
	_ "fgame/fgame/game/chuangshi/handler"
	_ "fgame/fgame/game/chuangshi/player" 
)

//创世之战
type chuangShiModule struct {
	r runner.GoRunner
}

const (
	taskTime = 5 * time.Second
)

func (m *chuangShiModule) InitTemplate() (err error) {
	err = chuangshitemplate.Init()
	return
}

func (m *chuangShiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = chuangshi.Init()
	if err != nil {
		return
	}

	// m.r = runner.NewGoRunner("chuangshi", chuangshi.GetChuangShiService().Heartbeat, taskTime)
	return
}

func (m *chuangShiModule) Start() {
	chuangshi.GetChuangShiService().Star()
	// m.r.Start()
}

func (m *chuangShiModule) Stop() {
	// m.r.Stop()
}

func (m *chuangShiModule) String() string {
	return "chuangshi"
}

var (
	m = &chuangShiModule{}
)

func init() {
	module.Register(m)
}
