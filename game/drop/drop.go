package drop

import (
	"fgame/fgame/core/module"

	droptemplate "fgame/fgame/game/drop/template"
)

import (
	_ "fgame/fgame/game/drop/event/listener"
)

type dropModule struct {
}

func (m *dropModule) InitTemplate() (err error) {
	err = droptemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *dropModule) Init() (err error) {
	return
}

func (m *dropModule) Start() {

}

func (m *dropModule) Stop() {

}

func (m *dropModule) String() string {
	return "drop"
}

var (
	m = &dropModule{}
)

func init() {
	module.RegisterBase(m)
}
