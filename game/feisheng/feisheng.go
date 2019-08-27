package feisheng

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/feisheng/dao"
	"fgame/fgame/game/feisheng/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/feisheng/drop_handler"
	_ "fgame/fgame/game/feisheng/event/listener"
	_ "fgame/fgame/game/feisheng/handler"
	_ "fgame/fgame/game/feisheng/player"
	_ "fgame/fgame/game/feisheng/use"
)

//飞升
type feishengModule struct {
}

func (m *feishengModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *feishengModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *feishengModule) Start() {

}

func (m *feishengModule) Stop() {

}

func (m *feishengModule) String() string {
	return "feisheng"
}

var (
	m = &feishengModule{}
)

func init() {
	module.Register(m)
}
