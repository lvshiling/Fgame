package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_STOP_MATCH_TYPE), dispatch.HandlerFunc(handleArenaStopMatch))
}

//处理跨服3v3停止匹配
func handleArenaStopMatch(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3停止匹配")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isArenaStopMatch := msg.(*crosspb.ISArenaStopMatch)
	success := isArenaStopMatch.GetSuccess()
	err = arenaStopMatch(tpl, success)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3停止匹配,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3停止匹配")
	return nil

}

//3v3匹配
func arenaStopMatch(pl player.Player, success bool) (err error) {
	if !success {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Info("arena:处理跨服3v3停止匹配,失败")
		return
	}
	//停止匹配
	team.GetTeamService().ArenaStopMatch(pl)

	return
}
