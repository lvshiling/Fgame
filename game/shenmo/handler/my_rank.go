package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenmo/pbutil"
	"fgame/fgame/game/shenmo/shenmo"
	"fgame/fgame/game/shenmo/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENMO_MY_RANK_TYPE), dispatch.HandlerFunc(handleShenMoMyRank))
}

//处理神魔战场我的周排名
func handleShenMoMyRank(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenMoMyRank := msg.(*uipb.CSShenMoMyRank)
	isThis := csShenMoMyRank.GetIsThis()

	err = shenMoMyRank(tpl, isThis)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isThis":   isThis,
			}).Error("shenmo:处理神魔战场取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理神魔战场取消排队")
	return nil

}

//处理神魔战场取消排队
func shenMoMyRank(pl player.Player, isThis bool) (err error) {
	allianceId := pl.GetAllianceId()
	rankType := types.RankTimeTypeLast
	if isThis {
		rankType = types.RankTimeTypeThis
	}

	rankTime := shenmo.GetShenMoService().GetRankTime(rankType)
	if allianceId == 0 {
		scShenMoMyRank := pbutil.BuildSCShenMoMyRank(isThis, 0, rankTime)
		pl.SendMsg(scShenMoMyRank)
		return
	}
	serverId := global.GetGame().GetServerIndex()
	pos, _ := shenmo.GetShenMoService().GetMyRank(rankType, serverId, allianceId)
	scShenMoMyRank := pbutil.BuildSCShenMoMyRank(isThis, pos, rankTime)
	pl.SendMsg(scShenMoMyRank)
	return
}
