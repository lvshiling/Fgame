package mount

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/mount/dao"
	"fgame/fgame/game/mount/mount"
)

import (
	//注册管理器
	_ "fgame/fgame/game/mount/active_check"
	_ "fgame/fgame/game/mount/cross_handler"
	_ "fgame/fgame/game/mount/event/listener"
	_ "fgame/fgame/game/mount/event/listener/common"
	_ "fgame/fgame/game/mount/guaji"
	_ "fgame/fgame/game/mount/guaji_check"
	_ "fgame/fgame/game/mount/handler"
	_ "fgame/fgame/game/mount/player"
	_ "fgame/fgame/game/mount/systemskill_handler"
	_ "fgame/fgame/game/mount/use"
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
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
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
