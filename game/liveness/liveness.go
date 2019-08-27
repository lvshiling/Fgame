package template

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/global"
	"fgame/fgame/game/liveness/dao"
	"fgame/fgame/game/liveness/template"

	//注册管理器
	_ "fgame/fgame/game/liveness/event/listener"
	_ "fgame/fgame/game/liveness/handler"
	_ "fgame/fgame/game/liveness/player"
)

type livenessModule struct {
}

func (m *livenessModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *livenessModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *livenessModule) Start() {

}

func (m *livenessModule) Stop() {

}

func (cm *livenessModule) String() string {
	return "liveness"
}

var (
	m = &livenessModule{}
)

func init() {
	module.Register(m)
}
