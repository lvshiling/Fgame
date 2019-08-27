package goldequip

import (
	"fgame/fgame/core/module"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
)

//活动大厅
type goldequipModule struct {
}

func (acModule *goldequipModule) InitTemplate() (err error) {
	err = goldequiptemplate.Init()
	return
}

func (acModule *goldequipModule) Init() (err error) {

	return
}

func (acModule *goldequipModule) Start() {

}

func (acModule *goldequipModule) Stop() {

}

func (acModule *goldequipModule) String() string {
	return "goldequip"
}

var (
	m = &goldequipModule{}
)

func init() {
	module.Register(m)
}
