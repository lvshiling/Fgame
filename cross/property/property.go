package property

import (
	"fgame/fgame/core/module"

	_ "fgame/fgame/game/property/npc/effect"
)

//模块化
type propertyModule struct {
}

func (m *propertyModule) InitTemplate() (err error) {

	return
}
func (m *propertyModule) Init() (err error) {
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
