package notice

import (
	"fgame/fgame/core/module"
)

import (
	_ "fgame/fgame/cross/notice/handler"
)

type noticeModule struct {
}

func (m *noticeModule) InitTemplate() (err error) {

	return
}

func (m *noticeModule) Init() (err error) {

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
