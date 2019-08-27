package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	powertemplate "fgame/fgame/game/welfare/advancedrew/power/template"
	advancedrewpowertypes "fgame/fgame/game/welfare/advancedrew/power/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 升阶战力数据更新(随功能开启)
func UpdateAdvancedPowerData(pl player.Player, power int64, advancedType welfaretypes.AdvancedType) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypePower
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		//刷新
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		err := welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			continue
		}

		obj := welfareManager.GetOpenActivity(groupId)
		if obj == nil {
			continue
		}
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*powertemplate.GroupTemplatePower)
		groupAdvancedType := groupTemp.GetAdvancedType()
		if advancedType != groupAdvancedType {
			continue
		}

		info := obj.GetActivityData().(*advancedrewpowertypes.AdvancedPowerInfo)
		info.Power = power
		welfareManager.UpdateObj(obj)
	}
}
