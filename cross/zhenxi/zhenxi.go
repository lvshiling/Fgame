package zhenxi

import (
	"fgame/fgame/core/module"
	zhenxitemplate "fgame/fgame/game/zhenxi/template"

	//注册
	_ "fgame/fgame/cross/zhenxi/boss_handler"
	_ "fgame/fgame/cross/zhenxi/login_handler"

	_ "fgame/fgame/cross/zhenxi/event/listener"

	_ "fgame/fgame/cross/zhenxi/relive_handler"
)

//藏经阁boss
type zhenxiModule struct {
}

func (m *zhenxiModule) InitTemplate() (err error) {
	err = zhenxitemplate.Init()
	if err != nil {
		return
	}

	return
}
func (m *zhenxiModule) Init() (err error) {

	return
}

func (m *zhenxiModule) Start() {

	return
}

func (m *zhenxiModule) Stop() {

}

func (m *zhenxiModule) String() string {
	return "zhenxi"
}

var (
	m = &zhenxiModule{}
)

func init() {
	module.Register(m)
}
