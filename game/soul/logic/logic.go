package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//帝魂属性
func SoulPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeSoul.Mask())
	return
}

//帝魂强化判断
func SoulStrengthen(curTimesNum int32, curBless int32, strengthenTemplate *gametemplate.SoulLevelUpTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := strengthenTemplate.TimesMin
	timesMax := strengthenTemplate.TimesMax
	updateRate := strengthenTemplate.UpdateWfb
	blessMax := strengthenTemplate.ZhufuMax
	addMin := strengthenTemplate.AddMin
	addMax := strengthenTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//处理帝魂激活信息逻辑
func HandleSoulActive(pl player.Player, soulTag soultypes.SoulType) (err error) {
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := soulTag.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = soulManager.IfSoulTagExist(soulTag)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:该帝魂已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.SoulRepeatActive)
		return
	}

	soulActiveTemplate := soul.GetSoulService().GetSoulActiveTemplate(soulTag)
	//激活的前置帝魂条件
	preSoulCond := soulActiveTemplate.GetPreSoulCond()
	if preSoulCond != nil {
		preSoulTag := preSoulCond.GetSoulType()
		flag := soulManager.IfPreSoul(preSoulTag, preSoulCond.Level)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
			}).Warn("soul:激活该帝魂的前置条件不足")
			playerlogic.SendSystemMessage(pl, lang.SoulActiveNotPreCond)
			return
		}
	}

	//激活需要物品
	items := soulActiveTemplate.GetNeedItemMap()
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
			}).Warn("soul:您的物品不足，无法激活帝魂")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonSoulActive
		reasonText := fmt.Sprintf(inventoryReason.String(), soulTag)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soul: soulActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	soulObj, autoEmbed, flag := soulManager.SoulActive(soulTag)
	if !flag {
		panic(fmt.Errorf("soul: soulActive should be ok"))
	}
	//同步属性
	SoulPropertyChanged(pl)
	//同步属性
	propertylogic.SnapChangedProperty(pl)
	scSoulActive := pbutil.BuildSCSoulActive(soulObj)
	pl.SendMsg(scSoulActive)

	if autoEmbed {
		soulId, flag := soulManager.GetSoulIdByOrder(soulTag)
		if !flag {
			panic(fmt.Errorf("soul: soulActive GetSoulIdByOrder should be ok"))
		}
		scSoulEmbed := pbutil.BuildSCSoulEmbed(soulId)
		pl.SendMsg(scSoulEmbed)
	}
	return
}

//帝魂镶嵌逻辑
func HandleSoulEmbed(pl player.Player, soulTag soultypes.SoulType) (err error) {
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := soulTag.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = soulManager.IfSoulTagExist(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:未激活的帝魂,无法镶嵌")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotEmbed)
		return
	}

	flag = soulManager.IfSoulTagEmemded(soulTag)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:该帝魂已镶嵌,重复镶嵌")
		playerlogic.SendSystemMessage(pl, lang.SoulRepeatEmbed)
		return
	}

	//镶嵌
	flag = soulManager.Embed(soulTag)
	if !flag {
		panic(fmt.Errorf("soul: soulEmbed should be ok"))
	}

	soulId, flag := soulManager.GetSoulIdByOrder(soulTag)
	if !flag {
		panic(fmt.Errorf("soul: soulEmbed GetSoulIdByOrder should be ok"))
	}

	scSoulEmbed := pbutil.BuildSCSoulEmbed(soulId)
	pl.SendMsg(scSoulEmbed)
	return
}
