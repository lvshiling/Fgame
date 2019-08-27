package systemcompensate

import (
	"fgame/fgame/core/module"
	systemcompensatetemplate "fgame/fgame/game/systemcompensate/template"

	//注册管理器
	_ "fgame/fgame/game/systemcompensate/event/listener"
)

type systemcompensateModule struct {
}

func (m *systemcompensateModule) InitTemplate() (err error) {
	err = systemcompensatetemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *systemcompensateModule) Init() (err error) {
	return
}

func (m *systemcompensateModule) Start() {

}

func (m *systemcompensateModule) Stop() {

}

func (m *systemcompensateModule) String() string {
	return "systemcompensate"
}

var (
	m = &systemcompensateModule{}
)

func init() {
	module.Register(m)
}
