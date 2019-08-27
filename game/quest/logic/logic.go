package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	secretcardlogic "fgame/fgame/game/secretcard/logic"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func GiveQuestCommitReward(pl player.Player, questId int32) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	//发送奖励
	rewardItemMap := questTemplate.GetRewardItemMap(pl.GetRole())
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	if len(rewardItemMap) != 0 {
		//添加物品
		flag := inventoryManager.HasEnoughSlots(rewardItemMap)
		if !flag {
			emailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.QuestFinishTitle), questTemplate.Name)
			emailContent := lang.GetLangService().ReadLang(lang.QuestFinishContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewardItemMap)
		} else {
			reasonText := fmt.Sprintf(commonlog.InventoryLogReasonQuestReward.String(), questId)
			flag = inventoryManager.BatchAdd(rewardItemMap, commonlog.InventoryLogReasonQuestReward, reasonText)
			if !flag {
				panic("quest:处理任务交付,发送物品应该ok")
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}
	expPoint := int64(math.Ceil(float64(questTemplate.RewExpPoint) * pl.GetWallowState().Rate()))
	//添加经验点(先加经验点,跟等级挂钩,先加经验可能等级会升级)
	if expPoint != 0 {
		reasonText := fmt.Sprintf(commonlog.LevelLogReasonQuestReward.String(), questTemplate.TemplateId())
		propertyManager.AddExpPoint(expPoint, commonlog.LevelLogReasonQuestReward, reasonText)
	}
	expBase := int64(math.Ceil(float64(questTemplate.RewXp) * pl.GetWallowState().Rate()))
	//添加经验
	if expBase != 0 {
		reasonText := fmt.Sprintf(commonlog.LevelLogReasonQuestReward.String(), questTemplate.TemplateId())
		propertyManager.AddExp(expBase, commonlog.LevelLogReasonQuestReward, reasonText)
	}

	silver := int64(math.Ceil(float64(questTemplate.RewSilver) * pl.GetWallowState().Rate()))
	//添加银两
	if silver != 0 {
		reasonText := fmt.Sprintf(commonlog.SilverLogReasonQuestReward.String(), questTemplate.TemplateId())
		propertyManager.AddSilver(silver, commonlog.SilverLogReasonQuestReward, reasonText)
	}
	gold := int64(math.Ceil(float64(questTemplate.RewGold) * pl.GetWallowState().Rate()))
	//添加元宝
	if gold != 0 {
		reasonText := fmt.Sprintf(commonlog.GoldLogReasonQuestReward.String(), questTemplate.TemplateId())
		propertyManager.AddGold(gold, false, commonlog.GoldLogReasonQuestReward, reasonText)
	}
	bindGold := int64(math.Ceil(float64(questTemplate.RewBindGold) * pl.GetWallowState().Rate()))
	//添加绑定元宝
	if bindGold != 0 {
		reasonText := fmt.Sprintf(commonlog.GoldLogReasonQuestReward.String(), questTemplate.TemplateId())
		propertyManager.AddGold(bindGold, true, commonlog.GoldLogReasonQuestReward, reasonText)
	}

	//修改转数
	if questTemplate.RewZhuanshu != 0 && questTemplate.RewZhuanshu > propertyManager.GetZhuanSheng() {
		reasonText := fmt.Sprintf(commonlog.ZhuanShengLogReasonQuestReward.String(), questTemplate.TemplateId())
		propertyManager.SetZhuanSheng(questTemplate.RewZhuanshu, commonlog.ZhuanShengLogReasonQuestReward, reasonText)
	}

	propertylogic.SnapChangedProperty(pl)
	return
}

//检查任务是否完成
func CheckQuestIfFinish(pl player.Player, questId int32) (qu *playerquest.PlayerQuestObject, err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	tQuest := questManager.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	level := propertyManager.GetLevel()
	zhuanSheng := propertyManager.GetZhuanSheng()

	//TODO 判断前置条件
	//判断级数
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if level < questTemplate.ReqLevel {
		return
	}
	//判断转生
	if zhuanSheng < questTemplate.ReqZhuanshu {
		return
	}

	//校正数据
	quest.CheckHandle(pl, questTemplate)

	//判断需求
	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	flag := true
	for k, v := range questDemandMap {
		val, ok := tQuest.QuestDataMap[k]
		if !ok {
			flag = false
			break
		}
		if val < v {
			flag = false
			break
		}
	}
	//没满足
	if !flag {
		return
	}
	reqItemMap := questTemplate.GetReqItemMap(pl.GetRole())
	flag = true
	for k, v := range reqItemMap {
		val, ok := tQuest.CollectItemDataMap[k]
		if !ok {
			flag = false
			break
		}
		if val < v {
			flag = false
			break
		}
	}
	//没满足
	if !flag {
		return
	}

	flag = questManager.FinishQuest(questId)
	if !flag {
		panic(fmt.Errorf("quest:完成任务应该ok"))
	}
	if questTemplate.AutoCommit() {
		flag = questManager.CommitQuest(questId, false)
		if !flag {
			panic(fmt.Errorf("quest:提交任务应该ok"))
		}
		//TODO发送奖励
		GiveQuestCommitReward(pl, questId)
	}
	qu = tQuest
	return
}

//检查任务是否可以激活
func CheckQuestIfActive(pl player.Player, questId int32) (qu *playerquest.PlayerQuestObject, err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	level := propertyManager.GetLevel()
	zhuanSheng := propertyManager.GetZhuanSheng()
	tquest := questManager.GetQuestByIdAndState(questtypes.QuestStateInit, questId)
	if tquest == nil {
		return
	}
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	if level < questTemplate.MinLevel || level > questTemplate.MaxLevel {
		return
	}

	//判断转生
	if zhuanSheng < questTemplate.MinZhuanshu || zhuanSheng > questTemplate.MaxZhuanshu {
		return
	}
	//TODO 判断特殊任务
	flag := questManager.ActiveQuest(questId)
	if !flag {
		panic(fmt.Errorf("quest:激活任务应该ok"))
	}

	if questTemplate.AutoAccept() {
		flag = questManager.AcceptQuest(questId)
		if !flag {
			panic(fmt.Errorf("quest:接受任务应该ok"))
		}
	}
	qu = tquest
	return
}

//检查初始化的任务
func CheckInitQuest(pl player.Player, qu *playerquest.PlayerQuestObject) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	level := propertyManager.GetLevel()
	zhuanSheng := propertyManager.GetZhuanSheng()
	//TODO 判断前置条件
	allCommit := true
	//gm使用
	//检查是否前置任务都完成了
	nextQuestTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
	for _, prevQuestId := range nextQuestTemplate.GetPrevQuestIds() {
		if questManager.GetQuestByIdAndState(questtypes.QuestStateCommit, prevQuestId) == nil {
			allCommit = false
			break
		}
	}

	if !allCommit {
		return
	}

	//判断级数
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
	if level < questTemplate.MinLevel || level > questTemplate.MaxLevel {
		return
	}
	if questTemplate.GetQuestType() == questtypes.QuestTypeDaily {
		return
	}
	//判断转生
	if zhuanSheng < questTemplate.MinZhuanshu || zhuanSheng > questTemplate.MaxZhuanshu {
		return
	}

	//TODO 判断特殊任务
	flag := questManager.ActiveQuest(qu.QuestId)
	if !flag {
		panic(fmt.Errorf("quest:激活任务应该ok"))
	}
	if questTemplate.AutoAccept() {
		flag = questManager.AcceptQuest(qu.QuestId)
		if !flag {
			panic(fmt.Errorf("quest:接受任务应该ok"))
		}
	}
}

//检查初始化的任务列表
func CheckInitQuestList(pl player.Player) (questList []*playerquest.PlayerQuestObject) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	initQuests := questManager.GetQuestMap(questtypes.QuestStateInit)
	level := propertyManager.GetLevel()
	zhuanSheng := propertyManager.GetZhuanSheng()
	for _, qu := range initQuests {
		//TODO 判断前置条件
		allCommit := true
		//gm使用
		//检查是否前置任务都完成了
		nextQuestTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		for _, prevQuestId := range nextQuestTemplate.GetPrevQuestIds() {
			if questManager.GetQuestByIdAndState(questtypes.QuestStateCommit, prevQuestId) == nil {
				allCommit = false
				break
			}
		}

		if !allCommit {
			continue
		}

		//判断级数
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		if level < questTemplate.MinLevel || level > questTemplate.MaxLevel {
			continue
		}
		//判断转生
		if zhuanSheng < questTemplate.MinZhuanshu || zhuanSheng > questTemplate.MaxZhuanshu {
			continue
		}
		if questTemplate.GetQuestType() == questtypes.QuestTypeDaily || questTemplate.GetQuestType() == questtypes.QuestTypeDailyAlliance {
			continue
		}

		//TODO 判断特殊任务
		flag := questManager.ActiveQuest(qu.QuestId)

		if !flag {
			panic(fmt.Errorf("quest:激活任务应该ok"))
		}
		if questTemplate.AutoAccept() {
			flag = questManager.AcceptQuest(qu.QuestId)
			if !flag {
				panic(fmt.Errorf("quest:接受任务应该ok"))
			}
		}
		questList = append(questList, qu)
	}
	return
}

func CheckAcceptQuestList(pl player.Player) (questList []*playerquest.PlayerQuestObject) {
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	acceptQuests := questManager.GetQuestMap(questtypes.QuestStateAccept)
	for _, qu := range acceptQuests {
		qu, err := CheckQuestIfFinish(pl, qu.QuestId)
		if err != nil {
			return
		}
		if qu != nil {
			questList = append(questList, qu)
		}
	}
	return
}

//提交任务后
func CheckQuestIfInit(pl player.Player, questId int32) (questList []*playerquest.PlayerQuestObject) {
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	qu := questManager.GetQuestByIdAndState(questtypes.QuestStateCommit, questId)
	if qu == nil {
		return
	}
	nextQuestIds := questTemplate.GetNextQuestIds()
	questList = make([]*playerquest.PlayerQuestObject, 0, 16)
	for _, nextQuestId := range nextQuestIds {
		nextQuest := questManager.GetQuestById(nextQuestId)
		if nextQuest != nil {
			continue
		}
		allCommit := true
		//检查是否前置任务都完成了
		nextQuestTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(nextQuestId)
		for _, prevQuestId := range nextQuestTemplate.GetPrevQuestIds() {
			if questManager.GetQuestByIdAndState(questtypes.QuestStateCommit, prevQuestId) == nil {
				allCommit = false
				break
			}
		}
		if !allCommit {
			continue
		}
		_, flag := questManager.AddQuest(nextQuestId)
		if !flag {
			panic(fmt.Errorf("quest:任务添加应该成功"))
		}

		//检查是否所有条件满足
		// nextQuest, err = CheckQuestIfActive(pl, nextQuestId)
		// if err != nil {
		// 	return nil, err
		// }
		// if nextQuest != nil {
		// 	questList = append(questList, nextQuest)
		// }
	}
	//TODO gm需要 但是效率不高
	questList = CheckInitQuestList(pl)
	return
}

//检查所有提交的任务
func CheckCommitQuestList(pl player.Player) (err error) {
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := questManager.GetQuestMap(questtypes.QuestStateCommit)
	for _, quest := range questMap {
		CheckQuestIfInit(pl, quest.QuestId)
	}
	return
}

//累计任务条件数据
func IncreaseQuestData(pl player.Player, questSubType questtypes.QuestSubType, demandId int32, num int32) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}

		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		_, ok := questDemandMap[demandId]
		if !ok {
			continue
		}

		flag := manager.IncreaseQuestData(qu.QuestId, demandId, num)
		if !flag {
			panic("quest:设置 IncreaseQuestData 应该成功")
		}
		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

// 设置任务条件数据（无条件）
func SetQuestData(pl player.Player, questSubType questtypes.QuestSubType, demandId int32, num int32) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}
		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		_, ok := questDemandMap[demandId]
		if !ok {
			continue
		}

		flag := manager.SetQuestData(qu.QuestId, demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

func SetQuestDataFinish(pl player.Player, questSubType questtypes.QuestSubType, demandId int32, num int32) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}
		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		curNum, ok := questDemandMap[demandId]
		if !ok {
			continue
		}
		if num < curNum {
			continue
		}

		flag := manager.SetQuestData(qu.QuestId, demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

// 设置任务条件数据(大于原本已完成的条件)
func SetQuestDataSurpass(pl player.Player, questSubType questtypes.QuestSubType, demandId int32, num int32) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}
		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		_, ok := questDemandMap[demandId]
		if !ok {
			continue
		}

		// 是否替换已完成的条件
		finishNum := qu.QuestDataMap[demandId]
		if num <= finishNum {
			continue
		}

		flag := manager.SetQuestData(qu.QuestId, demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

//填充任务数据(直接完成)
func FillQuestData(pl player.Player, questSubType questtypes.QuestSubType, demandId int32) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}
		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		num := questDemandMap[demandId]
		flag := manager.SetQuestData(qu.QuestId, demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

//收集物品
func SetQuestCollectData(pl player.Player, questSubType questtypes.QuestSubType) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}
		questCollectMap := questTemplate.GetReqItemMap(pl.GetRole())
		if len(questCollectMap) == 0 {
			continue
		}
		for reqItemId, _ := range questCollectMap {
			num := inventoryManager.NumOfItems(reqItemId)
			flag := manager.SetCollectItemData(qu.QuestId, reqItemId, num)
			if !flag {
				panic("quest:设置 SetCollectItemData 应该成功")
			}
		}

		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

func SetQuestEmbedData(pl player.Player, questSubType questtypes.QuestSubType) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		if questTemplate == nil {
			continue
		}
		subType := questTemplate.GetQuestSubType()
		if subType != questSubType {
			continue
		}
		//TODO 检测任务是否完成
		CheckQuestIfFinish(pl, qu.QuestId)
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

//仅GM使用
func GMIncreaseQuestData(pl player.Player, questSubType questtypes.QuestSubType, demandId int32, num int32) (err error) {
	IncreaseQuestData(pl, questSubType, demandId, num)
	return
}

func CommitQuest(pl player.Player, questId int32, double bool) (err error) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	//TODO 记录恶意刷的
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"double":   double,
			}).Warn("quest:处理任务交付,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest := manager.GetQuestByIdAndState(questtypes.QuestStateFinish, questId)
	if quest == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"double":   double,
			}).Warn("quest:处理任务交付,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}

	if questTemplate.GetQuestType() != questtypes.QuestTypeDaily {
		double = false
	}

	switch questTemplate.GetQuestType() {
	case questtypes.QuestTypeTianJiPai:
		{
			err = secretcardlogic.FinishSecretCardQuest(pl, questId)
			if err != nil {
				return
			}
		}
	case questtypes.QuestTypeDaily:
		{
			isReturn, err := FinishDailyQuestReward(pl, questId, double)
			if err != nil {
				return err
			}
			if isReturn {
				return nil
			}
		}
	}

	flag := manager.CommitQuest(questId, double)
	if !flag {
		panic("quest:处理任务交付,应该成功")
	}

	GiveQuestCommitReward(pl, questId)
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	scQuestCommit := pbutil.BuildSCQuestCommit(questId, double)
	pl.SendMsg(scQuestCommit)
	return
}

//完成任务
func FinishQuest(pl player.Player, questId int32) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest := manager.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if quest == nil {
		return
	}

	flag := manager.FinishQuest(questId)
	if !flag {
		panic("quest:处理完成任务,应该成功")
	}

	if questTemplate.AutoCommit() {
		flag = manager.CommitQuest(questId, false)
		if !flag {
			panic(fmt.Errorf("quest:提交任务应该ok"))
		}
		//TODO发送奖励
		GiveQuestCommitReward(pl, questId)
	}

	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)

}
