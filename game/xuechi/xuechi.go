package xuechi

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xuechi/dao"
)

import (
	//注册管理器
	_ "fgame/fgame/game/xuechi/cross_handler"
	_ "fgame/fgame/game/xuechi/event/listener"
	_ "fgame/fgame/game/xuechi/guaji"
	_ "fgame/fgame/game/xuechi/handler"
	_ "fgame/fgame/game/xuechi/player"
	_ "fgame/fgame/game/xuechi/use"
)

type xuechiModule struct {
}

func (m *xuechiModule) InitTemplate() (err error) {

	return
}

func (m *xuechiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *xuechiModule) Start() {

}

func (m *xuechiModule) Stop() {

}

func (m *xuechiModule) String() string {
	return "xuechi"
}

var (
	m = &xuechiModule{}
)

func init() {
	module.Register(m)
}
