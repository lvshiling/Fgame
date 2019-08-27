package ring

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/ring/dao"
	ringtemplate "fgame/fgame/game/ring/template"
)

import (
	_ "fgame/fgame/game/ring/event/listener"
	_ "fgame/fgame/game/ring/handler"
	_ "fgame/fgame/game/ring/player"
)

type ringModule struct {
}

func (m *ringModule) InitTemplate() (err error) {
	err = ringtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *ringModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *ringModule) Start() {
}

func (m *ringModule) Stop() {
}

func (m *ringModule) String() string {
	return "ring"
}

var (
	m = &ringModule{}
)

func init() {
	module.Register(m)
}
