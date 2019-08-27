package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	fourgodlogic "fgame/fgame/game/fourgod/logic"
	pbuitl "fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	"fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"
)

//四神遗迹宝箱采集完成(废弃，使用通用的)
func fourGodCollectBoxFinish(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	npcId, ok := data.(int64)
	if !ok {
		return
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	fourGodWarSceneData, ok := sd.(fourgodscene.FourGodWarSceneData)
	if !ok {
		return
	}
	npc := fourGodWarSceneData.GetNpc(npcId)
	if npc == nil {
		return
	}

	//获取玩家身上钥匙
	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	keyNum := manager.GetKeyNum()
	// keyMax := template.GetFourGodTemplateService().GetFourGodConstTemplate().KeyMax
	// useKeyNum := keyNum
	// if keyNum > keyMax {
	// 	useKeyNum = keyMax
	// }
	// boxTemplate := template.GetFourGodTemplateService().GetFourGodBoxTemplate(useKeyNum)
	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	boxTemplate := template.GetFourGodTemplateService().GetFourGodBoxTemplateByBiologyId(biologyId)
	if boxTemplate == nil {
		//清空采集记录
		fourGodWarSceneData.CollectBoxInterrupt(npcId)
		keyNumStr := fmt.Sprintf("%d", boxTemplate.UseItemCount)
		playerlogic.SendSystemMessage(pl, lang.FourGodKeyNoEnough, keyNumStr)
		return
	}
	useKeyNum := boxTemplate.UseItemCount
	if keyNum < useKeyNum {
		useKeyNum = keyNum
	}

	//获取奖励
	//template校验必掉落
	dropIdList := boxTemplate.GetDropIdList()
	dropItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)

	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(dropItemList)
	}

	//奖励物品
	isReturn := fourgodlogic.GiveFourGodOpenBoxReward(pl, boxTemplate, rewItemList, resMap, keyNum)
	if isReturn {
		//清空采集记录
		fourGodWarSceneData.CollectBoxInterrupt(npcId)
		return
	}
	//清空钥匙
	manager.OpenBox(rewItemList, useKeyNum)
	//采集完成
	fourGodWarSceneData.FinishCollectBox(npc)

	//推送采集完成
	scFourGodOpenBoxFinish := pbuitl.BuildSCFourGodOpenBoxFinish(npcId)
	pl.SendMsg(scFourGodOpenBoxFinish)

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
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FourGodOpenBox), playerName, keyNum, itemContent)
		//TODO 跑马灯
		noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
		//系统频道
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	}
	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodCollectBoxFinish, event.EventListenerFunc(fourGodCollectBoxFinish))
}
