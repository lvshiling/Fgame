package yinglingpu

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/yinglingpu/dao"
	ylptemplate "fgame/fgame/game/yinglingpu/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/yinglingpu/event/listener"
	_ "fgame/fgame/game/yinglingpu/handler"
	_ "fgame/fgame/game/yinglingpu/player"
)

type yingLingPuModule struct {
}

func (m *yingLingPuModule) InitTemplate() (err error) {
	err = ylptemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *yingLingPuModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *yingLingPuModule) Start() {
}

func (m *yingLingPuModule) Stop() {
}

func (m *yingLingPuModule) String() string {
	return "yinglingpu"
}

var (
	m = &yingLingPuModule{}
)

func init() {
	module.Register(m)
}
