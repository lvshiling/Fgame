package qixue

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/qixue/cross_handler"
	_ "fgame/fgame/cross/qixue/event/listener"
	qixuetemplate "fgame/fgame/game/qixue/template"
)

type qixueModule struct {
}

func (m *qixueModule) InitTemplate() (err error) {
	err = qixuetemplate.Init()
	if err != nil {
		return
	}
	return
}
func (m *qixueModule) Init() (err error) {

	return
}

func (m *qixueModule) Start() {

}

func (m *qixueModule) Stop() {

}

func (m *qixueModule) String() string {
	return "qixue"
}

var (
	m = &qixueModule{}
)

func init() {
	module.Register(m)
}
