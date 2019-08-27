package itemskill

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/itemskill/dao"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/itemskill/event/listener"
	_ "fgame/fgame/game/itemskill/handler"
	_ "fgame/fgame/game/itemskill/player"
	_ "fgame/fgame/game/itemskill/use"
)

type itemskillModule struct {
}

func (m *itemskillModule) InitTemplate() (err error) {
	err = itemskilltemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *itemskillModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *itemskillModule) Start() {

}

func (m *itemskillModule) Stop() {

}

func (m *itemskillModule) String() string {
	return "itemskill"
}

var (
	m = &itemskillModule{}
)

func init() {
	module.Register(m)
}
