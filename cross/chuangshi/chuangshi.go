package chuangshi

// import (
// 	"fgame/fgame/core/module"
// 	"fgame/fgame/core/runner"
// 	"fgame/fgame/cross/chuangshi/chuangshi"
// 	"fgame/fgame/cross/chuangshi/dao"
// 	chuangshitemplate "fgame/fgame/game/chuangshi/template"
// 	"fgame/fgame/game/global"
// 	"time"

// 	_ "fgame/fgame/cross/chuangshi/battle"
// 	_ "fgame/fgame/cross/chuangshi/cross_handler"
// 	_ "fgame/fgame/cross/chuangshi/login_handler"
// )

// type chuangShiModule struct {
// 	r runner.GoRunner
// }

// func (m *chuangShiModule) InitTemplate() (err error) {
// 	err = chuangshitemplate.Init()
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// func (m *chuangShiModule) Init() (err error) {
// 	db := global.GetGame().GetDB()
// 	rs := global.GetGame().GetRedisService()
// 	err = dao.Init(db, rs)
// 	if err != nil {
// 		return
// 	}

// 	err = chuangshi.Init()
// 	if err != nil {
// 		return
// 	}

// 	m.r = runner.NewGoRunner("chuangshi", chuangshi.GetChuangShiService().Heartbeat, time.Second)
// 	return
// }

// func (m *chuangShiModule) Start() {

// 	chuangshi.GetChuangShiService().Start()
// 	m.r.Start()

// }

// func (m *chuangShiModule) Stop() {
// 	m.r.Stop()
// 	chuangshi.GetChuangShiService().Stop()
// }

// func (m *chuangShiModule) String() string {
// 	return "chuangshi"
// }

// var (
// 	m = &chuangShiModule{}
// )

// func init() {
// 	module.Register(m)
// }
