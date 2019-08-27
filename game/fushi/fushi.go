package fushi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/fushi/dao"
	"fgame/fgame/game/fushi/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/fushi/event/listener"
	_ "fgame/fgame/game/fushi/handler"
	_ "fgame/fgame/game/fushi/player"
)

type fuShiModule struct {
}

func (m *fuShiModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *fuShiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *fuShiModule) Start() {

}

func (m *fuShiModule) Stop() {

}

func (m *fuShiModule) String() string {
	return "fushi"
}

var (
	m = &fuShiModule{}
)

func init() {
	module.Register(m)
}
