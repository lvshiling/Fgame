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
	advancedblessfeedbacktypes "fgame/fgame/game/welfare/advanced/bless_feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ACTIVITY_ADVANCED_BLESS_REWARDS_TYPE), dispatch.HandlerFunc(handlerAdvancedBlessReceive))
}

// 由advanced_rew_receive代替
//处理领取升阶祝福丹放送
func handlerAdvancedBlessReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理升阶祝福丹放送领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMergeActivityAdvancedBlessRewards)
	rewId := csMsg.GetRewId()

	err = receiveAdvancedBless(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理升阶祝福丹放送领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理升阶祝福丹放送领取奖励请求完成")

	return
}

//升阶祝福丹放送领取奖励请求逻辑
func receiveAdvancedBless(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAdvanced
	subType := welfaretypes.OpenActivityAdvancedSubTypeBlessFeedback
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取升阶祝福丹放送奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupId := openTemp.Group

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
			}).Warn("welfare:领取升阶祝福丹放送奖励请求，活动不存在")
		return
	}

	davancedDay := openTemp.Value1
	curAdvancedDay := welfarelogic.CountCurActivityDay(groupId)
	if davancedDay != curAdvancedDay {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"rewId":          rewId,
				"curAdvancedDay": curAdvancedDay,
			}).Warn("welfare:领取升阶祝福丹放送奖励请求,升级祝福丹放送类型错误")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
		return
	}

	//领取条件
	needAdvancedNum := openTemp.Value2
	info := obj.GetActivityData().(*advancedblessfeedbacktypes.BlessAdvancedInfo)
	if !info.IsCanReceiveRewards(needAdvancedNum) {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"needAdvancedNum": needAdvancedNum,
				"rewId":           rewId,
			}).Warn("welfare:领取升阶祝福丹放送奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needAdvancedNum)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	// //公告
	// plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	// itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	// if len(itemNameLinkStr) > 0 {
	// 	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MergeActivityAdvancedRewardsNotice), plName, itemNameLinkStr)
	// 	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	// 	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	// }

	scMsg := pbutil.BuildSCMergeActivityAdvancedBlessRewards(totalRewData, rewItemMap)
	pl.SendMsg(scMsg)

	return
}
