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
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/pbutil"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_RANK_FEATHER_TYPE), dispatch.HandlerFunc(handleRankFeatherList))
}

//处理护体仙羽排行榜信息
func handleRankFeatherList(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取护体仙羽排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankFeatherGet := msg.(*uipb.CSOpenActivityRankFeatherList)
	page := csRankFeatherGet.GetPage()
	groupId := csRankFeatherGet.GetGroupId()

	err = rankFeatherList(tpl, page, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"error":    err,
			}).Error("welfare:处理获取护体仙羽排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理获取护体仙羽排行榜消息完成")
	return nil

}

//获取护体仙羽排行榜界面信息的逻辑
func rankFeatherList(pl player.Player, page int32, groupId int32) (err error) {
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
			}).Warn("welfare:参数无效,活动时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	featherList, rankTime := rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeFeather, ranktypes.RankClassTypeLocalActivity, groupId, page)
	if page == 0 {
		nextPageList, _ := rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeFeather, ranktypes.RankClassTypeLocalActivity, groupId, page+1)
		if len(nextPageList) > 0 {
			featherList = append(featherList, nextPageList[0])
		}
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scRankFeatherGet := pbutil.BuildSCOpenActivityRankFeatherList(page, featherList, rankTime, startTime, endTime)
	pl.SendMsg(scRankFeatherGet)
	return
}
