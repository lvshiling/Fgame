package moonlove

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/moonlove/dao"
	moonlovetemplate "fgame/fgame/game/moonlove/template"
)

import (
	_ "fgame/fgame/game/moonlove/activity_handler"
	_ "fgame/fgame/game/moonlove/check_enter"
	_ "fgame/fgame/game/moonlove/collect_handler"
	_ "fgame/fgame/game/moonlove/event/listener"
	_ "fgame/fgame/game/moonlove/found_handler"
	_ "fgame/fgame/game/moonlove/guaji"
	_ "fgame/fgame/game/moonlove/handler"
	_ "fgame/fgame/game/moonlove/player"
)

//月下情缘
type moonloveModule struct {
}

func (m *moonloveModule) InitTemplate() (err error) {
	err = moonlovetemplate.Init()

	if err != nil {
		return
	}
	return
}

func (m *moonloveModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return err
	}

	return
}

func (m *moonloveModule) Start() {

}

func (m *moonloveModule) Stop() {

}

func (mlModule *moonloveModule) String() string {
	return "moonlove"
}

var (
	m = &moonloveModule{}
)

func init() {
	module.Register(m)
}
