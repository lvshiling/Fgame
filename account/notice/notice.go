package notice

import (
	"fgame/fgame/account/notice/notice"
	_ "fgame/fgame/account/serverlist/handler"
	"fgame/fgame/core/module"
	corerunner "fgame/fgame/core/runner"
)

type noticeModule struct {
	r corerunner.GoRunner
}

func (m *noticeModule) InitTemplate() (err error) {
	return
}

func (m *noticeModule) Init() (err error) {
	err = notice.Init()
	if err != nil {
		return
	}
	return
}

func (m *noticeModule) Start() {

}

func (m *noticeModule) Stop() {
}

func (m *noticeModule) String() string {
	return "notice"
}

var (
	m = &noticeModule{}
)

func init() {
	module.Register(m)
}
