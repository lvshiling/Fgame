package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	consttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	questtemplate "fgame/fgame/game/quest/template"
	"fmt"
	"math"
)

//屠魔任务直接完成奖励
func GiveQuestTuMoImmediateFinishReward(pl player.Player, questId int32) (itemMap map[int32]int32, err error) {
	itemMap = make(map[int32]int32)
	expBase := int64(0)
	silver := int64(0)
	gold := int32(0)
	bindGold := int32(0)
	zhuanshu := int32(0)
	rewItemMap := make(map[int32]int32)

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	rewDataMap := make(map[int32]int32)
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	silver += int64(math.Ceil(float64(questTemplate.RewSilver) * pl.GetWallowState().Rate()))
	gold += int32(math.Ceil(float64(questTemplate.RewGold) * pl.GetWallowState().Rate()))
	bindGold += int32(math.Ceil(float64(questTemplate.RewBindGold) * pl.GetWallowState().Rate()))

	rewData := questTemplate.GetRewData()
	if rewData != nil {
		rewDataMap = propertylogic.GetRewDataMap(rewData, pl.GetLevel())
		exp, exist := rewDataMap[consttypes.ExpItem]
		if exist {
			expBase += int64(exp)
		}
	}

	if questTemplate.RewZhuanshu != 0 && questTemplate.RewZhuanshu > propertyManager.GetZhuanSheng() && questTemplate.RewZhuanshu > zhuanshu {
		zhuanshu = questTemplate.RewZhuanshu
	}

	//奖励
	if pl.GetWallowState() == playertypes.WallowStateNone {
		rewardItemMap := questTemplate.GetRewardItemMap(pl.GetRole())
		if len(rewardItemMap) != 0 {
			for itemId, num := range rewardItemMap {
				tempNum, exist := itemMap[itemId]
				if exist {
					tempNum += num
				} else {
					tempNum = num
				}
				itemMap[itemId] = tempNum
				curNum, exist := rewItemMap[itemId]
				if exist {
					curNum += num
				} else {
					curNum = num
				}
				rewItemMap[itemId] = curNum
			}
		}
	}

	//固定奖励(展示用)
	for itemId, num := range rewDataMap {
		curNum, exist := itemMap[itemId]
		if exist {
			curNum += num
		}
		itemMap[itemId] = num
	}

	//奖励物品
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(rewItemMap) != 0 {
		//添加物品
		if !inventoryManager.HasEnoughSlots(rewItemMap) {
			//写邮件
			emailTitle := lang.GetLangService().ReadLang(lang.QuestTuMoImmediateFinishTitle)
			emailContent := lang.GetLangService().ReadLang(lang.QuestTuMoImmediateFinishContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			reasonText := commonlog.InventoryLogReasonTuMoImmediateFinishRew.String()
			flag := inventoryManager.BatchAdd(rewItemMap, commonlog.InventoryLogReasonTuMoImmediateFinishRew, reasonText)
			if !flag {
				panic("quest: GiveQuestTuMoImmediateFinishReward BatchAdd  should be ok")
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//奖励属性
	if expBase != 0 || silver != 0 || gold != 0 || bindGold != 0 {
		reasonGold := commonlog.GoldLogReasonTuMoImmediateFinish

		reasonSilver := commonlog.SilverLogReasonTuMoImmediateFinish

		reasonLevel := commonlog.LevelLogReasonTuMoImmediateFinish
		reasonGoldText := reasonGold.String()
		reasonSliverText := reasonSilver.String()
		reasonlevelText := reasonLevel.String()

		totalRewData := propertytypes.CreateRewData(int32(expBase), 0, int32(silver), gold, bindGold)
		flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("quest: GiveQuestTuMoImmediateFinishReward AddRewData  should be ok"))
		}
	}

	//同步属性
	propertylogic.SnapChangedProperty(pl)
	return
}

//屠魔任务一键完成奖励
func GiveQuestTuMoFinishAllReward(pl player.Player, questIdList []int32, needNum int32) (itemMapList []map[int32]int32, err error) {
	//奖励
	itemMapList = make([]map[int32]int32, 0, needNum)
	expBase := int64(0)
	silver := int64(0)
	gold := int32(0)
	bindGold := int32(0)
	zhuanshu := int32(0)
	rewItemMap := make(map[int32]int32)

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	for _, questId := range questIdList {
		itemMap := make(map[int32]int32)
		rewDataMap := make(map[int32]int32)
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
		if questTemplate == nil {
			continue
		}
		silver += int64(math.Ceil(float64(questTemplate.RewSilver) * pl.GetWallowState().Rate()))
		gold += int32(math.Ceil(float64(questTemplate.RewGold) * pl.GetWallowState().Rate()))
		bindGold += int32(math.Ceil(float64(questTemplate.RewBindGold) * pl.GetWallowState().Rate()))

		rewData := questTemplate.GetRewData()
		if rewData != nil {
			rewDataMap = propertylogic.GetRewDataMap(rewData, pl.GetLevel())
			exp, exist := rewDataMap[consttypes.ExpItem]
			if exist {
				expBase += int64(exp)
			}
		}

		if questTemplate.RewZhuanshu != 0 && questTemplate.RewZhuanshu > propertyManager.GetZhuanSheng() && questTemplate.RewZhuanshu > zhuanshu {
			zhuanshu = questTemplate.RewZhuanshu
		}

		//奖励
		if pl.GetWallowState() == playertypes.WallowStateNone {
			rewardItemMap := questTemplate.GetRewardItemMap(pl.GetRole())
			if len(rewardItemMap) != 0 {
				for itemId, num := range rewardItemMap {
					tempNum, exist := itemMap[itemId]
					if exist {
						tempNum += num
					} else {
						tempNum = num
					}
					itemMap[itemId] = tempNum
					curNum, exist := rewItemMap[itemId]
					if exist {
						curNum += num
					} else {
						curNum = num
					}
					rewItemMap[itemId] = curNum
				}
			}
		}

		//固定奖励(展示用)
		for itemId, num := range rewDataMap {
			curNum, exist := itemMap[itemId]
			if exist {
				curNum += num
			}
			itemMap[itemId] = num
		}

		itemMapList = append(itemMapList, itemMap)
	}

	//奖励物品
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(rewItemMap) != 0 {
		//添加物品
		if !inventoryManager.HasEnoughSlots(rewItemMap) {
			//写邮件
			emailTitle := lang.GetLangService().ReadLang(lang.QuestTuMoFinishAllTitle)
			emailContent := lang.GetLangService().ReadLang(lang.QuestTuMoFinishAllContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			reasonText := fmt.Sprintf(commonlog.InventoryLogReasonTuMoFinishAllRew.String(), needNum)
			flag := inventoryManager.BatchAdd(rewItemMap, commonlog.InventoryLogReasonTuMoFinishAllRew, reasonText)
			if !flag {
				panic("quest: GiveQuestTuMoFinishAllReward BatchAdd  should be ok")
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//奖励属性
	if expBase != 0 || silver != 0 || gold != 0 || bindGold != 0 {
		reasonGold := commonlog.GoldLogReasonTuMoFinishAll
		reasonSilver := commonlog.SilverLogReasonTuMoFinishAll
		reasonLevel := commonlog.LevelLogReasonTuMoFinishAll
		reasonGoldText := fmt.Sprintf(reasonGold.String(), needNum)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), needNum)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), needNum)

		totalRewData := propertytypes.CreateRewData(int32(expBase), 0, int32(silver), gold, bindGold)
		flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("quest: questTuMoFinishAll AddRewData  should be ok"))
		}
	}

	//同步属性
	propertylogic.SnapChangedProperty(pl)
	return
}
