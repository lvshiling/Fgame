package dummy

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/dummy/template"
)

//假名
type dummyModule struct {
}

func (m *dummyModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *dummyModule) Init() (err error) {
	return
}

func (m *dummyModule) Start() {
	return
}

func (m *dummyModule) Stop() {
}

func (m *dummyModule) String() string {
	return "dummy"
}

var (
	m = &dummyModule{}
)

func init() {
	module.RegisterBase(m)
}
