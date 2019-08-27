package xianfu

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xianfu/dao"
	xianfutemplate "fgame/fgame/game/xianfu/template"
)

import (
	_ "fgame/fgame/game/xianfu/event/listener"
	_ "fgame/fgame/game/xianfu/found_handler"
	_ "fgame/fgame/game/xianfu/guaji"

	_ "fgame/fgame/game/xianfu/handler"
	_ "fgame/fgame/game/xianfu/player"
)

//秘境仙府
type xianfuModule struct {
}

func (m *xianfuModule) InitTemplate() (err error) {
	err = xianfutemplate.Init()
	if err != nil {
		return err
	}
	return
}
func (m *xianfuModule) Init() error {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err := dao.Init(ds, rs)
	if err != nil {
		return err
	}

	return err
}

func (m *xianfuModule) Start() {

}

func (m *xianfuModule) Stop() {

}

func (m *xianfuModule) String() string {
	return "xianfu"
}

var (
	m = &xianfuModule{}
)

func init() {
	module.Register(m)
}
