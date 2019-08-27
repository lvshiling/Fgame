package chess

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/dao"
	chesstemplate "fgame/fgame/game/chess/template"
	"fgame/fgame/game/global"
	"time"
)

import (
	//注册管理器
	_ "fgame/fgame/game/chess/event/listener"
	_ "fgame/fgame/game/chess/handler"
	_ "fgame/fgame/game/chess/player"
)

//苍龙棋局
type chessModule struct {
	r runner.GoRunner
}

func (m *chessModule) InitTemplate() (err error) {
	err = chesstemplate.Init()
	return
}

func (m *chessModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = chess.Init()
	if err != nil {
		return
	}

	m.r = runner.NewGoRunner("chess", chess.GetChessService().Heartbeat, 3*time.Second)
	return
}

func (m *chessModule) Start() {
	m.r.Start()
}

func (m *chessModule) Stop() {
	m.r.Stop()
}

func (m *chessModule) String() string {
	return "chess"
}

var (
	m = &chessModule{}
)

func init() {
	module.Register(m)
}
