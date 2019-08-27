package feedbackfee

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/feedbackfee/dao"
	"fgame/fgame/game/feedbackfee/feedbackfee"
	feedbackfeetemplate "fgame/fgame/game/feedbackfee/template"
	"fgame/fgame/game/global"
	"time"

	//注册
	_ "fgame/fgame/game/feedbackfee/event/listener"

	_ "fgame/fgame/game/feedbackfee/handler"

	_ "fgame/fgame/game/feedbackfee/player"
)

//逆付费
type feedbackfeeModule struct {
	r runner.GoRunner
}

func (m *feedbackfeeModule) InitTemplate() (err error) {
	err = feedbackfeetemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *feedbackfeeModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = feedbackfee.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("feedbackfee", feedbackfee.GetFeedbackFeeService().Heartbeat, time.Second*5)
	return
}

func (m *feedbackfeeModule) Start() {
	feedbackfee.GetFeedbackFeeService().Start()
	m.r.Start()
	return
}

func (m *feedbackfeeModule) Stop() {
	m.r.Stop()
	feedbackfee.GetFeedbackFeeService().Stop()
}

func (m *feedbackfeeModule) String() string {
	return "feedbackfee"
}

var (
	m = &feedbackfeeModule{}
)

func init() {
	module.Register(m)
}
