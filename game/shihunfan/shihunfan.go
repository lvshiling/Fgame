package shihunfan

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/shihunfan/dao"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/shihunfan/event/listener"
	_ "fgame/fgame/game/shihunfan/guaji"
	_ "fgame/fgame/game/shihunfan/handler"
	_ "fgame/fgame/game/shihunfan/player"
	_ "fgame/fgame/game/shihunfan/systemskill_handler"
	_ "fgame/fgame/game/shihunfan/use"
)

//噬魂幡
type shihunfanModule struct {
}

func (m *shihunfanModule) InitTemplate() (err error) {
	err = shihunfantemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *shihunfanModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *shihunfanModule) Start() {
}

func (m *shihunfanModule) Stop() {
}

func (m *shihunfanModule) String() string {
	return "shihunfan"
}

var (
	m = &shihunfanModule{}
)

func init() {
	module.Register(m)
}
