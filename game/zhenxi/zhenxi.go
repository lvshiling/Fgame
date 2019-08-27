package cangjingge

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/zhenxi/dao"
	zhenxitemplate "fgame/fgame/game/zhenxi/template"

	//注册
	_ "fgame/fgame/game/zhenxi/boss_handler"

	_ "fgame/fgame/game/zhenxi/dao"

	_ "fgame/fgame/game/zhenxi/event/listener"
	_ "fgame/fgame/game/zhenxi/handler"
	_ "fgame/fgame/game/zhenxi/player"
)

//藏经阁boss
type zhenxiModule struct {
}

func (m *zhenxiModule) InitTemplate() (err error) {
	err = zhenxitemplate.Init()
	if err != nil {
		return
	}

	return
}
func (m *zhenxiModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	return
}

func (m *zhenxiModule) Start() {

	return
}

func (m *zhenxiModule) Stop() {

}

func (m *zhenxiModule) String() string {
	return "zhenxi"
}

var (
	m = &zhenxiModule{}
)

func init() {
	module.Register(m)
}
