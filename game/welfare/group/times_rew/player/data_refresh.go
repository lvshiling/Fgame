package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	grouptimesrewtypes "fgame/fgame/game/welfare/group/times_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeTimesRew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(timesRewRefreshInfo))
}

//累计次数奖励-刷新
func timesRewRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
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
	info := obj.GetActivityData().(*grouptimesrewtypes.TimesRewInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("feedbackGoldPig:活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	timesRewEnd(obj)
	return
}

//累抽结束
func timesRewEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*grouptimesrewtypes.TimesRewInfo)
	times := info.Times
	if times == 0 {
		info.IsEmail = true
	}

	if !info.IsEmail {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			return
		}
		temp := groupInterface.GetFirstOpenTemp()
		timesTempList := welfaretemplate.GetWelfareTemplateService().GetTimesRewTemplateByGorup(groupId)
		for _, timesTemp := range timesTempList {
			if timesTemp.VipLevel > pl.GetVip() {
				continue
			}

			//是否领取
			if info.IsCanReceiveRewards(timesTemp.DrawTimes, timesTemp.VipLevel) {
				title := timeTemp.Name
				acName := chatlogic.FormatMailKeyWordNoticeStr(timeTemp.Name)
				econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityTimesRewContent), acName, times, pl.GetVip())
				itemDataList := welfarelogic.ConvertToItemData(timesTemp.GetRewItemMap(), temp.GetExpireType(), temp.GetExpireTime())
				emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, itemDataList)

				info.AddRecord(timesTemp.DrawTimes, timesTemp.VipLevel)
			}
		}
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
