package xuedun

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xuedun/dao"
	"fgame/fgame/game/xuedun/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/xuedun/event/listener"
	_ "fgame/fgame/game/xuedun/handler"
	_ "fgame/fgame/game/xuedun/player"
)

type xueDunModule struct {
}

func (cm *xueDunModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (cm *xueDunModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *xueDunModule) Start() {

}

func (cm *xueDunModule) Stop() {

}

func (cm *xueDunModule) String() string {
	return "xuedun"
}

var (
	m = &xueDunModule{}
)

func init() {
	module.Register(m)
}
