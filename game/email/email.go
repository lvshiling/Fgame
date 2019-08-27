package email

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/email/dao"
	"fgame/fgame/game/global"
)

import (
	_ "fgame/fgame/game/email/event/listener"
	_ "fgame/fgame/game/email/guaji"
	_ "fgame/fgame/game/email/handler"
	_ "fgame/fgame/game/email/player"
)

type emailModule struct {
}

func (m *emailModule) InitTemplate() (err error) {
	return
}

func (m *emailModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return err
	}

	return
}

func (m *emailModule) Start() {
}

func (m *emailModule) Stop() {

}

func (m *emailModule) String() string {
	return "email"
}

var (
	m = &emailModule{}
)

func init() {
	module.RegisterBase(m)
}
