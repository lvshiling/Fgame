package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	playeralliance "fgame/fgame/game/alliance/player"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func getDailyTagByQuestId(questId int32) (dailyTag questtypes.QuestDailyTag, flag bool) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	questType := questTemplate.GetQuestType()
	return questType.GetDailyTag()
}

//判断是否可以双倍领取
func IfQuestCanCommitDouble(pl player.Player, questId int32) (flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyTag, flag := getDailyTagByQuestId(questId)
	if !flag {
		return
	}
	dailyObj := manager.GetDailyObj(dailyTag)
	if dailyObj == nil {
		return
	}
	seqId := dailyObj.GetSeqId()
	dailyTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
	if dailyTemplate.GetQuestId() != questId {
		return false
	}
	costBindGold := int64(0)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	costBindGold = int64(questtemplate.GetQuestTemplateService().GetQuestDailyCommitDouble(dailyTag))
	flag = propertyManager.HasEnoughGold(int64(costBindGold), true)
	if !flag {
		return
	}

	return
}

//完成日环任务奖励
func FinishDailyQuestReward(pl player.Player, questId int32, double bool) (isReturn bool, err error) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyTag, flag := getDailyTagByQuestId(questId)
	if !flag {
		return
	}
	dailyObj := manager.GetDailyObj(dailyTag)
	if dailyObj == nil {
		return
	}
	seqId := dailyObj.GetSeqId()
	dailyTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
	if dailyTemplate == nil {
		return
	}
	if dailyTemplate.GetQuestId() != questId {
		isReturn = true
		return
	}
	costBindGold := int64(0)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if double {
		costBindGold = int64(questtemplate.GetQuestTemplateService().GetQuestDailyCommitDouble(dailyTag))
		flag := propertyManager.HasEnoughGold(int64(costBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			isReturn = true
			return
		}
	}

	rewItemMap := make(map[int32]int32)
	var rewData *propertytypes.RewData
	if double {
		rewData = dailyTemplate.GetDoubleRewData()
		dailyRewItemMap := dailyTemplate.GetDoubleRewItemMap()
		rewItemMap = coreutils.MergeMap(rewItemMap, dailyRewItemMap)
	} else {
		rewData = dailyTemplate.GetRewData()
		dailyRewItemMap := dailyTemplate.GetRewItemMap()
		rewItemMap = coreutils.MergeMap(rewItemMap, dailyRewItemMap)
	}

	//个人日环处理
	if dailyTag == questtypes.QuestDailyTagPerson {
		ratio := int32(1)
		if double {
			ratio = 2
		}
		itemId, num := droptemplate.GetDropTemplateService().GetDropItem(dailyTemplate.GetDropId())
		if itemId != 0 && num != 0 {
			rewItemMap[itemId] += (num * ratio)
		}
	}

	var addItemMap map[int32]int32
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	addItemMap, resMap = droplogic.SeperateItems(rewItemMap)

	if costBindGold != 0 {
		reasonGoldText := commonlog.GoldLogReasonDailyQuestCommitDouble.String()
		flag := propertyManager.CostGold(costBindGold, true, commonlog.GoldLogReasonDailyQuestCommitDouble, reasonGoldText)
		if !flag {
			panic(fmt.Errorf("quest: FinishDailyQuestReward CostGold should be ok"))
		}
	}

	if len(addItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughSlots(addItemMap)
		if !flag {
			//TODO 发送邮件
			emailTitle := lang.GetLangService().ReadLang(lang.QuestDailyEmailTitle)
			emailContent := lang.GetLangService().ReadLang(lang.QuestCommitContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, addItemMap)
		} else {
			reason := commonlog.InventoryLogReasonDailyQuestReward
			reasonText := fmt.Sprintf(reason.String(), dailyTag.String())
			flag = inventoryManager.BatchAdd(addItemMap, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("quest: commit  BatchAdd should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	reasonGold := commonlog.GoldLogReasonDailyQuestReward
	reasonSilver := commonlog.SilverLogResonDailyQuestReward
	reasonLevel := commonlog.LevelLogReasonDailyQuestReward
	reasonGoldText := reasonGold.String()
	reasonSliverText := reasonSilver.String()
	reasonlevelText := reasonLevel.String()
	if rewData != nil {
		flag := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("quest: FinishDailyQuestReward AddRewData  should be ok"))
		}
	}

	if len(resMap) > 0 {
		droplogic.AddRes(pl, resMap, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
	}

	//
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCQuestDailyCommitRew(questId, int32(dailyTag), rewItemMap, rewData)
	pl.SendMsg(scMsg)
	return
}

//仙盟日常任务处理
func dailyAllianceQuest(dailyTag questtypes.QuestDailyTag, pl player.Player, times int32, allianceExp int32, jiFen int32, rewItemMap map[int32]int32, rd *propertytypes.RewData) {
	if dailyTag != questtypes.QuestDailyTagAlliance {
		return
	}
	allianceId := pl.GetAllianceId()
	name := pl.GetName()
	//仙盟仓库增加积分
	manager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	manager.AddDepotPoint(jiFen)

	if allianceId == 0 {
		return
	}

	//增加仙盟boss经验
	alliance.GetAllianceService().AllianceBossAddExp(allianceId, allianceExp)

	//物品奖励
	if len(rewItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			emailTitle := lang.GetLangService().ReadLang(lang.QuestDailyAllianceEmailTitle)
			emailContent := lang.GetLangService().ReadLang(lang.QuestDailyAllianceEmailContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			reason := commonlog.InventoryLogReasonDailyQuestReward
			reasonText := fmt.Sprintf(reason.String(), dailyTag.String())
			flag = inventoryManager.BatchAdd(rewItemMap, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("quest: commit  BatchAdd should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	reasonGold := commonlog.GoldLogReasonDailyQuestReward
	reasonSilver := commonlog.SilverLogResonDailyQuestReward
	reasonLevel := commonlog.LevelLogReasonDailyQuestReward
	reasonGoldText := fmt.Sprintf(reasonGold.String(), dailyTag.String())
	reasonSliverText := fmt.Sprintf(reasonSilver.String(), dailyTag.String())
	reasonlevelText := fmt.Sprintf(reasonLevel.String(), dailyTag.String())
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.AddRewData(rd, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
	if !flag {
		panic(fmt.Errorf("quest: questDailyFinishAll AddRewData  should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	//仙盟频道
	if times != 0 && times%5 == 0 {
		playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(name))
		numStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", times))
		expStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", allianceExp))
		chatContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.QuestDailyAllianceFinishChat), playerName, numStr, expStr)
		al := alliance.GetAllianceService().GetAlliance(allianceId)
		if al != nil {
			chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(chatContent))
		}
	}
}

//推送下一个日环任务
func GetNextDailyQuest(pl player.Player, dailyTag questtypes.QuestDailyTag) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyObj := manager.GetDailyObj(dailyTag)
	if dailyObj == nil {
		return
	}
	dailyTempalte := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, dailyObj.GetSeqId())
	if dailyTempalte != nil {
		qu := manager.GetQuestById(dailyTempalte.GetQuestId())
		if qu != nil && qu.QuestState != questtypes.QuestStateCommit {
			panic(fmt.Errorf("quest: 前一个日环任务应该是commit"))
		}

		//仙盟日常任务处理
		dailyAllianceQuest(dailyTag, pl, dailyObj.GetTimes(), dailyTempalte.GetBossExp(), dailyTempalte.GetUnionStorageJiFen(), dailyTempalte.GetRewItemMap(), dailyTempalte.GetRewData())
	}

	//接受下一个日环任务
	maxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(dailyTag)
	if dailyObj.GetTimes() < maxTimes {
		if dailyTag == questtypes.QuestDailyTagAlliance && pl.GetAllianceId() == 0 {
			return
		}
		quest := manager.GetNextDailyQuest(dailyTag)
		if quest != nil {
			scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
			pl.SendMsg(scQuestDailSeq)
			scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
			pl.SendMsg(scQuestUpdate)
		}
	}
}
