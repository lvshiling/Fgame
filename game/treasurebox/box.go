package treasurebox

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"

	"fgame/fgame/game/treasurebox/template"
	"fgame/fgame/game/treasurebox/treasurebox"

	"time"
)

import (
	_ "fgame/fgame/game/treasurebox/handler"
	_ "fgame/fgame/game/treasurebox/use"
)

//宝箱
type boxModule struct {
	r runner.GoRunner
}

func (m *boxModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *boxModule) Init() (err error) {
	err = treasurebox.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("treasurebox", treasurebox.GetTreasureBoxService().Heartbeat, time.Minute)
	return
}

func (m *boxModule) Start()  {
	treasurebox.GetTreasureBoxService().Star()

	m.r.Start()

}

func (m *boxModule) Stop() {
	m.r.Stop()

}

func (m *boxModule) String() string {
	return "box"
}

var (
	m = &boxModule{}
)

func init() {
	module.Register(m)
}
