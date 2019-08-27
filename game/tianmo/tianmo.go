package tianmo

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/tianmo/dao"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/tianmo/event/listener"
	_ "fgame/fgame/game/tianmo/guaji"
	_ "fgame/fgame/game/tianmo/handler"
	_ "fgame/fgame/game/tianmo/player"
	_ "fgame/fgame/game/tianmo/systemskill_handler"
	_ "fgame/fgame/game/tianmo/use"
)

//天魔体
type tianmoModule struct {
}

func (m *tianmoModule) InitTemplate() (err error) {
	err = tianmotemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *tianmoModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *tianmoModule) Start() {
}

func (m *tianmoModule) Stop() {
}

func (m *tianmoModule) String() string {
	return "tianmo"
}

var (
	m = &tianmoModule{}
)

func init() {
	module.Register(m)
}
