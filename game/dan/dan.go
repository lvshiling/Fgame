package dan

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/dan/dan"
	"fgame/fgame/game/dan/dao"
	"fgame/fgame/game/global"
)

//注册管理器
import (
	_ "fgame/fgame/game/dan/handler"
	_ "fgame/fgame/game/dan/player"
)

type danModule struct {
}

func (m *danModule) InitTemplate() (err error) {
	err = dan.Init()
	if err != nil {
		return
	}
	return
}

func (m *danModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *danModule) Start() {

}

func (m *danModule) Stop() {

}

func (m *danModule) String() string {
	return "dan"
}

var (
	m = &danModule{}
)

func init() {
	module.Register(m)
}
