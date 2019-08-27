package treasurebox

import (
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/cross/treasurebox/dao"
	"fgame/fgame/cross/treasurebox/treasurebox"
	"fgame/fgame/game/global"
	"fgame/fgame/game/treasurebox/template"
	"time"
)

type treasureBoxModule struct {
	r runner.GoRunner
}

func (m *treasureBoxModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *treasureBoxModule) Init() (err error) {
	if global.GetGame().GetServerType() != centertypes.GameServerTypePlatform {
		return
	}
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = treasurebox.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("treasurebox", treasurebox.GetTreasureBoxService().Heartbeat, 2*time.Second)
	return
}

func (m *treasureBoxModule) Start() {
	if global.GetGame().GetServerType() != centertypes.GameServerTypePlatform {
		return
	}
	m.r.Start()

}

func (m *treasureBoxModule) Stop() {
	if global.GetGame().GetServerType() != centertypes.GameServerTypePlatform {
		return
	}
	m.r.Stop()
}

func (m *treasureBoxModule) String() string {
	return "treasureBox"
}

var (
	m = &treasureBoxModule{}
)

func init() {
	module.Register(m)
}
