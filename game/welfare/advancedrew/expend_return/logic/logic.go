package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewexpendreturntypes "fgame/fgame/game/welfare/advancedrew/expend_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 升阶消耗返还
func AdvancedExpendReturn(pl player.Player, danNum int32, advancedType welfaretypes.AdvancedType) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*advancedrewexpendreturntypes.AdvancedExpendReturnInfo)
		if advancedType != info.RewType {
			continue
		}

		info.DanNum += danNum
		welfareManager.UpdateObj(obj)
	}

}
