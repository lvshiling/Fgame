package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	playerfuncopen "fgame/fgame/game/funcopen/player"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"
	"fgame/fgame/game/secretcard/secretcard"
	secretcardtypes "fgame/fgame/game/secretcard/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//天机牌任务完成
func FinishSecretCardQuest(pl player.Player, questId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	cardObj := manager.GetSecretCard()
	cardId := cardObj.CardId
	if cardId == 0 {
		return
	}
	curQuestId, questName, _ := secretcard.GetSecretCardService().GetQuestIdByCardId(cardId)
	if curQuestId != questId {
		panic(fmt.Errorf("secretcard: FinishSecretCardQuest questId be equal"))
	}
	manager.SecretCardFinish(cardObj.Star)

	var dropItemMap []*droptemplate.DropItemData
	dropId := secretcard.GetSecretCardService().GetDropIdByNum(cardId, int32(cardObj.TotalNum))
	if dropId != 0 {
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData != nil {
			dropItemMap = append(dropItemMap, dropData)

			//公告
			itemId := dropData.GetItemId()
			num := dropData.GetNum()
			inventorylogic.SecretCardPrecioustemBroadcast(pl, itemId, num)
		}
	}

	secretCardTemplate := secretcard.GetSecretCardService().GetSecretCardTemplate(cardId)
	if secretCardTemplate != nil && secretCardTemplate.RewSilver != 0 {
		newData := droptemplate.CreateItemData(constanttypes.SilverItem, secretCardTemplate.RewSilver, 0, itemtypes.ItemBindTypeUnBind)
		dropItemMap = append(dropItemMap, newData)
	}

	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemMap) > 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(dropItemMap)
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//判断背包是否足够
	if len(newItemList) != 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
			//写邮件
			emailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ScecretCardFinishTitle), questName)
			emailContent := lang.GetLangService().ReadLang(lang.ScecretCardFinishContent)
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, emailTitle, emailContent, now, newItemList)
		} else {
			reasonText := commonlog.InventoryLogReasonSecretCardFinish.String()
			flag := inventoryManager.BatchAddOfItemLevel(newItemList, commonlog.InventoryLogReasonSecretCardFinish, reasonText)
			if !flag {
				panic(fmt.Errorf("secretcard: FinishSecretCardQuest BatchAdd should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//添加掉落资源
	if len(resMap) != 0 {
		reasonGoldText := commonlog.GoldLogReasonSecretCardFinish.String()
		reasonSliverText := commonlog.SilverLogReasonSecretCardFinish.String()
		reasonLevelText := commonlog.LevelLogReasonSecretCardFinish.String()
		err = droplogic.AddRes(pl, resMap, commonlog.GoldLogReasonSecretCardFinish, reasonGoldText, commonlog.SilverLogReasonSecretCardFinish, reasonSliverText, commonlog.LevelLogReasonSecretCardFinish, reasonLevelText)
		if err != nil {
			return
		}
		propertylogic.SnapChangedProperty(pl)
	}
	scQuestSecretFinish := pbutil.BuildSCQuestSecretFinish(dropItemMap, false)
	err = pl.SendMsg(scQuestSecretFinish)
	return
}

//获取天机任务
func GetSecretCardSpy(pl player.Player) (cardMap map[int32]int32) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	opendManager := pl.GetPlayerDataManager(types.PlayerFuncOpenDataManagerType).(*playerfuncopen.PlayerFuncOpenDataManager)
	cardPoolMap := secretcard.GetSecretCardService().GetQuestPool()
	usedCardsPool := manager.GetSecretCard().UsedCardList
	lowSize := secretcardtypes.SpyLowPoolNum
	totalSize := secretcardtypes.SpyTotalNum
	level := pl.GetLevel()
	tempcardPoolMap := make(map[secretcardtypes.SecretCardPoolType][]int32)
	for poolTyp, questPool := range cardPoolMap {
		for _, tempTamplate := range questPool {
			cardId := int32(tempTamplate.TemplateId())
			levelMin := tempTamplate.LevelMin
			levelMax := tempTamplate.LevelMax
			if level < levelMin || level > levelMax {
				continue
			}
			flag := utils.ContainInt32(usedCardsPool, cardId)
			if flag {
				continue
			}
			if tempTamplate.ModuleOpenedId != 0 {
				if !opendManager.IsOpen(tempTamplate.GetFuncOpenTyp()) {
					continue
				}
			}

			poolList := tempcardPoolMap[poolTyp]
			poolList = append(poolList, cardId)
			tempcardPoolMap[poolTyp] = poolList
		}
	}
	cardIdList := make([]int32, 0, totalSize)
	lowList, existLow := tempcardPoolMap[secretcardtypes.SecretCardPoolTypeNormal]
	if existLow {
		//低级池取两个
		cardIdList = mathutils.RandomList(lowList, lowSize)
	}
	curSize := int32(len(cardIdList))
	//高级池
	seniorList, existSenior := tempcardPoolMap[secretcardtypes.SecretCardPoolTypeSenior]
	if existSenior {
		listSenior := mathutils.RandomList(seniorList, totalSize-curSize)
		for i := 0; i < len(listSenior); i++ {
			cardIdList = append(cardIdList, listSenior[i])
		}
	}
	//轮询池
	pollList := tempcardPoolMap[secretcardtypes.SecretCardPoolTypePoll]
	for curSize = int32(len(cardIdList)); curSize < totalSize; {
		listPoll := mathutils.RandomList(pollList, totalSize-curSize)
		for i := 0; i < len(listPoll); i++ {
			cardIdList = append(cardIdList, listPoll[i])
		}
		curSize = int32(len(cardIdList))
	}
	cardMap = secretcard.GetSecretCardService().GetRandomStar(cardIdList)
	return
}

//获取一键完成天机任务
func GetSecretCardSpyByGold(pl player.Player, num int32) (cardIdList []int32) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	opendManager := pl.GetPlayerDataManager(types.PlayerFuncOpenDataManagerType).(*playerfuncopen.PlayerFuncOpenDataManager)
	cardPoolMap := secretcard.GetSecretCardService().GetQuestPool()
	usedCardsPool := manager.GetSecretCard().UsedCardList
	tempcardPoolList := make([]int32, 0, 8)
	pollPoolList := make([]int32, 0, 8)
	level := pl.GetLevel()
	for poolTyp, questPool := range cardPoolMap {
		for _, tempTamplate := range questPool {
			cardId := int32(tempTamplate.TemplateId())
			levelMin := tempTamplate.LevelMin
			levelMax := tempTamplate.LevelMax
			if level < levelMin || level > levelMax {
				continue
			}
			flag := utils.ContainInt32(usedCardsPool, cardId)
			if flag {
				continue
			}
			if tempTamplate.ModuleOpenedId != 0 {
				if !opendManager.IsOpen(tempTamplate.GetFuncOpenTyp()) {
					continue
				}
			}
			if poolTyp == secretcardtypes.SecretCardPoolTypePoll {
				pollPoolList = append(pollPoolList, cardId)
			} else {
				tempcardPoolList = append(tempcardPoolList, cardId)
			}
		}
	}
	if len(tempcardPoolList) != 0 {
		cardIdList = mathutils.RandomList(tempcardPoolList, num)
	}
	curSize := int32(len(cardIdList))
	if curSize == num {
		return
	}
	for curSize < num {
		listPoll := mathutils.RandomList(tempcardPoolList, num-curSize)
		for i := 0; i < len(listPoll); i++ {
			cardIdList = append(cardIdList, listPoll[i])
		}
		curSize = int32(len(cardIdList))
	}
	return
}

//天机牌宝箱奖励
func GiveSecretCardBoxReward(pl player.Player, starTemplate *gametemplate.TianJiPaiStarTemplate) (isReturn bool) {
	rewData := starTemplate.GetRewData()
	rewItemMap := starTemplate.GetRewItemMap()
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//奖励物品
	if len(rewItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			isReturn = true
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		inventoryLogReason := commonlog.InventoryLogReasonSecretCardStarRew
		reasonText := inventoryLogReason.String()
		flag = inventoryManager.BatchAdd(rewItemMap, inventoryLogReason, reasonText)
		if !flag {
			panic(fmt.Errorf("secretcard: secretStarRew BatchAdd should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//添加奖励属性
	if rewData != nil {
		goldLog := commonlog.GoldLogReasonSecretCardStarRew
		silverLog := commonlog.SilverLogReasonSecretCardStarRew
		levelReason := commonlog.LevelLogReasonSecretCardStarRew
		goldReasonText := goldLog.String()
		silverReasonText := silverLog.String()
		reasonText := levelReason.String()
		flag := propertyManager.AddRewData(rewData, goldLog, goldReasonText, silverLog, silverReasonText, levelReason, reasonText)
		if !flag {
			panic(fmt.Errorf("secretcard: secretStarRew AddRewData should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}
	return
}

//天机牌直接完成奖励
func GiveSecretCardImmediateFinishReward(pl player.Player, cardObj *playersecretcard.PlayerSecretCardObject) (dropItemMap []*droptemplate.DropItemData, err error) {
	totalNum := cardObj.TotalNum
	cardId := cardObj.CardId

	//获取dropId
	dropId := secretcard.GetSecretCardService().GetDropIdByNum(cardId, int32(totalNum))
	if dropId == 0 {
		return
	}
	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
	if dropData != nil {
		dropItemMap = append(dropItemMap, dropData)
		//公告
		itemId := dropData.GetItemId()
		num := dropData.GetNum()
		inventorylogic.SecretCardPrecioustemBroadcast(pl, itemId, num)

	}

	secretCardTemplate := secretcard.GetSecretCardService().GetSecretCardTemplate(cardId)
	if secretCardTemplate != nil && secretCardTemplate.RewSilver != 0 {
		newData := droptemplate.CreateItemData(constanttypes.SilverItem, secretCardTemplate.RewSilver, 0, itemtypes.ItemBindTypeUnBind)
		dropItemMap = append(dropItemMap, newData)
	}

	dropNewData, resMap := droplogic.SeperateItemDatas(dropItemMap)

	//判断背包是否足够
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if dropNewData != nil {
		itemList := make([]*droptemplate.DropItemData, 0, 8)
		itemList = append(itemList, dropNewData...)
		if !inventoryManager.HasEnoughSlotsOfItemLevel(itemList) {
			//写邮件
			emailTitle := lang.GetLangService().ReadLang(lang.SecretCardFinishAllTitle)
			emailContent := lang.GetLangService().ReadLang(lang.SecretCardFinishAllContent)
			now := global.GetGame().GetTimeService().Now()

			emaillogic.AddEmailItemLevel(pl, emailTitle, emailContent, now, itemList)
		} else {
			reasonText := commonlog.InventoryLogReasonSecretCardFinishAll.String()
			flag := inventoryManager.BatchAddOfItemLevel(itemList, commonlog.InventoryLogReasonSecretCardFinishAll, reasonText)
			if !flag {
				panic(fmt.Errorf("secretcard: secretFinish BatchAdd should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//添加掉落资源
	if len(resMap) != 0 {
		reasonGoldText := commonlog.GoldLogReasonSecretCardFinishAll.String()
		reasonSliverText := commonlog.SilverLogReasonSecretCardFinishAll.String()
		reasonLevelText := commonlog.LevelLogReasonSecretCardFinishAll.String()
		err = droplogic.AddRes(pl, resMap, commonlog.GoldLogReasonSecretCardFinishAll, reasonGoldText, commonlog.SilverLogReasonSecretCardFinishAll, reasonSliverText, commonlog.LevelLogReasonSecretCardFinishAll, reasonLevelText)
		if err != nil {
			return
		}
	}
	return
}

//天机牌一键完成奖励
func GiveSecretCardFinishAllReward(pl player.Player, cardObj *playersecretcard.PlayerSecretCardObject, leftNum int32) (dropItemMap []*droptemplate.DropItemData, addStar int32, leftBoxList []int32, err error) {
	totalNum := cardObj.TotalNum
	addStar = 0
	curCardId := cardObj.CardId
	star := cardObj.Star
	cardIdList := make([]int32, 0, 8)

	//获取随机的cardId
	if leftNum != 0 {
		cardIdList = GetSecretCardSpyByGold(pl, leftNum)
	}
	for cardId, cruStar := range cardObj.CardMap {
		star += cruStar
		curCardId = cardId
		break
	}
	if curCardId != 0 {
		cardIdList = append(cardIdList, curCardId)
	}

	silverData := droptemplate.CreateItemData(constanttypes.SilverItem, 0, 0, itemtypes.ItemBindTypeUnBind)
	//随机掉落
	for _, cardId := range cardIdList {
		to := secretcard.GetSecretCardService().GetSecretCardTemplate(cardId)
		if cardId != curCardId {
			//星数统计
			addStar += to.StarMax
		} else {
			addStar += star
		}
		totalNum++
		//获取dropId
		dropId := secretcard.GetSecretCardService().GetDropIdByNum(cardId, int32(totalNum))
		if dropId == 0 {
			continue
		}
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData != nil {
			dropItemMap = append(dropItemMap, dropData)
			//公告
			itemId := dropData.GetItemId()
			num := dropData.GetNum()
			inventorylogic.SecretCardPrecioustemBroadcast(pl, itemId, num)
		}

		//银两奖励
		secretCardTemplate := secretcard.GetSecretCardService().GetSecretCardTemplate(cardId)
		if secretCardTemplate != nil && secretCardTemplate.RewSilver != 0 {
			silverData.Num += secretCardTemplate.RewSilver
		}
	}
	if len(dropItemMap) != 0 {
		dropItemMap = droplogic.MergeItemLevel(dropItemMap)
	}

	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	//未开启的宝箱 默认领取赠送物品
	starTemplateList := secretcard.GetSecretCardService().GetLeftBoxTemplate(cardObj.OpenBoxList)
	if len(starTemplateList) != 0 {
		for _, starTemplate := range starTemplateList {
			totalRewData.RewSilver += starTemplate.AwardSilver
			totalRewData.RewGold += starTemplate.AwardGold
			totalRewData.RewBindGold += starTemplate.AwardBindGold
			leftBoxList = append(leftBoxList, int32(starTemplate.TemplateId()))
			for itemId, num := range starTemplate.GetRewItemMap() {
				//未开启宝箱奖励物品
				newData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeUnBind)
				dropItemMap = append(dropItemMap, newData)

				//公告
				inventorylogic.SecretCardPrecioustemBroadcast(pl, itemId, num)
			}
		}
	}

	var newItemMap []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(dropItemMap) != 0 {
		newItemMap, resMap = droplogic.SeperateItemDatas(dropItemMap)
	}
	//判断背包是否足够
	if len(newItemMap) != 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemMap) {
			//写邮件
			emailTitle := lang.GetLangService().ReadLang(lang.SecretCardFinishAllTitle)
			emailContent := lang.GetLangService().ReadLang(lang.SecretCardFinishAllContent)
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, emailTitle, emailContent, now, newItemMap)
		} else {
			reasonText := commonlog.InventoryLogReasonSecretCardFinishAll.String()
			flag := inventoryManager.BatchAddOfItemLevel(newItemMap, commonlog.InventoryLogReasonSecretCardFinishAll, reasonText)
			if !flag {
				panic(fmt.Errorf("secretcard: secretFinish BatchAdd should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//添加掉落资源
	if len(resMap) != 0 {
		reasonGoldText := commonlog.GoldLogReasonSecretCardFinishAll.String()
		reasonSliverText := commonlog.SilverLogReasonSecretCardFinishAll.String()
		reasonLevelText := commonlog.LevelLogReasonSecretCardFinishAll.String()
		err = droplogic.AddRes(pl, resMap, commonlog.GoldLogReasonSecretCardFinishAll, reasonGoldText, commonlog.SilverLogReasonSecretCardFinishAll, reasonSliverText, commonlog.LevelLogReasonSecretCardFinishAll, reasonLevelText)
		if err != nil {
			return
		}
	}

	//剩余宝箱属性奖励
	silverNum := silverData.Num
	silverData.Num += totalRewData.RewSilver
	totalRewData.RewSilver += silverNum
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if totalRewData.RewSilver != 0 || totalRewData.RewGold != 0 || totalRewData.RewBindGold != 0 {
		goldLog := commonlog.GoldLogReasonSecretCardFinishAllStarRew
		silverLog := commonlog.SilverLogReasonSecretCardFinishAllStarRew
		levelReason := commonlog.LevelLogReasonSecretCardFinishAllStarRew

		goldReasonText := goldLog.String()
		silverReasonText := silverLog.String()
		reasonText := levelReason.String()
		flag := propertyManager.AddRewData(totalRewData, goldLog, goldReasonText, silverLog, silverReasonText, levelReason, reasonText)
		if !flag {
			panic(fmt.Errorf("secretcard: secretFinish AddRewData should be ok"))
		}
	}

	if silverData.Num != 0 {
		dropItemMap = append(dropItemMap, silverData)
	}

	if totalRewData.RewGold != 0 {
		newData := droptemplate.CreateItemData(constanttypes.GoldItem, totalRewData.RewGold, 0, itemtypes.ItemBindTypeUnBind)
		dropItemMap = append(dropItemMap, newData)
	}

	if totalRewData.RewBindGold != 0 {
		newData := droptemplate.CreateItemData(constanttypes.BindGoldItem, totalRewData.RewBindGold, 0, itemtypes.ItemBindTypeUnBind)
		dropItemMap = append(dropItemMap, newData)
	}
	return
}
