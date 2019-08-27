package buff

import (
	"fgame/fgame/core/module"

	bufftemplate "fgame/fgame/game/buff/template"
)

import (
	_ "fgame/fgame/cross/buff/cross_handler"
	_ "fgame/fgame/game/buff/event/listener"
)

//buff模块
type buffModule struct {
}

func (m *buffModule) InitTemplate() (err error) {
	err = bufftemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *buffModule) Init() (err error) {

	return
}

func (m *buffModule) Start() {

}

func (m *buffModule) Stop() {

}

func (m *buffModule) String() string {
	return "buff"
}

var (
	m = &buffModule{}
)

func init() {
	module.Register(m)
}
