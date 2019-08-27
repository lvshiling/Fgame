package shenyu

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/shenyu/dao"
	"fgame/fgame/game/shenyu/template"
)

import (
	"fgame/fgame/game/shenyu/shenyu"
	//注册管理器
	_ "fgame/fgame/game/shenyu/activity_handler"
	_ "fgame/fgame/game/shenyu/check_enter"
	_ "fgame/fgame/game/shenyu/drop_handler"
	_ "fgame/fgame/game/shenyu/event/listener"
	_ "fgame/fgame/game/shenyu/foe_notice_handler"
	_ "fgame/fgame/game/shenyu/found_handler"
	_ "fgame/fgame/game/shenyu/handler"
)

//神域
type shenyuModule struct {
}

func (m *shenyuModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *shenyuModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	err = shenyu.Init()
	if err != nil {
		return
	}
	return
}

func (m *shenyuModule) Start() {
}

func (m *shenyuModule) Stop() {
}

func (m *shenyuModule) String() string {
	return "shenyu"
}

var (
	m = &shenyuModule{}
)

func init() {
	module.Register(m)
}
