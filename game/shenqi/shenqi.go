package shenqi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/shenqi/dao"
	shenqitemplate "fgame/fgame/game/shenqi/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/shenqi/drop_handler"
	_ "fgame/fgame/game/shenqi/event/listener"
	_ "fgame/fgame/game/shenqi/handler"
	_ "fgame/fgame/game/shenqi/player"
)

type shenQiModule struct {
}

func (m *shenQiModule) InitTemplate() (err error) {
	err = shenqitemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *shenQiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *shenQiModule) Start() {

}

func (m *shenQiModule) Stop() {

}

func (m *shenQiModule) String() string {
	return "shenqi"
}

var (
	m = &shenQiModule{}
)

func init() {
	module.Register(m)
}
