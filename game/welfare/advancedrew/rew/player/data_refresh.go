package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedrewrewtypes "fgame/fgame/game/welfare/advancedrew/rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(rewAdvancedRefreshInfo))
}

//进阶奖励-刷新
func rewAdvancedRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
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
	info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发未领取奖励邮件
	advancedRewEnd(obj)
	return
}

//进阶奖励结束
func advancedRewEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)
	lastUpdateTime := obj.GetUpdateTime()

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needAdvancedNum := temp.Value2
		needChargeNum := temp.Value3
		if info.IsCanReceiveRewards(needAdvancedNum, needChargeNum) {
			title := temp.Label
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			moduleName := chatlogic.FormatMailKeyWordNoticeStr(info.RewType.String())
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityAdvancedRewContent), acName, moduleName, needAdvancedNum)
			itemDataList := welfarelogic.ConvertToItemData(temp.GetEmailRewItemMap(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, itemDataList)
			info.AddRecord(needAdvancedNum)
		}
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)

	return
}
