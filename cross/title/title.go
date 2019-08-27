package title

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/title/title"
)

type titleModule struct {
}

func (m *titleModule) InitTemplate() (err error) {
	err = title.Init()
	if err != nil {
		return
	}
	return
}

func (m *titleModule) Init() (err error) {

	return
}
func (m *titleModule) Start() {

}

func (m *titleModule) Stop() {

}

func (m *titleModule) String() string {
	return "title"
}

var (
	m = &titleModule{}
)

func init() {
	module.Register(m)
}
