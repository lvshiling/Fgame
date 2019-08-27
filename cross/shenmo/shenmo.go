package shenmo

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/cross/shenmo/dao"
	crossshenmo "fgame/fgame/cross/shenmo/shenmo"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	gameshenmo "fgame/fgame/game/shenmo/shenmo"
	"fgame/fgame/game/shenmo/template"

	//注册管理器
	centertypes "fgame/fgame/center/types"
	_ "fgame/fgame/cross/shenmo/cross_handler"
	_ "fgame/fgame/cross/shenmo/event/listener"
	_ "fgame/fgame/cross/shenmo/handler"
	_ "fgame/fgame/cross/shenmo/login_handler"

	//
	_ "fgame/fgame/game/shenmo/event/listener/common"
	_ "fgame/fgame/game/shenmo/relive_handler"
	"time"
)

type shenMoModule struct {
	r runner.GoRunner
}

func (m *shenMoModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *shenMoModule) Init() (err error) {
	if global.GetGame().GetServerType() != centertypes.GameServerTypePlatform {
		return
	}
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	err = crossshenmo.Init()
	if err != nil {
		return
	}
	err = gameshenmo.Init(activitytypes.ActivityTypeShenMoWar)
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("crossshenmo", ShenMoHeartbeat, 3*time.Second)
	return
}

func ShenMoHeartbeat() {
	crossshenmo.GetShenMoService().Heartbeat()
	gameshenmo.GetShenMoService().Heartbeat()
}

func (m *shenMoModule) Start() {
	if global.GetGame().GetServerType() != centertypes.GameServerTypePlatform {
		return
	}
	m.r.Start()

}

func (m *shenMoModule) Stop() {
	if global.GetGame().GetServerType() != centertypes.GameServerTypePlatform {
		return
	}
	m.r.Stop()
}

func (m *shenMoModule) String() string {
	return "shenmo"
}

var (
	m = &shenMoModule{}
)

func init() {
	module.Register(m)
}
