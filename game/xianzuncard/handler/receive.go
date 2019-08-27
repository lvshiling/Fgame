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
	"fgame/fgame/game/xianzuncard/pbutil"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RECEIVE_XIAN_ZUN_CARD_REWARD_TYPE), dispatch.HandlerFunc(handleXianZunCardReceive))
}

func handleXianZunCardReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("xianzuncard: 开始处理购买仙尊特权卡请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csReceiveXianZunCardReward := msg.(*uipb.CSReceiveXianZunCardReward)
	typ := xianzuncardtypes.XianZunCardType(csReceiveXianZunCardReward.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"typ":      typ.String(),
			}).Warn("xianzuncard: 仙尊卡类型不符")
		return
	}

	err = xianZunCardReceive(tpl, typ)
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

func xianZunCardReceive(pl player.Player, typ xianzuncardtypes.XianZunCardType) (err error) {
	playerId := pl.GetId()
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	// 判断是否已经激活
	if !xianZunManager.IsActivite(typ) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("xianzuncard: 仙尊卡未激活")
		playerlogic.SendSystemMessage(pl, lang.XianZunCardNotActivite)
		return
	}

	// 判断是否已经领取
	if xianZunManager.IsReceive(typ) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("xianzuncard: 仙尊卡每日奖励已经领取")
		playerlogic.SendSystemMessage(pl, lang.XianZunCardAlreadyReceive)
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

	// 判断背包空间是否足够
	rewItem := xianZunTemp.GetDayReceiveItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(rewItem) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItem)
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

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	// 添加奖励钱
	silverGetReason := commonlog.SilverLogReasonXianZunCardReceiveRew
	silverGetReasonText := fmt.Sprintf(silverGetReason.String(), typ.String())
	goldGetReason := commonlog.GoldLogReasonXianZunReceiveAdd
	goldGetReasonText := fmt.Sprintf(goldGetReason.String(), typ.String())
	flag := propertyManager.AddMoney(int64(xianZunTemp.DayRewBindGold), int64(xianZunTemp.DayRewGold), goldGetReason, goldGetReasonText, int64(xianZunTemp.DayRewSilver), silverGetReason, silverGetReasonText)
	if !flag {
		panic("xianzuncard: 领取每日仙尊特权卡奖励添加钱应该成功")
	}

	// 添加奖励物品
	reason := commonlog.InventoryLogReasonXianZunCardReceiveAdd
	reasonText := fmt.Sprintf(reason.String(), typ.String())
	flag = inventoryManager.BatchAdd(rewItem, reason, reasonText)
	if !flag {
		panic("xianzuncard: 领取每日仙尊特权卡奖励添加物品应该成功")
	}

	//同步钱
	propertylogic.SnapChangedProperty(pl)

	// 物品改变推送
	inventorylogic.SnapInventoryChanged(pl)

	// 刷新数据
	flag = xianZunManager.ReceiveSuccess(typ)
	if !flag {
		panic("领取每日仙尊特权卡奖励,刷新数据应该成功")
	}

	rd := propertytypes.CreateRewData(0, 0, xianZunTemp.DayRewSilver, xianZunTemp.DayRewGold, xianZunTemp.DayRewBindGold)
	scReceiveXianZunCardReward := pbutil.BuildSCReceiveXianZunCardReward(int32(typ), rd, rewItem)
	pl.SendMsg(scReceiveXianZunCardReward)

	return
}
