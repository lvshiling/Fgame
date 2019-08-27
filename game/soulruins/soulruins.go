package soulruins

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/soulruins/dao"
	"fgame/fgame/game/soulruins/soulruins"
)

import (
	//注册管理器
	_ "fgame/fgame/game/soulruins/event/listener"
	_ "fgame/fgame/game/soulruins/found_handler"
	_ "fgame/fgame/game/soulruins/guaji_check"
	_ "fgame/fgame/game/soulruins/handler"
	_ "fgame/fgame/game/soulruins/player"
)

//帝魂遗迹
type soulRuinsModule struct {
}

func (m *soulRuinsModule) InitTemplate() (err error) {
	err = soulruins.Init()
	if err != nil {
		return
	}
	return
}
func (m *soulRuinsModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *soulRuinsModule) Start() {

}

func (m *soulRuinsModule) Stop() {

}

func (m *soulRuinsModule) String() string {
	return "soulruins"
}

var (
	m = &soulRuinsModule{}
)

func init() {
	module.Register(m)
}
