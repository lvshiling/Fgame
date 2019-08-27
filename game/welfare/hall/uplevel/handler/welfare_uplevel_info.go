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
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_UPLEVEL_INFO_TYPE), dispatch.HandlerFunc(handlerUplevelGetInfo))
}

//处理获取活动信息
func handlerUplevelGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取信息请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityUplevelInfo := msg.(*uipb.CSOpenActivityUplevelInfo)
	groupId := csOpenActivityUplevelInfo.GetGroupId()

	err = getUplevelInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取信息请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取信息请求完成")

	return
}

//获取信息请求逻辑
func getUplevelInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeUpLevel

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	openTime := welfare.GetWelfareService().GetServerStartTime()
	mergeTime := welfare.GetWelfareService().GetServerMergeTime()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	level := pl.GetLevel()
	uplevelObj := welfareManager.GetOpenActivity(groupId)
	start, end := groupInterface.GetActivityTime(openTime, mergeTime)
	scMsg := pbutil.BuildSCOpenActivityUplevelInfo(uplevelObj, groupId, level, start, end)
	pl.SendMsg(scMsg)

	return
}
