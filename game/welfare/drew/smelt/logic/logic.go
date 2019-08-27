package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	smelttemplate "fgame/fgame/game/welfare/drew/smelt/template"
	smelttypes "fgame/fgame/game/welfare/drew/smelt/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func ListenEventAddItemNum(pl player.Player, itemMap map[int32]int32) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeSmelt

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*smelttemplate.GroupTemplateSmelt)
		itemNum := itemMap[groupTemp.GetItemId()]
		if itemNum == 0 {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*smelttypes.SmeltInfo)
		info.Num += itemNum
		welfareManager.UpdateObj(obj)
	}
	return
}
