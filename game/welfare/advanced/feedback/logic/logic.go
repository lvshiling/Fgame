package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedfeedbacktypes "fgame/fgame/game/welfare/advanced/feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 升阶返利数据更新（废弃）
func UpdateAdvancedActivityData(pl player.Player, danNum int32, advancedType welfaretypes.AdvancedType) {
	typ := welfaretypes.OpenActivityTypeAdvanced
	subType := welfaretypes.OpenActivityAdvancedSubTypeFeedback
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		//进阶类型
		advancedDay := welfarelogic.CountCurActivityDay(groupId)
		if advancedDay != int32(advancedType) {
			return
		}

		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		//刷新
		err := welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			return
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)

		info := obj.GetActivityData().(*advancedfeedbacktypes.AdvancedInfo)
		info.DanNum += danNum
		welfareManager.UpdateObj(obj)
	}
}
