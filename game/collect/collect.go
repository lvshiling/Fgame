package collect

import (
	//注册管理器
	"fgame/fgame/core/module"
	_ "fgame/fgame/game/collect/cross_handler"
	_ "fgame/fgame/game/collect/event/listener"
	_ "fgame/fgame/game/collect/event/listener/common"
	_ "fgame/fgame/game/collect/handler"
	collecttemplate "fgame/fgame/game/collect/template"
)

type collectModule struct {
}

func (m *collectModule) InitTemplate() (err error) {
	err = collecttemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *collectModule) Init() (err error) {

	return
}

func (m *collectModule) Start() {
}

func (m *collectModule) Stop() {
}

func (m *collectModule) String() string {
	return "collect"
}

var (
	m = &collectModule{}
)

func init() {
	module.Register(m)
}
