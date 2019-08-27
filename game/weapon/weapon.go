package weapon

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/weapon/dao"
	"fgame/fgame/game/weapon/weapon"
)

import (
	//注册管理器
	_ "fgame/fgame/game/weapon/active_check"
	_ "fgame/fgame/game/weapon/event/listener"
	_ "fgame/fgame/game/weapon/handler"
	_ "fgame/fgame/game/weapon/player"
	_ "fgame/fgame/game/weapon/use"
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
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
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
