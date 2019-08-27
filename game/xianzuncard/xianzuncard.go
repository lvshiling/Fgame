package xianzunprivilege

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xianzuncard/dao"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
)

import (
	_ "fgame/fgame/game/xianzuncard/event/listener"
	_ "fgame/fgame/game/xianzuncard/handler"
	_ "fgame/fgame/game/xianzuncard/player"
)

// 仙尊特权卡
type xianZunCardModule struct {
}

func (m *xianZunCardModule) InitTemplate() (err error) {
	err = xianzuncardtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *xianZunCardModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *xianZunCardModule) Start() {
}

func (m *xianZunCardModule) Stop() {
}

func (m *xianZunCardModule) String() string {
	return "xianzuncard"
}

var (
	m = &xianZunCardModule{}
)

func init() {
	module.Register(m)
}
