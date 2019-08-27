package shopdiscount

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/shopdiscount/dao"
	shopdiscounttemplate "fgame/fgame/game/shopdiscount/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/shopdiscount/event/listener"
	_ "fgame/fgame/game/shopdiscount/handler"
	_ "fgame/fgame/game/shopdiscount/player"
)

type shopdiscountModule struct {
}

func (m *shopdiscountModule) InitTemplate() (err error) {
	err = shopdiscounttemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *shopdiscountModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *shopdiscountModule) Start() {

}

func (m *shopdiscountModule) Stop() {

}

func (m *shopdiscountModule) String() string {
	return "shopdiscount"
}

var (
	m = &shopdiscountModule{}
)

func init() {
	module.Register(m)
}
