package player

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/lingtong/handler"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
)

type lingTongDevModule struct {
}

func (m *lingTongDevModule) InitTemplate() (err error) {
	return
}
func (m *lingTongDevModule) Init() (err error) {
	err = lingtongdevtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *lingTongDevModule) Start() {

}

func (m *lingTongDevModule) Stop() {

}

func (m *lingTongDevModule) String() string {
	return "lingtongdev"
}

var (
	m = &lingTongDevModule{}
)

func init() {
	module.Register(m)
}
