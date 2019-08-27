package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droptemplate "fgame/fgame/game/drop/template"
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
	"fgame/fgame/game/vip/pbutil"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RECEIVE_FREE_GIFT_TYPE), dispatch.HandlerFunc(handleFreeGiftReceive))
}

//处理领取免费礼包
func handleFreeGiftReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理领取免费礼包消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSReceiveFreeGift)
	giftLevel := csMsg.GetGiftLevel()
	star := csMsg.GetGiftStar()

	err = receiveFreeGift(tpl, giftLevel, star)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("vip:处理领取免费礼包消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("vip:处理领取免费礼包消息完成")
	return nil

}

//领取免费礼包界面逻辑
func receiveFreeGift(pl player.Player, giftLevel, star int32) (err error) {
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	curLevel, _ := vipManager.GetVipLevel()
	vipTemplate := viptemplate.GetVipTemplateService().GetVipTemplate(giftLevel, star)
	if vipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("vip:处理领取免费礼包消息,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 是否领取
	if curLevel < giftLevel {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"curLevel":  curLevel,
				"giftLevel": giftLevel,
			}).Warn("vip:处理领取免费礼包消息,vip等级不足")
		playerlogic.SendSystemMessage(pl, lang.VipLevelToLow)
		return
	}

	if !vipManager.IsCanReceiveFreeGift(giftLevel) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"giftLevel": giftLevel,
			}).Warn("vip:处理领取免费礼包消息, 不能重复领取")
		playerlogic.SendSystemMessage(pl, lang.VipHadBuyGift)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemMap := vipTemplate.GetFreeGiftItemMap()
	itemDataList := droptemplate.ConvertToItemDataList(itemMap, itemtypes.ItemBindTypeBind)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemMap":  itemMap,
			}).Warn("vip:处理领取免费礼包消息, 背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	silverRew := vipTemplate.FreeGiftSilver
	if silverRew > 0 {
		getSilverReason := commonlog.SilverLogReasonVipFreeGiftRew
		propertyManager.AddSilver(silverRew, getSilverReason, getSilverReason.String())
	}

	if len(itemMap) > 0 {
		itemGetReason := commonlog.InventoryLogReasonVipFreeGiftGet
		flag := inventoryManager.BatchAddOfItemLevel(itemDataList, itemGetReason, itemGetReason.String())
		if !flag {
			panic(fmt.Errorf("vip: 批量添加物品应该成功"))
		}
	}

	//更新领取记录
	vipManager.ReceiveFreeGift(giftLevel)

	// 同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	showItemMap := vipTemplate.GetEmailFreeGiftItemMap()
	scMsg := pbutil.BuildSCReceiveFreeGift(giftLevel, star, showItemMap)
	pl.SendMsg(scMsg)
	return
}
