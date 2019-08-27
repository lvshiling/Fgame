package lingtongdev

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtongdev/dao"
	"fgame/fgame/game/lingtongdev/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/lingtongdev/event/listener"
	_ "fgame/fgame/game/lingtongdev/guaji"
	_ "fgame/fgame/game/lingtongdev/handler"
	_ "fgame/fgame/game/lingtongdev/player"
	_ "fgame/fgame/game/lingtongdev/systemskill_handler"
	_ "fgame/fgame/game/lingtongdev/use"
)

type lingTongDevModule struct {
}

func (m *lingTongDevModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *lingTongDevModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *lingTongDevModule) Start() {

}

func (m *lingTongDevModule) Stop() {

}

func (m *lingTongDevModule) String() string {
	return "lingTongDev"
}

var (
	m = &lingTongDevModule{}
)

func init() {
	module.Register(m)
}
