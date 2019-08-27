package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	huntlogic "fgame/fgame/game/hunt/logic"
	"fgame/fgame/game/hunt/pbutil"
	playerhunt "fgame/fgame/game/hunt/player"
	hunttemplate "fgame/fgame/game/hunt/template"
	hunttypes "fgame/fgame/game/hunt/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_HUNT_XUNBAO_TYPE), dispatch.HandlerFunc(handleHuntXunBao))
}

//处理寻宝
func handleHuntXunBao(s session.Session, msg interface{}) (err error) {
	log.Debug("hunt:处理寻宝")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSHuntXunBao)
	huntInt := csMsg.GetHuntType()
	attendTimes := csMsg.GetXunBaoTimes()

	huntType := hunttypes.HuntType(huntInt)
	if !huntType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"huntType": huntType,
			}).Warn("inventory:处理寻宝,寻宝类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = huntXunBao(tpl, huntType, attendTimes)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"huntType": huntType,
				"error":    err,
			}).Error("hunt:处理寻宝,错误")
		return err
	}
	log.Debug("hunt:处理寻宝,完成")
	return nil
}

//寻宝
func huntXunBao(pl player.Player, huntType hunttypes.HuntType, attendTimes int32) (err error) {
	huntTemp := hunttemplate.GetHuntTemplateService().GetHuntTemplat(huntType)
	if huntTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"huntType":    huntType,
				"attendTimes": attendTimes,
			}).Warn("hunt:寻宝失败,寻宝模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	huntManager := pl.GetPlayerDataManager(playertypes.PlayerHuntDataManagerType).(*playerhunt.PlayerHuntDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	huntInfo := huntManager.GetHuntInfo(huntType)

	// 免费次数
	isCost := true
	if attendTimes == 1 && huntManager.IsFreeTimes(huntType) {
		isCost = false
	}

	needItemId := huntTemp.UseItemId
	needItemCount := huntTemp.UseItemCount * attendTimes
	needGold := int64(huntTemp.GoldUse * attendTimes)
	needBindGold := int64(huntTemp.BindGoldUse * attendTimes)

	if isCost {
		// 物品是否足够
		if needItemCount > 0 && !inventoryManager.HasEnoughItem(needItemId, needItemCount) {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"huntType":      huntType,
					"attendTimes":   attendTimes,
					"needItemCount": needItemCount,
				}).Warn("hunt:寻宝失败,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//元宝
		if !propertyManager.HasEnoughGold(needGold, false) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"huntType":    huntType,
					"attendTimes": attendTimes,
					"needGold":    needGold,
				}).Warn("hunt:寻宝失败,元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
		//是否足够绑元
		if !propertyManager.HasEnoughGold(needBindGold, true) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"huntType":     huntType,
					"attendTimes":  attendTimes,
					"needBindGold": needBindGold,
				}).Warn("hunt:寻宝失败,绑元不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	curAttendTimes := huntInfo.GetTotalHuntCount()
	rewList := huntlogic.CountHuntDropItemList(huntType, curAttendTimes, attendTimes)

	var itemDataList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(rewList) > 0 {
		itemDataList, resMap = droplogic.SeperateItemDatas(rewList)
	}

	// 背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"huntType":    huntType,
				"attendTimes": attendTimes,
			}).Warn("hunt:寻宝失败,背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗物品
	if isCost && needItemCount > 0 {
		itemUseReason := commonlog.InventoryLogReasonHuntUse
		reasonText := fmt.Sprintf(itemUseReason.String(), huntType)
		flag := inventoryManager.UseItem(needItemId, needItemCount, itemUseReason, reasonText)
		if !flag {
			panic(fmt.Errorf("hunt:add item should be success"))
		}
	}

	//消耗元宝
	goldUseReason := commonlog.GoldLogReasonHuntCost
	goldUseReasonText := fmt.Sprintf(goldUseReason.String(), huntType)
	if isCost && needGold > 0 {
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("chess:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if isCost && needBindGold > 0 {
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("chess:消耗元宝应该成功")
		}
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonChessGet
		silverReason := commonlog.SilverLogReasonChessGet
		levelReason := commonlog.LevelLogReasonChessGet
		goldReasonText := fmt.Sprintf(goldReason.String(), huntType)
		silverReasonText := fmt.Sprintf(silverReason.String(), huntType)
		levelReasonText := fmt.Sprintf(levelReason.String(), huntType)

		err = droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
	}

	//添加物品
	if len(itemDataList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonHuntRew
		reasonText := fmt.Sprintf(itemGetReason.String(), huntType)
		flag := inventoryManager.BatchAddOfItemLevel(itemDataList, itemGetReason, reasonText)
		if !flag {
			panic("hunt:增加物品应该成功")
		}
	}

	if !isCost {
		huntManager.UpdateFreeHunt(huntType)
	}

	huntManager.UpdateHuntCount(huntType, attendTimes)

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCHuntXunBao(rewList, int32(huntType), huntInfo.GetLastHuntTime(), huntInfo.GetFreeHuntCount())
	pl.SendMsg(scMsg)
	return
}
