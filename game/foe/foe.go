package foe

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/foe/dao"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/foe/event/listener"
	_ "fgame/fgame/game/foe/for_notice_handler"
	_ "fgame/fgame/game/foe/handler"
	_ "fgame/fgame/game/foe/player"
)

//仇人
type foeModule struct {
}

func (m *foeModule) InitTemplate() (err error) {

	return
}
func (rm *foeModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (rm *foeModule) Start() {

}

func (rm *foeModule) String() string {
	return "foe"
}

func (rm *foeModule) Stop() {

}

var (
	m = &foeModule{}
)

func init() {
	module.Register(m)
}
