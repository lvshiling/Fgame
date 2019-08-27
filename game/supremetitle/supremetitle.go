package supremetitle

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/supremetitle/dao"
	"fgame/fgame/game/supremetitle/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/supremetitle/event/listener"
	_ "fgame/fgame/game/supremetitle/handler"
	_ "fgame/fgame/game/supremetitle/player"
	_ "fgame/fgame/game/supremetitle/use"
)

type supremeTitleModule struct {
}

func (m *supremeTitleModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *supremeTitleModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *supremeTitleModule) Start() {

}

func (m *supremeTitleModule) Stop() {

}

func (m *supremeTitleModule) String() string {
	return "suprenetitle"
}

var (
	m = &supremeTitleModule{}
)

func init() {
	module.Register(m)
}
