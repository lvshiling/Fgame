package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RESET_RELIVETIMES_TYPE), dispatch.HandlerFunc(handleArenapvpResetRelive))
}

//处理重置复活次数
func handleArenapvpResetRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服pvp重置复活次数")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = arenapvpResetRelive(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服pvp重置复活次数,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服pvp重置复活次数,完成")
	return nil

}

//pvp复活
func arenapvpResetRelive(pl player.Player) (err error) {
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.ResetRelive()
	return
}
