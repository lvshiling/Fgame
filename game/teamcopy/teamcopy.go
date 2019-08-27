package teamcopy

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/global"
	"fgame/fgame/game/teamcopy/dao"
	"fgame/fgame/game/teamcopy/template"

	//注册管理
	_ "fgame/fgame/game/teamcopy/cross_handler"
	_ "fgame/fgame/game/teamcopy/cross_loginflow"
	_ "fgame/fgame/game/teamcopy/event/listener"
	_ "fgame/fgame/game/teamcopy/found_handler"
	_ "fgame/fgame/game/teamcopy/handler"
)

//组队副本
type teamCoypModule struct {
}

func (m *teamCoypModule) InitTemplate() (err error) {
	err = template.Init()
	return
}

func (m *teamCoypModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *teamCoypModule) Start() {

}
func (m *teamCoypModule) Stop() {

}

func (m *teamCoypModule) String() string {
	return "teamcopy"
}

var (
	m = &teamCoypModule{}
)

func init() {
	module.Register(m)
}
