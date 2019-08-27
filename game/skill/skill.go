package skill

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/skill/dao"
	skilltemplate "fgame/fgame/game/skill/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/skill/event/listener"
	_ "fgame/fgame/game/skill/guaji"
	_ "fgame/fgame/game/skill/guaji_check"
	_ "fgame/fgame/game/skill/handler"
	_ "fgame/fgame/game/skill/player"
	_ "fgame/fgame/game/skill/use"
)

type skillModule struct {
}

func (m *skillModule) InitTemplate() (err error) {
	err = skilltemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *skillModule) Init() (err error) {

	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *skillModule) Start() {

}
func (m *skillModule) Stop() {

}

func (m *skillModule) String() string {
	return "skill"
}

var (
	m = &skillModule{}
)

func init() {
	module.Register(m)
}
