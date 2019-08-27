package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	commontypes "fgame/fgame/game/common/types"
	gamevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	item "fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_SHENZHU_BODY_TYPE), dispatch.HandlerFunc(handleAdditionSysShenZhuBody))
}

//处理附加系统神铸部位
func handleAdditionSysShenZhuBody(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsys:处理附加系统神铸部位")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAdditionSysShenZhuBody := msg.(*uipb.CSAdditionSysShenZhuBody)
	sysTypeInt := csAdditionSysShenZhuBody.GetSysType()
	SlotIdInt := csAdditionSysShenZhuBody.GetSlotId()
	sysType := additionsystypes.AdditionSysType(sysTypeInt)
	pos := additionsystypes.SlotPositionType(SlotIdInt)

	//参数不对
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("shenzhu:系统类型神铸部位,错误")
		return
	}
	if !pos.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,错误")
		return
	}

	err = additionSysShenZhuBody(tpl, sysType, pos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
				"pos":      pos.String(),
				"error":    err,
			}).Warn("shenzhu:系统类型神铸部位,错误")
		return
	}

	log.Debug("shenzhu:系统类型神铸部位,完成")
	return nil

}

//系统类型神铸部位逻辑
func additionSysShenZhuBody(pl player.Player, typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) (err error) {
	if !additionsyslogic.GetAdditionSysShenZhuFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("shenzhu:系统类型神铸部位,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	equipBag := additionsysManager.GetAdditionSysEquipBagByType(typ)

	slotObject := equipBag.GetByPosition(pos)
	if slotObject == nil || slotObject.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	//判断槽位是否可以升
	nextTemp := equipBag.GetNextShenZhuTemplate(pos)
	if nextTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,已经满级")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysShenZhuHighest)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(slotObject.ItemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	if itemTemplate.GetQualityType() < itemtypes.ItemQualityTypeOrange {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,装备品质不足")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysShenZhuQualityLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	useTemp := additionsystemplate.GetAdditionSysTemplateService().GetShenZhuUseByType(typ)
	//进阶需要消耗的银两
	costSilver := int64(nextTemp.UseMoney)
	//进阶需要消耗的元宝
	costGold := int32(0)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	needItemCount := nextTemp.ItemCount
	needItemId := int32(useTemp.GetUseItemByPos(pos).TemplateId())
	if !inventoryManager.HasEnoughItem(needItemId, needItemCount) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ.String(),
			"pos":      pos.String(),
		}).Warn("shenzhu:系统类型神铸部位,物品不足无法升级")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysShenZhuItemNoEnough)
		return
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"sysType":  typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,银两不足无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"sysType":  typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,元宝不足无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"sysType":  typ.String(),
				"pos":      pos.String(),
			}).Warn("shenzhu:系统类型神铸部位,元宝不足无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonAdditionSysShenZhuCost
	silverUseReason := commonlog.SilverLogReasonAdditionSysShenZhuCost
	goldUseReasonText := fmt.Sprintf(goldUseReason.String(), typ.String(), pos.String(), slotObject.ShenZhuLev, slotObject.ShenZhuPro, slotObject.ShenZhuNum)
	silverUseReasonText := fmt.Sprintf(silverUseReason.String(), typ.String(), pos.String(), slotObject.ShenZhuLev, slotObject.ShenZhuPro, slotObject.ShenZhuNum)
	useCostflag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReasonText, costSilver, silverUseReason, silverUseReasonText)
	if !useCostflag {
		panic(fmt.Errorf("shenzhu: additionsys shenzhu Cost should be ok"))
	}

	//消耗物品
	inventoryReason := commonlog.InventoryLogReasonAdditionSysShenZhuCost
	reasonText := fmt.Sprintf(inventoryReason.String(), typ.String(), pos.String(), slotObject.ShenZhuLev, slotObject.ShenZhuPro, slotObject.ShenZhuNum)
	useItemflag := inventoryManager.UseItem(int32(needItemId), needItemCount, inventoryReason, reasonText)
	if !useItemflag {
		panic(fmt.Errorf("shenzhu: shenzhu use item should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//进阶判断
	beforeLev := slotObject.ShenZhuLev
	sucess, pro, _, addTimes := additionsyslogic.AdditionSysShenZhuJudge(pl, slotObject.ShenZhuNum, slotObject.ShenZhuPro, nextTemp)
	equipBag.ShenZhuLevel(pos, pro, addTimes, sucess)

	if sucess {
		//更新属性
		additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)
		//日志
		additionsysReason := commonlog.AdditionSysLogReasonShenZhu
		reasonText := fmt.Sprintf(additionsysReason.String(), typ.String(), pos.String(), commontypes.AdvancedTypeAdditionsysEquip.String())
		data := additionsyseventtypes.CreatePlayerAdditionSysShenZhuLevLogEventData(typ, pos, beforeLev, additionsysReason, reasonText)
		gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysShenZhuLog, pl, data)
	}

	scMsg := pbutil.BuildSCAdditionSysShenZhuBody(typ, pos, slotObject.ShenZhuLev, slotObject.ShenZhuPro)
	pl.SendMsg(scMsg)
	return
}
