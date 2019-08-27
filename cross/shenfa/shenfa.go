package shenfa

import (
	"fgame/fgame/core/module"
	shenfatemplate "fgame/fgame/game/shenfa/template"
)

//身法
type shenfaModule struct {
}

func (m *shenfaModule) InitTemplate() (err error) {
	err = shenfatemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *shenfaModule) Init() (err error) {
	return
}

func (m *shenfaModule) Start() {

}

func (m *shenfaModule) Stop() {

}

func (m *shenfaModule) String() string {
	return "shenfa"
}

var (
	m = &shenfaModule{}
)

func init() {
	module.Register(m)
}
