package worldboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/worldboss/dao"
	worldbosstemplate "fgame/fgame/game/worldboss/template"
	"fgame/fgame/game/worldboss/worldboss"

	//注册
	_ "fgame/fgame/game/worldboss/boss_handler"
	_ "fgame/fgame/game/worldboss/cross_handler"

	_ "fgame/fgame/game/worldboss/check_enter"

	_ "fgame/fgame/game/worldboss/event/listener"

	_ "fgame/fgame/game/worldboss/guaji"

	_ "fgame/fgame/game/worldboss/handler"

	_ "fgame/fgame/game/worldboss/player"
)

//世界boss
type worldBossModule struct {
}

func (m *worldBossModule) InitTemplate() (err error) {
	err = worldbosstemplate.Init()
	if err != nil {
		return
	}

	return
}
func (m *worldBossModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	err = worldboss.Init()
	if err != nil {
		return
	}
	return
}

func (m *worldBossModule) Start() {
	worldboss.GetWorldBossService().Start()
	return
}

func (m *worldBossModule) Stop() {

}

func (m *worldBossModule) String() string {
	return "worldboss"
}

var (
	m = &worldBossModule{}
)

func init() {
	module.Register(m)
}
