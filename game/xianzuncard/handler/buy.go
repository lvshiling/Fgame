package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"
	"fgame/fgame/game/xianzuncard/pbutil"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIAN_ZUN_CARD_BUY_TYPE), dispatch.HandlerFunc(handleXianZunCardBuy))
}

func handleXianZunCardBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("xianzuncard: 开始处理购买仙尊特权卡请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXianZunCardBuy := msg.(*uipb.CSXianZunCardBuy)
	typ := xianzuncardtypes.XianZunCardType(csXianZunCardBuy.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"typ":      typ.String(),
			}).Warn("xianzuncard: 仙尊卡类型不符")
		return
	}

	err = xianZunCardBuy(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"typ":      typ.String(),
				"err":      err,
			}).Error("xianzuncard: 处理购买仙尊特权卡请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("xianzuncard: 处理购买仙尊特权卡请求消息,成功")

	return
}

func xianZunCardBuy(pl player.Player, typ xianzuncardtypes.XianZunCardType) (err error) {
	playerId := pl.GetId()
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	// 判断是否已经激活
	if xianZunManager.IsActivite(typ) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("xianzuncard: 仙尊卡已经激活")
		playerlogic.SendSystemMessage(pl, lang.XianZunCardAlreadyActivite)
		return
	}

	xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(typ)
	if xianZunTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("xianzuncard: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.XianZunCardTemplateNotExist)
		return
	}

	// 判断元宝是否足够
	needGold := xianZunTemp.NeedGold
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if needGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("xianzuncard: 元宝不足,不能购买")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	// 判断背包空间是否足够
	receiveItem := xianZunTemp.GetReceiveItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(receiveItem) != 0 {
		flag := inventoryManager.HasEnoughSlots(receiveItem)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ.String(),
				}).Warn("xianzuncard: 背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	// 激活仙尊特权卡
	goldUseReason := commonlog.GoldLogReasonXianZunCardBuy
	goldReason := fmt.Sprintf(goldUseReason.String(), typ.String())
	flag := propertyManager.CostGold(int64(needGold), false, goldUseReason, goldReason)
	if !flag {
		panic(fmt.Errorf("xianzuncard: 购买仙尊特权卡消耗元宝应该成功"))
	}

	// 添加奖励物品
	reason := commonlog.InventoryLogReasonXianZunCardActiviteAdd
	reasonText := fmt.Sprintf(reason.String(), typ.String())
	flag = inventoryManager.BatchAdd(receiveItem, reason, reasonText)
	if !flag {
		panic("xianzuncard: 购买仙尊特权卡添加物品应该成功")
	}

	// 添加奖励钱
	silverGetReason := commonlog.SilverLogReasonXianZunCardActiviteRew
	silverGetReasonText := fmt.Sprintf(silverGetReason.String(), typ.String())
	goldGetReason := commonlog.GoldLogReasonXianZunActiviteAdd
	goldGetReasonText := fmt.Sprintf(goldGetReason.String(), typ.String())
	flag = propertyManager.AddMoney(int64(xianZunTemp.JiHuoRewBindGold), int64(xianZunTemp.JiHuoRewGold), goldGetReason, goldGetReasonText, int64(xianZunTemp.JiHuoRewSilver), silverGetReason, silverGetReasonText)
	if !flag {
		panic("xianzuncard: 购买仙尊特权卡添加钱应该成功")
	}

	// 物品改变推送
	inventorylogic.SnapInventoryChanged(pl)

	// 刷新数据
	flag = xianZunManager.BuySuccess(typ)
	if !flag {
		panic("购买仙尊特权卡,应该成功")
	}

	// 刷新属性
	xianzuncardlogic.PropertyChanged(pl)
	//同步钱
	propertylogic.SnapChangedProperty(pl)

	rd := propertytypes.CreateRewData(0, 0, xianZunTemp.JiHuoRewSilver, xianZunTemp.JiHuoRewGold, xianZunTemp.JiHuoRewBindGold)
	scXianZunCardBuy := pbutil.BuildSCXianZunCardBuy(int32(typ), rd, receiveItem)
	pl.SendMsg(scXianZunCardBuy)

	return
}
