package massacre

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/massacre/dao"
	massacretemplate "fgame/fgame/game/massacre/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/massacre/cross_handler"
	_ "fgame/fgame/game/massacre/drop_handler"
	_ "fgame/fgame/game/massacre/event/listener"
	_ "fgame/fgame/game/massacre/guaji"
	_ "fgame/fgame/game/massacre/handler"
	_ "fgame/fgame/game/massacre/player"
)

//戮仙刃
type massacreModule struct {
}

func (m *massacreModule) InitTemplate() (err error) {
	err = massacretemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *massacreModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *massacreModule) Start() {
}

func (m *massacreModule) Stop() {
}

func (m *massacreModule) String() string {
	return "massacre"
}

var (
	m = &massacreModule{}
)

func init() {
	module.Register(m)
}
