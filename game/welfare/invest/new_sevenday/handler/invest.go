package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/global"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	investnewsevendaytemplate "fgame/fgame/game/welfare/invest/new_sevenday/template"
	investnewsevendaytypes "fgame/fgame/game/welfare/invest/new_sevenday/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_CS_NEW_INVEST_DAY_BUY_TYPE), dispatch.HandlerFunc(handlerBuyNewInvestDay))
}

//处理购买新七日投资
func handlerBuyNewInvestDay(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理购买新七日投资请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityNewInvestDayBuy := msg.(*uipb.CSOpenActivityNewInvestDayBuy)
	groupId := csOpenActivityNewInvestDayBuy.GetGroupId()
	investType := csOpenActivityNewInvestDayBuy.GetTyp()

	err = buyNewInvestDay(tpl, groupId, investType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理购买新七日投资请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理购买新七日投资请求完成")

	return
}

//购买新七日投资请求逻辑
func buyNewInvestDay(pl player.Player, groupId int32, investType int32) (err error) {
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewServenDay

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:购买新七日投资请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:新七日投资领取奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)
	if info.IsAlreadyBuy(investType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:购买新七日投资请求，已购买")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityHadBuyInvest)
		return
	}

	// 检查投资类型
	groupTemp := groupInterface.(*investnewsevendaytemplate.GroupTemplateNewInvestDay)
	if !groupTemp.CheckInvestType(investType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:投资类型不正确")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNewSevenInvestTypeWorry)
		return
	}
	// 检查元宝是否足够
	needGold := investType

	// 检查是否足够资格买
	if !info.IsEnoughMaxSigleChargeNum(groupTemp.GetMaxSigleChargeNum(investType)) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:最大单笔充值不够")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNewSevenInvestNotEnoughMaxSigle)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:购买七日投资请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyInvest
	goldReasonText := fmt.Sprintf(goldReason.String(), subType)
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic(fmt.Errorf("welfare: buy new seven day invest use gold should be ok"))
	}

	//更新信息
	now := global.GetGame().GetTimeService().Now()
	info.InitNewSevenDayInvest(investType, now)
	welfareManager.UpdateObj(obj)

	//公告
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityInvestDayRewardsNotice), playerName, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scOpenActivityNewInvestDayBuy := pbutil.BuildSCOpenActivityNewInvestDayBuy(groupId, investType)
	pl.SendMsg(scOpenActivityNewInvestDayBuy)

	return
}
