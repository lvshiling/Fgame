package skill

import (
	"fgame/fgame/core/module"
	skilltemplate "fgame/fgame/game/skill/template"
)

type skillModule struct {
}

func (m *skillModule) InitTemplate() (err error) {
	err = skilltemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *skillModule) Init() (err error) {

	return
}

func (m *skillModule) Start() {

}
func (m *skillModule) Stop() {

}

func (m *skillModule) String() string {
	return "skill"
}

var (
	m = &skillModule{}
)

func init() {
	module.Register(m)
}
