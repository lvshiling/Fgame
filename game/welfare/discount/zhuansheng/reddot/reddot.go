package reddot

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/reddot/reddot"
	discountzhuanshengtemplate "fgame/fgame/game/welfare/discount/zhuansheng/template"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeZhuanSheng, reddot.HandlerFunc(handleRedDotDiscountZhuanSheng))
}

//限时礼包-转生大礼包红点
func handleRedDotDiscountZhuanSheng(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemplate := groupInterface.(*discountzhuanshengtemplate.GroupTemplateZhaunSheng)

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)

	buyRecord := info.BuyRecord
	_, leftPoint := groupTemplate.GetTotalAndRemainPoint(info.ChargeNum, info.UsePoint)

	curGold := propertyManager.GetGold()

	// 礼包配置
	zhuanShnegGroupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(groupId)
	if zhuanShnegGroupTemp == nil {
		return
	}
	discountZhuanShengTempList := zhuanShnegGroupTemp.GetDiscountZhuanShengTemplate(pl.GetRole(), pl.GetSex())
	for _, temp := range discountZhuanShengTempList {
		if len(temp.GetGiftItemMap()) != 0 && info.IsCanReceiveGift(temp.Type) {
			isNotice = true
			return
		}

		if len(temp.GetUseItemMap()) > 0 || temp.UsePoint > 0 {
			//可物品兑换和积分兑换需要显示红点
			// 购买次数
			curCnt, ok := buyRecord[temp.Type]
			if ok && curCnt >= temp.BuyMax {
				continue
			}

			// 所需元宝
			if curGold < int64(temp.UseGold) {
				continue
			}

			// 所需积分
			if leftPoint < temp.UsePoint {
				continue
			}

			// 所需物品
			if len(temp.GetUseItemMap()) > 0 && !inventoryManager.HasEnoughItems(temp.GetUseItemMap()) {
				continue
			}

			isNotice = true
			return
		}

	}
	return
}
