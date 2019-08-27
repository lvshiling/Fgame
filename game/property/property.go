package property

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/property/dao"
)

import (
	_ "fgame/fgame/game/property/drop_handler"
	_ "fgame/fgame/game/property/event/listener"
	_ "fgame/fgame/game/property/npc/effect"
	_ "fgame/fgame/game/property/player/effect"
	_ "fgame/fgame/game/property/player/event/listener"
	_ "fgame/fgame/game/property/use"
)

//模块化
type propertyModule struct {
}

func (m *propertyModule) InitTemplate() (err error) {

	return
}
func (m *propertyModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *propertyModule) Start() {

}

func (m *propertyModule) Stop() {

}

func (m *propertyModule) String() string {
	return "property"
}

var (
	m = &propertyModule{}
)

func init() {
	module.Register(m)
}
