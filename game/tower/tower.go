package tower

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/tower/dao"
	"fgame/fgame/game/tower/template"
	"fgame/fgame/game/tower/tower"
)

import (
	//注册管理器
	_ "fgame/fgame/game/tower/check_enter"
	_ "fgame/fgame/game/tower/event/listener"
	_ "fgame/fgame/game/tower/guaji"
	_ "fgame/fgame/game/tower/handler"
	_ "fgame/fgame/game/tower/use"
)

//打宝塔
type towerModule struct {
}

func (m *towerModule) InitTemplate() (err error) {
	err = template.Init()

	return
}

func (m *towerModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = tower.Init()
	if err != nil {
		return
	}

	return
}

func (m *towerModule) Start() {
	tower.GetTowerService().Star()

}

func (m *towerModule) Stop() {
}

func (m *towerModule) String() string {
	return "tower"
}

var (
	m = &towerModule{}
)

func init() {
	module.Register(m)
}
