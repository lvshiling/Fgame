package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedblessfeedbacktypes "fgame/fgame/game/welfare/advanced/bless_feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 升阶祝福丹放送数据更新（每天类型变更，顺序随AdvancedType）
func UpdateAdvancedBlessActivityData(pl player.Player, advancedNum int32, advancedType welfaretypes.AdvancedType) {
	typ := welfaretypes.OpenActivityTypeAdvanced
	subType := welfaretypes.OpenActivityAdvancedSubTypeBlessFeedback
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		//进阶类型
		advancedDay := welfarelogic.CountCurActivityDay(groupId)
		if advancedDay != int32(advancedType) {
			continue
		}

		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		//刷新
		err := welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*advancedblessfeedbacktypes.BlessAdvancedInfo)
		info.AdvancedNum = advancedNum
		welfareManager.UpdateObj(obj)
	}
}
