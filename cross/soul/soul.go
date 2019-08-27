package soul

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/soul/soul"
)

type soulModule struct {
}

func (m *soulModule) InitTemplate() (err error) {
	err = soul.Init()
	if err != nil {
		return
	}

	return
}
func (m *soulModule) Init() (err error) {
	return
}

func (m *soulModule) Start() {

}

func (m *soulModule) Stop() {

}

func (m *soulModule) String() string {
	return "soul"
}

var (
	m = &soulModule{}
)

func init() {
	module.Register(m)
}
