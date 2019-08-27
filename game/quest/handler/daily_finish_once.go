package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	constanttypes "fgame/fgame/game/constant/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"
	weeklogic "fgame/fgame/game/week/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_DAILY_FINISH_ONECE_TYPE), dispatch.HandlerFunc(handleQuestDailyFinishOnce))
}

//处理日环任务完成单次信息
func handleQuestDailyFinishOnce(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理日环任务完成单次消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSQuestDailyFinishOnce)
	dailyTagInt := csMsg.GetDailyTag()
	isDouble := csMsg.GetIsDoube()

	dailyTag := questtypes.QuestDailyTag(dailyTagInt)
	if !dailyTag.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dailyTag": dailyTag,
			}).Warn("quest:参数无效")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = questDailyFinishOnce(tpl, dailyTag, isDouble)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dailyTag": dailyTag,
				"error":    err,
			}).Error("quest:处理日环任务完成单次消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理日环任务完成单次消息完成")
	return nil

}

//日环任务完成单次信息的逻辑
func questDailyFinishOnce(pl player.Player, dailyTag questtypes.QuestDailyTag, isDoube bool) (err error) {
	funcOpen := dailyTag.GetFuncOpen()
	if !pl.IsFuncOpen(funcOpen) {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyObj := manager.GetDailyObj(dailyTag)
	seqId := dailyObj.GetSeqId()
	questId := int32(0)
	seqIdList := make([]int32, 0, 20)
	realNum := int32(0)
	immediateNum := int32(0)
	tatalItemMap := make(map[int32]int32)
	newItemMap := make(map[int32]int32)
	dailyTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
	if dailyTemplate != nil {
		questId = dailyTemplate.GetQuestId()
		seqIdList = append(seqIdList, seqId)
		realNum++
		immediateNum++
	}

	finishQuest := manager.GetQuestByIdAndState(questtypes.QuestStateFinish, questId)
	if finishQuest != nil {
		immediateNum--
	}

	if realNum <= 0 {
		return
	}

	if immediateNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest:日环任务已完成,直接领取奖励")
		playerlogic.SendSystemMessage(pl, lang.QuestDailyAlreadyFinishOnce, dailyTag.String())
		return
	}

	doubleCost := int64(0)
	if isDoube {
		doubleCost = int64(questtemplate.GetQuestTemplateService().GetQuestDailyCommitDouble(dailyTag) * realNum)
	}
	immediateCost := int64(0)
	if !weeklogic.IsSeniorWeek(pl) {
		immediateCost = int64(questtemplate.GetQuestTemplateService().GetQuestDailyImmediateFinish(dailyTag) * immediateNum)
	}
	needBindGold := doubleCost + immediateCost
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(needBindGold, true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"doubleCost":    doubleCost,
				"immediateCost": immediateCost,
				"needBindGold":  needBindGold,
			}).Warn("quest:日环单次完成，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}
	ratio := defaultRatio

	if isDoube {
		ratio = doubleRatio
	}

	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	for _, seqId := range seqIdList {
		seqTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
		if seqTemplate == nil {
			continue
		}
		itemMap := seqTemplate.GetRewItemMap()
		for itemId, num := range itemMap {
			newItemMap[itemId] += (num * ratio)
		}
		emailMap := seqTemplate.GetEmailItemMap()
		for itemId, num := range emailMap {
			tatalItemMap[itemId] += (num * ratio)
		}
		tempRewData := seqTemplate.GetRewData()
		rewData := propertytypes.CreateRewData(
			tempRewData.GetRewExp()*ratio,
			tempRewData.GetRewExpPoint()*ratio,
			tempRewData.GetRewSilver()*ratio,
			tempRewData.GetRewGold()*ratio,
			tempRewData.GetRewBindGold()*ratio)
		rewExpPoint := rewData.RewExpPoint
		if rewExpPoint != 0 {
			exp := int32(propertylogic.ExpPointConvertExp(rewExpPoint, pl.GetLevel()))
			tatalItemMap[constanttypes.ExpItem] += exp
		}

		// 个人日常额外处理
		if dailyTag == questtypes.QuestDailyTagPerson {
			itemId, num := droptemplate.GetDropTemplateService().GetDropItem(seqTemplate.GetDropId())
			if itemId != 0 && num != 0 {
				newItemMap[itemId] += (num * ratio)
				tatalItemMap[itemId] += (num * ratio)
			}
		}
		totalRewData.AddRewData(rewData)
	}

	//消耗绑元
	if needBindGold != 0 {
		reasonGold := commonlog.GoldLogReasonDailyQuestImmediateFinishAll
		reasonGoldText := fmt.Sprintf(reasonGold.String(), dailyTag.String(), realNum)
		flag = propertyManager.CostGold(needBindGold, true, commonlog.GoldLogReasonDailyQuestImmediateFinishAll, reasonGoldText)
		if !flag {
			panic(fmt.Errorf("quest: questDailyFinishAll CostGold should be ok"))
		}
	}

	var addItemMap map[int32]int32
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(newItemMap) != 0 {
		addItemMap, resMap = droplogic.SeperateItems(newItemMap)
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(addItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(addItemMap)
		if !flag {
			emailTitle := lang.GetLangService().ReadLang(lang.QuestDailyEmailTitle)
			emailContent := lang.GetLangService().ReadLang(lang.QuestDailyFinishAllContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, addItemMap)
		} else {
			reasonText := commonlog.InventoryLogReasonDailyQuestFinishAllReward.String()
			flag = inventoryManager.BatchAdd(addItemMap, commonlog.InventoryLogReasonDailyQuestFinishAllReward, reasonText)
			if !flag {
				panic(fmt.Errorf("quest: questDailyFinishAll  BatchAdd should be ok"))
			}
		}
	}

	qu := manager.GetQuestById(questId)
	if qu != nil {
		qu, flag := manager.QuestDailyImmediateFinish(questId)
		if !flag {
			return
		}
		scQuestListUpdate := pbutil.BuildSCQuestUpdate(qu)
		pl.SendMsg(scQuestListUpdate)
	}

	reasonGold := commonlog.GoldLogReasonDailyQuestFinishAll
	reasonSilver := commonlog.SilverLogReasonDailyQuestFinishAll
	reasonLevel := commonlog.LevelLogReasonDailyQuestFinishAll
	reasonGoldText := fmt.Sprintf(reasonGold.String(), dailyTag.String())
	reasonSliverText := fmt.Sprintf(reasonSilver.String(), dailyTag.String())
	reasonlevelText := fmt.Sprintf(reasonLevel.String(), dailyTag.String())
	flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
	if !flag {
		panic(fmt.Errorf("quest: questDailyFinishAll AddRewData  should be ok"))
	}
	if len(resMap) > 0 {
		droplogic.AddRes(pl, resMap, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
	}

	//同步物品
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	// 领取事件
	data := questeventtypes.CreateQuestFinishAllEventData(dailyTag.GetQuestType(), realNum)
	gameevent.Emit(questeventtypes.EventTypeQuestFinishAll, pl, data)

	// 与日环任务关联的任务
	for i := int32(0); i < realNum; i++ {
		relateQuestSubType, ok := dailyTag.GetQuestType().QuestNestedSubType()
		if !ok {
			continue
		}
		questlogic.IncreaseQuestData(pl, relateQuestSubType, 0, 1)
	}

	//增加仙盟boss经验
	allianceItemMap := addAllianceBossExp(pl, dailyTag, seqIdList)
	for itemId, itemNum := range allianceItemMap {
		tatalItemMap[itemId] = itemNum
	}

	scMsg := pbutil.BuildSCQuestDailyFinishOnce(int32(dailyTag), tatalItemMap)
	pl.SendMsg(scMsg)

	//获取下一个日常
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
	return
}
