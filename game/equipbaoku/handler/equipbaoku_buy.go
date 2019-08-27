package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commontypes "fgame/fgame/game/common/types"
	equipbaokueventtypes "fgame/fgame/game/equipbaoku/event/types"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutemplate "fgame/fgame/game/equipbaoku/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_POINTS_EXCHANGE_TYPE), dispatch.HandlerFunc(handleEquipBaoKuShopBuy))
}

//处理宝库兑换商店购买道具
func handleEquipBaoKuShopBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("baokujifen:处理宝库兑换商店购买道具")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csEquipbaokuPointsExchange := msg.(*uipb.CSEquipbaokuPointsExchange)
	shopId := csEquipbaokuPointsExchange.GetShopId()
	num := csEquipbaokuPointsExchange.GetNum()
	typ := equipbaokutypes.BaoKuType(csEquipbaokuPointsExchange.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Warn("equipbaoku:处理探索宝库,宝库类型不合法")
		return
	}

	err = equipBaoKuShopBuy(tpl, shopId, num, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baokujifen:处理宝库兑换商店购买道具,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baokujifen:处理宝库兑换商店购买道具完成")
	return nil

}

//宝库兑换商店购买道具的逻辑
func equipBaoKuShopBuy(pl player.Player, shopId int32, num int32, typ equipbaokutypes.BaoKuType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeEquipBaoKu) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baokujifen:宝库积分兑换错误,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	//zrc: 临时处理,前端只能使用1
	num = 1
	shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
	if shopTemplate == nil || num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断级别
	if shopTemplate.LevelMin > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:宝库积分兑换错误,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断级别
	if shopTemplate.LevelMax < pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:宝库积分兑换错误,等级太高")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooHigh)
		return
	}

	//判断转数
	if shopTemplate.ZhuanshuMin > pl.GetZhuanSheng() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:宝库积分兑换错误,转数不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	//判断转数
	if shopTemplate.ZhuanshuMax < pl.GetZhuanSheng() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:宝库积分兑换错误,转数太高")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooHigh)
		return
	}

	if shopTemplate.MaxCount != 0 && num > shopTemplate.MaxCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:兑换数量大于最大兑换数量")
		playerlogic.SendSystemMessage(pl, lang.EquipBaoKuShopBuyNumInvalid)
		return
	}

	//兑换总数
	totalNum := int32(shopTemplate.BuyCount * num)

	manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	isLimitBuy, leftNum := manager.LeftDayCount(shopId)
	if isLimitBuy && leftNum < num {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:兑换次数，已达每日兑换数量")
		playerlogic.SendSystemMessage(pl, lang.EquipBaoKuShopBuyReacheLimit)
		return
	}

	//积分是否足够
	needJiFen := int32(shopTemplate.UseJiFen * num)
	equipBaoKuObj := manager.GetEquipBaoKuObj(typ)
	attendPoints := equipBaoKuObj.GetAttendPoints()
	if needJiFen > attendPoints {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"shopId":    shopId,
				"num":       num,
				"needJiFen": needJiFen,
			}).Warn("baokujifen:积分不足，无法完成兑换")
		playerlogic.SendSystemMessage(pl, lang.EquipBaoKuShopBuyJiFenNotEnough)
		return
	}

	itemId := shopTemplate.GetItemIdByRoleAndSex(pl.GetRole(), pl.GetSex())
	if itemId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"role":     int32(pl.GetRole()),
				"sex":      int32(pl.GetSex()),
			}).Warn("baokujifen:配置参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//判断背包空间
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag := inventoryManager.HasEnoughSlot(itemId, totalNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("baokujifen:背包空间不足，请清理后再购买")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗
	flag = manager.SubEquipBaoKuAttendPoints(needJiFen, typ)
	if !flag {
		panic("baokujifen: costjifem should be ok")
	}
	//宝库积分日志
	jiFenReason := commonlog.EquipBaoKuLogReasonAttendPointsChange
	jiFenReasonText := fmt.Sprintf(jiFenReason.String(),typ.GetBaoKuName(), commontypes.ChangeTypeExchange.String())
	data := equipbaokueventtypes.CreatePlayerEquipBaoKuAttendPointsLogEventData(attendPoints, equipBaoKuObj.GetAttendPoints(), itemId, totalNum, jiFenReason, jiFenReasonText, typ)
	gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuAttendPointsLog, pl, data)
	//添加物品
	reasonText := commonlog.InventoryLogReasonShopBuy.String()
	flag = inventoryManager.AddItem(itemId, totalNum, commonlog.InventoryLogReasonShopBuy, reasonText)
	if !flag {
		panic(fmt.Errorf("baokujifen: shopBuy add item should be ok"))
	}

	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	//更新当日购买次数
	dayCount := int32(0)
	manager.UpdateEquipBaoKuShopObject(shopId, num)
	if shopTemplate.MaxCount != 0 {
		dayCount = manager.GetEquipBaoKuShopBuyByShopId(shopId).DayCount
	}
	leftAttendPoints := manager.GetEquipBaoKuObj(typ).GetAttendPoints()
	scMsg := pbutil.BuildSCEquipBaoKuPointsExchange(shopId, num, leftAttendPoints, dayCount, int32(typ))
	pl.SendMsg(scMsg)
	return
}
