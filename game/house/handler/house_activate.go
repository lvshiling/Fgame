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
	housetypes "fgame/fgame/game/house/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_ACTIVATE_TYPE), dispatch.HandlerFunc(handleHouseActivate))
}

//处理房子激活信息
func handleHouseActivate(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子激活消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMgs := msg.(*uipb.CSHouseActivate)
	houseIndex := csMgs.GetHouseIndex()
	typeInt := csMgs.GetHouseType()
	logTime := csMgs.GetLogTime()

	houseType := housetypes.HouseType(typeInt)
	if !houseType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"houseType": houseType,
			}).Warn("house:房子激活失败，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = houseActivate(tpl, houseIndex, houseType, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"error":      err,
			}).Error("house:处理获取房子激活消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"houseIndex": houseIndex,
		}).Debug("house:处理获取房子激活消息完成")
	return nil
}

func houseActivate(pl player.Player, houseIndex int32, houseType housetypes.HouseType, logTime int64) (err error) {
	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)
	// 是否能激活
	if !houseManager.IsCanActivate(houseIndex) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
			}).Warn("house:房子激活失败，激活下一个房子需要上一个房子历史装修等级达到最高")
		playerlogic.SendSystemMessage(pl, lang.HouseActivateFali)
		return
	}

	// 初始激活、默认类型
	if houseIndex == housetypes.InitHouseIndex {
		houseType = housetemplate.GetHouseTemplateService().GetHouseConstantTemplate().GetInitHouseType()
	}

	// 默认等级
	level := int32(housetypes.InitHouseLevel)
	houseTemplate := housetemplate.GetHouseTemplateService().GetHouseTemplate(houseIndex, houseType, level)
	if houseTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": level,
			}).Warn("house:房子激活失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	// 物品数量
	itemsMap := houseTemplate.GetUseItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItems(itemsMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": level,
				"itemsMap":   itemsMap,
			}).Warn("house:房子激活失败，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	if len(itemsMap) > 0 {
		itemUseReason := commonlog.InventoryLogReasonHouseActivateUse
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), houseIndex, houseType, level)
		flag := inventoryManager.BatchRemove(itemsMap, itemUseReason, itemUseReasonText)
		if !flag {
			panic(fmt.Errorf("house: house activate use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag := houseManager.HouseActivate(houseIndex, houseType)
	if !flag {
		panic(fmt.Errorf("house:维修房子应该成功,房子序号：%d", houseIndex))
	}

	logList := house.GetHouseService().GetLogByTime(logTime)
	scMsg := pbutil.BuildSCHouseActivate(int32(houseType), houseIndex, logList)
	pl.SendMsg(scMsg)
	return
}
