package scene

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/robot/template"
)

import (
	_ "fgame/fgame/game/robot/ai"
	_ "fgame/fgame/game/robot/ai/client_test"
	_ "fgame/fgame/game/robot/ai/model"
	_ "fgame/fgame/game/robot/event/listener"
)

type robotModule struct {
}

func (m *robotModule) InitTemplate() (err error) {

	return
}
func (m *robotModule) Init() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	err = robot.Init()
	return
}

func (m *robotModule) Start() {

}

func (m *robotModule) Stop() {

}

func (m *robotModule) String() string {
	return "robot"
}

var (
	m = &robotModule{}
)

func init() {
	module.Register(m)
}
