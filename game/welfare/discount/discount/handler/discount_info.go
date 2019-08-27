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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ACTIVITY_DISCOUNT_INFO_TYPE), dispatch.HandlerFunc(handlerDiscountGetInfo))
}

//处理获取限时礼包信息
func handlerDiscountGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取限时礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMergeActivityDiscountInfo := msg.(*uipb.CSMergeActivityDiscountInfo)
	groupId := csMergeActivityDiscountInfo.GetGroupId()

	err = getDiscountInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取限时礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取限时礼包请求完成")

	return
}

//获取限时礼包请求逻辑
func getDiscountInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeCommon
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参数无效,活动时间模板不存在")
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

	timesList := welfare.GetWelfareService().GetDiscountTimes(groupId)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	discountObj := welfareManager.GetOpenActivity(groupId)
	dicountDay := welfarelogic.CountCurActivityDay(groupId)
	scMsg := pbutil.BuildSCMergeActivityDiscountInfo(discountObj, groupId, dicountDay, timesList, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
