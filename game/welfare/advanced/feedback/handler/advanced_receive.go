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
	advancedfeedbacktypes "fgame/fgame/game/welfare/advanced/feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ACTIVITY_ADVANCED_REWARDS_TYPE), dispatch.HandlerFunc(handlerAdvancedReceive))
}

//处理领取升阶返利（升阶消耗）
func handlerAdvancedReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理升阶返利领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMergeActivityAdvancedRewards := msg.(*uipb.CSMergeActivityAdvancedRewards)
	rewId := csMergeActivityAdvancedRewards.GetRewId()

	err = receiveAdvanced(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理升阶返利领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理升阶返利领取奖励请求完成")

	return
}

//升阶返利领取奖励请求逻辑
func receiveAdvanced(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAdvanced
	subType := welfaretypes.OpenActivityAdvancedSubTypeFeedback
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取升阶返利奖励请求，模板不存在")
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
			}).Warn("welfare:领取升阶返利奖励请求，活动不存在")
		return
	}

	davancedDay := openTemp.Value1
	curAdvancedDay := welfarelogic.CountCurActivityDay(openTemp.Group)
	if davancedDay != curAdvancedDay {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"rewId":          rewId,
				"curAdvancedDay": curAdvancedDay,
			}).Warn("welfare:领取升阶返利奖励请求,升级返利类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	info := obj.GetActivityData().(*advancedfeedbacktypes.AdvancedInfo)
	needDanNum := openTemp.Value2

	//领取条件
	if !info.IsCanReceiveRewards(needDanNum) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"needDanNum": needDanNum,
				"rewId":      rewId,
			}).Warn("welfare:领取升阶返利奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needDanNum)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	//公告
	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(openTemp.Group)
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MergeActivityAdvancedRewardsNotice), plName, itemNameLinkStr, link)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	scMergeActivityAdvancedRewards := pbutil.BuildSCMergeActivityAdvancedRewards(totalRewData, rewItemMap)
	pl.SendMsg(scMergeActivityAdvancedRewards)

	return
}
