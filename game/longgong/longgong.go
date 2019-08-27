package longgong

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/longgong/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/longgong/activity_handler"
	_ "fgame/fgame/game/longgong/check_enter"
	_ "fgame/fgame/game/longgong/collect_handler"
	_ "fgame/fgame/game/longgong/event/listener"
	_ "fgame/fgame/game/longgong/handler"
)

type longGongModule struct {
}

func (m *longGongModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *longGongModule) Init() (err error) {
	return
}

func (m *longGongModule) Start() {

}

func (m *longGongModule) Stop() {

}

func (m *longGongModule) String() string {
	return "longgong"
}

var (
	m = &longGongModule{}
)

func init() {
	module.Register(m)
}
