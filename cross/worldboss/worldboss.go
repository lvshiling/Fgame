package worldboss

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/worldboss/cross_handler"
	_ "fgame/fgame/cross/worldboss/event/listener"
	_ "fgame/fgame/cross/worldboss/handler"
)

type worldbossModule struct {
}

func (m *worldbossModule) InitTemplate() (err error) {

	return
}
func (m *worldbossModule) Init() (err error) {

	return
}

func (m *worldbossModule) Start() {

}

func (m *worldbossModule) Stop() {

}

func (m *worldbossModule) String() string {
	return "worldboss"
}

var (
	m = &worldbossModule{}
)

func init() {
	module.Register(m)
}
