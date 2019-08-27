package transportation

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/global"
	"fgame/fgame/game/transportation/dao"
	"fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
)

import (
	//注册
	_ "fgame/fgame/game/transportation/event/listener"
	_ "fgame/fgame/game/transportation/guaji"
	_ "fgame/fgame/game/transportation/handler"
	_ "fgame/fgame/game/transportation/npc"
	_ "fgame/fgame/game/transportation/player"
)

//镖车
type transportationModule struct {
}

func (m *transportationModule) InitTemplate() (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	return
}
func (m *transportationModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = transpotation.Init()

	return
}

func (m *transportationModule) Start() {
	transpotation.GetTransportService().Start()
}

func (m *transportationModule) Stop() {

}

func (m *transportationModule) String() string {
	return "transportation"
}

var (
	m = &transportationModule{}
)

func init() {
	module.Register(m)
}
