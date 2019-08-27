package jieyi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/jieyi/dao"
	"fgame/fgame/game/jieyi/jieyi"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"time"

	// 注册管理器
	_ "fgame/fgame/game/jieyi/drop_handler"

	_ "fgame/fgame/game/jieyi/event/listener"

	_ "fgame/fgame/game/jieyi/cross_handler"
	_ "fgame/fgame/game/jieyi/handler"

	_ "fgame/fgame/game/jieyi/player"

	_ "fgame/fgame/game/jieyi/use"
)

const (
	tastTime = 15 * time.Second
)

// 结义
type jieYiModule struct {
	r runner.GoRunner // 定时器
}

func (m *jieYiModule) InitTemplate() (err error) {
	err = jieyitemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *jieYiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = jieyi.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("jieyi", jieyi.GetJieYiService().Heartbeat, tastTime)

	return
}

func (m *jieYiModule) Start() {
	m.r.Start()
}

func (m *jieYiModule) Stop() {
	m.r.Stop()
}

func (m *jieYiModule) String() string {
	return "jieyi"
}

var (
	m = &jieYiModule{}
)

func init() {
	module.Register(m)
}
