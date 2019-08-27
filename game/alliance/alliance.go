package alliance

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/alliance/alliance"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"time"
)

import (
	_ "fgame/fgame/game/alliance/activity_handler"
	_ "fgame/fgame/game/alliance/battle"
	_ "fgame/fgame/game/alliance/collect_handler"
	_ "fgame/fgame/game/alliance/drop_handler"
	_ "fgame/fgame/game/alliance/event/listener"
	_ "fgame/fgame/game/alliance/found_handler"
	_ "fgame/fgame/game/alliance/guaji"
	_ "fgame/fgame/game/alliance/handler"
	_ "fgame/fgame/game/alliance/npc/check_attack"
	_ "fgame/fgame/game/alliance/player"
	_ "fgame/fgame/game/alliance/use"
)

type allianceModule struct {
	r runner.GoRunner
}

func (m *allianceModule) InitTemplate() (err error) {
	err = alliancetemplate.Init()
	if err != nil {
		return
	}

	return
}

var (
	heartbeatTime = time.Second
)

func (m *allianceModule) Init() (err error) {
	err = alliance.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner(m.String(), m.heartbeat, heartbeatTime)
	return
}

func (m *allianceModule) heartbeat() {
	alliance.GetAllianceService().Heartbeat()
}

func (m *allianceModule) Start() {
	alliance.GetAllianceService().Start()
	m.r.Start()

}

func (m *allianceModule) Stop() {
	m.r.Stop()
}

func (m *allianceModule) String() string {
	return "alliance"
}

var (
	m = &allianceModule{}
)

func init() {
	module.Register(m)
}
