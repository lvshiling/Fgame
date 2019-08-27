package listener

import (
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	huhudroptypes "fgame/fgame/game/welfare/huhu/drop/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//捡取物品
func dropItemGet(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)
	pl := data.(player.Player)

	itemId := dropItem.GetItemId()
	num := dropItem.GetItemNum()

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if !itemTemp.IsYunYingItem() {
		return
	}

	typ := welfaretypes.OpenActivityTypeHuHu
	subType := welfaretypes.OpenActivitySpecialSubTypeDrop
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		if !itemTemp.IsRelationToGroup(groupId) {
			continue
		}

		welfareManager.RefreshActivityData(typ, subType)

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*huhudroptypes.HuHuInfo)
		info.CurDayDropNum += num
		welfareManager.UpdateObj(obj)
	}

	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemGet, event.EventListenerFunc(dropItemGet))
}
