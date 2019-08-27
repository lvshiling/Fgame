package chat

import (
	"fgame/fgame/core/module"
)

import (
	_ "fgame/fgame/cross/chat/handler"
)

type chatModule struct {
}

func (m *chatModule) InitTemplate() (err error) {

	return
}

func (m *chatModule) Init() (err error) {

	return
}

func (m *chatModule) Start() {

}

func (m *chatModule) Stop() {

}

func (m *chatModule) String() string {
	return "chat"
}

var (
	m = &chatModule{}
)

func init() {
	module.Register(m)
}
