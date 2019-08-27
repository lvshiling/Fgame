package major

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/major/dao"
	"fgame/fgame/game/major/major"
	"fgame/fgame/game/major/template"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/major/event/listener"
	_ "fgame/fgame/game/major/found_handler"
	_ "fgame/fgame/game/major/guaji_check"
	_ "fgame/fgame/game/major/handler"
	_ "fgame/fgame/game/major/player"
)

type majorModule struct {
	r runner.GoRunner
}

func (m *majorModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *majorModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = major.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("major", major.GetMajorService().Heartbeat, 5*time.Second)
	return
}

func (m *majorModule) Start() {
	m.r.Start()

}

func (m *majorModule) Stop() {
	m.r.Stop()
}

func (m *majorModule) String() string {
	return "major"
}

var (
	m = &majorModule{}
)

func init() {
	module.Register(m)
}
