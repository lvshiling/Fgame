package densewat

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/densewat/dao"
	"fgame/fgame/game/global"
)

//注册管理器
import (
	_ "fgame/fgame/game/densewat/activity_handler"
	_ "fgame/fgame/game/densewat/cross_handler"
	_ "fgame/fgame/game/densewat/event/listener"
	_ "fgame/fgame/game/densewat/found_handler"
	_ "fgame/fgame/game/densewat/player"
)

type denseWatModule struct {
}

func (m *denseWatModule) InitTemplate() (err error) {
	return
}

func (m *denseWatModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *denseWatModule) Start() {

}

func (m *denseWatModule) Stop() {

}

func (m *denseWatModule) String() string {
	return "densewat"
}

var (
	m = &denseWatModule{}
)

func init() {
	module.Register(m)
}
