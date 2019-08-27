package realm

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/realm/dao"
	"fgame/fgame/game/realm/realm"
	"fgame/fgame/game/realm/template"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/realm/event/listener"
	_ "fgame/fgame/game/realm/guaji"
	_ "fgame/fgame/game/realm/handler"
	_ "fgame/fgame/game/realm/player"
)

type realmModule struct {
	r runner.GoRunner
}

func (m *realmModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *realmModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = realm.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("realm", realm.GetRealmRankService().Heartbeat, 5*time.Second)
	return
}

func (m *realmModule) Start() {
	m.r.Start()

}

func (m *realmModule) Stop() {
	m.r.Stop()
}

func (m *realmModule) String() string {
	return "realm"
}

var (
	m = &realmModule{}
)

func init() {
	module.Register(m)
}
