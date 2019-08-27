package title

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/title/dao"
	"fgame/fgame/game/title/title"
)

import (
	//注册管理器
	_ "fgame/fgame/game/title/active_check"
	_ "fgame/fgame/game/title/event/listener"
	_ "fgame/fgame/game/title/handler"
	_ "fgame/fgame/game/title/player"
)

type titleModule struct {
}

func (m *titleModule) InitTemplate() (err error) {
	err = title.Init()
	if err != nil {
		return
	}
	return
}

func (m *titleModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *titleModule) Start() {

}

func (m *titleModule) Stop() {

}

func (m *titleModule) String() string {
	return "title"
}

var (
	m = &titleModule{}
)

func init() {
	module.Register(m)
}
