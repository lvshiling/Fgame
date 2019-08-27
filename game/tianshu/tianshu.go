package tianshu

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/tianshu/dao"
	"fgame/fgame/game/tianshu/template"

	//注册管理器
	_ "fgame/fgame/game/tianshu/event/listener"
	_ "fgame/fgame/game/tianshu/handler"
)

//天书
type tianshuModule struct {
}

func (m *tianshuModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *tianshuModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *tianshuModule) Start() {

}

func (m *tianshuModule) Stop() {

}

func (m *tianshuModule) String() string {
	return "tianshu"
}

var (
	m = &tianshuModule{}
)

func init() {
	module.Register(m)
}
