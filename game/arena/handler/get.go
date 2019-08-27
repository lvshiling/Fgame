package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PLAYER_ARENA_INFO_TYPE), dispatch.HandlerFunc(handleArenaInfo))
}

//处理3v3奖励次数
func handleArenaInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理3v3奖励次数")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理3v3奖励次数,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理3v3奖励次数,完成")
	return nil
}

//3v3奖励次数
func arenaInfo(pl player.Player) (err error) {
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaObj := arenaManager.GetPlayerArenaObjectByRefresh()
	scMsg := pbutil.BuildSCPlayerArenaInfo(arenaObj)
	pl.SendMsg(scMsg)
	return
}
