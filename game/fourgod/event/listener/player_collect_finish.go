package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	fourgodlogic "fgame/fgame/game/fourgod/logic"
	pbuitl "fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	"fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//通用采集 采集完成关联
func playerCollectFinishWith(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(scene.Player)
	if !ok {
		return
	}
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*collecteventtypes.CollectFinishWithEventData)
	if !ok {
		return
	}

	s := spl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}

	npc := eventData.GetCollectNpc()
	npcId := npc.GetId()
	biologyId := int32(npc.GetBiologyTemplate().TemplateId())

	//获取奖励
	dropItemList := eventData.GetItemDataList()
	var rewItemList []*droptemplate.DropItemData
	// var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) > 0 {
		rewItemList, _ = droplogic.SeperateItemDatas(dropItemList)
	}

	//消耗钥匙
	fourManager := pl.GetPlayerDataManager(playertypes.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	keyNum := fourManager.GetKeyNum()
	boxTemplate := template.GetFourGodTemplateService().GetFourGodBoxTemplateByBiologyId(biologyId)
	useKeyNum := boxTemplate.UseItemCount
	if keyNum < useKeyNum {
		useKeyNum = keyNum
	}
	fourManager.OpenBox(rewItemList, useKeyNum)

	//场景内广播副本生物变化信息
	allPlayer := s.GetAllPlayers()
	scFourGodBioBroadcast := pbuitl.BuildSCFourGodBioBroadcast(npcId, npc)
	fourgodlogic.BroadcastMsgInScene(allPlayer, scFourGodBioBroadcast)

	//判断物品品质
	isNeedBroast := false
	itemContent := ""
	for _, itemData := range rewItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()

		itemTemplate := item.GetItemService().GetItem(int(itemId))
		quality := itemtypes.ItemQualityType(itemTemplate.Quality)
		if quality > itemtypes.ItemQualityTypePurple {
			isNeedBroast = true

			itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(num)))
			linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
			itemNameLink := coreutils.FormatLink(itemName, linkArgs)
			itemContent += itemNameLink
		}
	}
	//物品品质大于3广播
	if isNeedBroast {
		playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FourGodOpenBox), playerName, useKeyNum, itemContent)
		//TODO 跑马灯
		noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
		//系统频道
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	}

	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinishWith, event.EventListenerFunc(playerCollectFinishWith))
}
