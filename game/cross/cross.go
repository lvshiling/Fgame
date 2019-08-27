package cross

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/cross/cross"
	"fgame/fgame/game/cross/dao"
	"fgame/fgame/game/global"
)

import (
	_ "fgame/fgame/game/cross/cross_handler"
	_ "fgame/fgame/game/cross/event/listener"
	_ "fgame/fgame/game/cross/loginflow"
	_ "fgame/fgame/game/cross/player"
)

type crossModule struct {
}

func (m *crossModule) InitTemplate() (err error) {
	return
}

func (m *crossModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = cross.Init()
	if err != nil {
		return
	}
	return
}

func (m *crossModule) Start() {
	cross.GetCrossService().Start()

	return
}

func (m *crossModule) Stop() {

}

func (m *crossModule) String() string {
	return "cross"
}

var (
	m = &crossModule{}
)

func init() {
	module.Register(m)
}
