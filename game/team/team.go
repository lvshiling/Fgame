package team

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/team/team"
)

import (
	//注册管理器
	_ "fgame/fgame/game/team/event/listener"
	_ "fgame/fgame/game/team/handler"
	_ "fgame/fgame/game/team/player"
)

type teamModule struct {
}

func (m *teamModule) InitTemplate() (err error) {

	return
}
func (cm *teamModule) Init() (err error) {
	err = team.Init()
	if err != nil {
		return
	}
	return
}

func (m *teamModule) Start() {

}

func (m *teamModule) Stop() {

}

func (m *teamModule) String() string {
	return "team"
}

var (
	m = &teamModule{}
)

func init() {
	module.Register(m)
}
