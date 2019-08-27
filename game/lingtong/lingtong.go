package lingtong

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtong/dao"
	"fgame/fgame/game/lingtong/template"

	//注册管理器
	_ "fgame/fgame/game/lingtong/action"
	_ "fgame/fgame/game/lingtong/event/listener"
	_ "fgame/fgame/game/lingtong/handler"
	_ "fgame/fgame/game/lingtong/player"
	_ "fgame/fgame/game/lingtong/player/effect"
	_ "fgame/fgame/game/lingtong/use"
)

type lingTongModule struct {
}

func (m *lingTongModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *lingTongModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *lingTongModule) Start() {

}

func (m *lingTongModule) Stop() {

}

func (m *lingTongModule) String() string {
	return "lingTong"
}

var (
	m = &lingTongModule{}
)

func init() {
	module.Register(m)
}
