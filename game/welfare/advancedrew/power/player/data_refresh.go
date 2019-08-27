package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedrewtemplate "fgame/fgame/game/welfare/advancedrew/power/template"
	advancedrewtypes "fgame/fgame/game/welfare/advancedrew/power/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypePower, playerwelfare.ActivityObjInfoRefreshHandlerFunc(advancedRewPowerRefreshInfo))
}

func advancedRewPowerRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	info := obj.GetActivityData().(*advancedrewtypes.AdvancedPowerInfo)
	endTime := info.ExpireTime
	if endTime <= 0 {
		return
	}

	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,活动未结束")
		return
	}

	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("已经发过邮件")
		return
	}

	sendPowerChangedRewardUseEmail(obj)
	return
}

func sendPowerChangedRewardUseEmail(obj *playerwelfare.PlayerOpenActivityObject) {
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	groupTempI := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
	if groupTempI == nil {
		return
	}

	info := obj.GetActivityData().(*advancedrewtypes.AdvancedPowerInfo)
	groupTemp := groupTempI.(*advancedrewtemplate.GroupTemplatePower)
	openTemplateList := groupTemp.GetOpenTemplateListAboutReward(info.Power, info.RewRecord)

	for _, openTemplate := range openTemplateList {
		power := int64(openTemplate.Value2)
		title := openTemplate.Label
		activityText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%s", title))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailAdvancedPowerRewardsContent), activityText, power)

		itemList := openTemplate.GetEmailRewItemDataList()
		itemData := welfarelogic.ConvertToItemDataWithWelfareData(itemList, openTemplate.GetExpireType(), openTemplate.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, content, obj.GetUpdateTime(), itemData)

		info.AddRecord(power)
	}

	info.IsEmail = true
	welfareManager.UpdateObj(obj)
}
