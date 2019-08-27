package xiantao

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xiantao/dao"
	"fgame/fgame/game/xiantao/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/xiantao/activity_handler"
	_ "fgame/fgame/game/xiantao/check_enter"
	_ "fgame/fgame/game/xiantao/collect_handler"
	_ "fgame/fgame/game/xiantao/drop_handler"
	_ "fgame/fgame/game/xiantao/event/listener"
	_ "fgame/fgame/game/xiantao/foe_notice_handler"
	_ "fgame/fgame/game/xiantao/found_handler"
	_ "fgame/fgame/game/xiantao/guaji"
	_ "fgame/fgame/game/xiantao/handler"
	_ "fgame/fgame/game/xiantao/player"
)

type xianTaoModule struct {
}

func (m *xianTaoModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *xianTaoModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *xianTaoModule) Start() {

}

func (m *xianTaoModule) Stop() {

}

func (m *xianTaoModule) String() string {
	return "xiantao"
}

var (
	m = &xianTaoModule{}
)

func init() {
	module.Register(m)
}
