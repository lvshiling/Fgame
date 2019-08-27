package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/house/pbutil"
	playerhouse "fgame/fgame/game/house/player"
	housetemplate "fgame/fgame/game/house/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_REPAIR_TYPE), dispatch.HandlerFunc(handleHouseRepair))
}

//处理房子维修
func handleHouseRepair(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子维修消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMgs := msg.(*uipb.CSHouseRepair)
	houseIndex := csMgs.GetHouseIndex()

	err = houseRepair(tpl, houseIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"error":      err,
			}).Error("house:处理获取房子维修消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"houseIndex": houseIndex,
		}).Debug("house:处理获取房子维修消息完成")
	return nil
}

func houseRepair(pl player.Player, houseIndex int32) (err error) {
	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)

	// 房子是否存在
	houseInfo := houseManager.GetHouse(houseIndex)
	if houseInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
			}).Warn("house:房子维修失败，房子不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	houseType := houseInfo.GetHouseType()
	houseLevel := houseInfo.GetHouseLevel()
	houseTemplate := housetemplate.GetHouseTemplateService().GetHouseTemplate(houseIndex, houseType, houseLevel)
	if houseTemplate == nil {
		// 控制台日志，写全引起错误的信息
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
			}).Warn("house:房子维修失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if !houseInfo.IsBroken() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
			}).Warn("house:房子维修失败，未损坏")
		playerlogic.SendSystemMessage(pl, lang.HouseHadBroken)
		return
	}

	// 物品数量
	itemsMap := houseTemplate.GetRepairItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItems(itemsMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
				"itemsMap":   itemsMap,
			}).Warn("house:房子激活失败，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	if len(itemsMap) > 0 {
		//物品变化日志
		itemUseReason := commonlog.InventoryLogReasonHouseRepairUse
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), houseIndex, houseType, houseLevel)
		// 批量使用物品
		flag := inventoryManager.BatchRemove(itemsMap, itemUseReason, itemUseReasonText)
		if !flag {
			panic(fmt.Errorf("house: house activate use item should be ok"))
		}

		// 同步物品信息到客户端
		inventorylogic.SnapInventoryChanged(pl)
	}

	// 房子维修
	flag := houseManager.HouseRepair(houseIndex)
	if !flag {
		panic(fmt.Errorf("house:维修房子应该成功,房子序号：%d", houseIndex))
	}

	scMsg := pbutil.BuildSCHouseRepair(houseIndex)
	pl.SendMsg(scMsg)
	return
}
