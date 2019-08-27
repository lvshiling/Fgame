package listener

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	guidereplicaeventtypes "fgame/fgame/game/guidereplica/event/types"
	"fgame/fgame/game/guidereplica/pbutil"
	guidereplicascene "fgame/fgame/game/guidereplica/scene"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

func guideSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	success, ok := data.(bool)
	if !ok {
		return
	}

	//猫狗奖励
	catDogFinish(pl, success)
	// 救援奖励
	rescureFinish(pl, success)
	return
}

func catDogFinish(pl player.Player, success bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeGuideReplicaCatDog {
		return
	}

	catSd, ok := s.SceneDelegate().(guidereplicascene.GuideCatDogSceneData)
	if !ok {
		return
	}

	var showItemList []*droptemplate.DropItemData
	if success {
		guideTemp := catSd.GetGuideTemp()
		killMap := catSd.GetKillMap()
		dropIdList := guideTemp.GetCatDogGuideTemp().GetDropId(killMap)
		showItemList = addFinishRew(pl, dropIdList, guideTemp.GetGuideType())
	}

	scMsg := pbutil.BuildSCGuideReplicaChallengeResult(success, showItemList)
	pl.SendMsg(scMsg)
	return
}

func rescureFinish(pl player.Player, success bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeGuideReplicaRescue {
		return
	}
	rescureSd, ok := s.SceneDelegate().(guidereplicascene.GuideRescureSceneData)
	if !ok {
		return
	}
	guideTemp := rescureSd.GetGuideTemp()
	rescureTemp := guideTemp.GetRescureGuideTemp()

	// 完成奖励
	var showItemList []*droptemplate.DropItemData
	if success {
		dropIdList := rescureTemp.GetDropId()
		showItemList = addFinishRew(pl, dropIdList, guideTemp.GetGuideType())
	}

	scMsg := pbutil.BuildSCGuideReplicaChallengeResult(success, showItemList)
	pl.SendMsg(scMsg)
	return
}

func addFinishRew(pl player.Player, dropIdList []int32, guideType guidereplicatypes.GuideReplicaType) (showItemList []*droptemplate.DropItemData) {
	var newItemDataList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	itemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
	newItemDataList, resMap = droplogic.SeperateItemDatas(itemList)

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		//写邮件
		emailTitle := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
		emailContent := lang.GetLangService().ReadLang(lang.GuideFinishNotSlotsMailContent)
		now := global.GetGame().GetTimeService().Now()

		emaillogic.AddEmailItemLevel(pl, emailTitle, emailContent, now, itemList)
	} else {
		itemReason := commonlog.InventoryLogReasonGuideRew
		itemReasonText := fmt.Sprintf(itemReason.String(), guideType.String())
		flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemReason, itemReasonText)
		if !flag {
			panic(fmt.Errorf("guide: guide BatchAdd should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//增加资源
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonGuideRew
		silverReason := commonlog.SilverLogReasonGuideRew
		levelReason := commonlog.LevelLogReasonGuideRew
		goldReasonText := fmt.Sprintf(goldReason.String(), guideType.String())
		silverReasonText := fmt.Sprintf(silverReason.String(), guideType.String())
		levelReasonText := fmt.Sprintf(levelReason.String(), guideType.String())
		err := droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
		propertylogic.SnapChangedProperty(pl)
	}

	showItemList = itemList
	return
}

func init() {
	gameevent.AddEventListener(guidereplicaeventtypes.EventTypeGuideFinish, event.EventListenerFunc(guideSceneFinish))
}
