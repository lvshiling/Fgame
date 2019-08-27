package arenaboss

import (
	//注册
	"fgame/fgame/core/module"
	"fgame/fgame/cross/arenaboss/arenaboss"
	_ "fgame/fgame/cross/arenaboss/boss_handler"
	"fgame/fgame/cross/arenaboss/dao"
	_ "fgame/fgame/cross/arenaboss/event/listener"
	_ "fgame/fgame/cross/arenaboss/login_handler"
	_ "fgame/fgame/cross/arenaboss/relive_handler"
	"fgame/fgame/game/global"
)

type arenaBossModule struct {
}

func (m *arenaBossModule) InitTemplate() (err error) {

	return
}

func (m *arenaBossModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}

	err = arenaboss.Init()
	if err != nil {
		return
	}

	return
}

func (m *arenaBossModule) Start() {
	arenaboss.GetArenaBossService().Start()

}

func (m *arenaBossModule) Stop() {

	arenaboss.GetArenaBossService().Stop()
}

func (m *arenaBossModule) String() string {
	return "arenaboss"
}

var (
	m = &arenaBossModule{}
)

func init() {
	module.Register(m)
}
