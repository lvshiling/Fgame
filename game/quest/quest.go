package quest

import (
	"fgame/fgame/core/module"

	"fgame/fgame/game/global"
	"fgame/fgame/game/quest/dao"
	"fgame/fgame/game/quest/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/quest/check"
	_ "fgame/fgame/game/quest/click"
	_ "fgame/fgame/game/quest/commitflow"
	_ "fgame/fgame/game/quest/event/listener"
	_ "fgame/fgame/game/quest/found_handler"
	_ "fgame/fgame/game/quest/guaji"
	_ "fgame/fgame/game/quest/handler"
	_ "fgame/fgame/game/quest/player"
	_ "fgame/fgame/game/quest/quest_guaji"
	_ "fgame/fgame/game/quest/robot/ai"
)

type questModule struct {
}

func (m *questModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *questModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}

func (m *questModule) Start() {

}

func (m *questModule) Stop() {

}

func (m *questModule) String() string {
	return "quest"
}

var (
	m = &questModule{}
)

func init() {
	module.Register(m)
}
