package types

import (
	droptemplate "fgame/fgame/game/drop/template"
)

//排行榜奖励邮件回调数据
type LongGongRankEmailData struct {
	Title           string
	Econtent        string
	EndTime         int64
	RewItemDataList []*droptemplate.DropItemData
}

func NewLongGongRankEmailData(title string, econtent string, endTime int64, rewItemDataList []*droptemplate.DropItemData) *LongGongRankEmailData {
	d := &LongGongRankEmailData{
		Title:           title,
		Econtent:        econtent,
		EndTime:         endTime,
		RewItemDataList: rewItemDataList,
	}
	return d
}
