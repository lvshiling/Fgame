package vip

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/vip/dao"
	"fgame/fgame/game/vip/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/vip/event/listener"
	_ "fgame/fgame/game/vip/handler"
	_ "fgame/fgame/game/vip/player"
)

//VIP
type vipModule struct {
}

func (m *vipModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *vipModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *vipModule) Start() {

}

func (m *vipModule) Stop() {

}

func (m *vipModule) String() string {
	return "vip"
}

var (
	m = &vipModule{}
)

func init() {
	module.Register(m)
}
