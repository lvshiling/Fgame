package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/rank/pbutil"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RANK_FORCE_GET_TYPE), dispatch.HandlerFunc(handleRankForceGet))
}

//处理战力排行榜信息
func handleRankForceGet(s session.Session, msg interface{}) (err error) {
	log.Debug("rank:处理获取战力排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankForceGet := msg.(*uipb.CSRankForceGet)
	page := csRankForceGet.GetPage()
	isArea := csRankForceGet.GetIsArea()

	err = rankForceGet(tpl, page, isArea)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"isArea":   isArea,
				"error":    err,
			}).Error("rank:处理获取战力排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("rank:处理获取战力排行榜消息完成")
	return nil

}

//获取战力排行榜界面信息的逻辑
func rankForceGet(pl player.Player, page int32, isArea bool) (err error) {
	if page < 0 {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"page":     page,
			"isArea":   isArea,
		}).Warn("rank:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	var classType ranktypes.RankClassType
	if isArea {
		classType = ranktypes.RankClassTypeArea
	} else {
		classType = ranktypes.RankClassTypeLocal
	}
	showServer := false
	if merge.GetMergeService().GetMergeTime() != 0 {
		showServer = true
	}
	forceList, rankTime := rank.GetRankService().GetForceListByPage(classType, 0, page)
	scMountGet := pbutil.BuildSCRankForceGet(showServer, isArea, page, forceList, rankTime)
	pl.SendMsg(scMountGet)
	return
}
