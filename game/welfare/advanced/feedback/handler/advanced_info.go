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
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ACTIVITY_ADVANCED_INFO_TYPE), dispatch.HandlerFunc(handlerAdvancedGetInfo))
}

// 由common_handler的advanced_expend_return代替
//处理获取升阶返利信息（升阶消耗）
func handlerAdvancedGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取升阶返利请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMergeActivityAdvancedInfo := msg.(*uipb.CSMergeActivityAdvancedInfo)
	groupId := csMergeActivityAdvancedInfo.GetGroupId()

	err = getAdvancedInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取升阶返利请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取升阶返利请求完成")

	return
}

//获取升阶返利请求逻辑
func getAdvancedInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAdvanced
	subType := welfaretypes.OpenActivityAdvancedSubTypeFeedback
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
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	advancedObj := welfareManager.GetOpenActivity(groupId)
	scMergeActivityAdvancedInfo := pbutil.BuildSCMergeActivityAdvancedInfo(advancedObj, groupId, startTime, endTime)
	pl.SendMsg(scMergeActivityAdvancedInfo)
	return
}
