package hunt

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/hunt/dao"
	hunttemplate "fgame/fgame/game/hunt/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/hunt/event/listener"
	_ "fgame/fgame/game/hunt/handler"
	_ "fgame/fgame/game/hunt/player"
)

//屠龙装备
type tuLongEquipModule struct {
}

func (m *tuLongEquipModule) InitTemplate() (err error) {
	hunttemplate.Init()
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
	return "hunt"
}

var (
	m = &tuLongEquipModule{}
)

func init() {
	module.Register(m)
}
