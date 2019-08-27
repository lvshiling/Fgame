package wing

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/xuechi/cross_handler"
	_ "fgame/fgame/cross/xuechi/event/listener"
	_ "fgame/fgame/cross/xuechi/handler"
)

type xueChiModule struct {
}

func (m *xueChiModule) InitTemplate() (err error) {

	return
}
func (m *xueChiModule) Init() (err error) {

	return
}

func (m *xueChiModule) Start() {

}

func (m *xueChiModule) Stop() {

}

func (m *xueChiModule) String() string {
	return "xuechi"
}

var (
	m = &xueChiModule{}
)

func init() {
	module.Register(m)
}
