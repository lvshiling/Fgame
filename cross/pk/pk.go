package common

import (
	"fgame/fgame/core/module"

	//注册管理器
	_ "fgame/fgame/cross/pk/handler"
	_ "fgame/fgame/game/pk/event/listener/common"
)

type pkModule struct {
}

func (cm *pkModule) InitTemplate() (err error) {

	return
}

func (m *pkModule) Init() (err error) {

	return
}

func (m *pkModule) Start() {

}

func (m *pkModule) Stop() {

}

func (m *pkModule) String() string {
	return "pk"
}

var (
	m = &pkModule{}
)

func init() {
	module.Register(m)
}
