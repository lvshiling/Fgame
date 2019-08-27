package myboss

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/myboss/dao"
	mybosstemplate "fgame/fgame/game/myboss/template"

	//注册管理器
	_ "fgame/fgame/game/myboss/event/listener"
	_ "fgame/fgame/game/myboss/handler"
)

//个人BOSS
type mybossModule struct {
}

func (m *mybossModule) InitTemplate() (err error) {
	err = mybosstemplate.Init()
	if err != nil {
		return
	}
	return
}

func (m *mybossModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *mybossModule) Start() {

}

func (m *mybossModule) Stop() {

}

func (m *mybossModule) String() string {
	return "myboss"
}

var (
	m = &mybossModule{}
)

func init() {
	module.Register(m)
}
