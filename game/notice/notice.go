package notice

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/notice/notice"
	"time"

	_ "fgame/fgame/game/notice/event/listener"
	_ "fgame/fgame/game/notice/handler"
)

const (
	hearbeatTimer = 5 * time.Second
)

//公告
type noticeModule struct {
	r runner.GoRunner
}

func (m *noticeModule) InitTemplate() (err error) {
	return
}

func (m *noticeModule) Init() (err error) {
	err = notice.Init()

	m.r = runner.NewGoRunner("notice", notice.GetNoticeService().Heartbeat, hearbeatTimer)
	return
}

func (m *noticeModule) Start() {
	m.r.Start()
}

func (m *noticeModule) Stop() {
	m.r.Stop()
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
