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
	cyclechargetypes "fgame/fgame/game/welfare/cycle/charge/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_CYCLE_CHARGE_REWARDS_TYPE), dispatch.HandlerFunc(handlerCycleReceive))
}

//处理领取每日首充
func handlerCycleReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理每日首充领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityCycleChargeRewards := msg.(*uipb.CSOpenActivityCycleChargeRewards)
	rewId := csOpenActivityCycleChargeRewards.GetRewId()

	err = receiveCycle(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理每日首充领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理每日首充领取奖励请求完成")

	return
}

//每日首充领取奖励请求逻辑
func receiveCycle(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeCharge
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取每日首充奖励请求，模板不存在")
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
			}).Warn("welfare:领取每日首充奖励请求，活动不存在")
		return
	}

	cycDay := openTemp.Value1
	curCycDay := welfarelogic.CountCycleDay(openTemp.Group)
	if cycDay != curCycDay {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"rewId":     rewId,
				"cycDay":    cycDay,
				"curCycDay": curCycDay,
			}).Warn("welfare:领取每日首充奖励请求,充值日类型错误")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
		return
	}

	info := obj.GetActivityData().(*cyclechargetypes.CycleChargeInfo)
	needGoldNum := openTemp.Value2

	//领取条件
	if !info.IsCanReceiveRewards(needGoldNum) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"needGoldNum": needGoldNum,
				"curCharge":   info.GoldNum,
				"record":      info.RewRecord,
			}).Warn("welfare:领取每日首充奖励请求，不满足领取条件")
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
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(openTemp.Group)
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		acName := chatlogic.FormatModuleNameNoticeStr(openTemp.Label)
		glodNum := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGoldNum))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleChargeRewardsNotice), plName, acName, glodNum, itemNameLinkStr, link)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scOpenActivityCycleChargeRewards := pbutil.BuildSCOpenActivityCycleChargeRewards(rewId, totalRewData, rewItemMap)
	pl.SendMsg(scOpenActivityCycleChargeRewards)

	return
}
