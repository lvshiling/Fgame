package scene

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/template"
)

import (
	//注册管理器
	_ "fgame/fgame/cross/scene/handler"
)

import (
	//事件监听共用
	_ "fgame/fgame/cross/scene/cross_handler"
	_ "fgame/fgame/cross/scene/event/listener"
	_ "fgame/fgame/game/scene/event/listener/common"
)

type sceneModule struct {
}

func (m *sceneModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}

func (m *sceneModule) Init() (err error) {

	err = scene.Init()
	if err != nil {
		return
	}

	return
}

func (m *sceneModule) Start() {

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
