package hongbao

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/hongbao/dao"
	"fgame/fgame/game/hongbao/hongbao"
	hongbaotemplate "fgame/fgame/game/hongbao/template"
	"time"

	//注册管理器
	_ "fgame/fgame/game/hongbao/event/listener"

	_ "fgame/fgame/game/hongbao/handler"

	_ "fgame/fgame/game/hongbao/player"
)

var (
	heartbeatTime = time.Minute
)

//红包
type hongBaoModule struct {
	r runner.GoRunner
}

func (m *hongBaoModule) InitTemplate() (err error) {
	err = hongbaotemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *hongBaoModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = hongbao.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner(m.String(), m.heartbeat, heartbeatTime)
	return
}

func (m *hongBaoModule) heartbeat() {
	hongbao.GetHongBaoService().Heartbeat()
}

func (m *hongBaoModule) Start() {
	m.r.Start()
}

func (m *hongBaoModule) Stop() {
	m.r.Stop()
}

func (m *hongBaoModule) String() string {
	return "hongbao"
}

var (
	m = &hongBaoModule{}
)

func init() {
	module.Register(m)
}
