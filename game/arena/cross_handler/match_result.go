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
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_MATCH_RESULT_TYPE), dispatch.HandlerFunc(handleArenaMatchResult))
}

//处理跨服3v3匹配结果
func handleArenaMatchResult(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3匹配结果")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isArenaMatchResult := msg.(*crosspb.ISArenaMatchResult)
	result := isArenaMatchResult.GetResult()
	err = arenaMatchResult(tpl, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3匹配结果,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3匹配结果,完成")
	return nil

}

//3v3匹配
func arenaMatchResult(pl player.Player, result bool) (err error) {
	if result {
		//通知其它队员
		err = team.GetTeamService().ArenaMatched(pl)
		if err != nil {
			return
		}
	} else {
		err = team.GetTeamService().ArenaMatchFailed(pl)
		if err != nil {
			return
		}
	}

	return
}
