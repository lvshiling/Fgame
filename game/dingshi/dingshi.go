package dingshi

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/dingshi/dao"
	"fgame/fgame/game/dingshi/dingshi"
	dingshitemplate "fgame/fgame/game/dingshi/template"
	"fgame/fgame/game/global"

	//注册
	_ "fgame/fgame/game/dingshi/boss_handler"

	_ "fgame/fgame/game/dingshi/check_enter"

	_ "fgame/fgame/game/dingshi/dingshi"

	_ "fgame/fgame/game/dingshi/event/listener"
)

type dingshiModule struct {
}

func (m *dingshiModule) InitTemplate() (err error) {
	err = dingshitemplate.Init()
	if err != nil {
		return
	}

	return
}
func (m *dingshiModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	err = dingshi.Init()
	if err != nil {
		return
	}
	return
}

func (m *dingshiModule) Start() {
	dingshi.GetDingShiService().Start()
	return
}

func (m *dingshiModule) Stop() {

}

func (m *dingshiModule) String() string {
	return "dingshi"
}

var (
	m = &dingshiModule{}
)

func init() {
	module.Register(m)
}
