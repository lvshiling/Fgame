package inventory

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/item/item"
	//注册管理器
	_ "fgame/fgame/cross/inventory/event/listener"
	_ "fgame/fgame/cross/inventory/handler"
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
