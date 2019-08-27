package welfarescene

import (
	"fgame/fgame/core/module"
	welfarescenetemplate "fgame/fgame/game/welfarescene/template"
)

import (
	"fgame/fgame/game/welfarescene/welfarescene"
	//注册管理器
	_ "fgame/fgame/game/welfarescene/check_enter"
	// _ "fgame/fgame/game/welfarescene/cross_handler"
	// _ "fgame/fgame/game/welfarescene/cross_loginflow"
	_ "fgame/fgame/game/welfarescene/event/listener"
	_ "fgame/fgame/game/welfarescene/handler"
	_ "fgame/fgame/game/welfarescene/scene"
)

//运营活动副本
type welfaresceneModule struct {
}

func (m *welfaresceneModule) InitTemplate() (err error) {
	err = welfarescenetemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *welfaresceneModule) Init() (err error) {
	err = welfarescene.Init()
	if err != nil {
		return
	}

	return
}

func (m *welfaresceneModule) Start() {

}

func (m *welfaresceneModule) Stop() {

}

func (m *welfaresceneModule) String() string {
	return "welfarescene"
}

var (
	m = &welfaresceneModule{}
)

func init() {
	module.Register(m)
}
