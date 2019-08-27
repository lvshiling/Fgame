package goldequip

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/goldequip/dao"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/goldequip/event/listener"
	_ "fgame/fgame/game/goldequip/guaji"
	_ "fgame/fgame/game/goldequip/handler"
	_ "fgame/fgame/game/goldequip/player"
	_ "fgame/fgame/game/goldequip/use"
)

//元神金装
type goldEquipModule struct {
}

func (m *goldEquipModule) InitTemplate() (err error) {
	goldequiptemplate.Init()
	return
}

func (m *goldEquipModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *goldEquipModule) Start() {

}

func (m *goldEquipModule) Stop() {

}

func (m *goldEquipModule) String() string {
	return "goldequip"
}

var (
	m = &goldEquipModule{}
)

func init() {
	module.Register(m)
}
