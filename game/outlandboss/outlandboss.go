package outlandboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/outlandboss/dao"
	"fgame/fgame/game/outlandboss/outlandboss"
	outlandbosstemplate "fgame/fgame/game/outlandboss/template"
	"time"
)

import (
	// 注册
	_ "fgame/fgame/game/outlandboss/boss_handler"
	_ "fgame/fgame/game/outlandboss/check_enter"
	_ "fgame/fgame/game/outlandboss/event/listener"
	_ "fgame/fgame/game/outlandboss/guaji"
	_ "fgame/fgame/game/outlandboss/handler"
	_ "fgame/fgame/game/outlandboss/npc/check_attack"
	_ "fgame/fgame/game/outlandboss/use"
)

//外域BOSS
type outlandbossModule struct {
	r runner.GoRunner
}

func (m *outlandbossModule) InitTemplate() (err error) {
	err = outlandbosstemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *outlandbossModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = outlandboss.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("outlandboss", outlandboss.GetOutlandBossService().Heartbeat, 3*time.Second)
	return
}

func (m *outlandbossModule) Start() {
	outlandboss.GetOutlandBossService().Start()
	m.r.Start()
}

func (m *outlandbossModule) Stop() {
	m.r.Stop()
}

func (m *outlandbossModule) String() string {
	return "outlandboss"
}

var (
	m = &outlandbossModule{}
)

func init() {
	module.Register(m)
}
