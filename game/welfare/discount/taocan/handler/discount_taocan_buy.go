package handler

/*   超值套餐 旭东要求 屏掉处理
import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantemplate "fgame/fgame/game/huiyuan/template"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	discounttaocantemplate "fgame/fgame/game/welfare/discount/taocan/template"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	investleveltemplate "fgame/fgame/game/welfare/invest/level/template"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_TAOCAN_BUY_TYPE), dispatch.HandlerFunc(handlerDiscountTaoCanBuy))
}

//处理购买超值套餐
func handlerDiscountTaoCanBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理购买超值套餐请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityTaoCanBuy)
	groupId := csMsg.GetGroupId()

	err = buyDiscountTaoCan(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理购买超值套餐请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理购买超值套餐请求完成")

	return
}

//购买超值套餐请求逻辑
func buyDiscountTaoCan(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeTaoCan

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:购买超值套餐请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//元宝是否足够
	groupTemp := groupInterface.(*discounttaocantemplate.GroupTemplateDiscountTaoCan)
	totalNeed := groupTemp.GetNeedGold()
	needDeductGold, addItemDataList, isTotalBuy := getCostGold(pl, groupId)
	if isTotalBuy {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:购买超值套餐请求，无套餐可购买")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCostNotEnoughCondition)
		return
	}

	if needDeductGold > 0 {
		totalNeed = needDeductGold
	}
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(int64(totalNeed), false) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"totalNeed": totalNeed,
			}).Warn("welfare:购买超值套餐请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//判断背包空间
	extralItemDataList := welfarelogic.ConvertToItemData(groupTemp.GetRewItemMap(), groupTemp.GetExpireType(), groupTemp.GetExpireTime())
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(extralItemDataList) > 0 || len(addItemDataList) > 0 {
		addItemDataList = append(addItemDataList, extralItemDataList...)
		if !inventoryManager.HasEnoughSlotsOfItemLevel(addItemDataList) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:背包空间不足，请清理后再购买")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyDiscountUse
	goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
	flag := propertyManager.CostGold(int64(totalNeed), false, goldReason, goldReasonText)
	if !flag {
		panic("welfare: buy taocan discount use gold should be ok")
	}

	// 全部打包额外物品
	if len(extralItemDataList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		itemReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		flag = inventoryManager.BatchAddOfItemLevel(extralItemDataList, itemGetReason, itemReasonText)
		if !flag {
			panic("welfare: buy taocan add item should be ok")
		}
	}

	//激活功能
	gameevent.Emit(welfareeventtypes.EventTypeDiscountBuyTaoCao, pl, groupId)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityTaoCanBuy(groupId)
	pl.SendMsg(scMsg)
	return
}

func getCostGold(pl player.Player, groupId int32) (optionalGold int32, itemList []*droptemplate.DropItemData, isTotalBuy bool) {
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*discounttaocantemplate.GroupTemplateDiscountTaoCan)

	// 会员
	deductGold := int32(0)
	totalGold := int32(0)
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	huiyuanTemp := huiyuantemplate.GetHuiYuanTemplateService().GetHuiYuanTemplate(huiyuantypes.HuiYuanTypePlus)
	if huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus) {
		deductGold += huiyuanTemp.NeedGold
	}
	totalGold += huiyuanTemp.NeedGold

	//等级投资、装备许愿礼包
	investType := welfaretypes.OpenActivityTypeInvest
	investSubType := welfaretypes.OpenActivityInvestSubTypeLevel
	discountType := welfaretypes.OpenActivityTypeDiscount
	discountSubType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
		relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
		if relateTimeTemp.GetOpenType() == investType && relateTimeTemp.GetOpenSubType() == investSubType {
			investLevelType := investleveltypes.InvesetLevelTypeJunior
			relateObj := welfareManager.GetOpenActivityIfNotCreate(investType, investSubType, relateGroupId)
			relateGroupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(relateGroupId)
			if relateGroupInterface == nil {
				continue
			}
			relateGroupTemp := relateGroupInterface.(*investleveltemplate.GroupTemplateInvestLevel)
			info := relateObj.GetActivityData().(*investleveltypes.InvestLevelInfo)
			totalGold += relateGroupTemp.GetInvestLevelNeedGold(investLevelType)
			if info.IsBuy(investLevelType) {
				// 初级投资的价格
				deductGold += relateGroupTemp.GetInvestLevelNeedGold(investLevelType)
			}
		}

		if relateTimeTemp.GetOpenType() == discountType && relateTimeTemp.GetOpenSubType() == discountSubType {
			giftIndex := groupTemp.GetEquipGiftIndex()
			relateGroupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(relateGroupId)
			discountTemp := relateGroupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftIndex)
			relateObj := welfareManager.GetOpenActivityIfNotCreate(discountType, discountSubType, relateGroupId)
			totalGold += discountTemp.UseGold
			info := relateObj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
			if info.IsBuy(giftIndex) {
				// 装备礼包的价格
				deductGold += discountTemp.UseGold
				//获得的物品
				groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(relateGroupId)
				if groupInterface == nil {
					continue
				}
				firsOpenTemp := groupInterface.GetFirstOpenTemp()
				itemList = welfarelogic.ConvertToItemData(discountTemp.GetItemMap(), firsOpenTemp.GetExpireType(), firsOpenTemp.GetExpireTime())
			}
		}
	}

	if totalGold <= deductGold {
		isTotalBuy = true
		return
	}

	if deductGold > 0 {
		optionalGold = int32(math.Ceil(float64(totalGold-deductGold) * float64(groupTemp.GetOptinalDiscountRate()) / float64(common.MAX_RATE)))
	}
	return
}
*/
