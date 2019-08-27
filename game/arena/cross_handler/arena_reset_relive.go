package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	playerarena "fgame/fgame/game/arena/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_RESET_RELIVE_TIMES_TYPE), dispatch.HandlerFunc(handleArenaResetRelive))
}

//处理重置复活次数
func handleArenaResetRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3重置复活次数")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = arenaResetRelive(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3重置复活次数,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3重置复活次数,完成")
	return nil

}

//3v3复活
func arenaResetRelive(pl player.Player) (err error) {
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaManager.ResetRelive()
	return
}
