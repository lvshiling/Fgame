package dragon

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/dragon/dao"
	"fgame/fgame/game/dragon/dragon"
	"fgame/fgame/game/global"
)
import (
	//注册管理器
	_ "fgame/fgame/game/dragon/event/listener"
	_ "fgame/fgame/game/dragon/handler"
	_ "fgame/fgame/game/dragon/player"
)

type dragonModule struct {
}

func (m *dragonModule) InitTemplate() (err error) {
	err = dragon.Init()
	if err != nil {
		return
	}
	return
}

func (m *dragonModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *dragonModule) Start() {

}

func (m *dragonModule) Stop() {

}

func (m *dragonModule) String() string {
	return "dragon"
}

var (
	m = &dragonModule{}
)

func init() {
	module.Register(m)
}
