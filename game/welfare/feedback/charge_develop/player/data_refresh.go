package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	feedbackchargedeveloptemplate "fgame/fgame/game/welfare/feedback/charge_develop/template"
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackDevelopRefreshInfo))
}

//返利培养-刷新
func feedbackDevelopRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//养鸡跨天
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	if !isSame {
		chargeDevelopCrossDay(obj)
	}

	// 刷新今日数据
	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		// 同步今日消费
		pl := obj.GetPlayer()
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if info.TodayCostNum != propertyManager.GetTodayCostNum() {
			info.TodayCostNum = propertyManager.GetTodayCostNum()
			welfareManager.UpdateObj(obj)
		}

		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {

			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			if info.ActivateChargeNum != chargeManager.GetTodayChargeNum() {
				info.ActivateChargeNum = chargeManager.GetTodayChargeNum()
				welfareManager.UpdateObj(obj)
			}
		}

	}

	//发送未领取奖励邮件
	feedbackChargeDevelopEnd(obj)
	return
}

//充值培养跨天
func chargeDevelopCrossDay(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	groupId := obj.GetGroupId()
	lastUpdateTime := obj.GetUpdateTime()
	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)

	// 未激活/死亡
	if !info.IsActivate || info.IsDead {
		return
	}

	//满足条件，不再喂养
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargedeveloptemplate.GroupTemplateDevelop)
	condition := groupTemp.GetDevelopNeedTotalTimes()
	if info.FeedTimes >= condition {
		return
	}

	isReachFeedCondition := false
	feedCondtion := groupTemp.GetDevelopFeedCondition(info.FeedTimes)
	if info.IsCanReceiveToday(feedCondtion) {
		rewTemp := groupTemp.GetDevelopDayRewTemp(info.FeedTimes)
		title := groupTemp.GetActivityName()
		acName := chatlogic.FormatMailKeyWordNoticeStr(groupTemp.GetActivityName())
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackDevelopCycleDayMailContent), acName)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(rewTemp.GetEmailRewItemDataList(), rewTemp.GetExpireType(), rewTemp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)

		isReachFeedCondition = true
		info.FeedTimes += 1
	}

	if !info.IsFeed && !isReachFeedCondition {
		info.IsDead = true
		info.ActivateChargeNum = 0
	}

	info.IsFeed = false
	info.TodayCostNum = 0
	welfareManager.UpdateObj(obj)
	return
}

//金鸡培养结束
func feedbackChargeDevelopEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
	groupId := obj.GetGroupId()
	now := global.GetGame().GetTimeService().Now()

	// 活动结束
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("feedbackGoldPig:活动结束发放邮件,活动未结束")
		return
	}

	if info.IsEndMail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	lastUpdateTime := obj.GetUpdateTime()

	// 未激活/死亡
	if !info.IsActivate || info.IsDead {
		return
	}

	//不满足条件
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargedeveloptemplate.GroupTemplateDevelop)
	needTimes := groupTemp.GetDevelopNeedTotalTimes()
	if !info.IsCanReceiveCountDay(needTimes) {
		return
	}

	rewTemp := groupTemp.GetDevelopEndRewTemp()
	title := groupTemp.GetActivityName()
	acName := chatlogic.FormatMailKeyWordNoticeStr(groupTemp.GetActivityName())
	econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackDevelopCycleDayMailContent), acName)
	newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(rewTemp.GetEmailRewItemDataList(), rewTemp.GetExpireType(), rewTemp.GetExpireTime())
	emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)

	info.IsEndMail = true
	info.IsReceiveRew = true
	welfareManager.UpdateObj(obj)
	return
}
