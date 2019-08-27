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

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_CYCLE_CHARGE_INFO_TYPE), dispatch.HandlerFunc(handlerCycleChargeGetInfo))
}

//处理获取每日首充信息
func handlerCycleChargeGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取每日首充请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityCycleChargeInfo := msg.(*uipb.CSOpenActivityCycleChargeInfo)
	groupId := csOpenActivityCycleChargeInfo.GetGroupId()

	err = getCycleInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取每日首充请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取每日首充请求完成")

	return
}

//获取每日首充请求逻辑
func getCycleInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeCharge

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

	scOpenActivityCycleChargeInfo := pbutil.BuildSCOpenActivityCycleChargeInfo(cycleObj, groupId)
	pl.SendMsg(scOpenActivityCycleChargeInfo)
	return
}
