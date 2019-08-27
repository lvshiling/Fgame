package lingyu

import (
	"fgame/fgame/core/module"
	lingyutemplate "fgame/fgame/game/lingyu/template"
)

//领域
type shenfaModule struct {
}

func (m *shenfaModule) InitTemplate() (err error) {
	err = lingyutemplate.Init()
	if err != nil {
		return
	}
	return
}
func (m *shenfaModule) Init() (err error) {

	return
}

func (m *shenfaModule) Start() {

}

func (m *shenfaModule) Stop() {

}

func (m *shenfaModule) String() string {
	return "lingyu"
}

var (
	m = &shenfaModule{}
)

func init() {
	module.Register(m)
}
