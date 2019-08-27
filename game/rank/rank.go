package rank

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/rank/dao"
	"fgame/fgame/game/rank/rank"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/rank/event/listener"
	_ "fgame/fgame/game/rank/handler"
)

type rankModule struct {
	r runner.GoRunner
}

func (m *rankModule) InitTemplate() (err error) {
	return
}

func (m *rankModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = rank.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("rank", rank.GetRankService().Heartbeat, time.Minute)

	return
}

func (m *rankModule) Start() {
	rank.GetRankService().Star()

	m.r.Start()

}

func (m *rankModule) Stop() {
	m.r.Stop()
}

func (m *rankModule) String() string {
	return "rank"
}

var (
	m = &rankModule{}
)

func init() {
	module.RegisterBase(m)
}
