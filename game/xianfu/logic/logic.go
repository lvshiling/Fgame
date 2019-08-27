package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	"fgame/fgame/game/xianfu/pbutil"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//升级请求逻辑
func HandleXianfuUpgrade(pl player.Player, xianfuType xianfutypes.XianfuType) (err error) {
	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)
	if xfTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府升级请求，模板不存在")

		playerlogic.SendSystemMessage(pl, lang.XianfuArgumentInvalid)
		return
	}

	//刷新数据
	now := global.GetGame().GetTimeService().Now()
	err = xianfuManager.RefreshData(now)
	if err != nil {
		return
	}

	if xianfuManager.IsUpgrading(xianfuType) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府升级请求，仙府升级中")

		playerlogic.SendSystemMessage(pl, lang.XianfuArgumentInvalid)
		return
	}

	//是否达到等级上限
	if xfTemplate.GetNextId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府升级请求，当前建筑已达最高级，无法升级")

		playerlogic.SendSystemMessage(pl, lang.XianfuReachMaxGrade)
		return
	}

	nextXfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xfTemplate.GetNextId(), xianfuType)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	upgradeNeedGold := nextXfTemplate.GetUpgradeGold()
	upgradeNeedBindGold := nextXfTemplate.GetUpgradeBindGold()
	totalUpgradeNeedGold := upgradeNeedGold + upgradeNeedBindGold
	upgradeNeedSilver := nextXfTemplate.GetUpgradeYinliang()
	upgradeNeedItemId := nextXfTemplate.GetUpgradeItemId()
	upgradeNeedItemNum := nextXfTemplate.GetUpgradeItemNum()

	//元宝是否足够
	if upgradeNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(upgradeNeedGold), false) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府升级请求，当前元宝不足，无法升级")

			playerlogic.SendSystemMessage(pl, lang.XianfuNotEnoughGold)
			return
		}
	}

	if totalUpgradeNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(totalUpgradeNeedGold), true) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府升级请求，当前绑定元宝不足，无法升级")

			playerlogic.SendSystemMessage(pl, lang.XianfuNotEnoughGold)
			return
		}
	}

	//银两是否足够
	if upgradeNeedSilver > 0 {
		if !propertyManager.HasEnoughSilver(upgradeNeedSilver) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
					"needSilver": upgradeNeedSilver,
				}).Warn("xianfu:秘境仙府升级请求，银两不足，无法升级")

			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//物品是否足够
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if upgradeNeedItemId > 0 && upgradeNeedItemNum > 0 {
		if !inventoryManager.HasEnoughItem(upgradeNeedItemId, upgradeNeedItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府升级请求，物品不足，无法升级")

			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonXianfuUpgrade
	goldReasonText := fmt.Sprintf(goldReason.String(), xianfuId, xianfuType)
	if upgradeNeedGold > 0 {
		flag := propertyManager.CostGold(int64(upgradeNeedGold), false, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("xianfu: xianfuUpgrade use gold should be ok"))
		}
	}
	if upgradeNeedBindGold > 0 {
		flag := propertyManager.CostGold(int64(upgradeNeedBindGold), true, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("xianfu: xianfuUpgrade use gold should be ok"))
		}
	}
	//消耗银两
	if upgradeNeedSilver > 0 {
		silverReason := commonlog.SilverLogReasonXianfuUpgrade
		silverReasonText := fmt.Sprintf(silverReason.String(), xianfuId, xianfuType)
		flag := propertyManager.CostSilver(upgradeNeedSilver, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("xianfu: xianfuUpgrade use silver should be ok"))
		}
	}
	//消耗材料
	if upgradeNeedItemId > 0 && upgradeNeedItemNum > 0 {
		itemReason := commonlog.InventoryLogReasonXianfuUpgrade
		itemReasonText := fmt.Sprintf(itemReason.String(), xianfuId, xianfuType)
		flag := inventoryManager.UseItem(upgradeNeedItemId, upgradeNeedItemNum, itemReason, itemReasonText)
		if !flag {
			panic(fmt.Errorf("xianfu: xianfuUpgrade use item should be ok"))
		}
	}

	//开始升级
	xianfuManager.UpgradePlayerXianfu(xianfuType, now)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	scXianfuUpgrade := pbutil.BuildSCXianfuUpgrade(xianfuId, xianfuType)
	pl.SendMsg(scXianfuUpgrade)

	return
}

//仙府挑战逻辑
func HandleXianfuChallenge(pl player.Player, xianfuType xianfutypes.XianfuType) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeXinaFu) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府挑战请求，功能未开始")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfutypes.XianfuType(xianfuType))

	if xfTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.XianfuArgumentInvalid)
		return
	}

	//刷新数据
	now := global.GetGame().GetTimeService().Now()
	err = xianfuManager.RefreshData(now)
	if err != nil {
		return
	}

	//是否免费次数
	freeTimes := FreeTimesCount(pl, xianfuType)
	if freeTimes < 1 {
		//挑战次数是否足够
		leftTimes := xianfuManager.GetChallengeTimes(xianfuType)
		if leftTimes < 1 {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府挑战请求，副本次数不足，无法挑战")
			playerlogic.SendSystemMessage(pl, lang.XinafuNotEnoughChallengeTimes)
			return
		}

		attendNeedItemId := xfTemplate.GetNeedItemId()
		attendNeedItemNum := xfTemplate.GetNeedItemCount()

		//挑战所需物品是否足够
		if !inventoryManager.HasEnoughItem(attendNeedItemId, attendNeedItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府挑战请求，道具不足，无法挑战")
			playerlogic.SendSystemMessage(pl, lang.XinafuNotEnoughChallengeItem)
			return
		}

		//消耗挑战所需物品
		itemReason := commonlog.InventoryLogReasonXianfuChallenge
		itemReasonText := fmt.Sprintf(itemReason.String(), xianfuId, xianfuType)
		if flag := inventoryManager.UseItem(attendNeedItemId, attendNeedItemNum, itemReason, itemReasonText); !flag {
			panic(fmt.Errorf("xianfu: xianfuChallenge use item should be ok"))
		}
	}

	//更新玩家挑战记录
	xianfuManager.UseTimes(xianfuType, 1, now)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	//进入场景
	flag := PlayerEnterXianFu(pl, xianfuType, xianfuId)
	if !flag {
		panic("enter xianfu scene should be ok!")
	}

	scXianfuChallenge := pbutil.BuildSCXianfuChallenge(xianfuId, xianfuType, flag)
	pl.SendMsg(scXianfuChallenge)
	return
}

func PlayerEnterXianFu(pl player.Player, typ xianfutypes.XianfuType, xianfuId int32) (flag bool) {
	xianFuTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, typ)
	if xianFuTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"xianFuId": xianfuId,
			}).Warn("xianfu:处理跳转仙府,仙府不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := createXianFuSceneData(pl.GetId(), xianFuTemplate)
	s := scene.CreateFuBenScene(int32(xianFuTemplate.GetMapTemplate().TemplateId()), sh)
	if s == nil {
		panic(fmt.Errorf("xianfu:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)

	flag = true
	return
}

//挑战完成奖励
func onXianFuFinish(pl player.Player, xianFuTemplate xianfutemplate.XianFuTemplate, success bool, useTime int64, resource int64, group int32) {
	successRew := make(map[int32]int32)
	xianfuId := int32(xianFuTemplate.TemplateId())
	xianfuType := xianFuTemplate.GetXianFuType()

	if success {
		//通关奖励
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		rewItemMap := xianFuTemplate.GetChallengeRewardsItemMap()
		successRew = rewItemMap
		if !inventoryManager.HasEnoughSlots(rewItemMap) {
			title := lang.GetLangService().ReadLang(lang.EmailXianfuChallengeRewTitle)
			content := lang.GetLangService().ReadLang(lang.EmailXianfuChallengeRewContent)
			emaillogic.AddEmail(pl, title, content, rewItemMap)
		} else {
			rewardItemReason := commonlog.InventoryLogReasonXianfuSaodangRewards
			rewardItemReasonText := fmt.Sprintf(rewardItemReason.String(), xianfuId, xianfuType)
			flag := inventoryManager.BatchAdd(rewItemMap, rewardItemReason, rewardItemReasonText)
			if !flag {
				panic("xianfu:challenge success add item should be ok")
			}

			inventorylogic.SnapInventoryChanged(pl)
		}

		attendTimes := int32(1)
		eventData := xianfueventtypes.CreateXianFuFinishEventData(xianFuTemplate.GetXianFuType(), attendTimes)
		gameevent.Emit(xianfueventtypes.EventTypeXianFuFinish, pl, eventData)
		//记录第几波
		if xianFuTemplate.GetXianFuType() == xianfutypes.XianfuTypeExp {
			s := pl.GetScene()
			if s != nil {
				gameevent.Emit(xianfueventtypes.EventTypeXianFuRefreshGroup, pl, group)
			}
		}
	}

	//发送提示
	scXianfuChallenge := pbutil.BuildSCXianfuChallengeResult(xianfuId, int32(xianfuType), success, resource, successRew, group)
	pl.SendMsg(scXianfuChallenge)
	return
}

//下发场景信息
func onPushSceneInfo(playerId int64, curTemp xianfutemplate.XianFuTemplate, killNum int32, resource int64, s scene.Scene) {
	if s == nil {
		return
	}
	pl := s.GetPlayer(playerId)
	if pl == nil {
		return
	}
	xianfuId := int32(curTemp.TemplateId())
	xianfuType := int32(curTemp.GetXianFuType())
	currentGroup := s.GetCurrentGroup()
	totalGroup := curTemp.GetMapTemplate().GetNumGroup()
	startTime := s.GetStartTime()

	scXianfuSceneInfo := pbutil.BuildSCXianfuSceneInfo(xianfuId, xianfuType, killNum, currentGroup, totalGroup, resource, startTime)
	pl.SendMsg(scXianfuSceneInfo)
	return
}

//下发怪物刷新波数
func onPushBiologyGroupInfo(playerId int64, groupIndex int32, s scene.Scene) {
	if s == nil {
		return
	}
	pl := s.GetPlayer(playerId)
	if pl == nil {
		return
	}
	showGroup := groupIndex + 1
	scXianfuNextBiologyGroup := pbutil.BuildSCXianfuNextBiologyGroup(showGroup)
	pl.SendMsg(scXianfuNextBiologyGroup)
	gameevent.Emit(xianfueventtypes.EventTypeXianFuRefreshGroup, pl, groupIndex)
}

//挑战奖励信息更新
func onPushResourceInfo(pl player.Player, resource int64) {
	scXianfuRewNotice := pbutil.BuildSCXianfuRewNotice(resource)
	pl.SendMsg(scXianfuRewNotice)

}

//下发怪物死亡数量
func onPushKillNum(playerId int64, num int32, s scene.Scene) {
	if s == nil {
		return
	}
	pl := s.GetPlayer(playerId)
	if pl == nil {
		return
	}
	scXianfuKillNumNotice := pbutil.BuildSCXianfuKillNumNotice(num)
	pl.SendMsg(scXianfuKillNumNotice)

}

//仙府扫荡
func GetSaoDangDrop(saoDangNum int32, xianfuId int32, xianfuType xianfutypes.XianfuType) (showItemList [][]*droptemplate.DropItemData, rewardsItemList []*droptemplate.DropItemData, rewardsResMap map[itemtypes.ItemAutoUseResSubType]int32) {
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)

	var totalItemList []*droptemplate.DropItemData
	dropIdList := xfTemplate.GetSaodangRewardDropArr()
	for index := int32(0); index < saoDangNum; index++ {
		var itemList []*droptemplate.DropItemData
		// 固定奖励
		for itemId, num := range xfTemplate.GetSaodangRewardItemMap(1) {
			newData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeUnBind)
			itemList = append(itemList, newData)
		}

		// 掉落包
		dropItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
		itemList = append(itemList, dropItemList...)

		showItemList = append(showItemList, itemList)
		totalItemList = append(totalItemList, itemList...)
	}

	if len(totalItemList) != 0 {
		rewardsItemList, rewardsResMap = droplogic.SeperateItemDatas(totalItemList)
	}

	return
}

// 剩余免费次数
func FreeTimesCount(pl player.Player, xianFuType xianfutypes.XianfuType) int32 {

	xianfuManager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	obj := xianfuManager.GetPlayerXianfuInfo(xianFuType)
	if obj == nil {
		return 0
	}

	//基础免费次数
	xianFuTemp := xianfutemplate.GetXianfuTemplateService().GetXianfu(obj.GetXianfuId(), xianFuType)
	if xianFuTemp == nil {
		return 0
	}
	maxFree := xianFuTemp.GetFree()
	// 仙尊免费次数
	maxFree += xianzuncardlogic.XianFuFreeTimes(pl, xianFuType)

	if obj.GetUseTimes() > maxFree {
		return 0
	}

	return maxFree - obj.GetUseTimes()
}

//仙府所有免费次数
func TotalFreeTimes(pl player.Player) int32 {

	total := int32(0)
	for initType := xianfutypes.MinXianfuType; initType < xianfutypes.MaxXianfuType; initType++ {
		total += FreeTimesCount(pl, initType)
	}

	return total
}
