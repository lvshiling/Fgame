package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	noticelogic "fgame/fgame/game/notice/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//Boss物品掉落
func battleObjectDrop(target event.EventTarget, data event.EventData) (err error) {
	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	dropItem := data.(scene.DropItem)
	s := n.GetScene()

	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeDingShiBoss {
		return
	}

	p := s.GetPlayer(dropItem.GetOwnerId())
	if p == nil {

		return
	}

	itemId := dropItem.GetItemId()
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {

		return
	}
	if !itemTemp.IsNotice() {
		return
	}

	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name))
	killName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(p.GetName()))
	dropItemName := coreutils.FormatColor(itemTemp.GetQualityType().GetColor(), coreutils.FormatNoticeStr(itemTemp.Name))

	args := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
	infoLink := coreutils.FormatLink(dropItemName, args)

	//系统广播
	format := lang.GetLangService().ReadLang(lang.WorldBossBeKilled)
	content := fmt.Sprintf(format, bossName, killName, infoLink)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	//跑马灯
	noticeFormat := lang.GetLangService().ReadLang(lang.WorldBossBeKilledNotice)
	noticeContent := fmt.Sprintf(noticeFormat, bossName, killName, infoLink)
	noticelogic.NoticeNumBroadcast([]byte(noticeContent), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDrop, event.EventListenerFunc(battleObjectDrop))
}
