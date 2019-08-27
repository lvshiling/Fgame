package common

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/pk/dao"
)

import (
	//注册管理器
	_ "fgame/fgame/game/pk/event/listener"
	_ "fgame/fgame/game/pk/event/listener/common"
	_ "fgame/fgame/game/pk/handler"
	_ "fgame/fgame/game/pk/player"
	_ "fgame/fgame/game/pk/player/event/listener"
	_ "fgame/fgame/game/pk/use"
)

type pkModule struct {
}

func (cm *pkModule) InitTemplate() (err error) {

	return
}

func (m *pkModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *pkModule) Start() {

}

func (m *pkModule) Stop() {

}

func (m *pkModule) String() string {
	return "pk"
}

var (
	m = &pkModule{}
)

func init() {
	module.Register(m)
}
