package weapon

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/weapon/weapon"
)

type weaponModule struct {
}

func (m *weaponModule) InitTemplate() (err error) {
	err = weapon.Init()
	if err != nil {
		return
	}
	return
}
func (m *weaponModule) Init() (err error) {

	return
}
func (m *weaponModule) Start() {

}
func (m *weaponModule) Stop() {

}

func (m *weaponModule) String() string {
	return "weapon"
}

var (
	m = &weaponModule{}
)

func init() {
	module.Register(m)
}
