package huiyuan

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/huiyuan/dao"
	huiyuantemplate "fgame/fgame/game/huiyuan/template"
)

import (
	//注册
	_ "fgame/fgame/game/huiyuan/event/listener"
	_ "fgame/fgame/game/huiyuan/handler"
	_ "fgame/fgame/game/huiyuan/player"
)

//会员
type huiyuanModule struct {
}

func (m *huiyuanModule) InitTemplate() (err error) {
	err = huiyuantemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *huiyuanModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *huiyuanModule) Start() {
	return
}

func (m *huiyuanModule) Stop() {
}

func (m *huiyuanModule) String() string {
	return "huiyuan"
}

var (
	m = &huiyuanModule{}
)

func init() {
	module.Register(m)
}
