package daliwan

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/daliwan/dao"
	daliwantemplate "fgame/fgame/game/daliwan/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/daliwan/event/listener"

	_ "fgame/fgame/game/daliwan/player"
	_ "fgame/fgame/game/daliwan/use"
)

type daLiWanModule struct {
}

func (m *daLiWanModule) InitTemplate() (err error) {
	err = daliwantemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *daLiWanModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *daLiWanModule) Start() {
}

func (m *daLiWanModule) Stop() {
}

func (m *daLiWanModule) String() string {
	return "daliwan"
}

var (
	m = &daLiWanModule{}
)

func init() {
	module.Register(m)
}
