package dianxing

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/dianxing/dao"
	dianxingtemplate "fgame/fgame/game/dianxing/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/dianxing/drop_handler"
	_ "fgame/fgame/game/dianxing/event/listener"
	_ "fgame/fgame/game/dianxing/handler"
	_ "fgame/fgame/game/dianxing/player"
)

//点星系统
type dianxingModule struct {
}

func (m *dianxingModule) InitTemplate() (err error) {
	err = dianxingtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *dianxingModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *dianxingModule) Start() {
}

func (m *dianxingModule) Stop() {
}

func (m *dianxingModule) String() string {
	return "dianxing"
}

var (
	m = &dianxingModule{}
)

func init() {
	module.Register(m)
}
