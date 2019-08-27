package player

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/player/event/listener"
	_ "fgame/fgame/cross/player/handler"
	"fgame/fgame/game/cd/template"
)

type playerModule struct {
}

func (m *playerModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *playerModule) Init() (err error) {

	return
}

func (m *playerModule) Start() {

}

func (m *playerModule) Stop() {

}

func (m *playerModule) String() string {
	return "player"
}

var (
	m = &playerModule{}
)

func init() {
	module.Register(m)
}
