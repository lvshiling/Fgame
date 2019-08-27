package secretcard

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/global"
	"fgame/fgame/game/secretcard/dao"
	"fgame/fgame/game/secretcard/secretcard"
)

import (
	//注册管理器
	_ "fgame/fgame/game/secretcard/event/listener"
	_ "fgame/fgame/game/secretcard/handler"
	_ "fgame/fgame/game/secretcard/player"
)

type secretModule struct {
}

func (m *secretModule) InitTemplate() (err error) {
	err = secretcard.Init()
	if err != nil {
		return
	}
	return
}

func (m *secretModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *secretModule) Start() {

}

func (m *secretModule) Stop() {

}

func (cm *secretModule) String() string {
	return "secretcard"
}

var (
	m = &secretModule{}
)

func init() {
	module.Register(m)
}
