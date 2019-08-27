package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardscosttypes "fgame/fgame/game/welfare/rewards/cost/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_COST_REWARDS_COST_INFO_TYPE), dispatch.HandlerFunc(handlerCostRewardsGetInfo))
}

//处理获取每消费奖励信息
func handlerCostRewardsGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取每消费奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityCostRewardsCostInfo)
	groupId := csMsg.GetGroupId()

	err = getCostRewardsInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取每消费奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取每消费奖励请求完成")

	return
}

//获取每消费奖励请求逻辑
func getCostRewardsInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeRewards
	subType := welfaretypes.OpenActivityRewardsSubTypeCost
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

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityDataByGroupId(groupId)

	costGold := int64(0)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj != nil {
		info := obj.GetActivityData().(*rewardscosttypes.CostRewInfo)
		costGold = info.GoldNum
	}

	scMsg := pbutil.BuildSCOpenActivityCostRewardsCostInfo(groupId, costGold)
	pl.SendMsg(scMsg)
	return
}
