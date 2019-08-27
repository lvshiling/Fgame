package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	rewextendedtemplate "fgame/fgame/game/welfare/advancedrew/rew_extended/template"
	advancedrewrewextendedtypes "fgame/fgame/game/welfare/advancedrew/rew_extended/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 进阶数据更新(随功能开启)
func UpdateAdvancedRewExtendedData(pl player.Player, advancedNum int32, advancedType welfaretypes.AdvancedType) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRewExtended
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		obj := welfareManager.GetOpenActivity(groupId)
		if obj == nil {
			continue
		}
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*rewextendedtemplate.GroupTemplateRewExtended)
		groupAdvancedType := groupTemp.GetAdvancedType()
		if advancedType != groupAdvancedType {
			continue
		}

		info := obj.GetActivityData().(*advancedrewrewextendedtypes.AdvancedRewExtendedInfo)
		info.AdvancedNum = advancedNum
		welfareManager.UpdateObj(obj)
	}
}
