package shenfa

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/shenfa/dao"
	shenfatemplate "fgame/fgame/game/shenfa/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/shenfa/active_check"
	_ "fgame/fgame/game/shenfa/event/listener"
	_ "fgame/fgame/game/shenfa/guaji"
	_ "fgame/fgame/game/shenfa/handler"
	_ "fgame/fgame/game/shenfa/player"
	_ "fgame/fgame/game/shenfa/systemskill_handler"
	_ "fgame/fgame/game/shenfa/use"
)

//身法
type shenfaModule struct {
}

func (m *shenfaModule) InitTemplate() (err error) {
	err = shenfatemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *shenfaModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *shenfaModule) Start() {

}

func (m *shenfaModule) Stop() {

}

func (m *shenfaModule) String() string {
	return "shenfa"
}

var (
	m = &shenfaModule{}
)

func init() {
	module.Register(m)
}
