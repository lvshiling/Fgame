package player

import (
	"fgame/fgame/core/module"
	_ "fgame/fgame/cross/lingtong/handler"
	_ "fgame/fgame/game/lingtong/action"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
)

type lingTongModule struct {
}

func (m *lingTongModule) InitTemplate() (err error) {
	return
}
func (m *lingTongModule) Init() (err error) {
	err = lingtongtemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *lingTongModule) Start() {

}

func (m *lingTongModule) Stop() {

}

func (m *lingTongModule) String() string {
	return "lingtong"
}

var (
	m = &lingTongModule{}
)

func init() {
	module.Register(m)
}
