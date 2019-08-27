package listener

/*   超值套餐 旭东要求 屏掉处理
import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	discounttaocantemplate "fgame/fgame/game/welfare/discount/taocan/template"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//购买超值套餐
func discountBuyTaoCan(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	groupId, ok := data.(int32)
	if !ok {
		return
	}

	//至尊会员
	buyHuiYuan(pl)
	//装备礼包
	buyEquipGift(pl, groupId)
	//等级投资
	buyInvestLevel(pl, groupId)
	return
}

func buyHuiYuan(pl player.Player) {
	// 会员
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	if !huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus) {
		huiyuanManager.BuyHuiYuan(huiyuantypes.HuiYuanTypePlus, 0)
	}
}

func buyEquipGift(pl player.Player, groupId int32) {
	// 装备许愿礼包
	discountType := welfaretypes.OpenActivityTypeDiscount
	discountSubType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*discounttaocantemplate.GroupTemplateDiscountTaoCan)
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
		if !welfarelogic.IsOnActivityTime(relateGroupId) {
			continue
		}

		//装备礼包
		relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
		if relateTimeTemp.GetOpenType() == discountType && relateTimeTemp.GetOpenSubType() == discountSubType {
			relateObj := welfareManager.GetOpenActivityIfNotCreate(discountType, discountSubType, relateGroupId)
			giftIndex := groupTemp.GetEquipGiftIndex()
			info := relateObj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
			if info.IsBuy(giftIndex) {
				continue
			}
			welfarelogic.BuyZhuanShengGift(pl, relateObj, map[int32]int32{giftIndex: 1})
		}

	}
}

//等级投资
func buyInvestLevel(pl player.Player, groupId int32) {
	investType := welfaretypes.OpenActivityTypeInvest
	investSubType := welfaretypes.OpenActivityInvestSubTypeLevel
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
		if !welfarelogic.IsOnActivityTime(relateGroupId) {
			continue
		}

		relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
		if relateTimeTemp.GetOpenType() == investType && relateTimeTemp.GetOpenSubType() == investSubType {
			relateObj := welfareManager.GetOpenActivityIfNotCreate(investType, investSubType, relateGroupId)
			investLevelType := investleveltypes.InvesetLevelTypeJunior
			info := relateObj.GetActivityData().(*investleveltypes.InvestLevelInfo)
			if info.IsBuy(investLevelType) {
				continue
			}
			info.InvestBuyInfoMap[investLevelType] = 0
			welfareManager.UpdateObj(relateObj)
		}
	}
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeDiscountBuyTaoCao, event.EventListenerFunc(discountBuyTaoCan))
}
*/
