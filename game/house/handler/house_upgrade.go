package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/house/house"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_UPGRADE_TYPE), dispatch.HandlerFunc(handleHouseUplevel))
}

//处理房子升级信息
func handleHouseUplevel(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子升级消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMgs := msg.(*uipb.CSHouseUpgrade)
	houseIndex := csMgs.GetHouseIndex()
	logTime := csMgs.GetLogTime()

	err = houseUplevel(tpl, houseIndex, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"error":      err,
			}).Error("house:处理获取房子升级消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"houseIndex": houseIndex,
		}).Debug("house:处理获取房子升级消息完成")
	return nil

}

//房子升级
func houseUplevel(pl player.Player, houseIndex int32, logTime int64) (err error) {
	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)

	// 是否存在
	houseInfo := houseManager.GetHouse(houseIndex)
	if houseInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
			}).Warn("house:房子升级失败，房子不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	houseType := houseInfo.GetHouseType()
	houseLevel := houseInfo.GetHouseLevel()
	nextLevel := houseLevel + 1

	// 是否满级
	nextHouseTemplate := housetemplate.GetHouseTemplateService().GetHouseTemplate(houseIndex, houseType, nextLevel)
	if nextHouseTemplate == nil {
		// 控制台日志，写全引起错误的信息
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"nextLevel":  nextLevel,
			}).Warn("house:房子升级失败，已经满级")
		playerlogic.SendSystemMessage(pl, lang.HouseFullLevel)
		return
	}

	// 每日次数
	if !houseManager.IsEnoughUplevelTimes(houseIndex) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
			}).Warn("house:房子升级失败，每日升级次数不足")
		playerlogic.SendSystemMessage(pl, lang.HouseNotEnoughTimes)
		return
	}

	// 物品数量
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemsMap := nextHouseTemplate.GetUseItemMap()
	if !inventoryManager.HasEnoughItems(itemsMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
				"itemsMap":   itemsMap,
			}).Warn("house:房子升级失败，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//背包空间
	rewItemsMap := nextHouseTemplate.GetRewardsItemMap()
	if !inventoryManager.HasEnoughSlots(rewItemsMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
				"itemsMap":   itemsMap,
			}).Warn("house:房子升级失败，背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	if len(itemsMap) > 0 {
		itemUseReason := commonlog.InventoryLogReasonHouseUplevelUse
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), houseIndex, houseType, houseLevel) //日志参数，写清楚变化的房子信息
		flag := inventoryManager.BatchRemove(itemsMap, itemUseReason, itemUseReasonText)
		if !flag {
			panic(fmt.Errorf("house: house activate use item should be ok"))
		}
	}

	if len(rewItemsMap) > 0 {
		itemGetReason := commonlog.InventoryLogReasonHouseUplevelGet
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), houseIndex, houseType, houseLevel) //日志参数，写清楚变化的房子信息
		flag := inventoryManager.BatchAdd(rewItemsMap, itemGetReason, itemGetReasonText)
		if !flag {
			panic(fmt.Errorf("house: house upgrade get item should be ok"))
		}
	}

	//物品变化，同步背包数据
	inventorylogic.SnapInventoryChanged(pl)

	//房子升级
	flag := houseManager.HouseUplevel(houseIndex)
	if !flag {
		panic(fmt.Errorf("house:升级应该成功,房子序号：%d", houseIndex))
	}

	logList := house.GetHouseService().GetLogByTime(logTime)
	scMsg := pbutil.BuildSCHouseUpgrade(houseIndex, rewItemsMap, logList)
	pl.SendMsg(scMsg)
	return
}
