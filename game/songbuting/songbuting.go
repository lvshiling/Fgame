package songbuting

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/global"
	"fgame/fgame/game/songbuting/dao"
	"fgame/fgame/game/songbuting/template"

	//注册管理器
	_ "fgame/fgame/game/songbuting/event/listener"
	_ "fgame/fgame/game/songbuting/handler"
	_ "fgame/fgame/game/songbuting/player"
)

type songBuTingModule struct {
}

func (m *songBuTingModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *songBuTingModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *songBuTingModule) Start() {

}

func (m *songBuTingModule) Stop() {

}

func (m *songBuTingModule) String() string {
	return "songbuting"
}

var (
	m = &songBuTingModule{}
)

func init() {
	module.Register(m)
}
