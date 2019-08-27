package template

import (
	"fgame/fgame/game/lingtongdev/types"
	propertytypes "fgame/fgame/game/property/types"
)

type LingTongDevUpstarTemplate interface {
	GetClassType() types.LingTongDevSysType
	TemplateId() int
	GetNextId() int32
	GetLevel() int32
	GetUpdateWfb() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
	GetZhuFuMax() int32
	GetItemMap() map[int32]int32
	GetUpstarPercent() int32
	GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetNext() LingTongDevUpstarTemplate
}

type LingTongDevPeiYangTemplate interface {
	GetClassType() types.LingTongDevSysType
	TemplateId() int
	GetNextId() int32
	GetLevel() int32
	GetUpdateWfb() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
	GetZhuFuMax() int32
	GetItemMap() map[int32]int32
	GetItemId() int32
	GetItemCount() int32
	GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetNext() LingTongDevPeiYangTemplate
}

type LingTongDevHuanHuaTemplate interface {
	GetClassType() types.LingTongDevSysType
	TemplateId() int
	GetNextId() int32
	GetLevel() int32
	GetUpdateWfb() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
	GetZhuFuMax() int32
	GetItemMap() map[int32]int32
	GetItemId() int32
	GetItemCount() int32
	GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetNext() LingTongDevHuanHuaTemplate
}

type LingTongDevTongLingTemplate interface {
	GetClassType() types.LingTongDevSysType
	TemplateId() int
	GetNextId() int32
	GetLevel() int32
	GetUpdateWfb() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
	GetZhuFuMax() int32
	GetItemMap() map[int32]int32
	GetItemCount() int32
	GetTongLingPercent() int32
	GetNext() LingTongDevTongLingTemplate
}

type LingTongDevTemplate interface {
	GetClassType() types.LingTongDevSysType
	GetName() string
	TemplateId() int
	GetIsClear() bool
	GetType() types.LingTongDevType
	GetNextId() int32
	GetUpdateWfb() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
	GetZhuFuMax() int32
	GetItemMap() map[int32]int32
	GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64
	GetMagicParamIMap() map[int32]int32
	GetMagicParamXUMap() map[int32]int32
	GetUpstarBeginId() int32
	GetGold() int64
	GetBindGold() int64
	GetSilver() int64
	GetNumber() int32
	GetShiDanLimit() int32
	GetCulDanLimit() int32
	GetNext() LingTongDevTemplate
	GetLingTongDevUpstarByLevel(level int32) LingTongDevUpstarTemplate
	GetLingTongDevPeiYangByLevel(level int32) LingTongDevPeiYangTemplate
	GetLingTongDevTongLingByLevel(level int32) LingTongDevTongLingTemplate
}
