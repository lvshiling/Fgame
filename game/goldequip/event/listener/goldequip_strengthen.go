package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/item/item"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fmt"
)

const (
	noticeLimit = 7
)

//玩家元神金装强化
func playerGoldEquipStrengSuccess(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	itemData := data.(*droptemplate.DropItemData)
	if itemData == nil {
		return
	}
	if itemData.GetLevel() < noticeLimit {
		return
	}
	goldEquipLevel := itemData.GetLevel()
	itemId := int(itemData.GetItemId())
	itemTemplate := item.GetItemService().GetItem(itemId)
	if itemTemplate == nil {
		return
	}
	if !itemTemplate.IsNotice() {
		return
	}

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	equipName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStrUnderline(itemTemplate.Name))
	args := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
	equipNameLink := coreutils.FormatLink(equipName, args)
	levelStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("+%d", goldEquipLevel))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.GoldEquipStrengthenNotice), playerName, equipNameLink, levelStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipStrengSuccess, event.EventListenerFunc(playerGoldEquipStrengSuccess))
}
