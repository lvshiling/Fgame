package lingyu

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingyu/dao"
	lingyutemplate "fgame/fgame/game/lingyu/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/lingyu/active_check"
	_ "fgame/fgame/game/lingyu/event/listener"
	_ "fgame/fgame/game/lingyu/guaji"
	_ "fgame/fgame/game/lingyu/handler"
	_ "fgame/fgame/game/lingyu/player"
	_ "fgame/fgame/game/lingyu/systemskill_handler"
	_ "fgame/fgame/game/lingyu/use"
)

//领域
type shenfaModule struct {
}

func (m *shenfaModule) InitTemplate() (err error) {
	err = lingyutemplate.Init()
	if err != nil {
		return
	}
	return
}
func (m *shenfaModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *shenfaModule) Start() {

}

func (m *shenfaModule) Stop() {

}

func (m *shenfaModule) String() string {
	return "lingyu"
}

var (
	m = &shenfaModule{}
)

func init() {
	module.Register(m)
}
