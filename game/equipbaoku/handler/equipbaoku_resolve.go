package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	equipbaokulogic "fgame/fgame/game/equipbaoku/logic"
	"fgame/fgame/game/equipbaoku/pbutil"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_RESOLVE_EQUIP_TYPE), dispatch.HandlerFunc(handleEquipBaoKuResolveEquip))
}

//处理分解宝库装备
func handleEquipBaoKuResolveEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:处理宝库装备分解")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEquipBaoKuResolveEquip := msg.(*uipb.CSEquipbaokuResolveEquip)
	itemIndexList := csEquipBaoKuResolveEquip.GetIndexList()

	err = equipBaoKuResolveEquip(tpl, itemIndexList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
				"error":     err,
			}).Error("equipbaoku:处理宝库装备分解,错误")

		return err
	}
	log.Debug("equipbaoku:处理宝库装备分解,完成")
	return nil
}

//分解
func equipBaoKuResolveEquip(pl player.Player, itemIndexList []int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeEquipBaoKu) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:处理分解宝库装备,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	if len(itemIndexList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("equipbaoku:处理分解宝库装备,没有装备")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentNotItemEat)
		return
	}

	if coreutils.IfRepeatElementInt32(itemIndexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("equipbaoku:处理分解宝库装备,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 计算分解物品
	totalExp, returnItemMap, flag := equipbaokulogic.CountResolveMiBaoDepotEquip(pl, itemIndexList)
	if !flag {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlots(returnItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
			}).Warn("equipbaoku:分解失败,背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗装备
	useReason := commonlog.InventoryLogReasonEquipBaoKuResolveUse
	flag, err = inventoryManager.BatchRemoveMiBaoDepotByIndex(itemIndexList, useReason, useReason.String(), equipbaokutypes.BaoKuTypeEquip)
	if err != nil {
		return
	}
	if !flag {
		panic(fmt.Errorf("equipbaoku:消耗物品应该成功"))
	}

	// 返还物品
	if len(returnItemMap) > 0 {
		addItemReason := commonlog.InventoryLogReasonEquipBaoKuResolveReturn
		flag := inventoryManager.BatchAdd(returnItemMap, addItemReason, addItemReason.String())
		if !flag {
			panic(fmt.Errorf("equipbaoku:添加物品应该成功"))
		}
	}

	if totalExp > 0 {
		goldLevelReason := commonlog.GoldYuanLevelLogReasonEatEquip
		propertyManager.AddGoldYuanExp(totalExp, goldLevelReason, goldLevelReason.String())
	}

	//整理
	//inventoryManager.MergeMiBaoDepot()

	inventorylogic.SnapMiBaoDepotChanged(pl, equipbaokutypes.BaoKuTypeEquip)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//发送事件
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipResolve, pl, int32(len(itemIndexList)))
	scMsg := pbutil.BuildSCEquipBaoKuResolveEquip(propertyManager.GetGoldYuanLevel(), propertyManager.GetGoldYuanExp(), returnItemMap)
	pl.SendMsg(scMsg)
	return
}
