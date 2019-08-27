package tulong

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"

	"fgame/fgame/game/tulong/template"
	"fgame/fgame/game/tulong/tulong"
	"time"
)

import (
	_ "fgame/fgame/game/tulong/activity_handler"
	_ "fgame/fgame/game/tulong/cross_handler"
	_ "fgame/fgame/game/tulong/cross_loginflow"
	_ "fgame/fgame/game/tulong/event/listener"
	_ "fgame/fgame/game/tulong/found_handler"
	_ "fgame/fgame/game/tulong/handler"
)

//屠龙
type tuLongModule struct {
	gr runner.GoRunner
}

func (m *tuLongModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *tuLongModule) Init() (err error) {
	err = tulong.Init()
	if err != nil {
		return
	}
	m.gr = runner.NewGoRunner("tulong", tulong.GetTuLongService().Heartbeat, 3*time.Second)
	return
}

func (m *tuLongModule) Start() {
	tulong.GetTuLongService().Star()

	m.gr.Start()

}
func (m *tuLongModule) Stop() {
	m.gr.Stop()
}

func (m *tuLongModule) String() string {
	return "tulong"
}

var (
	m = &tuLongModule{}
)

func init() {
	module.Register(m)
}
