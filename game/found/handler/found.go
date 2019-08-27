package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/found/pbutil"
	playerfound "fgame/fgame/game/found/player"
	foundtemplate "fgame/fgame/game/found/template"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/global"
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
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOUND_TYPE), dispatch.HandlerFunc(handlerFound))
}

//资源找回
func handlerFound(s session.Session, msg interface{}) (err error) {
	log.Debug("found:处理资源找回请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csFound := msg.(*uipb.CSFound)
	resType := csFound.GetResType()
	typ := csFound.GetTyp()

	foundResType := foundtypes.FoundResourceType(resType)
	foundType := foundtypes.FoundType(typ)
	if !foundResType.Valid() || !foundType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"resType":  resType,
			}).Warn("found:参数错误")

		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = foundRes(tpl, foundResType, foundType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"resType":  resType,
				"typ":      typ,
				"err":      err,
			}).Error("found:处理资源找回请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"resType":  resType,
			"typ":      typ,
		}).Debug("found：处理资源找回请求完成")

	return
}

func foundRes(pl player.Player, resType foundtypes.FoundResourceType, typ foundtypes.FoundType) (err error) {
	foundManager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	resLevel := foundManager.GetFoundResLevel(resType)
	foundTemp := foundtemplate.GetFoundTemplateService().GetFoundTemplateByType(resType, resLevel)
	if foundTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"resType":  resType,
			}).Warn("found:参数错误,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否能够找回
	flag := foundManager.IsCanFoundBack(resType)
	flag2 := foundManager.IsReceiveFound(resType)
	if !flag || flag2 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"resType":  resType,
			}).Warn("found:参数错误,该资源不满足找回条件")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch typ {
	case foundtypes.FoundTypeGold:
		{
			foundTimes := foundManager.GetFoundTimes(resType)
			totalNeedGold := int64(foundTemp.FoundUsing * foundTimes)
			//判断元宝是否足够
			if !propertyManager.HasEnoughGold(totalNeedGold, true) {
				log.WithFields(
					log.Fields{
						"playerId":      pl.GetId(),
						"typ":           typ,
						"resType":       resType,
						"totalNeedGold": totalNeedGold,
					}).Warn("found:元宝不足，无法完美找回")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}

			//消耗元宝
			goldUseReason := commonlog.GoldLogReasonFoundResourceCost
			goldUseReasonText := fmt.Sprintf(goldUseReason.String(), resType, foundTimes)
			flag := propertyManager.CostGold(totalNeedGold, true, goldUseReason, goldUseReasonText)
			if !flag {
				panic("found:消耗元宝应该成功")
			}
		}
	case foundtypes.FoundTypeFree:
		{
			foundTimes := foundManager.GetFoundTimes(resType)
			totalNeedSilver := int64(foundTemp.FoundUsingSilver * foundTimes)
			//判断银两是否足够
			if !propertyManager.HasEnoughSilver(totalNeedSilver) {
				log.WithFields(
					log.Fields{
						"playerId":        pl.GetId(),
						"typ":             typ,
						"resType":         resType,
						"totalNeedSilver": totalNeedSilver,
					}).Warn("found:银两不足，无法普通找回")
				playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
				return
			}

			//消耗银两
			if totalNeedSilver > 0 {
				silverUseReason := commonlog.SilverLogReasonFoundResourceUse
				silverUseReasonText := fmt.Sprintf(silverUseReason.String(), resType, foundTimes)
				flag := propertyManager.CostSilver(totalNeedSilver, silverUseReason, silverUseReasonText)
				if !flag {
					panic("found:消耗银两应该成功")
				}
			}
		}

	default:
		break
	}

	totalItemMap := addRes(pl, resType, foundTemp.GetFoundData(typ))

	scFound := pbutil.BuildSCFound(int32(resType), totalItemMap)
	pl.SendMsg(scFound)
	return
}

func addRes(pl player.Player, resType foundtypes.FoundResourceType, foundData foundtypes.FoundData) (totalItemMap map[int32]int32) {
	currentLevel := pl.GetLevel()
	foundManager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	foundTimes := foundManager.GetFoundTimes(resType)
	group := foundManager.GetFoundGroup(resType)

	//增加资源
	reasonGold := commonlog.GoldLogReasonFoundResource
	reasonSilver := commonlog.SilverLogReasonFoundResource
	reasonLevel := commonlog.LevelLogReasonFoundResource
	goldReasonText := fmt.Sprintf(reasonGold.String(), resType, foundTimes, group)
	silverReasonText := fmt.Sprintf(reasonSilver.String(), resType, foundTimes, group)
	expReasonText := fmt.Sprintf(reasonLevel.String(), resType, foundTimes, group)

	rewSilver := foundData.FoundSilver * foundTimes * group
	rewBindGold := foundData.FoundBindgold * foundTimes * group
	rewGold := foundData.FoundGold * foundTimes * group
	rewExp := foundData.FoundExp * foundTimes * group
	rewExpPoint := foundData.FoundExpPoint * foundTimes * group
	totalResData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)

	flag := propertyManager.AddRewData(totalResData, reasonGold, goldReasonText, reasonSilver, silverReasonText, reasonLevel, expReasonText)
	if !flag {
		panic("found:found add resource should be ok")
	}

	//增加物品
	newItemMap := make(map[int32]int32)
	var newItemDataList []*droptemplate.DropItemData
	for itemId, num := range foundData.FoundItemMap {
		itemNum := num * foundTimes * group
		level := int32(0)
		bind := itemtypes.ItemBindTypeUnBind

		newItemMap[itemId] = itemNum
		newData := droptemplate.CreateItemData(itemId, itemNum, level, bind)
		newItemDataList = append(newItemDataList, newData)
	}

	flag = inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList)
	if flag {
		inventoryReason := commonlog.InventoryLogReasonFoundResource
		inventoryReasonText := fmt.Sprintf(inventoryReason.String(), resType, foundTimes)
		flag = inventoryManager.BatchAddOfItemLevel(newItemDataList, inventoryReason, inventoryReasonText)
		if !flag {
			panic("found:found add item should be ok")
		}
	} else {
		now := global.GetGame().GetTimeService().Now()
		emaillogic.AddEmailItemLevel(pl, "背包空间不足", "资源找回", now, newItemDataList)
	}

	foundManager.ReceiveFound(resType)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	totalItemMap = propertylogic.CombineRewDataAndItemData(totalResData, currentLevel, newItemMap)
	return totalItemMap
}
