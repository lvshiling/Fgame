package wardrobe

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/wardrobe/dao"
	"fgame/fgame/game/wardrobe/template"
)

import (
	//注册管理器
	_ "fgame/fgame/game/wardrobe/event/listener"
	_ "fgame/fgame/game/wardrobe/handler"
	_ "fgame/fgame/game/wardrobe/player"
)

type wardrobeModule struct {
}

func (m *wardrobeModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *wardrobeModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs)
	if err != nil {
		return
	}
	return
}
func (m *wardrobeModule) Start() {

}
func (m *wardrobeModule) Stop() {

}

func (m *wardrobeModule) String() string {
	return "wardrobe"
}

var (
	m = &wardrobeModule{}
)

func init() {
	module.Register(m)
}
