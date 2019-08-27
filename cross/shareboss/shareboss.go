package shareboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/cross/shareboss/shareboss"
	sharebosstemplate "fgame/fgame/game/shareboss/template"

	//注册
	_ "fgame/fgame/cross/shareboss/boss_handler"
	_ "fgame/fgame/cross/shareboss/event/listener"
	_ "fgame/fgame/cross/shareboss/guaji"
	_ "fgame/fgame/cross/shareboss/handler"

	_ "fgame/fgame/cross/shareboss/login_handler"
)

import (
	//注册
	_ "fgame/fgame/cross/arenaboss/boss_handler"
	_ "fgame/fgame/cross/arenaboss/login_handler"
)

//跨服世界boss
type shareBossModule struct {
}

func (m *shareBossModule) InitTemplate() (err error) {
	err = sharebosstemplate.Init()
	if err != nil {
		return
	}

	return
}
func (m *shareBossModule) Init() (err error) {

	err = shareboss.Init()

	return
}

func (m *shareBossModule) Start() {
	shareboss.GetShareBossService().Start()
	return
}

func (m *shareBossModule) Stop() {

}

func (m *shareBossModule) String() string {
	return "shareboss"
}

var (
	m = &shareBossModule{}
)

func init() {
	module.Register(m)
}
