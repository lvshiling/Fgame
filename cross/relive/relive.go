package relive

import (
	"fgame/fgame/core/module"

	//注册管理器
	_ "fgame/fgame/cross/relive/cross_handler"
	_ "fgame/fgame/cross/relive/event/listener"
	_ "fgame/fgame/cross/relive/handler"
)

type reliveModule struct {
}

func (m *reliveModule) InitTemplate() (err error) {
	return
}

func (m *reliveModule) Init() (err error) {

	return
}

func (m *reliveModule) Start() {

}

func (m *reliveModule) Stop() {

}

func (m *reliveModule) String() string {
	return "relive"
}

var (
	m = &reliveModule{}
)

func init() {
	module.Register(m)
}
