package systemskill

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/systemskill/dao"
	systemskilltemplate "fgame/fgame/game/systemskill/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/systemskill/event/listener"
	_ "fgame/fgame/game/systemskill/handler"
	_ "fgame/fgame/game/systemskill/player"
)

type systemskillModule struct {
}

func (m *systemskillModule) InitTemplate() (err error) {
	err = systemskilltemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *systemskillModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *systemskillModule) Start() {

}

func (m *systemskillModule) Stop() {

}

func (m *systemskillModule) String() string {
	return "systemskill"
}

var (
	m = &systemskillModule{}
)

func init() {
	module.Register(m)
}
