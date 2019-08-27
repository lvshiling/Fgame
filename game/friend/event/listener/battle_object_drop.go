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

//Boss物品掉落
func battleObjectDrop(target event.EventTarget, data event.EventData) (err error) {
	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	dropItem := data.(scene.DropItem)

	if n.GetBiologyTemplate().GetBiologySetType() != scenetypes.BiologySetTypeWorldBoss {
		return
	}

	s := n.GetScene()
	p := s.GetPlayer(dropItem.GetOwnerId())
	if p == nil {
		return
	}
	pl, ok := p.(player.Player)
	if !ok {
		return
	}

	itemId := dropItem.GetItemId()
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return
	}
	noticeType := friendtypes.FriendNoticeTypeKillBoss
	noticeTempList := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplate(noticeType)
	if len(noticeTempList) < 1 {
		return
	}

	noticeTemp := noticeTempList[0]
	if itemTemp.Quality < noticeTemp.TiaoJian {
		return
	}

	if dropItem.GetBindType() == itemtypes.ItemBindTypeBind {
		return
	}

	// 推送消息
	args := coreutils.FormatParamsAsString(n.GetBiologyTemplate().Name, itemId, dropItem.GetItemNum())
	friendlogic.BroadcastFriendNotice(pl, noticeType, noticeTemp.TiaoJian, args)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDrop, event.EventListenerFunc(battleObjectDrop))
}
