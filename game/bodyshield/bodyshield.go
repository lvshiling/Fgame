package bodyshield

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/bodyshield/bodyshield"
	"fgame/fgame/game/bodyshield/dao"
	"fgame/fgame/game/global"
)
import (
	//注册管理器
	_ "fgame/fgame/game/bodyshield/event/listener"
	_ "fgame/fgame/game/bodyshield/guaji"
	_ "fgame/fgame/game/bodyshield/handler"
	_ "fgame/fgame/game/bodyshield/player"
	_ "fgame/fgame/game/bodyshield/system_handler"
	_ "fgame/fgame/game/bodyshield/use"
)

type bodyshieldModule struct {
}

func (m *bodyshieldModule) InitTemplate() (err error) {
	err = bodyshield.Init()
	if err != nil {
		return
	}
	return
}

func (m *bodyshieldModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *bodyshieldModule) Start() {

}

func (m *bodyshieldModule) Stop() {

}

func (m *bodyshieldModule) String() string {
	return "bodyshield"
}

var (
	m = &bodyshieldModule{}
)

func init() {
	module.Register(m)
}
