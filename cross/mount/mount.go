package mount

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/mount/event/listener"
	_ "fgame/fgame/cross/mount/handler"
	_ "fgame/fgame/game/mount/event/listener/common"
	"fgame/fgame/game/mount/mount"
)

type mountModule struct {
}

func (cm *mountModule) InitTemplate() (err error) {
	err = mount.Init()
	if err != nil {
		return
	}
	return
}

func (cm *mountModule) Init() (err error) {

	return
}

func (m *mountModule) Start() {

}

func (cm *mountModule) Stop() {

}

func (cm *mountModule) String() string {
	return "mount"
}

var (
	m = &mountModule{}
)

func init() {
	module.Register(m)
}
