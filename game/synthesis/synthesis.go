package synthesis

import (
	"fgame/fgame/core/module"
	synthesistemplate "fgame/fgame/game/synthesis/template"
)

import (
	_ "fgame/fgame/game/synthesis/handler"
)

//合成模块
type synthesisModule struct {
}

func (m *synthesisModule) InitTemplate() (err error) {
	if err = synthesistemplate.Init(); err != nil {
		return
	}

	return
}
func (m *synthesisModule) Init() (err error) {

	return
}

func (m *synthesisModule) Start() {

}

func (m *synthesisModule) Stop() {

}

func (m *synthesisModule) String() string {
	return "synthesis"
}

var (
	m = &synthesisModule{}
)

func init() {
	module.Register(m)
}
