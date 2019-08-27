package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/pbutil"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_RANK_PROPERTY_TYPE), dispatch.HandlerFunc(handleRankPropertyList))
}

//处理属性排行榜信息
func handleRankPropertyList(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取属性排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankPropertyGet := msg.(*uipb.CSOpenActivityRankPropertyList)
	page := csRankPropertyGet.GetPage()
	groupId := csRankPropertyGet.GetGroupId()

	err = rankPropertyList(tpl, page, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"error":    err,
			}).Error("welfare:处理获取属性排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理获取属性排行榜消息完成")
	return nil

}

//获取属性排行榜界面信息的逻辑
func rankPropertyList(pl player.Player, page int32, groupId int32) (err error) {
	if page < 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"page":     page,
			}).Warn("welfare:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:活动时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	typ := timeTemp.GetOpenType()
	subType := timeTemp.GetOpenSubType()
	h := welfare.GetRankPropertyDataHandler(typ, subType)
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
				"typ":      typ,
				"subType":  subType,
			}).Warn("welfare:参数无效,排行榜处理器不存在")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityHandlerNotExist)
		return
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	rankList, rankTime := h.GetRankList(groupId, page)
	scMsg := pbutil.BuildSCOpenActivityRankPropertyList(page, rankList, groupId, int32(typ), subType.SubType(), rankTime, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
