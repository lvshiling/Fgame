package xinfa

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/xinfa/dao"
	"fgame/fgame/game/xinfa/xinfa"
)

import (
	//注册管理器
	_ "fgame/fgame/game/xinfa/event/listener"
	_ "fgame/fgame/game/xinfa/handler"
	_ "fgame/fgame/game/xinfa/player"
)

type xinfaModule struct {
}

func (m *xinfaModule) InitTemplate() (err error) {
	err = xinfa.Init()
	if err != nil {
		return
	}
	return
}

func (m *xinfaModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *xinfaModule) Start() {

}

func (m *xinfaModule) Stop() {

}

func (m *xinfaModule) String() string {
	return "xinfa"
}

var (
	m = &xinfaModule{}
)

func init() {
	module.Register(m)
}
