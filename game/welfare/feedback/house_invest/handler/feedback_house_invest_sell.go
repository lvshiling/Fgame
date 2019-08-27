package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	constanttypes "fgame/fgame/game/constant/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	feedbackhouseinvesttemplate "fgame/fgame/game/welfare/feedback/house_invest/template"
	feedbackhouseinvesttypes "fgame/fgame/game/welfare/feedback/house_invest/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FEEDBACK_HOUSE_INVEST_SELL_TYPE), dispatch.HandlerFunc(handlerFeedbackHouseInvestSell))
}

//处理房产投资出售
func handlerFeedbackHouseInvestSell(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理房产投资出售领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityFeedbackHouseInvestSell := msg.(*uipb.CSOpenActivityFeedbackHouseInvestSell)
	groupId := csOpenActivityFeedbackHouseInvestSell.GetGroupId()

	err = sellFeedbackHouseInvest(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理房产投资出售领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理房产投资出售领取奖励请求完成")

	return
}

//房产投资出售领取奖励请求逻辑
func sellFeedbackHouseInvest(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeHouseInvest
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参数无效,活动时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:房产投资出售请求错误，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:房产投资出售请求错误，活动不存在")
		return
	}
	info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)

	//是否激活了
	if !info.IsActivity {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:房产投资出售请求，还未激活")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否出售了
	if info.IsSell {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:房产投资出售请求，房产已经出售了")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:房产投资出售请求，活动模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupTemp := groupInterface.(*feedbackhouseinvesttemplate.GroupTemplateHouseInvest)

	openTemp := groupTemp.GetOpenActivityHouseInvest(info.DecorDays)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:房产投资出售请求，活动模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	reasonGold := commonlog.GoldLogReasonOpenActivityRew
	reasonSilver := commonlog.SilverLogReasonOpenActivityRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), typ, subType)
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), typ, subType)

	rewSilver := int64(0)
	rewBindGold := int64(0)
	rewGold := int64(0)
	sellNum := groupTemp.GetOpenActivityHouseInvestSellNum(info.DecorDays)
	rewSilver += sellNum
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.AddMoney(rewBindGold, rewGold, reasonGold, reasonGoldText, rewSilver, reasonSilver, reasonSilverText)
	if !flag {
		panic(fmt.Errorf("welfare:房产投资出售请求, 加钱应该是ok的"))
	}

	totalRewData := propertytypes.CreateRewData(0, 0, int32(rewSilver), 0, 0)
	rewItemMap := make(map[int32]int32)
	rewItemMap[constanttypes.SilverItem] = int32(rewSilver)
	// totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	// if !flag {
	// 	return
	// }

	//更新信息
	info.IsSell = true
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityFeedbackHouseInvestSell(groupId, totalRewData, rewItemMap)
	pl.SendMsg(scMsg)

	return
}
