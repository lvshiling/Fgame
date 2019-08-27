package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	feedbackchargecycletemplate "fgame/fgame/game/welfare/feedback/charge_cycle/template"
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FEEDBACK_CYCLE_CHARGE_REWARDS_TYPE), dispatch.HandlerFunc(handlerFeedbackCycleChargeReceive))
}

//处理连续充值
func handlerFeedbackCycleChargeReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理连续充值领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityFeedbackCycleChargeRewards := msg.(*uipb.CSOpenActivityFeedbackCycleChargeRewards)
	rewId := csOpenActivityFeedbackCycleChargeRewards.GetRewId()

	err = receiveFeedbackCycleCharge(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理连续充值领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理连续充值领取奖励请求完成")

	return
}

//连续充值领取奖励请求逻辑
func receiveFeedbackCycleCharge(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeCycleCharge
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:连续充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, openTemp.Group)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(openTemp.Group)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  openTemp.Group,
			}).Warn("welfare:连续充值奖励请求，活动不存在")
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(openTemp.Group)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargecycletemplate.GroupTemplateCycleCharge)
	info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
	isCanReceive := false
	switch feedbackchargecycletypes.FeedbackCycleRewType(openTemp.Value1) {
	case feedbackchargecycletypes.FeedbackCycleRewTypeSingleDay:
		{
			//奖励日
			rewDay := openTemp.Value2
			if info.CycleDay != rewDay {
				log.WithFields(
					log.Fields{
						"playerId":  pl.GetId(),
						"rewId":     rewId,
						"rewDay":    rewDay,
						"curRewDay": info.CycleDay,
					}).Warn("welfare:连续充值奖励请求,奖励日错误")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
				return
			}

			needCharge := groupTemp.GetDayRewCondition(info.CycleDay)
			isCanReceive = info.IsCanReceiveToday(needCharge)
		}
	case feedbackchargecycletypes.FeedbackCycleRewTypeCountDay:
		{
			isCanReceive = info.IsCanReceiveCountDay(openTemp.Value2)
		}
	}

	//领取条件
	if !isCanReceive {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"chargeNumToday": info.CurDayChargeNum,
				"countDay":       info.DayNum,
				"rewId":          rewId,
			}).Warn("welfare:连续充值奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	if openTemp.Value1 == int32(feedbackchargecycletypes.FeedbackCycleRewTypeSingleDay) {
		info.ReceiveToday()
	} else {
		info.ReceiveCountDay(openTemp.Value2)
	}
	welfareManager.UpdateObj(obj)

	//公告
	// itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	// if len(itemNameLinkStr) > 0 {
	// 	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	// 	glodNum := coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("%d", needGold))
	// 	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MergeActivityCycleChargeRewardsNotice), plName, glodNum, itemNameLinkStr)
	// 	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	// 	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	// }

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityFeedbackCycleChargeRewards(totalRewData, rewItemMap)
	pl.SendMsg(scMsg)

	return
}
