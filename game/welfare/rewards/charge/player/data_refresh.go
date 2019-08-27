package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardschargetemplate "fgame/fgame/game/welfare/rewards/charge/template"
	rewardschargetypes "fgame/fgame/game/welfare/rewards/charge/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCharge, playerwelfare.ActivityObjInfoRefreshHandlerFunc(rewardsChargeRefreshInfo))
}

//每充值奖励-刷新
func rewardsChargeRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	groupId := obj.GetGroupId()
	// 同步今日充值
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			pl := obj.GetPlayer()
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			info := obj.GetActivityData().(*rewardschargetypes.ChargeRewInfo)
			//不相等才同步
			if info.GoldNum != int32(chargeManager.GetTodayChargeNum()) {
				info.GoldNum = int32(chargeManager.GetTodayChargeNum())
				groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
				if groupInterface == nil {
					return
				}
				groupTemp := groupInterface.(*rewardschargetemplate.GroupTemplateRewardsCharge)
				convertRate := groupTemp.GetChargeRewardsConvertRate()
				addTimes := info.CountLeftTimes(convertRate)
				welfareManager.UpdateObj(obj)
				// 发邮件
				if addTimes > 0 {
					openTemp := groupTemp.GetFirstOpenTemp()
					if openTemp == nil {
						return
					}

					newRewItemDataList := openTemp.GetEmailRewItemDataListWithRatio(addTimes)
					endTime := global.GetGame().GetTimeService().Now()
					title := openTemp.Label
					acName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
					goldNumText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), info.GoldNum))
					econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityChargeRewardsContent), acName, goldNumText, convertRate)
					newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(newRewItemDataList, openTemp.GetExpireType(), openTemp.GetExpireTime())
					emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
				}
			}
		}
	}

	return
}
