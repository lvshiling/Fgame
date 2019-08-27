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
	"fgame/fgame/game/shenmo/pbutil"
	"fgame/fgame/game/shenmo/shenmo"
	"fgame/fgame/game/shenmo/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENMO_RANK_GET_TYPE), dispatch.HandlerFunc(handleShenMoRankGet))
}

//处理神魔战场周排名
func handleShenMoRankGet(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场周排行榜")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenMoRankGet := msg.(*uipb.CSShenMoRankGet)
	isThis := csShenMoRankGet.GetIsThis()
	page := csShenMoRankGet.GetPage()

	err = shenMoRankGet(tpl, isThis, page)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isThis":   isThis,
				"page":     page,
			}).Error("shenmo:处理神魔战场周排行榜,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"isThis": isThis,
			"page":   page,
		}).Debug("shenmo:处理神魔战场周排行榜")
	return nil

}

//处理神魔战场周排行榜
func shenMoRankGet(pl player.Player, isThis bool, page int32) (err error) {
	if page < 0 {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"page":     page,
			"isThis":   isThis,
		}).Warn("shenmo:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	rankType := types.RankTimeTypeLast
	if isThis {
		rankType = types.RankTimeTypeThis
	}
	dataList, rankTime := shenmo.GetShenMoService().GetRankList(rankType, page)
	scShenMoRankGet := pbutil.BuildSCShenMoRankGet(isThis, page, rankTime, dataList)
	pl.SendMsg(scShenMoRankGet)
	return
}
