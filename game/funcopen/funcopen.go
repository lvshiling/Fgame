package funcopen

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/funcopen/dao"
	"fgame/fgame/game/funcopen/funcopen"
	"fgame/fgame/game/global"
)
import (
	////注册管理器
	_ "fgame/fgame/game/funcopen/event/listener"
	_ "fgame/fgame/game/funcopen/guaji"
	_ "fgame/fgame/game/funcopen/handler"
	_ "fgame/fgame/game/funcopen/player"
)

type funcOpenModule struct {
}

func (m *funcOpenModule) InitTemplate() (err error) {
	err = funcopen.Init()
	if err != nil {
		return
	}
	return
}

func (cm *funcOpenModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *funcOpenModule) Start() {

}

func (cm *funcOpenModule) Stop() {

}

func (cm *funcOpenModule) String() string {
	return "funcopen"
}

var (
	m = &funcOpenModule{}
)

func init() {
	module.Register(m)
}
