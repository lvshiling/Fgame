package shenfa

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lucky/dao"
)

import (
	//注册管理器
	_ "fgame/fgame/game/lucky/event/listener"
	_ "fgame/fgame/game/lucky/handler"
	_ "fgame/fgame/game/lucky/player"
	_ "fgame/fgame/game/lucky/use"
)

//幸运符
type luckyModule struct {
}

func (m *luckyModule) InitTemplate() (err error) {

	return
}

func (m *luckyModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *luckyModule) Start() {

}

func (m *luckyModule) Stop() {

}

func (m *luckyModule) String() string {
	return "lucky"
}

var (
	m = &luckyModule{}
)

func init() {
	module.Register(m)
}
