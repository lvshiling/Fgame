package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_MATCH_TYPE), dispatch.HandlerFunc(handleArenaMatch))
}

//处理跨服3v3匹配
func handleArenaMatch(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3匹配")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaMatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3匹配,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3匹配,完成")
	return nil

}

//3v3匹配
func arenaMatch(pl player.Player) (err error) {
	//TODO
	return
}
