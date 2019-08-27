package template

import (
	propertytypes "fgame/fgame/game/property/types"
	questtypes "fgame/fgame/game/quest/types"
)

type DailyTagTemplate interface {
	TemplateId() int
	GetNextId() int32
	GetLevelMin() int32
	GetLevelMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
	GetQuestId() int32
	GetPercent() int32
	GetRewExp() int32
	GetRewExpPoint() int32
	GetRewSilver() int32
	GetRewBindGold() int32
	GetRewGold() int32
	GetBossExp() int32
	GetRewData() *propertytypes.RewData
	GetDoubleRewData() *propertytypes.RewData
	GetRewItemMap() map[int32]int32
	GetDoubleRewItemMap() map[int32]int32
	GetDailyTimesType() questtypes.QuestDailyType
	GetEmailItemMap() map[int32]int32
	GetUnionStorageJiFen() int32
	GetDropId() int32
}
