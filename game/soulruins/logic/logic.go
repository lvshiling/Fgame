package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
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
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	soulruinseventtypes "fgame/fgame/game/soulruins/event/types"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//挑战帝陵
func SoulRuinsChallenge(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IsValid(chapter, typ, level)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
			}).Warn("soulruins:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断挑战次数是否足够
	flag = manager.HasEnoughChallengeNum(1)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
			}).Warn("soulruins:挑战次数不足")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsChallengeNumNotEnough)
		return
	}

	//判断前置通关
	flag = manager.IfPreSoulRuinsPassed(chapter, typ, level)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
			}).Warn("soulruins:请先通关前置关卡")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsNotReachPreLevel)
		return
	}

	//消耗挑战次数
	manager.UseChallengeNum(1)
	flag = PlayerEnterSoulRuins(pl, chapter, typ, level)
	if !flag {
		panic(fmt.Errorf("soulruins: soulRuinsChallenge should be ok"))
	}

	numObj := manager.GetSoulRuinsNum()
	scSoulRuinsChallenge := pbutil.BuildSCSoulRuinsChallenge(numObj)
	pl.SendMsg(scSoulRuinsChallenge)
	return
}

//进入场景
func PlayerEnterSoulRuins(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32) (flag bool) {
	soulRuinsTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	if soulRuinsTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("soulruins:处理跳转帝陵遗迹,帝陵遗迹不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	sh := CreateSoulRuinsSceneData(pl.GetId(), soulRuinsTemplate, soulruinstypes.SoulRuinsStageTypeKillMonster)
	s := scene.CreateFuBenScene(soulRuinsTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("soulruins:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)
	flag = true
	return
}

//触发特殊事件
func triggerSpecialEvent(pl player.Player, soulRuinsTemplate *gametemplate.SoulRuinsTemplate) (eventType soulruinstypes.SoulRuinsEventType) {
	if soulRuinsTemplate == nil {
		return
	}
	chapter := soulRuinsTemplate.Chapter
	typ := soulRuinsTemplate.GetType()
	firstEvent := soulRuinsTemplate.FirstEvent
	level := soulRuinsTemplate.Level
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IfSoulRuinsExist(chapter, typ, level)
	//首次打 && firstEvent !=0
	if !flag && firstEvent != 0 {
		eventType = soulRuinsTemplate.GetEventType()
	} else {
		//随机触发
		eventType = soulruins.GetSoulRuinsService().GetSoulRuinsTriggerSpecialEvent(chapter, typ, level)
	}
	return
}

//触发出事件
func onSpecialEventType(pl player.Player, eventType soulruinstypes.SoulRuinsEventType) (err error) {
	//推送给前端
	scSoulRuinsEvent := pbutil.BuildSCSoulRuinsEvent(int32(eventType))
	err = pl.SendMsg(scSoulRuinsEvent)
	return
}

//下发场景信息
func onPushSceneInfo(pl player.Player, state soulruinstypes.SoulRuinsStageType, eventType soulruinstypes.SoulRuinsEventType, starTime int64, chapter int32, typ int32, level int32) (err error) {
	//推送给前端
	scEventType := int32(eventType)
	if state != soulruinstypes.SoulRuinsStageTypeEvent {
		scEventType = 0
	}
	scState := int32(state)
	scSoulRuinsScene := pbutil.BuildSCSoulRuinsScene(scState, scEventType, starTime, chapter, typ, level)
	err = pl.SendMsg(scSoulRuinsScene)
	return
}

//TODO 奖励
func onSoulRuinsFinish(pl player.Player, soulRuinsTemplate *gametemplate.SoulRuinsTemplate, successful bool, useTime int64, dropMap map[int32]int32) (err error) {
	if soulRuinsTemplate == nil {
		return
	}
	err = challengeFinished(pl, soulRuinsTemplate, successful, int32(useTime), dropMap)
	return
}

//挑战结束
func challengeFinished(pl player.Player, soulRuinsTemplate *gametemplate.SoulRuinsTemplate, successful bool, useTime int32, dropMap map[int32]int32) (err error) {
	chapter := soulRuinsTemplate.Chapter
	typ := soulRuinsTemplate.GetType()
	level := soulRuinsTemplate.Level
	star := int32(0)
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	isExist := manager.IfSoulRuinsExist(chapter, typ, level)
	if successful {
		star = soulruins.GetSoulRuinsService().GetSoulRuinsStarByTime(chapter, typ, level, int32(useTime))
		//刷新帝陵遗迹
		manager.RefreshSoulRuins(chapter, typ, level, star, true)
		//发送事件
		soulRuinsId := int32(soulRuinsTemplate.TemplateId())
		eventData := soulruinseventtypes.CreateSoulRuinsFinishEventData(soulRuinsId, 1)
		gameevent.Emit(soulruinseventtypes.EventTypeSoulruinsFinish, pl, eventData)
	}

	rewData := soulRuinsTemplate.GetRewData()
	//奖励属性
	if successful && rewData != nil {
		reasonGold := commonlog.GoldLogReasonSoulRuins
		reasonSilver := commonlog.SilverLogReasonSoulRuins
		reasonLevel := commonlog.LevelLogReasonSoulRuins
		reasonGoldText := fmt.Sprintf(reasonGold.String(), chapter, typ, level)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), chapter, typ, level)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), chapter, typ, level)

		rewDataMap := propertylogic.GetRewDataMap(rewData, pl.GetLevel())
		for itemId, num := range rewDataMap {
			curNum, exist := dropMap[itemId]
			if exist {
				curNum += num
			}
			dropMap[itemId] = curNum
		}
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		flag := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("soulruins: ChallengeFinished AddRewData  should be ok"))
		}
		//同步
		propertylogic.SnapChangedProperty(pl)
	}

	//挑战失败 或再次挑战成功
	if star == 0 || isExist {
		scSoulRuinsResult := pbutil.BuildSCSoulRuinsResult(chapter, int32(typ), level, useTime, dropMap, successful)
		err = pl.SendMsg(scSoulRuinsResult)
		return
	}
	//首次挑战成功
	numObj := manager.GetSoulRuinsNum()
	scSoulRuinsResult := pbutil.BuildSCSoulRuinsFirstPass(numObj, chapter, int32(typ), level, useTime, dropMap, successful)
	err = pl.SendMsg(scSoulRuinsResult)
	return
}

//上交银两给马贼
func GiveSilverToRobber(pl player.Player, sceneData *SoulRuinsSceneData) (isReturn bool) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	soulRuinsTemplate := sceneData.GetSoulRuinsTemplate()
	silver := soulRuinsTemplate.GetRobberSilver()
	flag := propertyManager.HasEnoughSilver(int64(silver))
	isReturn = false
	if !flag {
		isReturn = true
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}
	chapter := soulRuinsTemplate.Chapter
	typ := soulRuinsTemplate.GetType()
	level := soulRuinsTemplate.Level
	reasonSilver := commonlog.SilverLogReasonSoulRuinsRobber
	reasonSliverText := fmt.Sprintf(reasonSilver.String(), chapter, typ, level)
	flag = propertyManager.CostSilver(int64(silver), reasonSilver, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("soulruins: GiveSilverToRobber CostSilver should be ok"))
	}
	//同步元宝
	propertylogic.SnapChangedProperty(pl)
	return
}

//帝陵遗迹帝魂降临传功奖励
func GiveSoulRuinsSoulForceReward(pl player.Player, sceneData *SoulRuinsSceneData) {
	soulRuinsTemplate := sceneData.GetSoulRuinsTemplate()
	specialEventMap := soulRuinsTemplate.GetSpecialEventMap()
	rewItemMap := specialEventMap[soulruinstypes.SoulRuinsEventTypeSoul]
	if len(rewItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断背包是否足够
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			//写邮件
			emailtitle := lang.GetLangService().ReadLang(lang.SoulRuinsSoulArriveTitle)
			emailContent := lang.GetLangService().ReadLang(lang.SoulRuinsSoulArriveContent)
			emaillogic.AddEmail(pl, emailtitle, emailContent, rewItemMap)
			return
		}
		chapter := soulRuinsTemplate.Chapter
		typ := soulRuinsTemplate.GetType()
		level := soulRuinsTemplate.Level
		reasonInventory := commonlog.InventoryLogReasonSoulRuinsForceGet
		reasonInventoryText := fmt.Sprintf(reasonInventory.String(), chapter, typ, level)
		flag = inventoryManager.BatchAdd(rewItemMap, reasonInventory, reasonInventoryText)
		if !flag {
			panic(fmt.Errorf("soulruins: GiveSoulRuinsSoulForceReward BatchAdd should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}
	return
}

func getSoulRuinsDropRes(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32, num int32) (showDropItemList [][]*droptemplate.DropItemData, dropItemList []*droptemplate.DropItemData, rewData *propertytypes.RewData) {
	soulRuinsTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	if soulRuinsTemplate == nil {
		return
	}
	sweepDropList := soulRuinsTemplate.GetSweepDropList()
	sweepEventDropMap := soulRuinsTemplate.GetSweepEventDropMap()

	//奖励属性
	rewDataMap := make(map[int32]int32)
	rewData = soulRuinsTemplate.GetRewData()
	if rewData != nil {
		rewDataMap = propertylogic.GetRewDataMap(rewData, pl.GetLevel())
	}

	//扫荡掉落
	for i := int32(0); i < num; i++ {
		var rewItemList []*droptemplate.DropItemData
		//固定物品掉落
		if len(sweepDropList) != 0 {
			sweepDropList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(sweepDropList)
			rewItemList = append(rewItemList, sweepDropList...)
		}

		//特殊事件触发物品掉落
		eventType := soulruins.GetSoulRuinsService().GetSoulRuinsTriggerSpecialEvent(chapter, typ, level)
		if eventType != soulruinstypes.SoulRuinsEventTypeNot {
			eventDropId := sweepEventDropMap[eventType]
			dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(eventDropId)
			if dropData != nil {
				rewItemList = append(rewItemList, dropData)
			}
		}

		dropItemList = append(dropItemList, rewItemList...)

		//固定奖励(展示用)
		for itemId, num := range rewDataMap {
			newData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeUnBind)
			rewItemList = append(rewItemList, newData)
		}
		showDropItemList = append(showDropItemList, droplogic.MergeItemLevel(rewItemList))
	}
	return
}

//帝陵遗迹扫荡奖励
func GiveSoulRuinsSweepReward(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32, num int32) (showDropItemList [][]*droptemplate.DropItemData, rewData *propertytypes.RewData, isReturn bool, err error) {
	showDropItemList, dropItemList, rewData := getSoulRuinsDropRes(pl, chapter, typ, level, num)
	//区分资源
	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) != 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(dropItemList)
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//判断背包是否足够
	isReturn = false
	if len(newItemList) != 0 {
		flag := inventoryManager.HasEnoughSlotsOfItemLevel(newItemList)
		if !flag {
			isReturn = true
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		inventoryLogReason := commonlog.InventoryLogReasonSoulRuinsSweepDrop
		reasonText := fmt.Sprintf(inventoryLogReason.String(), chapter, typ, level, num)
		flag = inventoryManager.BatchAddOfItemLevel(newItemList, inventoryLogReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soulruins: GiveSoulRuinsSweepReward BatchAdd should be ok"))
		}
	}

	//添加扫荡获取资源
	if rewData != nil {
		totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
		totalRewData.RewSilver = num * rewData.RewSilver
		totalRewData.RewBindGold = num * rewData.RewBindGold
		totalRewData.RewGold = num * rewData.RewGold
		totalRewData.RewExp = num * rewData.RewExp

		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		reasonGold := commonlog.GoldLogReasonSoulRuinsSweep
		reasonSilver := commonlog.SilverLogReasonSoulRuinsSweep
		reasonLevel := commonlog.LevelLogReasonSoulRuinsSweep
		reasonGoldText := fmt.Sprintf(reasonGold.String(), chapter, typ, level, num)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), chapter, typ, level, num)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), chapter, typ, level, num)
		flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("soulruins: GiveSoulRuinsSweepReward AddRewData  should be ok"))
		}
	}

	//添加扫荡掉落获取资源
	if len(resMap) != 0 {
		reasonGoldDrop := commonlog.GoldLogReasonSoulRuinsSweepDrop
		reasonSliverDrop := commonlog.SilverLogReasonSoulRuinsSweepDrop
		reasonLevelDrop := commonlog.LevelLogReasonSoulRuinsSweepDrop
		reasonGoldDropText := fmt.Sprintf(reasonGoldDrop.String(), chapter, typ, level, num)
		reasonSliverDropText := fmt.Sprintf(reasonSliverDrop.String(), chapter, typ, level, num)
		reasonLevelDropText := fmt.Sprintf(reasonLevelDrop.String(), chapter, typ, level, num)
		err = droplogic.AddRes(pl, resMap, reasonGoldDrop, reasonGoldDropText, reasonSliverDrop, reasonSliverDropText, reasonLevelDrop, reasonLevelDropText)
		if err != nil {
			return
		}
	}

	//同步属性
	if len(resMap) != 0 || rewData != nil {
		propertylogic.SnapChangedProperty(pl)
	}
	return
}

//帝陵遗迹一键完成奖励
func GiveSoulRuinsFinishAllReward(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32, num int32) (showDropItemList [][]*droptemplate.DropItemData, rewData *propertytypes.RewData, err error) {
	showDropItemList, dropItemList, rewData := getSoulRuinsDropRes(pl, chapter, typ, level, num)
	//区分资源
	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) != 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(dropItemList)
	}

	//添加掉落物品
	if len(newItemList) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughSlotsOfItemLevel(newItemList)
		if !flag {
			emailTitle := lang.GetLangService().ReadLang(lang.SoulRuinsFinishAllTitle)
			emailContent := lang.GetLangService().ReadLang(lang.SoulRuinsFinishAllContent)
			//写邮件
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, emailTitle, emailContent, now, newItemList)
		} else {
			inventoryLogReason := commonlog.InventoryLogReasonSoulRuinsFinishAll
			reasonText := fmt.Sprintf(inventoryLogReason.String(), chapter, typ, level, num)
			flag := inventoryManager.BatchAddOfItemLevel(newItemList, inventoryLogReason, reasonText)
			if !flag {
				panic(fmt.Errorf("soulruins: GiveSoulRuinsFinishAllReward BatchAdd should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//添加一键完成资源
	if rewData != nil {
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
		totalRewData.RewSilver = num * rewData.RewSilver
		totalRewData.RewBindGold = num * rewData.RewBindGold
		totalRewData.RewGold = num * rewData.RewGold
		totalRewData.RewExp = num * rewData.RewExp

		reasonGold := commonlog.GoldLogReasonSoulRuinsFinishAll
		reasonSilver := commonlog.SilverLogReasonSoulRuinsFinishAll
		reasonLevel := commonlog.LevelLogReasonSoulRuinsFinishAll
		reasonGoldText := fmt.Sprintf(reasonGold.String(), chapter, typ, level, num)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), chapter, typ, level, num)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), chapter, typ, level, num)
		flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("soulruins: GiveSoulRuinsFinishAllReward AddRewData  should be ok"))
		}
	}

	//添加一键完成掉落资源
	if len(resMap) != 0 {
		reasonGoldDrop := commonlog.GoldLogReasonSoulRuinsFinishAllDrop
		reasonSliverDrop := commonlog.SilverLogReasonSoulRuinsFinishAllDrop
		reasonLevelDrop := commonlog.LevelLogReasonSoulRuinsFinishAllDrop
		reasonGoldDropText := fmt.Sprintf(reasonGoldDrop.String(), chapter, typ, level, num)
		reasonSliverDropText := fmt.Sprintf(reasonSliverDrop.String(), chapter, typ, level, num)
		reasonLevelDropText := fmt.Sprintf(reasonLevelDrop.String(), chapter, typ, level, num)
		err = droplogic.AddRes(pl, resMap, reasonGoldDrop, reasonGoldDropText, reasonSliverDrop, reasonSliverDropText, reasonLevelDrop, reasonLevelDropText)
		if err != nil {
			return
		}
	}
	return
}
