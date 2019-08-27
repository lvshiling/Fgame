package friend

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/friend/dao"
	"fgame/fgame/game/friend/friend"
	friendtemplate "fgame/fgame/game/friend/template"
	"fgame/fgame/game/global"
)

import (
	_ "fgame/fgame/game/friend/event/listener"
	_ "fgame/fgame/game/friend/handler"
	_ "fgame/fgame/game/friend/player"
)

type friendModule struct {
}

func (m *friendModule) InitTemplate() (err error) {
	err = friendtemplate.Init()
	return
}

func (m *friendModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = friend.Init()
	if err != nil {
		return
	}

	return
}

func (m *friendModule) Start() {

}

func (m *friendModule) Stop() {

}

func (m *friendModule) String() string {
	return "friend"
}

var (
	m = &friendModule{}
)

func init() {
	module.Register(m)
}
