package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_CYCLE_SINGLE_CHARGE_INFO_TYPE), dispatch.HandlerFunc(handlerCycleSingleChargeGetInfo))
}

//处理获取每日单笔充值信息
func handlerCycleSingleChargeGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取每日单笔充值请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityCycleSingleChargeInfo)
	groupId := csMsg.GetGroupId()

	err = getCycleSingleInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取每日单笔充值请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取每日单笔充值请求完成")

	return
}

//获取每日单笔充值请求逻辑
func getCycleSingleInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeSingleCharge

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	cycleObj := welfareManager.GetOpenActivity(groupId)
	if cycleObj != nil {
		err = welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			return
		}
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityCycleSingleChargeInfo(cycleObj, groupId, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
