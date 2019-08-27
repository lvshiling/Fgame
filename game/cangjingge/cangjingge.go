package cangjingge

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/cangjingge/cangjingge"
	cangjinggetemplate "fgame/fgame/game/cangjingge/template"
)
import (
	//注册
	_ "fgame/fgame/game/cangjingge/boss_handler"
	_ "fgame/fgame/game/cangjingge/check_enter"
	_ "fgame/fgame/game/cangjingge/event/listener"
	_ "fgame/fgame/game/cangjingge/guaji"
	_ "fgame/fgame/game/cangjingge/handler"
)

//藏经阁boss
type cangJingGeModule struct {
}

func (m *cangJingGeModule) InitTemplate() (err error) {
	err = cangjinggetemplate.Init()
	if err != nil {
		return
	}

	return
}
func (m *cangJingGeModule) Init() (err error) {

	err = cangjingge.Init()

	return
}

func (m *cangJingGeModule) Start() {
	cangjingge.GetCangJingGeService().Start()
	return
}

func (m *cangJingGeModule) Stop() {

}

func (m *cangJingGeModule) String() string {
	return "cangjingge"
}

var (
	m = &cangJingGeModule{}
)

func init() {
	module.Register(m)
}
