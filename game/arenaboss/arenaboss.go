package arenaboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"

	arenabosstemplate "fgame/fgame/game/arenaboss/template"
	"time"

	//注册
	_ "fgame/fgame/game/arenaboss/boss_handler"
	_ "fgame/fgame/game/arenaboss/check_enter"
)

//跨服世界boss
type arenaBossModule struct {
	r runner.GoRunner
}

func (m *arenaBossModule) InitTemplate() (err error) {
	err = arenabosstemplate.Init()
	if err != nil {
		return
	}

	return
}

const (
	arenaBossTimer = time.Second * 5
)

func (m *arenaBossModule) Init() (err error) {

	return
}

func (m *arenaBossModule) Start() {

	return
}

func (m *arenaBossModule) Stop() {

}

func (m *arenaBossModule) String() string {
	return "arenaboss"
}

var (
	m = &arenaBossModule{}
)

func init() {
	module.Register(m)
}
