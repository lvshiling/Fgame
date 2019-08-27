package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	materiallogic "fgame/fgame/game/material/logic"
	"fgame/fgame/game/material/pbutil"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	weeklogic "fgame/fgame/game/week/logic"

	playermaterial "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MATERIAL_SAO_DANG_TYPE), dispatch.HandlerFunc(handlerMaterialSaodang))
}

//材料副本扫荡请求
func handlerMaterialSaodang(s session.Session, msg interface{}) (err error) {
	log.Debug("material:处理材料副本扫荡请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMaterialSaoDang)
	typ := csMsg.GetMaterialType()
	num := csMsg.GetNum()

	materialType := materialtypes.MaterialType(typ)
	//验证参数
	if !materialType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
			}).Warn("material:材料副本扫荡请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
				"num":          num,
			}).Warn("material:材料副本扫荡请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = materialSaodang(tpl, materialType, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     tpl.GetId(),
				"materialType": materialType,
				"err":          err,
			}).Error("material:处理材料副本扫荡请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     tpl.GetId(),
			"materialType": materialType,
		}).Debug("material：处理材料副本扫荡请求完成")

	return
}

//材料副本扫荡逻辑
func materialSaodang(pl player.Player, materialType materialtypes.MaterialType, saoDangNum int32) (err error) {
	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(materialType)
	if materialTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
			}).Warn("material:材料副本扫荡请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	materialInfo := materialManager.GetPlayerMaterialInfo(materialType)

	//波数限制
	if materialInfo.GetGroup() < materialTemplate.GroupLimit {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
				"Group":        materialInfo.GetGroup(),
			}).Warn("material:材料副本扫荡请求，副本波数0不能扫荡")
		playerlogic.SendSystemMessage(pl, lang.MaterialGroupNotEnough)
		return
	}

	// 玩家等级
	if pl.GetLevel() < materialTemplate.NeedLevel {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
				"Group":        materialInfo.GetGroup(),
			}).Warn("material:材料副本扫荡请求，玩家等级不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//刷新数据
	materialManager.RefreshData()

	//挑战次数是否足够
	if !materialManager.IsEnoughAttendTimes(materialType, saoDangNum) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
				"saoDangNum":   saoDangNum,
			}).Warn("material:材料副本扫荡请求，副本次数不足，无法扫荡")
		playerlogic.SendSystemMessage(pl, lang.MaterialNotEnoughChallengeTimes)
		return
	}

	needItemId := materialTemplate.NeedItemId
	needItemNum := int32(0)
	saodangNeedGold := int64(0)
	freeTime := materialManager.GetFreeAttendTimes(materialType)
	if saoDangNum > freeTime {
		needItemNum = materialTemplate.NeedItemCount * (saoDangNum - freeTime)
		saodangNeedGold = int64(materialTemplate.SaodangNeedGold * (saoDangNum - freeTime))
	}

	//挑战所需物品是否足够
	if needItemNum > 0 {
		if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"materialType": materialType,
					"needItemNum":  needItemNum,
					"saoDangNum":   saoDangNum,
				}).Warn("material:材料副本扫荡请求，副本挑战令不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//免费扫荡
	saodangNeedItemMap := make(map[int32]int32)
	if !weeklogic.IsSeniorWeek(pl) {
		saodangNeedItemMap = materialTemplate.GetSaodangItemMap(saoDangNum)
	}

	//扫荡所需物品是否足够
	if len(saodangNeedItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(saodangNeedItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"materialType": materialType,
					"saoDangNum":   saoDangNum,
				}).Warn("material:材料副本扫荡请求，当前扫荡券不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//预留字段校验：扫荡所需元宝是否足够
	if saodangNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(saodangNeedGold), true) {
			log.WithFields(
				log.Fields{
					"playerId":        pl.GetId(),
					"materialType":    materialType,
					"saoDangNum":      saoDangNum,
					"saodangNeedGold": saodangNeedGold,
				}).Warn("material:材料副本扫荡请求，当前元宝不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.XianfuNotEnoughGold)
			return
		}
	}

	showItemList, rewardsItemList, rewardsResMap, totalRewData, flag := materiallogic.GetSaoDangDrop(pl, saoDangNum, materialType, materialInfo.GetGroup())
	if !flag {
		return
	}

	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewardsItemList) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
				"saoDangNum":   saoDangNum,
			}).Warn("material:材料副本扫荡请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//扣除扫荡物品
	useItemReason := commonlog.InventoryLogReasonMaterialSaodangUse
	useItemReasonText := fmt.Sprintf(useItemReason.String(), saoDangNum, materialType.String())
	if len(saodangNeedItemMap) > 0 {
		if flag := inventoryManager.BatchRemove(saodangNeedItemMap, useItemReason, useItemReasonText); !flag {
			panic("material: materialSaodang use item should be ok")
		}
	}

	//扣除挑战物品
	if needItemNum > 0 {
		if flag := inventoryManager.UseItem(needItemId, needItemNum, useItemReason, useItemReasonText); !flag {
			panic("material: materialSaodang use item should be ok")
		}
	}

	//消耗元宝
	if saodangNeedGold > 0 {
		goldReason := commonlog.GoldLogReasonMaterialSaoDangUse
		goldReasonText := fmt.Sprintf(goldReason.String(), saoDangNum, materialType.String())
		flag := propertyManager.CostGold(saodangNeedGold, true, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("material: materialSaodang use gold should be ok"))
		}
	}

	//增加物品
	getItemReason := commonlog.InventoryLogReasonMaterialSaodangGet
	getItemReasonText := fmt.Sprintf(getItemReason.String(), saoDangNum, materialType.String())
	flag = inventoryManager.BatchAddOfItemLevel(rewardsItemList, getItemReason, getItemReasonText)
	if !flag {
		panic("material:materialSaodang add item should be ok")
	}

	//获取扫荡固定资源
	reasonGold := commonlog.GoldLogReasonMaterialSaodangGet
	reasonSilver := commonlog.SilverLogReasonMaterialSaodangGet
	reasonLevel := commonlog.LevelLogReasonMaterialSaodangGet
	saodangGoldReasonText := fmt.Sprintf(reasonGold.String(), saoDangNum, materialType.String())
	saodangSilverReasonText := fmt.Sprintf(reasonSilver.String(), saoDangNum, materialType.String())
	expReasonText := fmt.Sprintf(reasonLevel.String(), saoDangNum, materialType.String())
	flag = propertyManager.AddRewData(totalRewData, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
	if !flag {
		panic("material:materialSaodang add RewData should be ok")
	}

	//增加资源
	if len(rewardsResMap) > 0 {
		err = droplogic.AddRes(pl, rewardsResMap, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
		if err != nil {
			return
		}
	}

	//完成扫荡
	materialManager.UseTimes(materialType, saoDangNum)
	materialManager.EmitFinishEvent(materialType, saoDangNum)
	// materialManager.EmitSweepEvent(materialType, saoDangNum)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCMaterialSaoDang(materialType, saoDangNum, showItemList)
	pl.SendMsg(scMsg)
	return
}
