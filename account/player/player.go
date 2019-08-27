package player

import (
	"fgame/fgame/account/player/player"
	_ "fgame/fgame/account/serverlist/handler"
	"fgame/fgame/core/module"
	corerunner "fgame/fgame/core/runner"
)

type playerModule struct {
	r corerunner.GoRunner
}

func (m *playerModule) InitTemplate() (err error) {
	return
}

func (m *playerModule) Init() (err error) {
	err = player.Init()
	if err != nil {
		return
	}
	return
}

func (m *playerModule) Start() {

}

func (m *playerModule) Stop() {
}

func (m *playerModule) String() string {
	return "player"
}

var (
	m = &playerModule{}
)

func init() {
	module.Register(m)
}
