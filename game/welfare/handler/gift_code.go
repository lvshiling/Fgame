package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/coupon/coupon"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_GIFT_CODE_TYPE), dispatch.HandlerFunc(handlerGiftCode))
}

//处理领取礼包
func handlerGiftCode(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理领取礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityGiftCode := msg.(*uipb.CSOpenActivityGiftCode)
	code := csOpenActivityGiftCode.GetCode()

	err = giftCode(tpl, code)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理领取礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理领取礼包请求完成")

	return
}

//领取礼包逻辑
func giftCode(pl player.Player, code string) (err error) {
	coupon.GetCouponService().Exchange(pl, code)
	// inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	// codeInt, err := strconv.ParseInt(code, 10, 64)
	// if err != nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"code":     code,
	// 		}).Warn("welfare: 礼包兑换，激活码错误")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	// 	return nil
	// }
	// giftTemp := welfaretemplate.GetWelfareTemplateService().GetCodeGift(int32(codeInt))
	// if giftTemp == nil {
	// 	log.WithFieEventTypeCrossDaylds(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"code":     code,
	// 		}).Warn("welfare: 礼包兑换，激活码不存在")
	// 	playerlogic.SendSystemMessage(pl, lang.OpenActivityWelfareGiftCode)
	// 	return
	// }

	// rewItemMap := giftTemp.GetItemMap()
	// newItemDataList := welfarelogic.ConvertToItemData(rewItemMap)
	// //背包空间
	// if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 		}).Warn("welfare:礼包兑换请求，背包空间不足")
	// 	playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
	// 	return
	// }

	// //增加物品
	// itemGetReason := commonlog.InventoryLogReasonFirstCharge
	// flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReason.String())
	// if !flag {
	// 	panic("welfare:first charge add item should be ok")
	// }

	// reasonGold := commonlog.GoldLogReasonGiftCode
	// reasonSilver := commonlog.SilverLogReasonGiftCode
	// reasonLevel := commonlog.LevelLogReasonGiftCode
	// rewSilver := int32(0)
	// rewBindGold := int32(0)
	// rewGold := int32(0)
	// rewExp := int32(0)
	// rewExpPoint := int32(0)
	// totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	// flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGold.String(), reasonSilver, reasonSilver.String(), reasonLevel, reasonLevel.String())
	// if !flag {
	// 	panic("welfare:gift add RewData should be ok")
	// }

	// //同步资源
	// propertylogic.SnapChangedProperty(pl)
	// inventorylogic.SnapInventoryChanged(pl)

	// scOpenActivityGiftCode := pbutil.BuildSCOpenActivityGiftCode(totalRewData, rewItemMap)
	// pl.SendMsg(scOpenActivityGiftCode)
	return
}
