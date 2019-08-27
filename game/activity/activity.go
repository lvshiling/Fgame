package activity

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/activity/activity"
	"fgame/fgame/game/activity/dao"
	activitytemplate "fgame/fgame/game/activity/template"
	"fgame/fgame/game/global"
	"time"
)

import (
	_ "fgame/fgame/game/activity/cross_handler"
	_ "fgame/fgame/game/activity/event/listener"

	_ "fgame/fgame/game/activity/guaji"
	_ "fgame/fgame/game/activity/handler"
	_ "fgame/fgame/game/activity/player"
)

var (
	heartbeatTime = time.Second
)

//活动大厅
type activeModule struct {
	r runner.GoRunner
}

func (m *activeModule) InitTemplate() (err error) {
	err = activitytemplate.Init()
	return
}

func (m *activeModule) Init() (err error) {
	rs := global.GetGame().GetRedisService()
	db := global.GetGame().GetDB()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}

	err = activity.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner(m.String(), m.heartbeat, heartbeatTime)
	return
}

func (m *activeModule) heartbeat() {
	activity.GetActivityService().Heartbeat()
}

func (m *activeModule) Start() {
	m.r.Start()
}

func (m *activeModule) Stop() {
	m.r.Stop()
}

func (acModule *activeModule) String() string {
	return "active"
}

var (
	m = &activeModule{}
)

func init() {
	module.Register(m)
}
