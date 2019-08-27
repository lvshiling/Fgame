package found

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/found/dao"
	"fgame/fgame/game/found/found"
	foundtemplate "fgame/fgame/game/found/template"
	"fgame/fgame/game/global"
)

import (
	//注册管理器
	_ "fgame/fgame/game/found/event/listener"
	_ "fgame/fgame/game/found/handler"
	_ "fgame/fgame/game/found/player"
)

//资源找回
type foundModule struct {
}

func (m *foundModule) InitTemplate() (err error) {

	err = foundtemplate.Init()
	if err != nil {
		return
	}

	return
}
func (rm *foundModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = found.Init()
	return
}

func (rm *foundModule) Start() {

}

func (rm *foundModule) String() string {
	return "found"
}

func (rm *foundModule) Stop() {

}

var (
	m = &foundModule{}
)

func init() {
	module.Register(m)
}
