package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//新等级投资计划
type InvestNewLevelInfo struct {
	InvestBuyInfoMap map[InvestNewLevelType][]int32 `json:"investBuyInfoMap"` //投资计划购买信息
}

func (info InvestNewLevelInfo) IsBuy() bool {
	if len(info.InvestBuyInfoMap) > 0 {
		return true
	}
	return false
}

func (info InvestNewLevelInfo) IsCanReceiveRew(typ InvestNewLevelType, lev int32) bool {
	receiveRewList, isBuy := info.InvestBuyInfoMap[typ]
	if !isBuy {
		return false
	}
	for _, val := range receiveRewList {
		if val == lev {
			return false
		}
	}
	return true
}

func (info InvestNewLevelInfo) IsCanUpLevel(typ InvestNewLevelType) bool {
	if len(info.InvestBuyInfoMap) == 0 {
		return false
	}

	for lastTyp, _ := range info.InvestBuyInfoMap {
		if int32(lastTyp) >= int32(typ) {
			return false
		}
	}

	return true
}

func (info InvestNewLevelInfo) GetInvestNewLevelType() (typ InvestNewLevelType, exist bool) {
	if len(info.InvestBuyInfoMap) == 0 {
		return
	}
	for key, _ := range info.InvestBuyInfoMap {
		typ = key
		exist = true
	}

	return
}

// 更新投资等级类型
func (info *InvestNewLevelInfo) UpdateInvestLevelType(typ InvestNewLevelType) {
	lastInvestType, exist := info.GetInvestNewLevelType()
	if !exist {
		return
	}

	recordList := info.InvestBuyInfoMap[lastInvestType]
	info.InvestBuyInfoMap[typ] = recordList
	delete(info.InvestBuyInfoMap, lastInvestType)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewLevel, (*InvestNewLevelInfo)(nil))
}
