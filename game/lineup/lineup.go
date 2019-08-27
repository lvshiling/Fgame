package lineup

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/game/lineup/cross_handler"
	_ "fgame/fgame/game/lineup/handler"
)

type lineupModule struct {
}

func (m *lineupModule) InitTemplate() (err error) {

	return
}

func (m *lineupModule) Init() (err error) {
	return
}

func (m *lineupModule) Start() {
	return
}

func (m *lineupModule) Stop() {
}

func (m *lineupModule) String() string {
	return "lineup"
}

var (
	m = &lineupModule{}
)

func init() {
	module.Register(m)
}
