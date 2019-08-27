package login

import (
	_ "fgame/fgame/account/login/handler"
	"fgame/fgame/account/login/login"
	_ "fgame/fgame/account/login/login_handler"
	"fgame/fgame/core/module"
	corerunner "fgame/fgame/core/runner"
)

type loginModule struct {
	r corerunner.GoRunner
}

func (m *loginModule) InitTemplate() (err error) {
	return
}

func (m *loginModule) Init() (err error) {
	err = login.Init()
	if err != nil {
		return
	}
	return
}

func (m *loginModule) Start() {

}

func (m *loginModule) Stop() {
}

func (m *loginModule) String() string {
	return "login"
}

var (
	m = &loginModule{}
)

func init() {
	module.Register(m)
}
