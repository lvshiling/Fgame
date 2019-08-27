package buff

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/buff/dao"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/buff/event/listener"
	_ "fgame/fgame/game/buff/player"
	_ "fgame/fgame/game/buff/player/event/listener"
)

//buff模块
type buffModule struct {
}

func (m *buffModule) InitTemplate() (err error) {
	err = bufftemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *buffModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *buffModule) Start() {

}

func (m *buffModule) Stop() {

}

func (m *buffModule) String() string {
	return "buff"
}

var (
	m = &buffModule{}
)

func init() {
	module.Register(m)
}
