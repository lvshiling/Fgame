package serverlist

import (
	_ "fgame/fgame/account/serverlist/handler"
	"fgame/fgame/account/serverlist/serverlist"
	"fgame/fgame/core/module"
	corerunner "fgame/fgame/core/runner"
)

type serverListModule struct {
	r corerunner.GoRunner
}

func (m *serverListModule) InitTemplate() (err error) {
	return
}

func (m *serverListModule) Init() (err error) {
	err = serverlist.Init()
	if err != nil {
		return
	}
	return
}

func (m *serverListModule) Start() {

}

func (m *serverListModule) Stop() {
}

func (m *serverListModule) String() string {
	return "serverlist"
}

var (
	m = &serverListModule{}
)

func init() {
	module.Register(m)
}
