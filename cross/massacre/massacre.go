package massacre

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/massacre/cross_handler"
	_ "fgame/fgame/cross/massacre/event/listener"
)

type massacreModule struct {
}

func (m *massacreModule) InitTemplate() (err error) {

	return
}
func (m *massacreModule) Init() (err error) {

	return
}

func (m *massacreModule) Start() {

}

func (m *massacreModule) Stop() {

}

func (m *massacreModule) String() string {
	return "massacre"
}

var (
	m = &massacreModule{}
)

func init() {
	module.Register(m)
}
