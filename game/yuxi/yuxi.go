package yuxi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/yuxi/dao"
	"fgame/fgame/game/yuxi/template"

	// _ "fgame/fgame/game/yuxi/found_handler"

	//注册管理器
	_ "fgame/fgame/game/yuxi/activity_handler"
	_ "fgame/fgame/game/yuxi/battle"
	_ "fgame/fgame/game/yuxi/event/listener"
	_ "fgame/fgame/game/yuxi/guaji"
	_ "fgame/fgame/game/yuxi/handler"
	_ "fgame/fgame/game/yuxi/player"
)

// 玉玺之战
type yuXiModule struct {
}

func (m *yuXiModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *yuXiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *yuXiModule) Start() {
}

func (m *yuXiModule) Stop() {

}

func (m *yuXiModule) String() string {
	return "yuxi"
}

var (
	m = &yuXiModule{}
)

func init() {
	module.Register(m)
}
