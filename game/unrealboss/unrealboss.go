package unrealboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/unrealboss/dao"
	unrealbosstemplate "fgame/fgame/game/unrealboss/template"
	"fgame/fgame/game/unrealboss/unrealboss"

	//注册管理器
	_ "fgame/fgame/game/unrealboss/boss_handler"
	_ "fgame/fgame/game/unrealboss/check_enter"
	_ "fgame/fgame/game/unrealboss/event/listener"
	_ "fgame/fgame/game/unrealboss/guaji"
	_ "fgame/fgame/game/unrealboss/handler"
	_ "fgame/fgame/game/unrealboss/npc/check_attack"
	_ "fgame/fgame/game/unrealboss/use"
)

//幻境BOSS
type unrealbossModule struct {
}

func (m *unrealbossModule) InitTemplate() (err error) {
	err = unrealbosstemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *unrealbossModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = unrealboss.Init()
	if err != nil {
		return
	}

	return
}

func (m *unrealbossModule) Start() {
	unrealboss.GetUnrealBossService().Start()
}

func (m *unrealbossModule) Stop() {

}

func (m *unrealbossModule) String() string {
	return "unrealboss"
}

var (
	m = &unrealbossModule{}
)

func init() {
	module.Register(m)
}
