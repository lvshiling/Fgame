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
	gamesession "fgame/fgame/game/session"
	discountyunyintypes "fgame/fgame/game/welfare/discount/yunyin/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_YUN_YIN_BUY_TYPE), dispatch.HandlerFunc(handleYunYinShopBuy))
}

func handleYunYinShopBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare: 开始处理购买云隐商店商品消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityYunYinBuy)
	groupId := csMsg.GetGroupId()
	itemTyp := csMsg.GetType()
	num := csMsg.GetNum()

	err = yunYinShopBuy(tpl, groupId, itemTyp, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"groupId":  groupId,
				"err":      err,
			},
		).Error("welfare: 处理购买云隐商店商品消息,错误")
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"err":      err,
		},
	).Debug("welfare: 处理购买云隐商店商品消息,成功")

	return
}

func yunYinShopBuy(pl player.Player, groupId int32, itemTyp int32, num int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeYunYin

	// 检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	info := obj.GetActivityData().(*discountyunyintypes.YunYinInfo)

	groupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(groupId)
	if groupTemp == nil {
		return
	}
	yunYinTemp := groupTemp.GetDiscountZhuanShengTemplateByType(0, 0, itemTyp)
	if yunYinTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"itemType": itemTyp,
			}).Warn("welfare: 购买云隐商店商品请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if num > yunYinTemp.BuyCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("welfare: 购买云隐商店商品请求,超过单次最大购买数量")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	buyNum := info.GetBuyNum(itemTyp) + num
	if buyNum > yunYinTemp.BuyMax {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      buyNum,
			}).Warn("welfare: 购买云隐商店商品请求,该商品不足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityYunYinShopNotEnough)
		return
	}

	// 判断元宝是否足够
	needGold := num * yunYinTemp.UseGold
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if needGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needGold), false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": needGold,
				}).Warn("welfare:购买云隐商店商品请求，当前元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	// 计算购买的商品
	itemMap := make(map[int32]int32)
	for id, count := range yunYinTemp.GetItemMap() {
		_, ok := itemMap[id]
		if !ok {
			itemMap[id] = count * num
		} else {
			itemMap[id] += count * num
		}
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(itemMap) > 0 {
		if !inventoryManager.HasEnoughSlots(itemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": needGold,
				}).Warn("welfare:购买云隐商店商品请求,背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	// 消耗元宝
	goldReason := commonlog.GoldLogReasonYunYinShopBuy
	goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType, itemTyp)
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic("welfare: 购买云隐商店商品消耗元宝应该成功")
	}

	// 添加物品
	inventoryReason := commonlog.InventoryLogReasonYunYinShopBuy
	inventoryReasonText := fmt.Sprintf(inventoryReason.String(), typ, subType, itemTyp)
	flag = inventoryManager.BatchAdd(itemMap, inventoryReason, inventoryReasonText)
	if !flag {
		panic("welfare: 购买云隐商店商品获得物品应该成功")
	}

	//推送变化
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	// 更新数据
	info.AddBuyRecord(itemTyp, num)
	info.AddGoldNum(needGold)
	welfareManager.UpdateObj(obj)

	scOpenActivityYunYunShopBuy := pbutil.BuildSCOpenActivityYunYunShopBuy(groupId, itemTyp, num, itemMap, info.GoldNum)
	pl.SendMsg(scOpenActivityYunYunShopBuy)

	return
}
