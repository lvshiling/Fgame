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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_RANK_ADVANCED_TYPE), dispatch.HandlerFunc(handleRankAdvancedList))
}

//处理进阶排行榜信息
func handleRankAdvancedList(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取进阶排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankAdvancedGet := msg.(*uipb.CSOpenActivityRankAdvancedList)
	page := csRankAdvancedGet.GetPage()
	groupId := csRankAdvancedGet.GetGroupId()

	err = rankAdvancedList(tpl, page, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"error":    err,
			}).Error("welfare:处理获取进阶排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理获取进阶排行榜消息完成")
	return nil

}

//获取进阶排行榜界面信息的逻辑
func rankAdvancedList(pl player.Player, page int32, groupId int32) (err error) {
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
	h := welfare.GetRankSystemDataHandler(typ, subType)
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
	scMsg := pbutil.BuildSCOpenActivityRankAdvancedList(page, rankList, groupId, int32(typ), subType.SubType(), rankTime, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
