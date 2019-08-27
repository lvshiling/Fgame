package shengtan

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/shengtan/shengtan"
	shengtantemplate "fgame/fgame/game/shengtan/template"
)
import (
	_ "fgame/fgame/game/shengtan/activity_handler"
	_ "fgame/fgame/game/shengtan/check_enter"
	_ "fgame/fgame/game/shengtan/event/listener"
	_ "fgame/fgame/game/shengtan/found_handler"
	_ "fgame/fgame/game/shengtan/guaji"
	_ "fgame/fgame/game/shengtan/handler"
	_ "fgame/fgame/game/shengtan/use"
)

//圣坛
type shengTanModule struct {
	r runner.GoRunner
}

func (m *shengTanModule) InitTemplate() (err error) {
	err = shengtantemplate.Init()
	if err != nil {
		return
	}

	return
}

func (m *shengTanModule) Init() (err error) {
	err = shengtan.Init()
	if err != nil {
		return
	}

	return
}

func (m *shengTanModule) Start() {
	return
}

func (m *shengTanModule) Stop() {

}

func (m *shengTanModule) String() string {
	return "shengtan"
}

var (
	m = &shengTanModule{}
)

func init() {
	module.Register(m)
}
