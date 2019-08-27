package qixue

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/qixue/dao"
	qixuetemplate "fgame/fgame/game/qixue/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/qixue/cross_handler"
	_ "fgame/fgame/game/qixue/drop_handler"
	_ "fgame/fgame/game/qixue/event/listener"
	_ "fgame/fgame/game/qixue/handler"
	_ "fgame/fgame/game/qixue/player"
)

//泣血枪
type qiXueModule struct {
}

func (m *qiXueModule) InitTemplate() (err error) {
	err = qixuetemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *qiXueModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *qiXueModule) Start() {
}

func (m *qiXueModule) Stop() {
}

func (m *qiXueModule) String() string {
	return "qixue"
}

var (
	m = &qiXueModule{}
)

func init() {
	module.Register(m)
}
