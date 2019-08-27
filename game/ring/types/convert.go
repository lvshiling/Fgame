package types

import (
	viptypes "fgame/fgame/game/vip/types"
)

// 宝库类型到后台规则系统
var (
	convertToCostLevelRuleTypeMap = map[BaoKuType]viptypes.CostLevelRuleType{
		BaoKuTypeRing: viptypes.CostLevelRuleTypeRingBaoKu,
	}
)

func RingTypeToCostLevelRuleType(baoKuType BaoKuType) (costLevelRuleType viptypes.CostLevelRuleType, isExist bool) {
	costLevelRuleType, isExist = convertToCostLevelRuleTypeMap[baoKuType]
	return
}
