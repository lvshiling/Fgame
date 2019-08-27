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
	advancedrewrewtypes "fgame/fgame/game/welfare/advancedrew/rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_ADVANCED_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerAdvancedRewReceive))
}

//处理领取升阶奖励
func handlerAdvancedRewReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理升阶奖励领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityAdvancedReceiveRew)
	rewId := csMsg.GetRewId()

	err = receiveAdvancedRew(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理升阶奖励领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理升阶奖励领取奖励请求完成")

	return
}

//升阶奖励领取奖励请求逻辑
func receiveAdvancedRew(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRew
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取升阶奖励奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupId := openTemp.Group

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"groupId":  groupId,
			}).Warn("welfare:运营活动,不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取升阶奖励奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)

	advancedDay := welfaretypes.AdvancedType(openTemp.Value1)
	if advancedDay != info.RewType {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"rewId":           rewId,
				"advancedDay":     advancedDay,
				"curAdvancedType": info.RewType,
			}).Warn("welfare:领取升阶奖励奖励请求,升级奖励类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//领取条件
	needAdvancedNum := openTemp.Value2
	needChargeNum := openTemp.Value3
	if !info.IsCanReceiveRewards(needAdvancedNum, needChargeNum) {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"needChargeNum":   needChargeNum,
				"needAdvancedNum": needAdvancedNum,
				"rewId":           rewId,
			}).Warn("welfare:领取升阶奖励奖励请求，不满足领取条件")
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

	scMsg := pbutil.BuildSCOpenActivityAdvancedReceiveRew(totalRewData, rewItemMap, rewId, info.RewRecord)
	pl.SendMsg(scMsg)

	return
}
