package logic

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardschargelimittypes "fgame/fgame/game/welfare/rewards/charge_limit/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"
	"fmt"
)

func CheckRewardsMail(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*rewardschargelimittypes.ChargeRewLimitInfo)
	groupId := obj.GetGroupId()
	pl := obj.GetPlayer()
	now := global.GetGame().GetTimeService().Now()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	openTempMap := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, openTemp := range openTempMap {
		mailTitle := openTemp.Label
		convertRate := openTemp.Value1
		globalTimesMax := openTemp.Value2
		playerTimesMax := openTemp.Value3

		//个人次数限制
		if !info.IsHadReceiveTimes(convertRate, playerTimesMax, 1) {
			// //无奖励提醒邮件
			// acName := coreutils.FormatNoticeStr(openTemp.Label)
			// econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityRewardsPlayerTimesNotEnoughContent), acName, convertRate)
			// emaillogic.AddEmailItemLevel(pl, mailTitle, econtent, now, nil)
			continue
		}

		addTimes := info.CountLeftTimes(convertRate, playerTimesMax)
		if addTimes <= 0 {
			continue
		}

		// 总次数限制
		if !welfare.GetWelfareService().IsHadReceiveTimes(groupId, convertRate, globalTimesMax, 1) {
			//无奖励提醒邮件
			acName := coreutils.FormatNoticeStr(openTemp.Label)
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityRewardsHadNoneContent), acName, convertRate)
			emaillogic.AddEmailItemLevel(pl, mailTitle, econtent, now, nil)
			continue
		}

		hadReceiveTimes := welfare.GetWelfareService().GetLeftReceiveTimes(groupId, convertRate)
		leftTimes := globalTimesMax - hadReceiveTimes
		if addTimes > leftTimes {
			addTimes = leftTimes
		}

		if addTimes <= 0 {
			continue
		}

		// 奖励邮件
		newRewItemDataList := openTemp.GetEmailRewItemDataListWithRatio(addTimes)
		endTime := global.GetGame().GetTimeService().Now()
		acName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
		goldNumText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), info.GoldNum))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityChargeRewardsContent), acName, goldNumText, convertRate)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(newRewItemDataList, openTemp.GetExpireType(), openTemp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, mailTitle, econtent, endTime, newItemDataList)

		welfare.GetWelfareService().AddReceiveTimes(groupId, convertRate, addTimes)
		info.ReceiveRewards(convertRate, addTimes)
	}

	welfareManager.UpdateObj(obj)
}
