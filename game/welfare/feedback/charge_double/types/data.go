package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//首充翻倍
type FeedbackChargeDoubleInfo struct {
	Record []int32 `json:"record"` //档次记录
}

func (info *FeedbackChargeDoubleInfo) IsDouble(chargeId int32) bool {
	for _, recordNum := range info.Record {
		if recordNum == chargeId {
			return true
		}
	}

	return false
}

func (info *FeedbackChargeDoubleInfo) AddRecord(chargeId int32) {
	info.Record = append(info.Record, chargeId)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDouble, (*FeedbackChargeDoubleInfo)(nil))
}
