package gem

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/chat/chat"
	"fgame/fgame/game/chat/dao"
	"fgame/fgame/game/chat/template"
	"fgame/fgame/game/global"

	_ "fgame/fgame/game/chat/event/listener"

	_ "fgame/fgame/game/chat/handler"
	_ "fgame/fgame/game/chat/player"
)

type chatModule struct {
}

func (m *chatModule) InitTemplate() (err error) {
	err = template.Init()

	return
}

func (m *chatModule) Init() (err error) {
	db := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(db, rs)
	if err != nil {
		return
	}
	err = chat.Init()
	if err != nil {
		return
	}

	return
}

func (m *chatModule) Start() {

}

func (m *chatModule) Stop() {

}

func (m *chatModule) String() string {
	return "chat"
}

var (
	m = &chatModule{}
)

func init() {
	module.RegisterBase(m)
}
