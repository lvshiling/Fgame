package teamcopy

import (
	"fgame/fgame/core/module"
	"fgame/fgame/cross/teamcopy/teamcopy"
	"fgame/fgame/game/teamcopy/template"

	//注册管理器
	_ "fgame/fgame/cross/teamcopy/cross_handler"
	_ "fgame/fgame/cross/teamcopy/event/listener"
	_ "fgame/fgame/cross/teamcopy/handler"
	_ "fgame/fgame/cross/teamcopy/login_handler"
	_ "fgame/fgame/cross/teamcopy/relive_handler"
	_ "fgame/fgame/cross/teamcopy/robot/ai"
)

type teamCopyModule struct {
}

func (m *teamCopyModule) InitTemplate() (err error) {
	err = template.Init()
	return
}

func (m *teamCopyModule) Init() (err error) {

	err = teamcopy.Init()
	if err != nil {
		return
	}

	return
}

func (m *teamCopyModule) Start() {

}

func (m *teamCopyModule) Stop() {
}

func (m *teamCopyModule) String() string {
	return "teamcopy"
}

var (
	m = &teamCopyModule{}
)

func init() {
	module.Register(m)
}
