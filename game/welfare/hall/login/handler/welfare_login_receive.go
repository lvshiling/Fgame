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
	halllogintypes "fgame/fgame/game/welfare/hall/login/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_WELFARE_LOGIN_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerWelfareLoginReceive))
}

//处理福利领奖
func handlerWelfareLoginReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理领取登录福利奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityWelfareLoginReceiveRew := msg.(*uipb.CSOpenActivityWelfareLoginReceiveRew)
	rewId := csOpenActivityWelfareLoginReceiveRew.GetRewId()

	err = welfareLoginReceiveRew(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理领取登录福利奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理领取登录福利奖励请求完成")

	return
}

//领取福利请求逻辑
func welfareLoginReceiveRew(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeLogin
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取登录福利奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	groupId := openTemp.Group
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取登录福利奖励请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	//领取条件
	rewDay := openTemp.Value1
	curLoginDay := welfarelogic.CountWelfareLoginDay(pl.GetCreateTime())
	isCondition := curLoginDay >= rewDay
	if !isCondition {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"rewDay":      rewDay,
				"curLoginDay": curLoginDay,
			}).Warn("welfare:领取登录福利奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*halllogintypes.WelfareLoginInfo)
	if info.IsReceive(rewDay) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewDay":   rewDay,
			}).Warn("welfare:领取登录福利奖励请求，已领取过奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新
	info.AddRecord(rewDay)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scWelfareLoginRew := pbutil.BuildSCOpenActivityWelfareLoginReceiveRew(totalRewData, rewItemMap)
	pl.SendMsg(scWelfareLoginRew) 
	return
}
