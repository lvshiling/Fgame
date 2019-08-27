package module

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/dao"
	"time"
)

import (
	_ "fgame/fgame/game/player"
	_ "fgame/fgame/game/player/event/listener"
	_ "fgame/fgame/game/player/handler"
	_ "fgame/fgame/game/player/player"
	_ "fgame/fgame/game/player/use"
	_ "fgame/fgame/game/property/player/effect"
)

//模块化
type playerModule struct {
	r runner.GoRunner
}

func (m *playerModule) InitTemplate() (err error) {

	return
}

const (
	playerTimer = time.Minute
)

func (m *playerModule) Init() (err error) {

	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	err = player.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("player", player.GetPlayerService().Heartbeat, playerTimer)
	return
}

func (m *playerModule) Start() {
	m.r.Start()

}
func (m *playerModule) Stop() {
	m.r.Stop()
}

func (m *playerModule) String() string {
	return "player"
}

var (
	m = &playerModule{}
)

func init() {
	module.Register(m)
}
