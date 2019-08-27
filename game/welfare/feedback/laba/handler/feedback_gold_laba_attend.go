package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	feedbacklabatypes "fgame/fgame/game/welfare/feedback/laba/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FEEDBACK_GOLD_LABA_ATTEND_TYPE), dispatch.HandlerFunc(handlerFeedbackGoldLabaAttend))
}

//参与元宝拉霸
func handlerFeedbackGoldLabaAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理参与元宝拉霸请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSOpenActivityFeedbackGoldLaBaAttend)
	groupId := csMsg.GetGroupId()
	lastLogTime := csMsg.GetLastLogTime()

	err = attendFeedbackGoldLaba(tpl, groupId, lastLogTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理参与元宝拉霸请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理参与元宝拉霸请求完成")

	return
}

//参与元宝拉霸请求逻辑
func attendFeedbackGoldLaba(pl player.Player, groupId int32, lastLogTime int64) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参与元宝拉霸请求，活动不存在")
		return
	}

	info := obj.GetActivityData().(*feedbacklabatypes.FeedbackGoldLaBaInfo)
	nextTimes := info.Times + 1
	// 拉霸模板
	labaTemp := welfaretemplate.GetWelfareTemplateService().GetGoldLabaTemplate(groupId, nextTimes)
	if labaTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参与元宝拉霸请求，次数已达上限")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityLaBaHadNotTimes)
		return
	}

	//参与条件
	if info.ChargeNum < labaTemp.InvestmentRecharge {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"curCharge":  info.ChargeNum,
				"needCharge": labaTemp.InvestmentRecharge,
			}).Warn("welfare:参与元宝拉霸请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	//元宝消耗
	needGold := labaTemp.Investment
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ChargeNum": info.ChargeNum,
			}).Warn("welfare:参与元宝拉霸请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗
	useGoldReason := commonlog.GoldLogReasonLaBaUse
	useGoldReasonText := fmt.Sprintf(useGoldReason.String(), nextTimes)
	flag := propertyManager.CostGold(int64(needGold), false, useGoldReason, useGoldReasonText)
	if !flag {
		panic("welfare:元宝拉霸消耗元宝应该成功")
	}
	//解决前端右下角元宝飘字显示差值的问题
	propertylogic.SnapChangedProperty(pl)

	// 计算随机元宝
	rewGold := labaTemp.RandomGold()
	if rewGold > 0 {
		getGoldReason := commonlog.GoldLogReasonLaBaGet
		getGoldReasonText := fmt.Sprintf(getGoldReason.String(), nextTimes)
		propertyManager.AddGold(rewGold, false, getGoldReason, getGoldReasonText)
	}

	// 添加日志
	addLogEventData := welfareeventtypes.CreateLaBaAddLogEventData(pl.GetName(), needGold, int32(rewGold))
	gameevent.Emit(welfareeventtypes.EventTypeLaBaAddLog, groupId, addLogEventData)
	// 后台玩家拉霸日志
	labaReason := commonlog.LaBaLogReasonAttend
	reasonText := fmt.Sprintf(labaReason.String(), nextTimes)
	logEventData := welfareeventtypes.CreatePlayerLaBaLogEventData(nextTimes, needGold, int32(rewGold), labaReason, reasonText)
	gameevent.Emit(welfareeventtypes.EventTypeLaBaAttendLog, pl, logEventData)

	//更新信息
	info.Times += 1
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	logList := welfare.GetWelfareService().GetLaBaLogByTime(groupId, lastLogTime)
	scMsg := pbutil.BuildSCOpenActivityFeedbackGoldLaBaAttend(groupId, logList, int32(rewGold), info.Times)
	pl.SendMsg(scMsg)
	return
}
