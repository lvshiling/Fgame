package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arena/arena"
	"fgame/fgame/game/arena/pbutil"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_RANK_GET_TYPE), dispatch.HandlerFunc(handleArenaRankGet))
}

//处理周排名
func handleArenaRankGet(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理周排行榜")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSArenaRankGet)
	typeInt := csMsg.GetTimeType()
	page := csMsg.GetPage()

	timeType := arenatypes.RankTimeType(typeInt)
	if !timeType.Vaild() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"timeType": timeType,
			}).Warn("arena:处理周排行榜,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = arenaRankGet(tpl, timeType, page)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"timeType": timeType,
				"page":     page,
			}).Error("arena:处理周排行榜,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"timeType": timeType,
			"page":     page,
		}).Debug("arena:处理周排行榜")
	return nil

}

//处理周排行榜
func arenaRankGet(pl player.Player, timeType arenatypes.RankTimeType, page int32) (err error) {
	if page < 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"page":     page,
				"timeType": timeType,
			}).Warn("arena:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	dataList, rankTime := arena.GetArenaService().GetRankList(timeType, page)
	scArenaRankGet := pbutil.BuildSCArenaRankGet(timeType, page, rankTime, dataList)
	pl.SendMsg(scArenaRankGet)
	return
}
