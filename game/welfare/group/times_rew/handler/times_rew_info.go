package handler

import (
	playerwelfare"fgame/fgame/game/welfare/player"
	playertypes"fgame/fgame/game/player/types"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_TIMES_REW_INFO_TYPE), dispatch.HandlerFunc(handlerTimesRewInfo))
}

//处理获取次数奖励信息
func handlerTimesRewInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取次数奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityTimesRewInfo)
	groupId := csMsg.GetGroupId()

	err = getTimesRewInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取次数奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取次数奖励请求完成")

	return
}

//获取次数奖励请求逻辑
func getTimesRewInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeGroup
	subType := welfaretypes.OpenActivityGroupSubTypeTimesRew
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
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	scMsg := pbutil.BuildSCOpenActivityTimesRewInfo(obj, groupId, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
