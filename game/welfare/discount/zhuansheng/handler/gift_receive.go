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
	gamesession "fgame/fgame/game/session"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_DISCOUNT_ZHUANSHENG_RECEIVE_GIFT_TYPE), dispatch.HandlerFunc(handlerZhuanShengGiftReceive))
}

//处理领取转生礼包赠品
func handlerZhuanShengGiftReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理领取转生礼包赠品请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityDiscountZhuanShengReceiveGift)
	groupId := csMsg.GetGroupId()
	giftType := csMsg.GetTyp()

	err = receiveZhuanShengGift(tpl, groupId, giftType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理领取转生礼包赠品请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理领取转生礼包赠品请求完成")

	return
}

//领取转生礼包赠品请求逻辑
func receiveZhuanShengGift(pl player.Player, groupId, giftType int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	zhuanShengGroupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(groupId)
	if zhuanShengGroupTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"giftType": giftType,
			}).Warn("welfare:领取充值返利奖励请求，转生模板组不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	discountTemp := zhuanShengGroupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
	if discountTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"giftType": giftType,
			}).Warn("welfare:领取充值返利奖励请求，转生模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取充值返利奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	newItemMap := discountTemp.GetGiftItemMap()
	if len(newItemMap) < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"giftType": giftType,
			}).Warn("welfare:领取充值返利奖励请求，该礼包没有赠品")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取充值返利奖励请求，活动不存在")
		return
	}

	//领取条件
	info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
	if !info.IsCanReceiveGift(giftType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"group":    groupId,
				"giftType": giftType,
			}).Warn("welfare:领取充值返利奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	firsOpenTemp := groupInterface.GetFirstOpenTemp()
	newItemDataList := welfarelogic.ConvertToItemData(newItemMap, firsOpenTemp.GetExpireType(), firsOpenTemp.GetExpireTime())
	//背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取活动奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonZhuanShengGiftReceive
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), groupId, giftType)
	flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReasonText)
	if !flag {
		panic("welfare:zhuansheng gift receive add item should be ok")
	}

	//更新信息
	info.AddGiftRecord(giftType)
	welfareManager.UpdateObj(obj)

	//同步资源
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityDiscountZhuanShengReceiveGift(groupId, giftType, newItemMap)
	pl.SendMsg(scMsg)
	return
}
