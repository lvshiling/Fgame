package lineup

import (
	"fgame/fgame/core/module"
	"fgame/fgame/cross/lineup/lineup"

	_ "fgame/fgame/cross/lineup/cross_handler"
	_ "fgame/fgame/cross/lineup/event/listener"
	_ "fgame/fgame/cross/lineup/handler"
)

type lineupModule struct {
}

func (m *lineupModule) InitTemplate() (err error) {

	return
}
func (m *lineupModule) Init() (err error) {
	err = lineup.Init()
	if err != nil {
		return
	}
	return
}

func (m *lineupModule) Start() {

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
