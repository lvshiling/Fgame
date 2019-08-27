package mingge

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/mingge/dao"
	"fgame/fgame/game/mingge/template"

	//注册管理器
	_ "fgame/fgame/game/mingge/event/listener"

	_ "fgame/fgame/game/mingge/handler"

	_ "fgame/fgame/game/mingge/player"
)

//命格
type mingGeModule struct {
}

func (m *mingGeModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return err
	}
	return
}
func (m *mingGeModule) Init() error {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err := dao.Init(ds, rs)
	if err != nil {
		return err
	}

	return err
}

func (m *mingGeModule) Start() {

}

func (m *mingGeModule) Stop() {

}

func (m *mingGeModule) String() string {
	return "mingge"
}

var (
	m = &mingGeModule{}
)

func init() {
	module.Register(m)
}
