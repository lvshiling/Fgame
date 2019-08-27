package christmas

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/christmas/christmas"
	_ "fgame/fgame/game/christmas/event/listener"
	"fgame/fgame/game/christmas/template"
	"time"

	_ "fgame/fgame/game/christmas/handler"
)

//圣诞采集
type christmasModule struct {
	r runner.GoRunner
}

const (
	heartBeatTime = time.Second * 5
)

func (m *christmasModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *christmasModule) Init() (err error) {
	err = christmas.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("christmas", christmas.GetCharistmasService().Heartbeat, heartBeatTime)
	return
}

func (m *christmasModule) Start() {
	christmas.GetCharistmasService().Star()

	m.r.Start()
	return
}

func (m *christmasModule) Stop() {
	m.r.Stop()
}

func (m *christmasModule) String() string {
	return "christmas"
}

var (
	m = &christmasModule{}
)

func init() {
	module.Register(m)
}
