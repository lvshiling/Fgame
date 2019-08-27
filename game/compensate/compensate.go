package shenfa

import (
	"fgame/fgame/core/module"
	"fgame/fgame/game/compensate/compensate"
	"fgame/fgame/game/compensate/dao"
	"fgame/fgame/game/global"

	//注册管理器
	_ "fgame/fgame/game/compensate/event/listener"
	_ "fgame/fgame/game/compensate/player"
)

//补偿邮件
type compensateModule struct {
}

func (m *compensateModule) InitTemplate() (err error) {
	return
}

func (m *compensateModule) Init() (err error) {
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = compensate.Init()
	if err != nil {
		return
	}

	return
}

func (m *compensateModule) Start() {
}

func (m *compensateModule) Stop() {

}

func (m *compensateModule) String() string {
	return "compensate"
}

var (
	m = &compensateModule{}
)

func init() {
	module.Register(m)
}
