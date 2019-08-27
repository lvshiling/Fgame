package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	playerlucky "fgame/fgame/game/lucky/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	synthesiseventtypes "fgame/fgame/game/synthesis/event/types"
	"fgame/fgame/game/synthesis/pbutil"
	synthesistemplate "fgame/fgame/game/synthesis/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SYNRHESIS_START_TYPE), dispatch.HandlerFunc(handlerSynthesisStart))
}

//处理合成请求
func handlerSynthesisStart(s session.Session, msg interface{}) error {
	log.Debug("synthesis:处理合成请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	cSSynthesisStart := msg.(*uipb.CSSynthesisStart)
	synthesisId := cSSynthesisStart.GetSynthesisId()
	num := cSSynthesisStart.GetNum()
	isUseAvoidBomb := cSSynthesisStart.GetIsAvoidBomb()

	err := synthesisStart(tpl, synthesisId, num, isUseAvoidBomb)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
				"num":         num,
				"error":       err,
			}).Error("synthesis:处理合成请求,错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId":    pl.GetId(),
			"synthesisId": synthesisId,
			"num":         num,
		}).Debug("synthesis:处理合成请求完成")

	return nil
}

func synthesisStart(pl player.Player, synthesisId int32, num int32, isUseAvoidBomb bool) (err error) {
	//校验入参数据
	synthesisTemplate := synthesistemplate.GetSynthesisTemplateService().GetSynthesis(synthesisId)
	if synthesisTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
				"num":         num,
			}).Warn("synthesis:处理合成请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if synthesisTemplate.NeedGender > 0 {
		if pl.GetSex() != synthesisTemplate.GetNeedGender() {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"synthesisId": synthesisId,
					"num":         num,
				}).Warn("synthesis:处理合成请求,性别不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
			return
		}
	}

	if synthesisTemplate.NeedProfession > 0 {
		if pl.GetRole() != synthesisTemplate.GetNeedRole() {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"synthesisId": synthesisId,
					"num":         num,
				}).Warn("synthesis:处理合成请求,职业不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
			return
		}
	}

	if synthesisTemplate == nil || num <= 0 || num > synthesisTemplate.MaxCount {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
				"num":         num,
			}).Warn("synthesis:处理合成请求,参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断银两是否足够
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	needSilver := int64(synthesisTemplate.NeedSilver * num)
	if needSilver > 0 {
		if !propertyManager.HasEnoughSilver(needSilver) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"synthesisId": synthesisId,
					"num":         num,
				}).Warn("synthesis:银两不足，无法合成")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//判断元宝是否足够
	needGold := synthesisTemplate.NeedGold * num
	needBindGoldNum := synthesisTemplate.NeedBindGold * num
	needTotalGold := needGold + needBindGoldNum
	if needGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needGold), false) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"synthesisId": synthesisId,
					"num":         num,
				}).Warn("synthesis:元宝不足，无法合成")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}
	if needTotalGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needTotalGold), true) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"synthesisId": synthesisId,
					"num":         num,
				}).Warn("synthesis:绑元不足，无法合成")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//判断背包材料是否足够
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	rewItemId := synthesisTemplate.ItemId
	rewItemNum := synthesisTemplate.ItemCount * num
	defaultLevel := int32(0)
	bindType := itemtypes.ItemBindTypeUnBind

	materialItemMap := make(map[int32]int32)
	for itemNeedId, itemNeedNum := range synthesisTemplate.GetSynthesisMap() {
		materialItemMap[itemNeedId] = itemNeedNum * num

		// 绑定属性
		if bindType == itemtypes.ItemBindTypeBind {
			continue
		}
		if inventoryManager.IsContainBindItem(itemNeedId) {
			bindType = itemtypes.ItemBindTypeBind
		}
	}
	//判断背包空间是否足够
	if !inventoryManager.HasEnoughSlotItemLevel(rewItemId, rewItemNum, defaultLevel, bindType) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
				"num":         num,
			}).Warn("synthesis:处理合成请求,背包空间不足，请清理后再来")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	if !inventoryManager.HasEnoughItems(materialItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"synthesisId":     synthesisId,
				"materialItemMap": materialItemMap,
				"num":             num,
			}).Warn("synthesis:材料不足，无法合成")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	// 防爆符
	if isUseAvoidBomb {
		needItemId := synthesisTemplate.ExplosionId
		needNum := synthesisTemplate.ExplosionCount * num
		if !inventoryManager.HasEnoughItem(needItemId, needNum) {
			log.WithFields(
				log.Fields{
					"playerId":        pl.GetId(),
					"synthesisId":     synthesisId,
					"materialItemMap": materialItemMap,
					"num":             num,
				}).Warn("synthesis:防爆符不足,无法合成")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//若成功，合成物放入背包
	luckManager := pl.GetPlayerDataManager(playertypes.PlayerLuckyDataManagerType).(*playerlucky.PlayerLuckyDataManager)
	luckuRate := luckManager.GetIncrSuccessRate(itemtypes.ItemTypeLuckyRate, synthesisTemplate.GetSynthesisType().GetLuckyItemSubType())
	rate := synthesisTemplate.SuccessRate + luckuRate
	isSuccessArr := mathutils.RandomHits(common.MAX_RATE, int(rate), num)
	successItemNum := int32(0)
	for _, isSuccess := range isSuccessArr {
		if isSuccess {
			successItemNum += synthesisTemplate.ItemCount
		}
	}

	//合成消耗银两元宝
	silverReason := commonlog.SilverLogReasonSynthesis
	goldReason := commonlog.GoldLogReasonSynthesis
	silverReasonText := fmt.Sprintf(silverReason.String(), synthesisId, num)
	goldReasonText := fmt.Sprintf(goldReason.String(), synthesisId, num)
	flag := propertyManager.Cost(int64(needBindGoldNum), int64(needGold), goldReason, goldReasonText, needSilver, silverReason, silverReasonText)
	if !flag {
		panic(fmt.Errorf("synthesis: synthesisStart use silver or gold should be ok"))
	}

	if isUseAvoidBomb {
		materialItemMap = map[int32]int32{}
		if successItemNum > 0 {
			for itemNeedId, itemNeedNum := range synthesisTemplate.GetSynthesisMap() {
				materialItemMap[itemNeedId] = itemNeedNum * successItemNum
			}
		}
		materialItemMap[synthesisTemplate.ExplosionId] = synthesisTemplate.ExplosionCount * num
	}

	//合成消耗物品
	reasonUseItem := commonlog.InventoryLogReasonSynthesisStart
	reasonUseItemText := fmt.Sprintf(reasonUseItem.String(), synthesisId, num)
	if flag := inventoryManager.BatchRemove(materialItemMap, reasonUseItem, reasonUseItemText); !flag {
		panic(fmt.Errorf("synthesis: synthesisStart use item should be ok"))
	}

	if successItemNum > 0 {
		itemGetReason := commonlog.InventoryLogReasonSynthesisReceive
		itemData := droptemplate.CreateItemData(rewItemId, successItemNum, defaultLevel, bindType)
		if flag := inventoryManager.AddItemLevel(itemData, itemGetReason, itemGetReason.String()); !flag {
			panic(fmt.Errorf("synthesis: synthesisStart add item should be ok"))
		}
	}

	//同步背包、资源
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//发送成功事件
	eventData := synthesiseventtypes.CreateSynthesisFinishEventData(synthesisTemplate.GetSynthesisType(), num)
	gameevent.Emit(synthesiseventtypes.EventTypeSynthesisFinish, pl, eventData)

	pl.SendMsg(pbutil.BuildSCSynthesisStart(synthesisId, successItemNum, isSuccessArr))
	return
}
