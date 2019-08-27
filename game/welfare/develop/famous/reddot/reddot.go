package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	developfamoustypes "fgame/fgame/game/welfare/develop/famous/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeDevelop, welfaretypes.OpenActivityDefaultSubTypeDefault, reddot.HandlerFunc(handleRedDotDevelopFame))
}

//名人普红点
func handleRedDotDevelopFame(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*developfamoustypes.DevelopFameInfo)

	// //培养
	// famousTemp := welfaretemplate.GetWelfareTemplateService().GetFamousTemplate(groupId)
	// inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// feedDataList := famousTemp.GetFameFeedInfoList()
	// for _, feedData := range feedDataList {
	// 	if !info.IfCanFeed(feedData.ItemId, 1, feedData.FeedLimit) {
	// 		continue
	// 	}
	// 	if !inventoryManager.HasEnoughItem(feedData.ItemId, feedData.ItemNum) {
	// 		continue
	// 	}
	// 	isNotice = true
	// 	return
	// }

	//奖励
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needFavorable := temp.Value1
		if !info.IsCanReceiveRewards(needFavorable) {
			continue
		}

		isNotice = true
		return
	}

	return
}
