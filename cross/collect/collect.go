package collect

import (
	"fgame/fgame/core/module"

	//注册管理器
	_ "fgame/fgame/cross/collect/cross_handler"
	_ "fgame/fgame/cross/collect/event/listener"
	_ "fgame/fgame/cross/collect/handler"
	_ "fgame/fgame/game/collect/event/listener/common"
)

type collectModule struct {
}

func (m *collectModule) InitTemplate() (err error) {

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
