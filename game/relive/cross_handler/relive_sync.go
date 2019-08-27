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
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_RELIVE_SYNC_TYPE), dispatch.HandlerFunc(handlePlayerReliveSync))
}

//玩家复活同步
func handlePlayerReliveSync(s session.Session, msg interface{}) (err error) {
	log.Debug("relive:玩家跨服复活同步")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isPlayerReliveSync := msg.(*crosspb.ISPlayerReliveSync)
	reliveTime := isPlayerReliveSync.GetPlayerReliveData().GetCulTime()
	lastReliveTime := isPlayerReliveSync.GetPlayerReliveData().GetLastReliveTime()
	err = playerSyncRelive(tpl, reliveTime, lastReliveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"reliveTime":     reliveTime,
				"lastReliveTime": lastReliveTime,
				"error":          err,
			}).Error("relive:玩家跨服复活同步,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":       pl.GetId(),
			"reliveTime":     reliveTime,
			"lastReliveTime": lastReliveTime,
		}).Debug("relive:玩家跨服复活同步")
	return nil
}

//处理设置血池线逻辑
func playerSyncRelive(pl player.Player, reliveTime int32, lastReliveTime int64) (err error) {
	pl.SyncRelive(reliveTime, lastReliveTime)

	return
}
