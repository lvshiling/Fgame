package guidereplica

import (
	"fgame/fgame/core/module"
	guidereplicatemplate "fgame/fgame/game/guidereplica/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/guidereplica/event/listener"
	_ "fgame/fgame/game/guidereplica/handler"
	_ "fgame/fgame/game/guidereplica/scene"
)

//引导副本
type guidereplicaModule struct {
}

func (m *guidereplicaModule) InitTemplate() (err error) {
	err = guidereplicatemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *guidereplicaModule) Init() (err error) {
	return
}

func (m *guidereplicaModule) Start() {

}

func (m *guidereplicaModule) Stop() {

}

func (m *guidereplicaModule) String() string {
	return "guidereplica"
}

var (
	m = &guidereplicaModule{}
)

func init() {
	module.Register(m)
}
