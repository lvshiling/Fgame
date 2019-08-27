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
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	newleveleventtypes "fgame/fgame/game/welfare/invest/new_level/event/types"
	investnewleveltemplate "fgame/fgame/game/welfare/invest/new_level/template"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_INVEST_NEW_LEVEL_BUY_TYPE), dispatch.HandlerFunc(handlerBuyInvestNewLevel))
}

//处理购买新等级投资计划
func handlerBuyInvestNewLevel(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理购买新等级投资计划请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	investNewLevelBuy := msg.(*uipb.CSOpenActivityInvestNewLevelBuy)
	typ := investNewLevelBuy.GetTyp()
	groupId := investNewLevelBuy.GetGroupId()

	investNewLevelType := investnewleveltypes.InvestNewLevelType(typ)
	if !investNewLevelType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("welfare:购买新等级投资计划请求，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = buyInvestNewLevel(tpl, investNewLevelType, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理购买新等级投资计划请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理购买新等级投资计划请求完成")

	return
}

//购买新等级投资计划请求逻辑
func buyInvestNewLevel(pl player.Player, investLevelType investnewleveltypes.InvestNewLevelType, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewLevel

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:购买新等级投资计划请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"investLevelType": investLevelType,
				"groupId":         groupId,
			}).Warn("welfare:购买投资计划请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//首充用户
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	if !welfareManager.IsFirstCharge() {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"investLevelType": investLevelType,
				"groupId":         groupId,
			}).Warn("welfare:购买投资计划请求，需要首充")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotFirstChargeUser)
		return
	}

	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)
	//是否购买
	if info.IsBuy() {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"groupId":         groupId,
				"investLevelType": investLevelType,
			}).Warn("welfare:购买投资计划请求，已购买")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityHadBuyInvest)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	groupTemp := groupInterface.(*investnewleveltemplate.GroupTemplateInvestNewLevel)
	needGold := groupTemp.GetInvestLevelNeedGold(investLevelType)
	//元宝是否足够
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:购买投资计划请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}
	//消耗日志
	goldReason, ok := investLevelType.ConvertToInvesetNewLevelCostLogType()
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"investLevelType": investLevelType,
			}).Warn("welfare:购买投资计划请求，消耗日志类型应该存在")
		return
	}
	goldReasonText := goldReason.String()
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic(fmt.Errorf("welfare: buy invest use gold should be ok"))
	}

	//更新信息
	investRewList := []int32{}
	info.InvestBuyInfoMap[investLevelType] = investRewList
	welfareManager.UpdateObj(obj)

	gameevent.Emit(newleveleventtypes.EventTypeInvestNewLevelBuy, pl, investLevelType)

	//公告
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityInvestLevelRewardsNotice), playerName, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scOpenActivityInvestNewLevelBuy := pbutil.BuildSCOpenActivityInvestNewLevelBuy()
	pl.SendMsg(scOpenActivityInvestNewLevelBuy)
	return
}
