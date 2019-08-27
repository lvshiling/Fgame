package baby

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/baby/dao"
	"fgame/fgame/game/baby/template"
	"fgame/fgame/game/global"
)

import (
	"fgame/fgame/game/baby/baby"
	//注册管理器
	_ "fgame/fgame/game/baby/event/listener"
	_ "fgame/fgame/game/baby/handler"
	_ "fgame/fgame/game/baby/player"
	_ "fgame/fgame/game/baby/use"
)

//宝宝
type babyModule struct {
}

func (m *babyModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}

	return
}

func (m *babyModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = baby.Init()
	if err != nil {
		return
	}

	return
}

func (m *babyModule) Start() {

}

func (m *babyModule) Stop() {

}

func (m *babyModule) String() string {
	return "baby"
}

var (
	m = &babyModule{}
)

func init() {
	module.Register(m)
}
