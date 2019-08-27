package fourgod

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/fourgod/dao"
	"fgame/fgame/game/fourgod/fourgod"
	"fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/fourgod/activity_handler"
	_ "fgame/fgame/game/fourgod/check_enter"
	_ "fgame/fgame/game/fourgod/collect_handler"
	_ "fgame/fgame/game/fourgod/drop_handler"
	_ "fgame/fgame/game/fourgod/event/listener"
	_ "fgame/fgame/game/fourgod/found_handler"
	_ "fgame/fgame/game/fourgod/guaji"
	_ "fgame/fgame/game/fourgod/handler"
	_ "fgame/fgame/game/fourgod/player"
)

type fourGodModule struct {
}

func (m *fourGodModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *fourGodModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = fourgod.Init()
	if err != nil {
		return
	}

	return
}

func (m *fourGodModule) Start() {

}

func (m *fourGodModule) Stop() {

}

func (m *fourGodModule) String() string {
	return "fourgod"
}

var (
	m = &fourGodModule{}
)

func init() {
	module.Register(m)
}
