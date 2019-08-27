package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	investleveltemplate "fgame/fgame/game/welfare/invest/level/template"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//加载完成后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeLevel
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	backGold := int32(0)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		obj := welfareManager.GetOpenActivity(groupId)
		if obj == nil {
			continue
		}
		info := obj.GetActivityData().(*investleveltypes.InvestLevelInfo)
		if info.IsBack {
			continue
		}
		groupTemp := groupInterface.(*investleveltemplate.GroupTemplateInvestLevel)
		for investLevelType, receiveLevelRecord := range info.InvestBuyInfoMap {
			maxRewards := groupTemp.GetInvestLevelMaxRewardsLevel(investLevelType)
			if receiveLevelRecord >= maxRewards {
				continue
			}
			backGold += groupTemp.GetInvestLevelNeedGold(investLevelType)
		}
		info.IsBack = true
		welfareManager.UpdateObj(obj)
	}

	if backGold > 0 {
		now := global.GetGame().GetTimeService().Now()
		title := lang.GetLangService().ReadLang(lang.OpenActivityInvestLevelBackGoldLabel)
		econtent := lang.GetLangService().ReadLang(lang.OpenActivityInvestLevelBackGoldContent)
		rewItemMap := make(map[int32]int32)
		rewItemMap[constanttypes.GoldItem] = backGold
		newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, inventorytypes.NewItemLimitTimeTypeNone, 0)
		emaillogic.AddEmailItemLevel(pl, title, econtent, now, newItemDataList)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
