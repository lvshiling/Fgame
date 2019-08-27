package week

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/week/dao"
	weektemplate "fgame/fgame/game/week/template"
)

import (
	//注册
	_ "fgame/fgame/game/week/event/listener"
	_ "fgame/fgame/game/week/handler"
	_ "fgame/fgame/game/week/player"
)

//周卡
type weekModule struct {
}

func (m *weekModule) InitTemplate() (err error) {
	err = weektemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *weekModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *weekModule) Start() {
	return
}

func (m *weekModule) Stop() {
}

func (m *weekModule) String() string {
	return "week"
}

var (
	m = &weekModule{}
)

func init() {
	module.Register(m)
}
