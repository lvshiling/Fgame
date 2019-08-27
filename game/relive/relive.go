package realm

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"

	"fgame/fgame/game/relive/dao"
)

import (
	//注册管理器
	_ "fgame/fgame/game/relive/cross_handler"
	_ "fgame/fgame/game/relive/event/listener"
	_ "fgame/fgame/game/relive/handler"
	_ "fgame/fgame/game/relive/player"
)

type reliveModule struct {
}

func (m *reliveModule) InitTemplate() (err error) {

	return
}

func (m *reliveModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *reliveModule) Start() {

}

func (m *reliveModule) Stop() {

}

func (m *reliveModule) String() string {
	return "relive"
}

var (
	m = &reliveModule{}
)

func init() {
	module.Register(m)
}
