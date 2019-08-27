package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/gem/gem"
	"fgame/fgame/game/gem/pbutil"
	gemtypes "fgame/fgame/game/gem/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"

	playergem "fgame/fgame/game/gem/player"

	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"

	log "github.com/Sirupsen/logrus"
)

//赌石奖励
func GiveGembleReward(pl player.Player, typ gemtypes.GemGambleType, curGambleNum int32, batchNum int32) (dropItemList []*droptemplate.DropItemData, isReturn bool, err error) {
	//掉落物品分类
	for i := int32(0); i < batchNum; i++ {
		curGambleNum += 1
		dropId, _ := gem.GetGemService().GetDropIdByNum(pl, typ, curGambleNum)
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData == nil {
			continue
		}

		dropItemList = append(dropItemList, dropData)
	}

	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) != 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(dropItemList)
	}

	//判断背包是否足够
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(newItemList) != 0 {
		flag := inventoryManager.HasEnoughSlotsOfItemLevel(newItemList)
		if !flag {
			isReturn = true
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}

		reasonText := commonlog.InventoryLogReasonGemGambleDrop.String()
		flag = inventoryManager.BatchAddOfItemLevel(newItemList, commonlog.InventoryLogReasonGemGambleDrop, reasonText)
		if !flag {
			panic(fmt.Errorf("gem: GiveGembleReward BatchAdd should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	//添加掉落资源
	if len(resMap) != 0 {
		reasonGoldText := commonlog.GoldLogReasonGemGambleDrop.String()
		reasonSliverText := commonlog.SilverLogReasonGemGambleDrop.String()
		reasonLevelText := commonlog.LevelLogReasonGemGambleDrop.String()
		err = droplogic.AddRes(pl, resMap, commonlog.GoldLogReasonGemGambleDrop, reasonGoldText, commonlog.SilverLogReasonGemGambleDrop, reasonSliverText, commonlog.LevelLogReasonGemGambleDrop, reasonLevelText)
		if err != nil {
			return
		}
	}
	return
}

//处理赌石信息逻辑
func HandleGemGamble(pl player.Player, typ gemtypes.GemGambleType, tenEven bool) (err error) {
	gemManager := pl.GetPlayerDataManager(playertypes.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	flag := typ.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"tenEven":  tenEven,
		}).Warn("gem:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	gamblingTemplate := gem.GetGemService().GetGambleTemplateByTyp(gemtypes.GemGambleType(typ))
	needYinLiang := int64(gamblingTemplate.NeedYinLiang)
	needGold := gamblingTemplate.NeedGold
	needYuanShi := gamblingTemplate.NeedYuanShi
	needItem := int32(0)
	needItemNum := int32(0)
	useItem := gamblingTemplate.GetUseItemTemplate()
	batchNum := int32(1)
	if useItem != nil {
		needItem = gamblingTemplate.NeedItem
		needItemNum = gamblingTemplate.NeedItemNum
	}
	//十连赌
	if tenEven {
		needYinLiang *= 10
		needGold *= 10
		needYuanShi *= 10
		needItemNum *= 10
		batchNum *= 10
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//判断银两
	if needYinLiang != 0 {
		flag = propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"tenEven":  tenEven,
			}).Warn("gem:银两不足，无法赌石")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//判断元宝
	if needGold != 0 {
		flag = propertyManager.HasEnoughGold(int64(needGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"tenEven":  tenEven,
			}).Warn("gem:元宝不足，无法赌石")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//判断原石
	if needYuanShi != 0 {
		flag = gemManager.HasEnoughYuanShi(needYuanShi)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"tenEven":  tenEven,
			}).Warn("gem:原石不足，无法赌石")
			playerlogic.SendSystemMessage(pl, lang.GemGambleNotYuanShi)
			return
		}
	}

	//物品判断
	if needItem != 0 {
		flag = inventoryManager.HasEnoughItem(needItem, needItemNum)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"tenEven":  tenEven,
			}).Warn("gem:道具不足，无法赌石")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	curGambleNum := gemManager.GetGambleNum(typ)
	dropItemList, isReturn, err := GiveGembleReward(pl, typ, curGambleNum, batchNum)
	if err != nil || isReturn {
		return
	}

	//消耗资源
	if needYinLiang != 0 || needGold != 0 {
		reasonGoldText := commonlog.GoldLogReasonGemGamble.String()
		reasonSliverText := commonlog.SilverLogReasonGemGamble.String()
		flag = propertyManager.Cost(0, int64(needGold), commonlog.GoldLogReasonGemGamble, reasonGoldText, needYinLiang, commonlog.SilverLogReasonGemGamble, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("gem: gemGamble Cost should be ok"))
		}
	}
	//同步元宝银两
	propertylogic.SnapChangedProperty(pl)

	//消耗原石
	flag = gemManager.GambleSubStone(needYuanShi)
	if !flag {
		panic(fmt.Errorf("gem: GambleSubStone should be ok"))
	}
	//更新赌石次数
	flag = gemManager.AddGambleNum(typ, batchNum)
	if !flag {
		panic(fmt.Errorf("gem: AddGambleNum should be ok"))
	}

	stone := gemManager.GetMineStone()
	scGemGamble := pbutil.BuildSCGemGamble(dropItemList, stone)
	pl.SendMsg(scGemGamble)
	return
}
