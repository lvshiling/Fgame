package additionsys

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/additionsys/dao"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/additionsys/event/listener"
	_ "fgame/fgame/game/additionsys/handler"
	_ "fgame/fgame/game/additionsys/player"
)

//附加系统
type additionSysModule struct {
}

func (m *additionSysModule) InitTemplate() (err error) {
	additionsystemplate.Init()
	return
}

func (m *additionSysModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *additionSysModule) Start() {

}

func (m *additionSysModule) Stop() {

}

func (m *additionSysModule) String() string {
	return "additionsys"
}

var (
	m = &additionSysModule{}
)

func init() {
	module.Register(m)
}
