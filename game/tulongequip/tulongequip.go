package tulongequip

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/tulongequip/dao"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/tulongequip/event/listener"
	_ "fgame/fgame/game/tulongequip/handler"
	_ "fgame/fgame/game/tulongequip/player"
)

//屠龙装备
type tuLongEquipModule struct {
}

func (m *tuLongEquipModule) InitTemplate() (err error) {
	err = tulongequiptemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *tuLongEquipModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *tuLongEquipModule) Start() {

}

func (m *tuLongEquipModule) Stop() {

}

func (m *tuLongEquipModule) String() string {
	return "tulongequip"
}

var (
	m = &tuLongEquipModule{}
)

func init() {
	module.Register(m)
}
