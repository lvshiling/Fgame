package activity

import (
	"fgame/fgame/core/module"
	activitytemplate "fgame/fgame/game/activity/template"
)
import (
	_ "fgame/fgame/cross/activity/cross_handler"
	_ "fgame/fgame/cross/activity/event/listener"
)

//活动大厅
type activeModule struct {
}

func (acModule *activeModule) InitTemplate() (err error) {
	err = activitytemplate.Init()
	return
}

func (acModule *activeModule) Init() (err error) {

	return
}

func (acModule *activeModule) Start() {

}

func (acModule *activeModule) Stop() {

}

func (acModule *activeModule) String() string {
	return "active"
}

var (
	m = &activeModule{}
)

func init() {
	module.Register(m)
}
