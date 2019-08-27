package quiz

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/quiz/dao"
	"fgame/fgame/game/quiz/quiz"
	quiztemplate "fgame/fgame/game/quiz/template"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/quiz/event/listener"
	_ "fgame/fgame/game/quiz/handler"
)

const (
	quizHeartBeatTime = 3 * time.Second
)

//仙尊问答
type quizModule struct {
	r runner.GoRunner
}

func (m *quizModule) InitTemplate() (err error) {
	err = quiztemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *quizModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = quiz.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("quiz", quiz.GetQuizService().Heartbeat, quizHeartBeatTime)
	return
}

func (m *quizModule) Start() {
	m.r.Start()
}

func (m *quizModule) Stop() {
	m.r.Stop()
}

func (m *quizModule) String() string {
	return "quiz"
}

var (
	m = &quizModule{}
)

func init() {
	module.Register(m)
}
