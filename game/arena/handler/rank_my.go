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
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_MY_RANK_TYPE), dispatch.HandlerFunc(handleArenaMyRank))
}

//处理我的周排名
func handleArenaMyRank(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSArenaMyRank)
	typeInt := csMsg.GetTimeType()

	timeType := arenatypes.RankTimeType(typeInt)
	if !timeType.Vaild() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"timeType": timeType,
			}).Warn("arena:处理取消排队")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = arenaMyRank(tpl, timeType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"timeType": timeType,
			}).Error("arena:处理取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理取消排队")
	return nil

}

//处理3v3我的排名
func arenaMyRank(pl player.Player, timeType arenatypes.RankTimeType) (err error) {
	rankTime := arena.GetArenaService().GetRankTime(timeType)
	serverId := global.GetGame().GetServerIndex()
	pos, val, _ := arena.GetArenaService().GetMyRank(timeType, serverId, pl.GetId())
	scMsg := pbutil.BuildSCArenaMyRank(int32(timeType), pos, val, rankTime)
	pl.SendMsg(scMsg)
	return
}
