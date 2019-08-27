package gem

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/gem/dao"
	"fgame/fgame/game/gem/gem"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/gem/drop_handler"
	_ "fgame/fgame/game/gem/event/listener"
	_ "fgame/fgame/game/gem/handler"
	_ "fgame/fgame/game/gem/player"
)

type gemModule struct {
}

func (m *gemModule) InitTemplate() (err error) {
	err = gem.Init()
	if err != nil {
		return
	}
	return
}

func (m *gemModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *gemModule) Start() {

}

func (m *gemModule) Stop() {

}

func (m *gemModule) String() string {
	return "gem"
}

var (
	m = &gemModule{}
)

func init() {
	module.Register(m)
}
