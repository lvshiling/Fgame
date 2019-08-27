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
	"fgame/fgame/game/rank/pbutil"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RANK_XIANTI_GET_TYPE), dispatch.HandlerFunc(handleRankXianTiGet))
}

//处理法宝排行榜信息
func handleRankXianTiGet(s session.Session, msg interface{}) (err error) {
	log.Debug("rank:处理获取法宝排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankXianTiGet := msg.(*uipb.CSRankXianTiGet)
	page := csRankXianTiGet.GetPage()
	isArea := csRankXianTiGet.GetIsArea()

	err = rankXianTiGet(tpl, page, isArea)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"isArea":   isArea,
				"error":    err,
			}).Error("rank:处理获取法宝排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("rank:处理获取法宝排行榜消息完成")
	return nil

}

//获取法宝排行榜界面信息的逻辑
func rankXianTiGet(pl player.Player, page int32, isArea bool) (err error) {
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

	xianTiList, rankTime := rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeXianTi, classType, 0, page)
	scRankXianTiGet := pbutil.BuildSCRankXianTiGet(isArea, page, xianTiList, rankTime)
	pl.SendMsg(scRankXianTiGet)
	return
}
