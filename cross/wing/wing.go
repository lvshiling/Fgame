package wing

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/wing/wing"
)

type wingModule struct {
}

func (m *wingModule) InitTemplate() (err error) {
	err = wing.Init()
	if err != nil {
		return
	}
	return
}
func (m *wingModule) Init() (err error) {

	return
}

func (m *wingModule) Start() {

}

func (m *wingModule) Stop() {

}

func (m *wingModule) String() string {
	return "wing"
}

var (
	m = &wingModule{}
)

func init() {
	module.Register(m)
}
