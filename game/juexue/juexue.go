package juexue

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/juexue/dao"
	"fgame/fgame/game/juexue/juexue"
)

import (
	//注册管理器
	_ "fgame/fgame/game/juexue/event/listener"
	_ "fgame/fgame/game/juexue/handler"
	_ "fgame/fgame/game/juexue/player"
)

type juexueModule struct {
}

func (m *juexueModule) InitTemplate() (err error) {
	err = juexue.Init()
	if err != nil {
		return
	}
	return
}

func (m *juexueModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *juexueModule) Start() {

}

func (m *juexueModule) Stop() {

}

func (m *juexueModule) String() string {
	return "juexue"
}

var (
	m = &juexueModule{}
)

func init() {
	module.Register(m)
}
