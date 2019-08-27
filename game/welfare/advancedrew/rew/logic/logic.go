package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewrewtypes "fgame/fgame/game/welfare/advancedrew/rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 升阶奖励数据更新(每个类型一个活动id，时间、顺序自由定义)
func UpdateAdvancedRewData(pl player.Player, advancedNum int32, advancedType welfaretypes.AdvancedType) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRew
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

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)
		if info.RewType != advancedType {
			continue
		}

		info.AdvancedNum = advancedNum
		welfareManager.UpdateObj(obj)

	}
}
