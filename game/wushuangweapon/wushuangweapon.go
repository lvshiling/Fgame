package wushuangweapon

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/wushuangweapon/dao"
	wushuangweapontemplate "fgame/fgame/game/wushuangweapon/template"
)

import (
	_ "fgame/fgame/game/wushuangweapon/event/listener"
	_ "fgame/fgame/game/wushuangweapon/handler"
	_ "fgame/fgame/game/wushuangweapon/player"
)

type wushuangWeaponModule struct {
}

func (m *wushuangWeaponModule) InitTemplate() (err error) {
	err = wushuangweapontemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *wushuangWeaponModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *wushuangWeaponModule) Start() {
	return
}

func (m *wushuangWeaponModule) Stop() {
}

func (m *wushuangWeaponModule) String() string {
	return "wushuangweapon"
}

var (
	m = &wushuangWeaponModule{}
)

func init() {
	module.Register(m)
}
