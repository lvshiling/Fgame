package gem

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/register/dao"
	"fgame/fgame/game/register/register"
	"time"
)

type registerModule struct {
	r runner.GoRunner
}

func (m *registerModule) InitTemplate() (err error) {

	return
}

const (
	registerTimer = time.Minute
)

func (m *registerModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	err = register.Init()
	if err != nil {
		return
	}
	m.r = runner.NewGoRunner("register", register.GetRegisterService().Heartbeat, registerTimer)
	return
}

func (m *registerModule) Start() {
	m.r.Start()
}

func (m *registerModule) Stop() {

}

func (m *registerModule) String() string {
	return "register"
}

var (
	m = &registerModule{}
)

func init() {
	module.Register(m)
}
