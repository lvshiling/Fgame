package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/cross/chat/logic"
	noticelogic "fgame/fgame/cross/notice/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//Boss物品掉落
func battleObjectDrop(target event.EventTarget, data event.EventData) (err error) {
	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	dropItem := data.(scene.DropItem)
	s := n.GetScene()

	// if !s.MapTemplate().IsCrossWorldBoss() {
	// 	return
	// }
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeArenaShengShou {
		return
	}

	p := s.GetPlayer(dropItem.GetOwnerId())
	if p == nil {
		return
	}

	itemId := dropItem.GetItemId()
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		log.Warningf("crossworldboss:物品模板不存在,itemId:%d", itemId)
		return
	}
	qualityType := itemtypes.ItemQualityType(itemTemp.Quality)
	if qualityType < itemtypes.ItemQualityTypeOrange {
		return
	}
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name))
	killName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(p.GetName()))
	dropItemName := coreutils.FormatColor(qualityType.GetColor(), fmt.Sprintf("[%s]", itemTemp.Name))
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
