package zhenfa

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/zhenfa/dao"
	"fgame/fgame/game/zhenfa/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/zhenfa/event/listener"
	_ "fgame/fgame/game/zhenfa/handler"
	_ "fgame/fgame/game/zhenfa/player"
)

//阵法
type zhenFaModule struct {
}

func (m *zhenFaModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return err
	}
	return
}

func (m *zhenFaModule) Init() error {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err := dao.Init(ds, rs)
	if err != nil {
		return err
	}

	return err
}

func (m *zhenFaModule) Start() {

}

func (m *zhenFaModule) Stop() {

}

func (m *zhenFaModule) String() string {
	return "zhenfa"
}

var (
	m = &zhenFaModule{}
)

func init() {
	module.Register(m)
}
