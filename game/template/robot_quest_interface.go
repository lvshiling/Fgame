package template

import (
	"fgame/fgame/core/template"
	propertytypes "fgame/fgame/game/property/types"
)

//任务机器人
type RobotQuestTemplateInterface interface {
	template.TemplateObject
	RandomProperty() map[propertytypes.BattlePropertyType]int64
	GetPortalTemplate() *PortalTemplate
	GetRefreshTime() int64
	GetPlayerLimitCount() int32
	GetQuestBeginId() int32
	GetQuestOverId() int32
}
