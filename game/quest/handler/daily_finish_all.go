package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	playeralliance "fgame/fgame/game/alliance/player"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_DAILY_FINISH_ALL_TYPE), dispatch.HandlerFunc(handleQuestDailyFinishAll))
}

//处理日环任务完成所有完成信息
func handleQuestDailyFinishAll(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理日环任务完成所有完成消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSQuestDailyFinishAll)
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

	err = questDailyFinishAll(tpl, dailyTag, isDouble)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dailyTag": dailyTag,
				"error":    err,
			}).Error("quest:处理日环任务完成所有完成消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理日环任务完成所有完成消息完成")
	return nil

}

//日环任务完成所有完成信息的逻辑
func questDailyFinishAll(pl player.Player, dailyTag questtypes.QuestDailyTag, isDoube bool) (err error) {
	funcOpen := dailyTag.GetFuncOpen()
	if !pl.IsFuncOpen(funcOpen) {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyObj := manager.GetDailyObj(dailyTag)
	seqId := dailyObj.GetSeqId()
	times := dailyObj.GetTimes()
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

	idList, lastSeqId := questtemplate.GetQuestTemplateService().GetQuestDailyFinishAll(pl, dailyTag, questtypes.QuestDailyType(dailyObj.GetTimes()))
	realNum += int32(len(idList))
	immediateNum += int32(len(idList))
	if realNum <= 0 {
		return
	}

	if immediateNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest:日环任务已全部完成")
		playerlogic.SendSystemMessage(pl, lang.QuestDailyAlreadyFinishAll, dailyTag.String())
		return
	}

	seqIdList = append(seqIdList, idList...)
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
			}).Warn("quest:日环全部完成，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}
	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	for _, seqId := range seqIdList {
		seqTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
		if seqTemplate == nil {
			continue
		}
		itemMap := seqTemplate.GetDoubleRewItemMap()
		for itemId, num := range itemMap {
			newItemMap[itemId] += num
		}
		emailMap := seqTemplate.GetEmailItemMap()
		rewData := seqTemplate.GetDoubleRewData()
		for itemId, num := range emailMap {
			tatalItemMap[itemId] += (num * 2)
		}
		if rewData.RewExpPoint != 0 {
			exp := int32(propertylogic.ExpPointConvertExp(rewData.RewExpPoint, pl.GetLevel()))
			tatalItemMap[constanttypes.ExpItem] += exp
		}

		// 个人日常额外处理
		if dailyTag == questtypes.QuestDailyTagPerson {
			itemId, num := droptemplate.GetDropTemplateService().GetDropItem(seqTemplate.GetDropId())
			if itemId != 0 && num != 0 {
				newItemMap[itemId] += (num * 2)
				tatalItemMap[itemId] += (num * 2)
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
	manager.QuestDailyFinishAll(dailyTag, lastSeqId)
	qu := manager.GetQuestById(questId)
	if qu != nil {
		qu, flag := manager.QuestDailyImmediateFinish(questId)
		if !flag {
			return
		}
		times = dailyObj.GetTimes()
		scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), seqId, times)
		pl.SendMsg(scQuestDailSeq)
		scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		pl.SendMsg(scQuestUpdate)
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
	//同步属性
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

	scQuestDailyFinishAll := pbutil.BuildSCQuestDailyFinishAll(int32(dailyTag), tatalItemMap)
	pl.SendMsg(scQuestDailyFinishAll)
	return
}

//zrc: 临时处理
const (
	bossExpItemId = 2860
	bossJiFen     = 2861
	defaultRatio  = int32(1)
	doubleRatio   = int32(2)
)

func addAllianceBossExp(pl player.Player, dailyTag questtypes.QuestDailyTag, seqIdList []int32) map[int32]int32 {
	itemMap := make(map[int32]int32)
	if dailyTag != questtypes.QuestDailyTagAlliance {
		return itemMap
	}

	if len(seqIdList) == 0 {
		return itemMap
	}

	totalBossExp := int32(0)
	totalJiFen := int32(0)
	for _, seqId := range seqIdList {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
		if questTemplate == nil {
			continue
		}
		bossExp := questTemplate.GetBossExp()
		jiFen := questTemplate.GetUnionStorageJiFen()
		totalBossExp += bossExp
		totalJiFen += jiFen
	}

	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		return itemMap
	}

	if totalBossExp != 0 {
		itemMap[bossExpItemId] = totalBossExp
		//增加仙盟boss经验
		alliance.GetAllianceService().AllianceBossAddExp(allianceId, totalBossExp)

		playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		expStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", totalBossExp))
		chatContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.QuestDailyAllianceFinishAllChat), playerName, expStr)
		al := alliance.GetAllianceService().GetAlliance(allianceId)
		if al != nil {
			chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(chatContent))
		}
	}

	if totalJiFen != 0 {
		itemMap[bossJiFen] = totalJiFen
		manager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
		manager.AddDepotPoint(totalJiFen)
	}
	return itemMap
}
