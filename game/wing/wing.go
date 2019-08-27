package wing

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/wing/dao"
	"fgame/fgame/game/wing/wing"
)

import (
	//注册管理器
	_ "fgame/fgame/game/wing/active_check"
	_ "fgame/fgame/game/wing/event/listener"
	_ "fgame/fgame/game/wing/guaji"
	_ "fgame/fgame/game/wing/handler"
	_ "fgame/fgame/game/wing/player"
	_ "fgame/fgame/game/wing/systemskill_handler"
	_ "fgame/fgame/game/wing/use"
)

type wingModule struct {
}

func (m *wingModule) InitTemplate() (err error) {
	err = wing.Init()
	if err != nil {
		return
	}
	return
}
func (m *wingModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *wingModule) Start() {

}

func (m *wingModule) Stop() {

}

func (m *wingModule) String() string {
	return "wing"
}

var (
	m = &wingModule{}
)

func init() {
	module.Register(m)
}
