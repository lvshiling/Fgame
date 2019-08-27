package anqi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/anqi/dao"
	anqitemplate "fgame/fgame/game/anqi/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/anqi/event/listener"
	_ "fgame/fgame/game/anqi/guaji"
	_ "fgame/fgame/game/anqi/handler"
	_ "fgame/fgame/game/anqi/player"
	_ "fgame/fgame/game/anqi/systemskill_handler"
	_ "fgame/fgame/game/anqi/use"
)

//暗器
type anqiModule struct {
}

func (m *anqiModule) InitTemplate() (err error) {
	err = anqitemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *anqiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *anqiModule) Start() {
}

func (m *anqiModule) Stop() {
}

func (m *anqiModule) String() string {
	return "anqi"
}

var (
	m = &anqiModule{}
)

func init() {
	module.Register(m)
}
