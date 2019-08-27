package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 进阶返利
type AdvancedInfo struct {
	DanNum      int32   `json:"danNum"`      //使用进阶丹数量
	RewRecord   []int32 `json:"rewRecord"`   //奖励领取记录
	AdvancedDay int32   `json:"advancedDay"` //当前返利日
}

func (info *AdvancedInfo) AddRecord(needDanNum int32) {
	info.RewRecord = append(info.RewRecord, needDanNum)
}

func (info *AdvancedInfo) IsCanReceiveRewards(needDanNum int32) bool {
	//条件
	if info.DanNum < needDanNum {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needDanNum {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeFeedback, (*AdvancedInfo)(nil))
}
