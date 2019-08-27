package fabao

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/fabao/dao"
	"fgame/fgame/game/fabao/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/fabao/active_check"
	_ "fgame/fgame/game/fabao/event/listener"
	_ "fgame/fgame/game/fabao/guaji"
	_ "fgame/fgame/game/fabao/handler"
	_ "fgame/fgame/game/fabao/player"
	_ "fgame/fgame/game/fabao/systemskill_handler"
	_ "fgame/fgame/game/fabao/use"
)

type faBaoModule struct {
}

func (m *faBaoModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *faBaoModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *faBaoModule) Start() {

}

func (m *faBaoModule) Stop() {

}

func (m *faBaoModule) String() string {
	return "fabao"
}

var (
	m = &faBaoModule{}
)

func init() {
	module.Register(m)
}
