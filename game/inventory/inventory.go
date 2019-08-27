package inventory

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/inventory/dao"
	"fgame/fgame/game/item/item"
)

import (
	//注册管理器
	_ "fgame/fgame/game/inventory/cross_handler"
	_ "fgame/fgame/game/inventory/event/listener"
	_ "fgame/fgame/game/inventory/gm/event/listener"
	_ "fgame/fgame/game/inventory/guaji"
	_ "fgame/fgame/game/inventory/use"

	_ "fgame/fgame/game/inventory/handler"
	_ "fgame/fgame/game/inventory/inventory"
	_ "fgame/fgame/game/inventory/player"
)

type inventoryModule struct {
}

func (m *inventoryModule) InitTemplate() (err error) {
	err = item.Init()
	if err != nil {
		return
	}
	return
}

func (m *inventoryModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *inventoryModule) Start() {

}

func (m *inventoryModule) Stop() {

}

func (m *inventoryModule) String() string {
	return "inventory"
}

var (
	m = &inventoryModule{}
)

func init() {
	module.RegisterBase(m)
}
