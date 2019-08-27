package emperor

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/emperor/dao"
	"fgame/fgame/game/emperor/emperor"
	emperortemplate "fgame/fgame/game/emperor/template"
	"fgame/fgame/game/global"
	"time"
)
import (
	//注册管理器
	_ "fgame/fgame/game/emperor/event/listener"
	_ "fgame/fgame/game/emperor/handler"
	_ "fgame/fgame/game/emperor/player"
)

type emperorModule struct {
	r runner.GoRunner
}

func (m *emperorModule) InitTemplate() (err error) {
	err = emperortemplate.Init()
	if err != nil {
		return
	}
	return
}

const (
	emperorTimer = 3 * time.Second
)

func (m *emperorModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = emperor.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("emperor", emperor.GetEmperorService().Heartbeat, emperorTimer)
	return
}

func (m *emperorModule) Start() {
	m.r.Start()

}

func (m *emperorModule) Stop() {
	m.r.Stop()
}

func (m *emperorModule) String() string {
	return "emperor"
}

var (
	m = &emperorModule{}
)

func init() {
	module.Register(m)
}
