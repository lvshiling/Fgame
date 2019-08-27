package welfare

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/welfare/dao"
	"fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"
	"time"
)

import (
	//注册
	_ "fgame/fgame/game/welfare/event/listener"
	_ "fgame/fgame/game/welfare/handler"
	_ "fgame/fgame/game/welfare/player"
	_ "fgame/fgame/game/welfare/welfare"
)

import (
	//注册
	_ "fgame/fgame/game/welfare/advanced"
	_ "fgame/fgame/game/welfare/advancedrew"
	_ "fgame/fgame/game/welfare/alliance"
	_ "fgame/fgame/game/welfare/boat_race"
	_ "fgame/fgame/game/welfare/cycle"
	_ "fgame/fgame/game/welfare/develop"
	_ "fgame/fgame/game/welfare/discount"
	_ "fgame/fgame/game/welfare/drew"
	_ "fgame/fgame/game/welfare/feedback"
	_ "fgame/fgame/game/welfare/group"
	_ "fgame/fgame/game/welfare/hall"
	_ "fgame/fgame/game/welfare/huhu"
	_ "fgame/fgame/game/welfare/invest"
	_ "fgame/fgame/game/welfare/longfeng"
	_ "fgame/fgame/game/welfare/made"
	_ "fgame/fgame/game/welfare/rank"
	_ "fgame/fgame/game/welfare/rewards"
	_ "fgame/fgame/game/welfare/shopdiscount"
	_ "fgame/fgame/game/welfare/system"
	_ "fgame/fgame/game/welfare/tongtianta"
	_ "fgame/fgame/game/welfare/xiuxianbook"
)

import (
	_ "fgame/fgame/game/welfare/invest"
)

//福利厅
type welfareModule struct {
	r runner.GoRunner
}

const (
	heartBeatTime = time.Second * 5
)

func (m *welfareModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *welfareModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = welfare.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("welfare", welfare.GetWelfareService().Heartbeat, heartBeatTime)
	return
}

func (m *welfareModule) Start() {
	welfare.GetWelfareService().Star()

	m.r.Start()
	return
}

func (m *welfareModule) Stop() {
	m.r.Stop()
}

func (m *welfareModule) String() string {
	return "welfare"
}

var (
	m = &welfareModule{}
)

func init() {
	module.Register(m)
}
