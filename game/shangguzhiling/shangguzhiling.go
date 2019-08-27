package shangguzhiling

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/shangguzhiling/dao"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/shangguzhiling/event/listener"
	_ "fgame/fgame/game/shangguzhiling/handler"
	_ "fgame/fgame/game/shangguzhiling/player"
)

type shangguzhilingModule struct {
}

func (m *shangguzhilingModule) InitTemplate() (err error) {
	err = shangguzhilingtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *shangguzhilingModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *shangguzhilingModule) Start() {

}

func (m *shangguzhilingModule) Stop() {

}

func (m *shangguzhilingModule) String() string {
	return "shangguzhiling"
}

var (
	m = &shangguzhilingModule{}
)

func init() {
	module.Register(m)
}
