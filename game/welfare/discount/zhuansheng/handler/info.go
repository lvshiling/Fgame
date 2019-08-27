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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_DISCOUNT_ZHUANSHENG_INFO_TYPE), dispatch.HandlerFunc(handlerDiscountZhuanShengGetInfo))
}

//处理获取转生大礼包信息
func handlerDiscountZhuanShengGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取转生大礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityDiscountZhuanShengInfo)
	groupId := csMsg.GetGroupId()

	err = getDiscountZhuanShengInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取转生大礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"groupId":  groupId,
		}).Info("welfare：处理获取转生大礼包请求完成")

	return
}

//获取转生大礼包请求逻辑
func getDiscountZhuanShengInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参数无效,活动时间模板不存在")
		return
	}

	if !welfarelogic.CheckGroupId(pl, typ, subType, groupId) {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参数无效,参数错误，不是转生礼包活动")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	discountObj := welfareManager.GetOpenActivity(groupId)
	scMsg := pbutil.BuildSCOpenActivityDiscountZhuanShengInfo(discountObj, groupId, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
