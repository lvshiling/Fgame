package tulong

import (
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/cross/tulong/dao"
	"fgame/fgame/cross/tulong/tulong"
	"fgame/fgame/game/global"
	tulongtemplate "fgame/fgame/game/tulong/template"
	"time"

	_ "fgame/fgame/cross/tulong/cross_handler"
	_ "fgame/fgame/cross/tulong/event/listener"
	_ "fgame/fgame/cross/tulong/guaji"
	_ "fgame/fgame/cross/tulong/handler"
	_ "fgame/fgame/cross/tulong/login_handler"
	_ "fgame/fgame/cross/tulong/relive_handler"
)

type tulongModule struct {
	r runner.GoRunner
}

func (m *tulongModule) InitTemplate() (err error) {
	err = tulongtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *tulongModule) Init() (err error) {
	if global.GetGame().GetServerType() != centertypes.GameServerTypeRegion {
		return
	}

	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	err = tulong.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("tulong", tulong.GetTuLongService().Heartbeat, 3*time.Second)
	return
}

func (m *tulongModule) Start() {
	if global.GetGame().GetServerType() != centertypes.GameServerTypeRegion {
		return
	}
	m.r.Start()

}

func (m *tulongModule) Stop() {
	if global.GetGame().GetServerType() != centertypes.GameServerTypeRegion {
		return
	}
	m.r.Stop()
}

func (m *tulongModule) String() string {
	return "tulong"
}

var (
	m = &tulongModule{}
)

func init() {
	module.Register(m)
}
