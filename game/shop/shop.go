package shop

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/global"
	"fgame/fgame/game/shop/dao"
	"fgame/fgame/game/shop/shop"
)

import (
	//注册管理器
	_ "fgame/fgame/game/shop/event/listener"
	_ "fgame/fgame/game/shop/handler"
	_ "fgame/fgame/game/shop/player"
)

type shopModule struct {
}

func (m *shopModule) InitTemplate() (err error) {
	err = shop.Init()
	if err != nil {
		return
	}
	return
}

func (m *shopModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *shopModule) Start() {

}

func (m *shopModule) Stop() {

}

func (m *shopModule) String() string {
	return "shop"
}

var (
	m = &shopModule{}
)

func init() {
	module.Register(m)
}
