package house

import (
	"fgame/fgame/core/module"
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/house/dao"
	"fgame/fgame/game/house/house"
	housetemplate "fgame/fgame/game/house/template"
	"time"

	//注册管理器
	_ "fgame/fgame/game/house/event/listener"
	_ "fgame/fgame/game/house/handler"
	_ "fgame/fgame/game/house/player"
)

//房子
type houseModule struct {
	r runner.GoRunner // service 使用定时器
}

func (m *houseModule) InitTemplate() (err error) {
	// 模板初始化
	err = housetemplate.Init()
	if err != nil {
		return
	}
	return
}

const (
	taskTime = 5 * time.Second
)

func (m *houseModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()

	err = dao.Init(ds, rs) //dao初始化
	if err != nil {
		return
	}

	err = house.Init() //serice初始化
	if err != nil {
		return
	}

	// service 使用定时器
	m.r = runner.NewGoRunner("house", house.GetHouseService().Heartbeat, taskTime)

	return
}

func (m *houseModule) Start() {
	m.r.Start() // service 使用定时器
}

func (m *houseModule) Stop() {
	m.r.Stop() // service 使用定时器
}

func (m *houseModule) String() string {
	return "house"
}

var (
	m = &houseModule{}
)

func init() {
	module.Register(m)
}
