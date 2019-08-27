package fashion

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/fashion/fashion"
)

type fashionModule struct {
}

func (m *fashionModule) InitTemplate() (err error) {
	err = fashion.Init()
	if err != nil {
		return
	}
	return
}
func (m *fashionModule) Init() (err error) {

	return
}

func (m *fashionModule) Start() {

}

func (m *fashionModule) Stop() {

}

func (m *fashionModule) String() string {
	return "fashion"
}

var (
	m = &fashionModule{}
)

func init() {
	module.Register(m)
}
