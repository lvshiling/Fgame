package xianti

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xianti/dao"
	"fgame/fgame/game/xianti/xianti"
)

import (
	//注册管理器
	_ "fgame/fgame/game/xianti/active_check"
	_ "fgame/fgame/game/xianti/event/listener"
	_ "fgame/fgame/game/xianti/guaji"
	_ "fgame/fgame/game/xianti/handler"
	_ "fgame/fgame/game/xianti/player"
	_ "fgame/fgame/game/xianti/systemskill_handler"
	_ "fgame/fgame/game/xianti/use"
)

type xiantiModule struct {
}

func (cm *xiantiModule) InitTemplate() (err error) {
	err = xianti.Init()
	if err != nil {
		return
	}
	return
}

func (cm *xiantiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *xiantiModule) Start() {

}

func (cm *xiantiModule) Stop() {

}

func (cm *xiantiModule) String() string {
	return "xianti"
}

var (
	m = &xiantiModule{}
)

func init() {
	module.Register(m)
}
