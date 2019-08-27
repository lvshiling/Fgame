package scene

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/dao"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
)
import (
	//注册管理器
	_ "fgame/fgame/game/scene/check_enter"
	_ "fgame/fgame/game/scene/cross_handler"
	_ "fgame/fgame/game/scene/event/listener"
	_ "fgame/fgame/game/scene/event/listener/common"
	_ "fgame/fgame/game/scene/guaji"
	_ "fgame/fgame/game/scene/handler"
	_ "fgame/fgame/game/scene/player"
	_ "fgame/fgame/game/scene/player/listener"
)

type sceneModule struct {
}

func (m *sceneModule) InitTemplate() (err error) {
	err = scenetemplate.Init()
	if err != nil {
		return
	}
	return
}
func (m *sceneModule) Init() (err error) {

	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	// err = scene.InitDingShi()
	// if err != nil {
	// 	return
	// }
	err = scene.Init()
	if err != nil {
		return
	}

	return
}

func (m *sceneModule) Start() {
	// scene.GetDingShiService().Start()
}

func (m *sceneModule) Stop() {

}

func (m *sceneModule) String() string {
	return "scene"
}

var (
	m = &sceneModule{}
)

func init() {
	module.Register(m)
}
