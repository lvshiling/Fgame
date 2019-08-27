package material

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/material/dao"
	materialtemplate "fgame/fgame/game/material/template"
)

import (
	_ "fgame/fgame/game/material/event/listener"
	_ "fgame/fgame/game/material/found_handler"
	_ "fgame/fgame/game/material/guaji"
	_ "fgame/fgame/game/material/handler"
	_ "fgame/fgame/game/material/player"
)

//材料副本
type materialModule struct {
}

func (m *materialModule) InitTemplate() (err error) {
	err = materialtemplate.Init()
	if err != nil {
		return err
	}
	return
}
func (m *materialModule) Init() error {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err := dao.Init(ds, rs)
	if err != nil {
		return err
	}

	return err
}

func (m *materialModule) Start() {

}

func (m *materialModule) Stop() {

}

func (m *materialModule) String() string {
	return "material"
}

var (
	m = &materialModule{}
)

func init() {
	module.Register(m)
}
