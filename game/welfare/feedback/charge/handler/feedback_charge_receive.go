package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FEEDBACK_CHARGE_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerFeedbackChargeReceive))
}

//处理领取充值返利
func handlerFeedbackChargeReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理充值返利领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityFeedbackChargeReceiveRew := msg.(*uipb.CSOpenActivityFeedbackChargeReceiveRew)
	rewId := csOpenActivityFeedbackChargeReceiveRew.GetRewId()

	err = receiveChargeFeedback(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理充值返利领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理充值返利领取奖励请求完成")

	return
}

//充值返利领取奖励请求逻辑
func receiveChargeFeedback(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeCharge
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取充值返利奖励请求，模板不存在")
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
			}).Warn("welfare:领取充值返利奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
	needGoldNum := openTemp.Value1

	//领取条件
	if !info.IsCanReceiveRewards(needGoldNum) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
				"goldNum":  needGoldNum,
				"group":    openTemp.Group,
			}).Warn("welfare:领取充值返利奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needGoldNum)
	welfareManager.UpdateObj(obj)

	//公告
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(openTemp.Group)
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		acName := chatlogic.FormatModuleNameNoticeStr(openTemp.Label)
		glodNum := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGoldNum))
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackChargeRewardsNotice), plName, acName, glodNum, itemNameLinkStr, link)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityFeedbackChargeReceiveRew(totalRewData, rewItemMap, rewId)
	pl.SendMsg(scMsg)

	return
}
