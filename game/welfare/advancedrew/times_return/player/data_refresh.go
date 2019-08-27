package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedrewtimesreturntemplate "fgame/fgame/game/welfare/advancedrew/times_return/template"
	advancedrewtimesreturntypes "fgame/fgame/game/welfare/advancedrew/times_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, playerwelfare.ActivityObjInfoRefreshHandlerFunc(timesReturnRefreshInfo))
}

//升阶次数返还
func timesReturnRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
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
	info := obj.GetActivityData().(*advancedrewtimesreturntypes.AdvancedTimesReturnInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发未领取奖励邮件
	advancedTimesReturnEnd(obj)
	return
}

//进阶次数返还结束
func advancedTimesReturnEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedrewtimesreturntypes.AdvancedTimesReturnInfo)
	lastUpdateTime := obj.GetUpdateTime()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewtimesreturntemplate.GroupTemplateAdvancedTimesReturn)
	for _, temp := range groupTemp.GetOpenTempMap() {

		needTimes := temp.Value2
		if !info.IsCanReceiveRewards(needTimes) {
			continue
		}

		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		moduleName := chatlogic.FormatMailKeyWordNoticeStr(info.RewType.String())
		timesText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", needTimes))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityAdvancedrewTimesReturnEndMailContent), acName, moduleName, timesText)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)

		info.AddRecord(needTimes)
	}

	//重置
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
