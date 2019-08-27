package listener

import (
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	gameevent "fgame/fgame/game/event"
	friendlogic "fgame/fgame/game/friend/logic"
	friendtemplate "fgame/fgame/game/friend/template"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//Boss物品掉落 直接入背包
func battleObjectDropIntoInventory(target event.EventTarget, data event.EventData) (err error) {
	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	eventData, ok := data.(*sceneeventtypes.BattleObjectDropIntoInventoryData)
	if !ok {
		return
	}

	if n.GetBiologyTemplate().GetBiologySetType() != scenetypes.BiologySetTypeWorldBoss {
		return
	}

	itemDataList := eventData.GetItemDataList()
	owerId := eventData.GetOwerId()

	s := n.GetScene()
	p := s.GetPlayer(owerId)
	if p == nil {
		return
	}
	pl := p.(player.Player)

	noticeType := friendtypes.FriendNoticeTypeKillBoss
	noticeTempList := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplate(noticeType)
	if len(noticeTempList) < 1 {
		return
	}

	noticeTemp := noticeTempList[0]
	for _, itemData := range itemDataList {
		itemId := itemData.GetItemId()
		itemTemp := item.GetItemService().GetItem(int(itemId))
		if itemTemp == nil {
			continue
		}

		if itemTemp.Quality < noticeTemp.TiaoJian {
			continue
		}

		if itemData.BindType == itemtypes.ItemBindTypeBind {
			continue
		}

		// 推送消息
		args := coreutils.FormatParamsAsString(n.GetBiologyTemplate().Name, itemId, itemData.Num)
		friendlogic.BroadcastFriendNotice(pl, noticeType, noticeTemp.TiaoJian, args)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDropIntoInventory, event.EventListenerFunc(battleObjectDropIntoInventory))
}
