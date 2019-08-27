package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 累计活动数据排行
func UpdateAddCountRankData(pl player.Player, attendGroupId, attendNum int32) {
	typ := welfaretypes.OpenActivityTypeRank
	subType := welfaretypes.OpenActivityRankSubTypeNumber
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		if !timeTemp.IsRelationToGroup(attendGroupId) {
			continue
		}

		welfareManager.AddActivityNumRecordRecord(groupId, attendNum)
	}
}

// 设置活动数据排行
func UpdateSetCountRankData(pl player.Player, attendGroupId, attendNum int32) {
	typ := welfaretypes.OpenActivityTypeRank
	subType := welfaretypes.OpenActivityRankSubTypeNumber
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		if !timeTemp.IsRelationToGroup(attendGroupId) {
			continue
		}

		welfareManager.SetActivityNumRecordRecord(groupId, attendNum)
	}
}

// 设置活动数据排行(多场排行对一场活动)
func UpdateSetDayCountRankData(pl player.Player, realteGroupId, adddNum int32) {
	typ := welfaretypes.OpenActivityTypeRank
	subType := welfaretypes.OpenActivityRankSubTypeNumberDay
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		if !timeTemp.IsRelationToGroup(realteGroupId) {
			continue
		}
		welfareManager.SetActivityNumRecordRecord(groupId, adddNum)
	}
}
