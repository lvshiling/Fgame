package soul

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/soul/dao"
	"fgame/fgame/game/soul/soul"
)

import (
	//注册管理器
	_ "fgame/fgame/game/soul/event/listener"
	_ "fgame/fgame/game/soul/guaji_check"
	_ "fgame/fgame/game/soul/handler"
	_ "fgame/fgame/game/soul/player"
)

type soulModule struct {
}

func (m *soulModule) InitTemplate() (err error) {
	err = soul.Init()
	if err != nil {
		return
	}

	return
}
func (m *soulModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *soulModule) Start() {

}

func (m *soulModule) Stop() {

}

func (m *soulModule) String() string {
	return "soul"
}

var (
	m = &soulModule{}
)

func init() {
	module.Register(m)
}
