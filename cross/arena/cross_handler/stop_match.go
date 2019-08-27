package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arena/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENA_STOP_MATCH_TYPE), dispatch.HandlerFunc(handleArenaStopMatch))
}

//停止匹配
func handleArenaStopMatch(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:停止3v3匹配")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenaStopMatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理3v3匹配,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理3v3匹配,完成")
	return nil

}

//停止3v3匹配
func arenaStopMatch(pl *player.Player) (err error) {
	//判断是否有组队
	success := arena.GetArenaService().ArenaStopMatch(pl.GetId())
	isArenaStopMatch := pbutil.BuildISArenaStopMatch(success)
	pl.SendMsg(isArenaStopMatch)
	return
}
