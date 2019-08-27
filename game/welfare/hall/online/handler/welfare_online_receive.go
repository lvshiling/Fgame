package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	hallonlinetypes "fgame/fgame/game/welfare/hall/online/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_WELFARE_ONLINE_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerWelfareOnlineReceive))
}

//处理福利领奖
func handlerWelfareOnlineReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理领取在线福利奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityWelfareOnlineReceiveRew := msg.(*uipb.CSOpenActivityWelfareOnlineReceiveRew)
	rewId := csOpenActivityWelfareOnlineReceiveRew.GetRewId()

	err = welfareOnlineReceiveRew(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理领取在线福利奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理领取在线福利奖励请求完成")

	return
}

//领取福利请求逻辑
func welfareOnlineReceiveRew(pl player.Player, rewId int32) (err error) {
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取在线福利奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//活动时间判断
	groupId := openTemp.Group
	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取在线福利奖励请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	//领取条件
	rewOnlineTimeSecond := openTemp.Value1
	onlineTime := pl.GetTodayOnlineTime()
	rewOnlineTime := int64(rewOnlineTimeSecond) * int64(common.SECOND)
	if onlineTime < rewOnlineTime {
		log.WithFields(
			log.Fields{
				"playerId":            pl.GetId(),
				"onlineTime":          onlineTime,
				"rewOnlineTimeSecond": rewOnlineTimeSecond,
			}).Warn("welfare:领取在线福利奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.PlayerOnlineTimeNoEnough)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取在线福利奖励请求，活动不存在")
		return
	}

	info := obj.GetActivityData().(*hallonlinetypes.WelfareOnlineInfo)
	if info.IsReceive(rewOnlineTimeSecond) {
		log.WithFields(
			log.Fields{
				"playerId":            pl.GetId(),
				"rewOnlineTimeSecond": rewOnlineTimeSecond,
			}).Warn("welfare:领取在线福利奖励请求，已领取过奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新
	info.AddRecord(rewOnlineTimeSecond)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scWelfareOnlineRew := pbutil.BuildSCOpenActivityWelfareOnlineReceiveRew(totalRewData, rewItemMap)
	pl.SendMsg(scWelfareOnlineRew)
	return
}
