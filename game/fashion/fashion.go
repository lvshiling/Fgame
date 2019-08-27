package fashion

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/fashion/dao"
	"fgame/fgame/game/fashion/fashion"
	"fgame/fgame/game/global"
)
import (
	//注册管理器
	_ "fgame/fgame/game/fashion/active_check"
	_ "fgame/fgame/game/fashion/event/listener"
	_ "fgame/fgame/game/fashion/handler"
	_ "fgame/fgame/game/fashion/player"
	_ "fgame/fgame/game/fashion/use"
)

type fashionModule struct {
}

func (m *fashionModule) InitTemplate() (err error) {
	err = fashion.Init()
	if err != nil {
		return
	}
	return
}
func (m *fashionModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *fashionModule) Start() {

}

func (m *fashionModule) Stop() {

}

func (m *fashionModule) String() string {
	return "fashion"
}

var (
	m = &fashionModule{}
)

func init() {
	module.Register(m)
}
