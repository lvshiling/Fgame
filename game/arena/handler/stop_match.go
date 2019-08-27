package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	arenalogic "fgame/fgame/game/arena/logic"
	"fgame/fgame/game/arena/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_STOP_MATCH_TYPE), dispatch.HandlerFunc(handleArenaStopMatch))
}

//处理3v3停止匹配
func handleArenaStopMatch(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理3v3停止匹配")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaStopMatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理3v3停止匹配,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理3v3停止匹配")
	return nil

}

//3v3匹配
func arenaStopMatch(pl player.Player) (err error) {
	//TODO 判断是否有组队
	arenalogic.ArenaStopMatchSend(pl)

	scArenaStopMatch := pbutil.BuildSCArenaStopMatch()
	pl.SendMsg(scArenaStopMatch)
	return
}
