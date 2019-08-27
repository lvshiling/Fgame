package equipbaoku

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/equipbaoku/dao"
	"fgame/fgame/game/equipbaoku/equipbaoku"
	equipbaokutemplate "fgame/fgame/game/equipbaoku/template"
	"fgame/fgame/game/global"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/equipbaoku/drop_handler"
	_ "fgame/fgame/game/equipbaoku/event/listener"
	_ "fgame/fgame/game/equipbaoku/handler"
	_ "fgame/fgame/game/equipbaoku/player"
)

//装备宝库
type equipBaoKuModule struct {
	r runner.GoRunner
}

func (m *equipBaoKuModule) InitTemplate() (err error) {
	err = equipbaokutemplate.Init()
	return
}

func (m *equipBaoKuModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = equipbaoku.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("equipbaoku", equipbaoku.GetEquipBaoKuService().Heartbeat, 3*time.Second)
	return
}

func (m *equipBaoKuModule) Start() {
	m.r.Start()
}

func (m *equipBaoKuModule) Stop() {
	m.r.Stop()
}

func (m *equipBaoKuModule) String() string {
	return "equipbaoku"
}

var (
	m = &equipBaoKuModule{}
)

func init() {
	module.Register(m)
}
