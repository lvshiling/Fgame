package reddot

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/reddot/reddot"
	drewchargedrewtemplate "fgame/fgame/game/welfare/drew/charge_drew/template"
	drewchargedrewtypes "fgame/fgame/game/welfare/drew/charge_drew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeChargeDrew, reddot.HandlerFunc(handleRedDotChargeDrew))
}

//充值抽奖
func handleRedDotChargeDrew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	// 有免费次数
	info := obj.GetActivityData().(*drewchargedrewtypes.LuckyChargeDrewInfo)
	if info.LeftTimes > 0 {
		isNotice = true
		return
	}

	// 有元宝
	groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInteface == nil {
		return
	}
	groupTemp := groupInteface.(*drewchargedrewtemplate.GroupTemplateChargeDrew)
	needGold := groupTemp.GetChargeDrewNeedGold()
	if needGold > 0 {
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if propertyManager.HasEnoughGold(needGold, false) {
			isNotice = true
			return
		}
	}

	// 有物品
	luckTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplate(groupId)
	needItemMap := luckTemp.GetUseItemMap()
	if len(needItemMap) > 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if inventoryManager.HasEnoughItems(needItemMap) {
			isNotice = true
			return
		}
	}

	return
}
