package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	goaltemplate "fgame/fgame/game/welfare/huhu/goal/template"
	goaltypes "fgame/fgame/game/welfare/huhu/goal/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, playerwelfare.ActivityObjInfoRefreshHandlerFunc(goalRefreshInfo))
}

//累计活动目标-刷新
func goalRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	initGoalQuest(obj)
	err = goalRefreshRew(obj)
	return
}

func goalRefreshRew(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	now := global.GetGame().GetTimeService().Now()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*goaltypes.GoalInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("feedbackGoldPig:活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	if info.GoalCount == 0 {
		info.IsEmail = true
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*goaltemplate.GroupTemplateGoal)
	for _, temp := range groupTemp.GetCanRewTemplate(info.GoalCount, info.RewRecordMap) {
		rewGoalCount := temp.Value2
		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityGoalMailContent), acName)
		itemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, content, endTime, itemDataList)

		info.AddRecord(rewGoalCount)
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}

func initGoalQuest(obj *playerwelfare.PlayerOpenActivityObject) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*goaltemplate.GroupTemplateGoal)
	//初始化运营目标任务
	if welfarelogic.IsOnActivityTime(groupId) {
		questlogic.InitYunYingGoalQuest(pl, groupTemp.GetGolaId())
	}
}
