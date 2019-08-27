package cd

import (
	"fgame/fgame/core/module"
	cdtemplate "fgame/fgame/game/cd/template"
)

//缓存模块
type cdModule struct {
}

func (m *cdModule) InitTemplate() (err error) {
	err = cdtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *cdModule) Init() (err error) {

	return
}

func (m *cdModule) Start() {

	return
}

func (m *cdModule) Stop() {

}

func (m *cdModule) String() string {
	return "cd"
}

var (
	m = &cdModule{}
)

func init() {
	module.Register(m)
}
